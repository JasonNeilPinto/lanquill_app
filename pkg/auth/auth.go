package auth

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Lanquill/Forge/pkg/db"
	"github.com/Lanquill/Forge/pkg/resp"
	"github.com/Lanquill/go-logger"
	"github.com/bitly/go-simplejson"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

var (
	Authenticate func(next http.Handler) http.Handler
)
var TOKENKEY = []byte(os.Getenv("TOKEN_KEY"))

func DecodeJWT(tokenStr string) (*JWTData, error) {

	jwtdata := &JWTData{}

	tkn, err := jwt.ParseWithClaims(tokenStr, jwtdata, func(token *jwt.Token) (interface{}, error) {
		return TOKENKEY, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return jwtdata, err
		}
		return jwtdata, errors.New("unauthorized")
	}
	if !tkn.Valid {
		return jwtdata, errors.New("unauthorized")
	}

	return jwtdata, nil
}

func UnverifiedJwt(tokenStr string) (*JWTData, error) {

	jwtdata := &JWTData{}

	_, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwtdata)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return jwtdata, nil
}

func RequestAuth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx, _ := logger.GetContextLogger(r.Context())

			token, err := r.Cookie("token")
			jsonData := simplejson.New()
			if err == http.ErrNoCookie {
				jsonData.Set("error", resp.NoCookie)
				resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
				return
			}
			_, err = DecodeJWT(token.Value)
			if err != nil {
				jsonData.Set("error", resp.InvalidToken)
				resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func Auth(next http.Handler) http.Handler {
	return Authenticate(next)
}

func init() {
	Authenticate = RequestAuth()
}

func RoleAdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, log := logger.GetContextLogger(r.Context())

		token, err := r.Cookie("token")
		jsonData := simplejson.New()
		if err == http.ErrNoCookie {
			jsonData.Set("error", resp.NoCookie)
			resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
			return
		}
		res, err := DecodeJWT(token.Value)
		if err != nil {
			jsonData.Set("error", resp.InvalidToken)
			resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
			return
		}

		entityReq := Entity{}

		contentType := r.Header.Get("Content-type")
		if contentType == "application/json" {
			req, err := io.ReadAll(r.Body)
			if err != nil {
				// TODO: log exception
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}

			reader := bytes.NewReader(req)
			err = json.NewDecoder(reader).Decode(&entityReq)

			if err != nil {
				// TODO: log exception
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}

			// Replace the body with a new reader after reading from the original
			// This is needed as the request body will become empty once it is read. So we have to rewrite the data back to request body so that it will be available for further operations in th api
			// For more info: https://stackoverflow.com/questions/43021058/golang-read-request-body-multiple-times

			r.Body = io.NopCloser(bytes.NewBuffer(req))

		} else {
			r.ParseMultipartForm(1 << 20) // Mb
			entityIdFormValue := r.Form.Get("entityId")
			if entityIdFormValue != "" {
				entityId, err := strconv.ParseInt(r.Form.Get("entityId"), 0, 64)
				if err != nil {
					// TODO: log exception
					log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
					jsonData.Set("error", resp.InternalServerError)
					resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
					return
				}
				entityReq.EntityId = &entityId
			}

			entityDivIdFormValue := r.Form.Get("entityDivId")
			if entityDivIdFormValue != "" && entityDivIdFormValue != "null" {
				entityDivId, err := strconv.ParseInt(r.Form.Get("entityDivId"), 0, 64)
				if err != nil {
					// TODO: log exception
					log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
					jsonData.Set("error", resp.InternalServerError)
					resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
					return
				}
				entityReq.EntityDivId = &entityDivId
			}

			entityLevelFormValue := r.Form.Get("entityLevel")
			if entityLevelFormValue != "" {
				entityLevel, err := strconv.Atoi(entityLevelFormValue)
				if err != nil {
					// TODO: log exception
					log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
					jsonData.Set("error", resp.InternalServerError)
					resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
					return
				}
				entityReq.EntityLevel = &entityLevel
			}
		}

		if entityReq.EntityLevel != nil && *entityReq.EntityLevel == 0 {
			entityReq.EntityLevel = nil
		}

		//Unauthorize the request if there is no entity information and the user is not root admin
		if res.UserType != 2 && (entityReq.EntityId == nil || *entityReq.EntityId == int64(0)) && (entityReq.EntityDivId == nil || *entityReq.EntityDivId == int64(0)) && entityReq.EntityLevel == nil {
			jsonData.Set("error", resp.InvalidReqBody)
			resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
			return
		}

		if entityReq.EntityLevel == nil && entityReq.EntityId != nil && *entityReq.EntityId != int64(0) {
			var userTypeLevel = 0
			err = db.MySqlDB.QueryRow(SELECT_USER_TYPE_LEVEL, res.UserType).Scan(&userTypeLevel)
			if err != nil {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}

			var entityUserLevel = 0
			err = db.MySqlDB.QueryRow(SELECT_HIGHER_ENTITY_LEVEL, entityReq.EntityId).Scan(&entityUserLevel)
			if err != nil && err != sql.ErrNoRows {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			} else if err == sql.ErrNoRows {
				err = db.MySqlDB.QueryRow(SELECT_LOWER_ENTITY_LEVEL, entityReq.EntityId).Scan(&entityUserLevel)
				if err != nil {
					log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
					jsonData.Set("error", resp.InternalServerError)
					resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
					return
				}
			}

			// Unauthorize the request if user level is lower than the entity level for the requested entityId
			if userTypeLevel < entityUserLevel {
				jsonData.Set("error", resp.InvalidReqBody)
				resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
				return
			}
		}

		if entityReq.EntityLevel == nil && entityReq.EntityDivId != nil && *entityReq.EntityDivId != int64(0) {
			var userTypeLevel = 0
			err = db.MySqlDB.QueryRow(SELECT_USER_TYPE_LEVEL, res.UserType).Scan(&userTypeLevel)
			if err != nil {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}

			var entityUserLevel = 0
			err = db.MySqlDB.QueryRow(SELECT_LOWER_ENTITY_LEVEL, entityReq.EntityDivId).Scan(&entityUserLevel)
			if err != nil {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}

			// Unauthorize the request if user level is lower than the entity level for the requested entityId
			if userTypeLevel < entityUserLevel {
				jsonData.Set("error", resp.InvalidReqBody)
				resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
				return
			}
		}

		if entityReq.EntityLevel != nil {
			var userLevel = 0
			err = db.MySqlDB.QueryRow(SELECT_USER_TYPE_LEVEL, res.UserType).Scan(&userLevel)
			if err != nil {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}

			// Unauthorize the request if user level is lower than the requested entity level
			if userLevel < *entityReq.EntityLevel {
				jsonData.Set("error", resp.InvalidReqBody)
				resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
				return
			}

			stmtStr := ""

			if *entityReq.EntityLevel == 7 || *entityReq.EntityLevel == 6 || *entityReq.EntityLevel == 5 {
				stmtStr = SELECT_HIGHER_ENTITY_LEVEL
			} else if *entityReq.EntityLevel <= 4 && *entityReq.EntityLevel >= 1 {
				stmtStr = SELECT_LOWER_ENTITY_LEVEL
				// Need to fetch using entityDivId for lower level entities
				entityReq.EntityId = entityReq.EntityDivId
			} else {
				jsonData.Set("error", resp.InvalidReqBody)
				resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
				return
			}

			var entityLevel = 0
			err = db.MySqlDB.QueryRow(stmtStr, entityReq.EntityId).Scan(&entityLevel)
			if err != nil {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}

			// Unauthorize the request if the entity level for the given entityId is lower than the requested entity level
			if entityLevel < *entityReq.EntityLevel {
				jsonData.Set("error", resp.InvalidReqBody)
				resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func RoleUserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, log := logger.GetContextLogger(r.Context())

		token, err := r.Cookie("token")
		jsonData := simplejson.New()
		if err == http.ErrNoCookie {
			jsonData.Set("error", resp.NoCookie)
			resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
			return
		}
		res, err := DecodeJWT(token.Value)
		if err != nil {
			jsonData.Set("error", resp.InvalidToken)
			resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
			return
		}

		userReq := User{}

		req, err := io.ReadAll(r.Body)
		if err != nil {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
			return
		}

		reader := bytes.NewReader(req)

		// Replace the body with a new reader after reading from the original
		r.Body = io.NopCloser(bytes.NewBuffer(req))

		err = json.NewDecoder(reader).Decode(&userReq)
		if err != nil {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
			return
		}

		// Check if userId exists for delete
		if userReq.UserId != 0 {

			var userTypeLevel = 0
			err = db.MySqlDB.QueryRow(SELECT_USER_TYPE_LEVEL, res.UserType).Scan(&userTypeLevel)
			if err != nil {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}

			var reqUserLevel = 0
			err = db.MySqlDB.QueryRow(SELECT_USER_LEVEL, userReq.UserId).Scan(&reqUserLevel)
			if err != nil {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}

			// Unauthorize the request if user level is lower than the requested user level for the requested entityId
			if userTypeLevel < reqUserLevel {
				jsonData.Set("error", resp.InvalidReqBody)
				resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func ResetPasswordAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, log := logger.GetContextLogger(r.Context())

		token, err := r.Cookie("token")
		jsonData := simplejson.New()
		if err == http.ErrNoCookie {
			jsonData.Set("error", resp.NoCookie)
			resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
			return
		}
		res, err := DecodeJWT(token.Value)
		if err != nil {
			jsonData.Set("error", resp.InvalidToken)
			resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
			return
		}

		resetReq := ResetReq{}

		req, err := io.ReadAll(r.Body)
		if err != nil {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
			return
		}

		reader := bytes.NewReader(req)

		// Replace the body with a new reader after reading from the original
		r.Body = io.NopCloser(bytes.NewBuffer(req))

		err = json.NewDecoder(reader).Decode(&resetReq)
		if err != nil {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
			return
		}

		if resetReq.UserEmail != "" {

			userEntity := UserEntityResp{}
			err = db.MySqlDB.QueryRow(SELECT_USER_ENTITY_DETAILS, res.UserID).Scan(&userEntity.UserType,
				&userEntity.Level6,
				&userEntity.Level5,
				&userEntity.Level4,
				&userEntity.Level3,
				&userEntity.Level2,
				&userEntity.Level1)
			if err != nil {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}

			userDetails := SELECT_USER_DETAILS
			var entityId = int64(0)
			if userEntity.UserType == 2 || userEntity.UserType == 3 {
				userDetails = fmt.Sprintf(userDetails, "")
			} else if userEntity.UserType == 4 || userEntity.UserType == 5 || userEntity.UserType == 6 {
				userDetails = fmt.Sprintf(userDetails, "AND Level_6 = ?")
				entityId = *userEntity.Level6
			} else if userEntity.UserType == 7 || userEntity.UserType == 8 || userEntity.UserType == 9 {
				userDetails = fmt.Sprintf(userDetails, "AND Level_5 = ?")
				entityId = *userEntity.Level5
			} else if userEntity.UserType == 10 || userEntity.UserType == 11 || userEntity.UserType == 12 {
				userDetails = fmt.Sprintf(userDetails, "AND Level_4 = ?")
				entityId = *userEntity.Level4
			} else if userEntity.UserType == 13 || userEntity.UserType == 14 || userEntity.UserType == 15 {
				userDetails = fmt.Sprintf(userDetails, "AND Level_3 = ?")
				entityId = *userEntity.Level3
			} else if userEntity.UserType == 16 || userEntity.UserType == 17 || userEntity.UserType == 18 {
				userDetails = fmt.Sprintf(userDetails, "AND Level_2 = ?")
				entityId = *userEntity.Level2
			} else if userEntity.UserType == 19 || userEntity.UserType == 20 {
				userDetails = fmt.Sprintf(userDetails, "AND Level_1 = ?")
				entityId = *userEntity.Level1
			}

			var userCount = 0
			if userEntity.UserType == 2 || userEntity.UserType == 3 {
				err = db.MySqlDB.QueryRow(userDetails, resetReq.UserEmail).Scan(&userCount)
				if err != nil {
					log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
					jsonData.Set("error", resp.InternalServerError)
					resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
					return
				}
			} else {
				err = db.MySqlDB.QueryRow(userDetails, resetReq.UserEmail, entityId).Scan(&userCount)
				if err != nil {
					log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
					jsonData.Set("error", resp.InternalServerError)
					resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
					return
				}
			}

			// Unauthorize the request if the user doesn't exist for that entity
			if userCount == 0 {
				jsonData.Set("error", resp.InvalidEmail)
				resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func ChangeExpiryDateAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, log := logger.GetContextLogger(r.Context())

		token, err := r.Cookie("token")
		jsonData := simplejson.New()
		if err == http.ErrNoCookie {
			jsonData.Set("error", resp.NoCookie)
			resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
			return
		}
		res, err := DecodeJWT(token.Value)
		if err != nil {
			jsonData.Set("error", resp.InvalidToken)
			resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
			return
		}

		changeExpiryReq := ChangeExpiryDateReq{}

		req, err := io.ReadAll(r.Body)
		if err != nil {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
			return
		}

		reader := bytes.NewReader(req)

		// Replace the body with a new reader after reading from the original
		r.Body = io.NopCloser(bytes.NewBuffer(req))

		err = json.NewDecoder(reader).Decode(&changeExpiryReq)
		if err != nil {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
			return
		}

		if changeExpiryReq.RequestType == "" && changeExpiryReq.UserEmail == "" && changeExpiryReq.EntityId == nil && res.UserType != 2 {
			jsonData.Set("error", resp.InvalidReqBody)
			resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
			return
		}

		if changeExpiryReq.EntityId != nil && *changeExpiryReq.EntityId != int64(0) {
			var userTypeLevel = 0
			err = db.MySqlDB.QueryRow(SELECT_USER_TYPE_LEVEL, res.UserType).Scan(&userTypeLevel)
			if err != nil {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}

			var entityUserLevel = 0
			err = db.MySqlDB.QueryRow(SELECT_HIGHER_ENTITY_LEVEL, changeExpiryReq.EntityId).Scan(&entityUserLevel)
			if err != nil && err != sql.ErrNoRows {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			} else if err == sql.ErrNoRows {
				err = db.MySqlDB.QueryRow(SELECT_LOWER_ENTITY_LEVEL, changeExpiryReq.EntityId).Scan(&entityUserLevel)
				if err != nil {
					log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
					jsonData.Set("error", resp.InternalServerError)
					resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
					return
				}
			}

			// Unauthorize the request if user level is lower than the entity level for the requested entityId
			if userTypeLevel < entityUserLevel {
				jsonData.Set("error", resp.InvalidReqBody)
				resp.SendResponse(w, http.StatusUnauthorized, jsonData, &ctx)
				return
			}
		}

		if changeExpiryReq.UserEmail != "" {

			userEntity := UserEntityResp{}
			err = db.MySqlDB.QueryRow(SELECT_USER_ENTITY_DETAILS, res.UserID).Scan(&userEntity.UserType,
				&userEntity.Level6,
				&userEntity.Level5,
				&userEntity.Level4,
				&userEntity.Level3,
				&userEntity.Level2,
				&userEntity.Level1)
			if err != nil {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}

			userDetails := SELECT_USER_DETAILS
			var entityId = int64(0)
			if userEntity.UserType == 2 || userEntity.UserType == 3 {
				userDetails = fmt.Sprintf(userDetails, "")
			} else if userEntity.UserType == 4 || userEntity.UserType == 5 || userEntity.UserType == 6 {
				userDetails = fmt.Sprintf(userDetails, "AND Level_6 = ?")
				entityId = *userEntity.Level6
			} else if userEntity.UserType == 7 || userEntity.UserType == 8 || userEntity.UserType == 9 {
				userDetails = fmt.Sprintf(userDetails, "AND Level_5 = ?")
				entityId = *userEntity.Level5
			} else if userEntity.UserType == 10 || userEntity.UserType == 11 || userEntity.UserType == 12 {
				userDetails = fmt.Sprintf(userDetails, "AND Level_4 = ?")
				entityId = *userEntity.Level4
			} else if userEntity.UserType == 13 || userEntity.UserType == 14 || userEntity.UserType == 15 {
				userDetails = fmt.Sprintf(userDetails, "AND Level_3 = ?")
				entityId = *userEntity.Level3
			} else if userEntity.UserType == 16 || userEntity.UserType == 17 || userEntity.UserType == 18 {
				userDetails = fmt.Sprintf(userDetails, "AND Level_2 = ?")
				entityId = *userEntity.Level2
			} else if userEntity.UserType == 19 || userEntity.UserType == 20 {
				userDetails = fmt.Sprintf(userDetails, "AND Level_1 = ?")
				entityId = *userEntity.Level1
			}

			var userCount = 0
			if userEntity.UserType == 2 || userEntity.UserType == 3 {
				err = db.MySqlDB.QueryRow(userDetails, changeExpiryReq.UserEmail).Scan(&userCount)
				if err != nil {
					log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
					jsonData.Set("error", resp.InternalServerError)
					resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
					return
				}
			} else {
				err = db.MySqlDB.QueryRow(userDetails, changeExpiryReq.UserEmail, entityId).Scan(&userCount)
				if err != nil {
					log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
					jsonData.Set("error", resp.InternalServerError)
					resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
					return
				}
			}

			// Unauthorize the request if the user doesn't exist for that entity
			if userCount == 0 {
				jsonData.Set("error", resp.InvalidEmail)
				resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Lanquill/Forge/pkg/activity"
	"github.com/Lanquill/Forge/pkg/auth"
	"github.com/Lanquill/Forge/pkg/entity"
	"github.com/Lanquill/Forge/pkg/questions"
	"github.com/Lanquill/Forge/pkg/reports"
	"github.com/Lanquill/Forge/pkg/resp"
	"github.com/Lanquill/Forge/pkg/user"
	"github.com/Lanquill/Forge/pkg/email"
	"github.com/Lanquill/Forge/pkg/utils"
	"github.com/Lanquill/go-logger"
	"github.com/bitly/go-simplejson"
	"github.com/go-chi/chi/v5"
	"github.com/go-gota/gota/dataframe"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func InitRoutes(r *chi.Mux) {
	r.Group(func(r chi.Router) {
		r.Use(auth.RoleAdminAuth)
		r.Post("/entities", Entities)
		r.Post("/users", Users)
		r.Post("/entity/create", CreateEntity)
		r.Post("/entity/update", UpdateEntity) // Name and Logo
		r.Post("/entity/delete", DeleteEntity)
		r.Post("/user/create", CreateUser)
		r.Post("/user/create-bulk-user", CreateBulkUser)
		r.Post("/report/entity", EntityReport)
		r.HandleFunc("/report/entity/export", ExportEntityReport)
		r.HandleFunc("/report/entity/entity-users-export", ExportEntityUserReport)
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.RoleUserAuth)
		r.Post("/user/delete", DeleteUser)
		r.Post("/user/update", UpdateUser) // Name and Email
		r.Post("/report/user", UserReport)
		r.HandleFunc("/report/user/export", ExportUserReport)
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.ResetPasswordAuth)
		r.Post("/user/get-user-detail", UserDetails)
		r.Post("/user/reset-password", ResetPassword)
		r.HandleFunc("/user/search-user", SearchUser)
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.ChangeExpiryDateAuth)
		r.Post("/user/get-entity-detail", EntityDetails)
		r.Post("/user/change-expiry-date", ChangeExpiryDate)
	})

	r.Post("/add/speechscript", AddSpeechscript)
	r.Post("/get/speechscript", GetSpeechScripts)
	r.Post("/update/speechscript", UpdateSpeechScripts)
	r.Post("/delete/speechscript", DeleteSpeechScript)

	r.Post("/add/readscript", AddReadingScript)
	r.Post("/get/readscript", GetReadScripts)
	r.Post("/update/readscript", UpdateReadScripts)
	r.Post("/delete/readscript", DeleteReadScript)

	r.Post("/add/listeningscript", AddListeningScript)
	r.Post("/get/listeningscript", GetListeningScripts)
	r.Post("/update/listeningscript", UpdateListeningScripts)
	r.Post("/delete/listeningscript", DeleteListeningScript)

	r.Post("/add/listeningAudio", AddListeningAudio)

	r.Post("/add/grammar", AddGrammarScript)
	r.Post("/get/grammar", GetGrammarScripts)
	r.Post("/update/grammar", UpdateGrammarScript)
	r.Post("/delete/grammar", DeleteGrammarScript)

	r.Post("/add/vocabulary", AddVocabularyScript)
	r.Post("/get/vocabulary", GetVocabularyScripts)
	r.Post("/update/vocabulary", UpdateVocabularyScript)
	r.Post("/delete/vocabulary", DeleteVocabularyScript)

	r.Post("/export/entity-users", ExportUsers)
	r.Post("/certification-info", GetCertificationInfo)
	r.Post("/active-users", GetMostActiveUsers)

	r.Post("/map/IptoEntity", MapIptoEntity)
	r.Post("/get/IptoEntity", GetIPAddress)
	r.Post("/update/IptoEntity", UpdateIPAddress)
	r.Post("/delete/IptoEntity", DeleteIPAddress)

	r.Post("/user-metrics", GetUserActivtyMetrics)
	r.Post("/payment-info", GetPaymentInfo)

	r.Post("/export/certification-info", ExportCertificationInfo)
	r.Post("/retail-users", GetRetailUsers)

	r.Post("/export/loggedUsers", ExportLoggedInUsers)
	r.Post("/export/retail-users", ExportRetailUsers)
	r.Post("/export/successful-payments", ExportSuccessFullPaymentReport)
	r.Post("/export/failed-payments", ExportFailedPaymentReport)
	r.Post("/export/pending-payments", ExportPendingPaymentReport)

	r.Post("/txn-details", TransactionDetails)
	r.Post("/invoice-details", InvoiceDetails)
}

func Entities(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	entityReq := reports.EntityUserReq{}
	err := json.NewDecoder(r.Body).Decode(&entityReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error decoding request: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	response, err := entityReq.GetEntities()
	if err != nil {
		// TODO: log exception
		log.Error("Error fetching entities: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", response)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func CreateEntity(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	err := r.ParseMultipartForm(1 << 20) // 1Mb
	if err != nil {
		if err.Error() == "http: request body too large" {
			jsonData.Set("error", resp.FileTooLarge)
			resp.SendResponse(w, http.StatusBadRequest, jsonData, &ctx)
		} else {
			// TODO: log exception
			log.Error("Error parsing request: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		}
		return
	}

	file, handler, err := r.FormFile("logo")
	if err != nil {
		if err != http.ErrMissingFile {
			log.Error("Error finding logo: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
			return
		}
	} else {
		defer file.Close()
	}

	entityReq := entity.CreateEntityReq{}

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
		entityReq.EntityId = entityId
	}

	entityDivIdFormValue := r.Form.Get("entityDivId")
	if entityDivIdFormValue != "" {
		entityDivId, err := strconv.ParseInt(r.Form.Get("entityDivId"), 0, 64)
		if err != nil {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
			return
		}
		entityReq.EntityDivId = entityDivId
	}

	numOfUsers, err := strconv.ParseInt(r.Form.Get("numOfUsers"), 0, 64)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	entityReq.EntityName = r.Form.Get("entityName")
	entityReq.EntityTypeName = r.Form.Get("entityTypeName")
	entityReq.NumOfUsers = numOfUsers
	entityReq.ContactName = r.Form.Get("contactName")
	entityReq.ContactEmail = r.Form.Get("contactEmail")
	entityReq.ContactMob = r.Form.Get("contactMob")

	filename := ""
	if handler != nil {
		fileType := handler.Header.Get("Content-Type")
		if ext, ok := utils.AcceptedEntityLogoType[fileType]; ok {
			date := time.Now().String()
			filename = strings.Replace(date, " ", "_", -1)
			filename = strings.Replace(filename, "-", "_", -1)
			filename = strings.Replace(filename, ":", "_", -1)
			filename = strings.Split(filename, ".")[0] + ext
		} else {
			jsonData.Set("error", resp.InvalidLogoFileType)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

	}

	if entityReq.LogoPath == "" {
		if file != nil {
			entityReq.LogoPath = "/static/logo/" + filename
			f, err := os.OpenFile(utils.LOGOPATH+"/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}
			defer f.Close()
			io.Copy(f, file)
		}
	} else {
		entityReq.LogoPath = ""
	}

	insertedId, err := entityReq.CreateEntity()
	if err != nil {

		if err.Error() == entity.ErrEntityExists {
			jsonData.Set("error", entity.EntityExists)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == user.ErrNoUserLicense {
			jsonData.Set("error", entity.NoUserLicense)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	if file != nil {
		newFileName := strconv.FormatInt(insertedId, 10)
		err = os.Rename(utils.LOGOPATH+"/"+filename, utils.LOGOPATH+"/"+newFileName)
		if err != nil {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
			return
		}
	}

	jsonData.Set("data", entity.EntityCreateSuccessful)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func Users(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	userReq := reports.EntityUserReq{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	response, err := userReq.GetUsers()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", response)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	userReq := user.ResetPassReq{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	if userReq.Password == "" {
		jsonData.Set("error", resp.InvalidReqBody)
		resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
		return
	}

	err = userReq.ResetPassword()
	if err != nil {
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.ResetPasswordResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	userReq := user.CreateUserReq{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	if !user.ValidEmail(userReq.UserEmail) {
		jsonData.Set("error", user.InvalidEmail)
		resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
		return
	}

	err = userReq.Create()
	if err != nil {

		if err.Error() == user.ErrEmailOrMobExists {
			jsonData.Set("error", user.EmailOrMobExists)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == user.ErrNoUserLicense {
			jsonData.Set("error", user.NoUserLicense)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", user.UserCreateSuccessful)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func CreateBulkUser(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	tokenData, _ := utils.UserTokenData(token)
	jsonData := simplejson.New()

	r.ParseMultipartForm(1 << 20) // Mb

	file, _, err := r.FormFile("csv_file")
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	userReq := user.CreateUserReq{}

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
		userReq.EntityId = entityId
	}

	entityDivIdFormValue := r.Form.Get("entityDivId")
	if entityDivIdFormValue != "" {
		entityDivId, err := strconv.ParseInt(r.Form.Get("entityDivId"), 0, 64)
		if err != nil {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
			return
		}
		userReq.EntityDivId = entityDivId
	}

	docValidTill := r.Form.Get("docValidTill")
	certValidTill := r.Form.Get("certValidTill")

	userReq.Password = r.Form.Get("password")
	userReq.DocQuantity = 10
	userReq.CertQuantity = 1
	userReq.DocValidTill = docValidTill
	userReq.CertValidTill = certValidTill

	data := dataframe.ReadCSV(file)

	if data.Nrow() > 1000 {
		jsonData.Set("error", user.ExceedsMaxUser)
		resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
		return
	}

	if !data.Col("User_Name(*)").HasNaN() && !data.Col("User_Email(*)").HasNaN() {

		if userReq.EntityId != 0 {
			license, err := user.GetLicenseDetails(userReq.EntityId)
			if err != nil {
				// TODO: log exception
				log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}
			if int(license.NumberOfLicences-license.LicenceConsumed) < data.Nrow() {
				jsonData.Set("error", user.NoUserLicense)
				resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
				return
			}
		}

		go uploadUsers(userReq, data, tokenData.UserEmail)

		jsonData.Set("data", user.UserBulkCreateProcessing)
		resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
		return

	} else {
		jsonData.Set("error", user.UserBulkCreateFailed)
		resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
		return
	}

	jsonData.Set("data", user.UserCreateSuccessful)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
	return
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	userReq := user.UpdateUserReq{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	if userReq.Email == "" || userReq.Name == "" {
		jsonData.Set("error", resp.InvalidReqBody)
		resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
		return
	}

	err = userReq.Update()
	if err != nil {

		if err.Error() == user.ErrEmailExists {
			jsonData.Set("error", user.EmailExists)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", user.UserInfoUpdateSuccessful)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	userReq := user.UserIdReq{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = userReq.Delete()
	if err != nil {
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", user.UserDeleteSuccessful)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func UpdateEntity(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	// Reject any request with payload more than 1.5MB
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20+512)

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	entityReq := entity.UpdateEntityReq{}

	err := r.ParseMultipartForm(1 << 20) // 1Mb
	if err != nil {
		if err.Error() == "http: request body too large" {
			jsonData.Set("error", resp.FileTooLarge)
			resp.SendResponse(w, http.StatusBadRequest, jsonData, &ctx)
		} else {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		}
		return
	}

	entityId, err := strconv.ParseInt(r.Form.Get("entityId"), 0, 64)
	if err != nil {
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}
	entityReq.EntityId = entityId

	entityLevel, err := strconv.ParseInt(r.Form.Get("entityLevel"), 0, 64)
	if err != nil {
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}
	entityReq.EntityLevel = int(entityLevel)

	entityReq.Name = r.Form.Get("name")

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
		// If it's a lower entity then consider entityDivId for updating
		if entityReq.EntityLevel <= 5 && entityReq.EntityLevel >= 2 {
			entityReq.EntityId = entityDivId
		}
	}

	if entityReq.Name == "" || entityReq.EntityId == 0 || entityReq.EntityLevel == 0 {
		jsonData.Set("error", resp.InvalidReqBody)
		resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
		return
	}

	// Verify/Update logo only for higher entities
	if _, ok := utils.LevelWithLogo[entityReq.EntityLevel]; ok {
		file, handler, err := r.FormFile("logo")
		if err != nil {
			if err != http.ErrMissingFile {
				fmt.Println("Error Retrieving the File", err)
				jsonData.Set("error", resp.InternalServerError)
				resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
				return
			}
		} else {
			defer file.Close()
		}

		filename := ""
		if handler != nil {
			fileType := handler.Header.Get("Content-Type")
			if ext, ok := utils.AcceptedEntityLogoType[fileType]; ok {
				filename = r.Form.Get("entityId") + ext
			} else {
				jsonData.Set("error", resp.InvalidLogoFileType)
				resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
				return
			}

		}

		err = entityReq.UpdateHigherEntity(file, filename)
		if err != nil {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
			return
		}

	} else if _, ok := utils.LevelWithoutLogo[entityReq.EntityLevel]; ok {
		err = entityReq.UpdateLowerEntity()
	}

	if err != nil {

		if err.Error() == entity.ErrEntityExists {
			jsonData.Set("error", entity.EntityExists)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", user.UserInfoUpdateSuccessful)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func DeleteEntity(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	entityReq := entity.DeleteEntityReq{}
	err := json.NewDecoder(r.Body).Decode(&entityReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	if _, ok := utils.HigherLevelTypes[entityReq.EntityLevel]; ok {
		err = entityReq.DeleteHigherEntity()
	} else if _, ok := utils.MiddleLevelTypes[entityReq.EntityLevel]; ok {
		err = entityReq.DeleteHigherEntity()
	} else if _, ok := utils.LowerLevelTypes[entityReq.EntityLevel]; ok {
		err = entityReq.DeleteLowerEntity()
	}

	if err != nil {
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = entityReq.DeleteEntityUser()
	if err != nil {
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", entity.EntityDeleteSuccessful)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func EntityReport(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	entityReq := reports.EntityReportReq{}
	err := json.NewDecoder(r.Body).Decode(&entityReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	response, err := entityReq.GetEntityReport()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", response)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func UserReport(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	userReq := reports.UserReportReq{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	response, err := userReq.GetUserReport()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", response)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func ExportEntityReport(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	entityReq := reports.EntityReportReq{}
	err := json.NewDecoder(r.Body).Decode(&entityReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	filePath, err := entityReq.ExportEntityReport()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	// Copy the relevant headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	w.Write(fileBytes)
}

func ExportUserReport(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	userReq := reports.UserReportReq{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	filePath, err := userReq.ExportUserReport()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	// Copy the relevant headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	w.Write(fileBytes)
}

func ExportEntityUserReport(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	entityReq := reports.EntityReportReq{}
	err := json.NewDecoder(r.Body).Decode(&entityReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	filePath, err := entityReq.ExportEntityUserReport()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	// Copy the relevant headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	w.Write(fileBytes)
}

func ChangeExpiryDate(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	jsonData := simplejson.New()

	changeExpiryDateReq := user.ChangeExpiryDateReq{}
	err := json.NewDecoder(r.Body).Decode(&changeExpiryDateReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = changeExpiryDateReq.Update()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", user.ExpiryDateUpdateSuccessful)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func UserDetails(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	userReq := user.UserDetailReq{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	response, err := user.GetUserDetails(userReq.UserEmail)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", response)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func EntityDetails(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	jsonData := simplejson.New()

	entityDetailsReq := user.EntityDetailsReq{}
	err := json.NewDecoder(r.Body).Decode(&entityDetailsReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	response, err := entityDetailsReq.GetEntityDetails()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", response)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

// Purpose: To Add, Retrieve, Update, Delete Speech Scripts

func AddSpeechscript(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	tknData, _ := utils.UserTokenData(token)
	jsonData := simplejson.New()

	addSpeechReq := questions.SpeechScripts{}
	r.Body = http.MaxBytesReader(w, r.Body, 3<<20)
	err := r.ParseMultipartForm(3 << 20) // 1Mb
	if err != nil {
		if err.Error() == "http: request body too large" {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		} else {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		}
		return
	}

	file, handler, err := r.FormFile("audioFile")
	if err != nil {
		if err != http.ErrMissingFile {
			log.Error("Error Retrieving the File", zap.Error(err), logger.LogUserId(ctx))
			return
		}
	} else {
		defer file.Close()
	}

	filename := ""
	AudioFor := "speech"
	if handler != nil {
		fileType := handler.Header.Get("Content-Type")
		if ext, ok := utils.AcceptedAudioType[fileType]; ok {
			filename = filename + ext
		} else {
			return
		}
	}

	filename, err = questions.AddAudioFile(file, filename, AudioFor)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		return
	}
	addSpeechReq.AudioFile = filename

	addSpeechReq.Title = r.FormValue("scriptTitle")
	addSpeechReq.ScriptText = r.FormValue("scriptText")
	addSpeechReq.Grade = r.FormValue("grade")
	addSpeechReq.Complexity = r.FormValue("complexity")

	words := questions.Get_words(addSpeechReq.ScriptText)
	err = questions.StoreSpeechscript(tknData, addSpeechReq, words)
	if err != nil {

		if err.Error() == questions.ErrWordLimit {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.WordLimit)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func GetSpeechScripts(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	speechResp, err := questions.GetSpeechData()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", speechResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func UpdateSpeechScripts(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	speechReq := questions.SpeechResp{}

	r.Body = http.MaxBytesReader(w, r.Body, 3<<20)
	err := r.ParseMultipartForm(3 << 20) // 3Mb
	if err != nil {
		if err.Error() == "http: request body too large" {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		} else {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		}
		return
	}

	file, handler, err := r.FormFile("audioFile")
	if err != nil {
		if err != http.ErrMissingFile {
			log.Error("Error Retrieving the File", zap.Error(err), logger.LogUserId(ctx))
			return
		}
	} else {
		defer file.Close()
	}

	filename := ""
	AudioFor := "speech"
	if handler != nil {
		fileType := handler.Header.Get("Content-Type")
		if ext, ok := utils.AcceptedAudioType[fileType]; ok {
			filename = filename + ext
		} else {
			return
		}
	}

	filename, err = questions.AddAudioFile(file, filename, AudioFor)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		return
	}
	speechReq.AudioFile = filename

	scriptIdFormValue := r.Form.Get("id")
	if scriptIdFormValue != "" {
		scriptId, err := primitive.ObjectIDFromHex(scriptIdFormValue)
		if err != nil {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", resp.InternalServerError)
			resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
			return
		}
		speechReq.ID = scriptId
	}

	speechReq.Title = r.FormValue("scriptTitle")
	speechReq.ScriptText = r.FormValue("scriptText")
	speechReq.Grade = r.FormValue("grade")
	speechReq.Complexity = r.FormValue("complexity")

	err = questions.UpdateSpeech(speechReq)
	if err != nil {

		if err.Error() == questions.ErrWordLimit {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.WordLimit)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == questions.ErrUpdateFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.UpdateFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func DeleteSpeechScript(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	speechReq := questions.SpeechResp{}
	err := json.NewDecoder(r.Body).Decode(&speechReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = questions.DeleteSpeech(speechReq.ID)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}
	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

// Purpose: To Add, Retrieve, Update, Delete Reading Scripts

func AddReadingScript(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	tknData, _ := utils.UserTokenData(token)
	jsonData := simplejson.New()

	readingScriptsReq := questions.PassageScripts{}

	err := json.NewDecoder(r.Body).Decode(&readingScriptsReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = questions.InsertReadScript(readingScriptsReq, tknData.UserEmail)
	if err != nil {

		if err.Error() == questions.ErrWordLimit {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.WordLimit)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == questions.ErrPassageExists {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.PassageExists)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == questions.ErrCreateFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.CreateFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func GetReadScripts(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	readData, err := questions.GetReadScripts()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", readData)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func UpdateReadScripts(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	readReq := questions.PassageInfo{}
	err := json.NewDecoder(r.Body).Decode(&readReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = readReq.UpdateRead()
	if err != nil {

		if err.Error() == questions.ErrWordLimit {
			jsonData.Set("error", questions.WordLimit)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == questions.ErrUpdateFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.UpdateFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func DeleteReadScript(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	readReq := questions.PassageInfo{}
	err := json.NewDecoder(r.Body).Decode(&readReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = questions.DeleteReading(readReq.ID)
	if err != nil {
		if err.Error() == questions.ErrDeleteFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.DeleteFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}
	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

// Purpose: To Add Listening audio

func AddListeningAudio(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	r.Body = http.MaxBytesReader(w, r.Body, 3<<20)
	err := r.ParseMultipartForm(3 << 20) // 1Mb
	if err != nil {
		if err.Error() == "http: request body too large" {
			log.Error("Error: ", zap.Error(errors.New(err.Error())))
		} else {
			// TODO: log exception
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		}
	}

	file, handler, err := r.FormFile("audioFile")
	if err != nil {
		if err != http.ErrMissingFile {
			log.Error("Error Retrieving the File", zap.Error(err), logger.LogUserId(ctx))
		}
	} else {
		defer file.Close()
	}

	filename := ""
	AudioFor := "listen"
	if handler != nil {
		fileType := handler.Header.Get("Content-Type")
		if ext, ok := utils.AcceptedAudioType[fileType]; ok {
			filename = filename + ext
		} else {
			return
		}
	}

	filename, err = questions.AddAudioFile(file, filename, AudioFor)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", filename)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)

}

//  Purpose: To Add, Retrieve, Update, Delete Listening Scripts

func AddListeningScript(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	tknData, _ := utils.UserTokenData(token)
	jsonData := simplejson.New()

	listeningReq := questions.PassageScripts{}

	err := json.NewDecoder(r.Body).Decode(&listeningReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = questions.InsertListeningScript(listeningReq, tknData.UserEmail)
	if err != nil {

		if err.Error() == questions.ErrWordLimit {
			jsonData.Set("error", questions.WordLimit)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == questions.ErrPassageExists {
			jsonData.Set("error", questions.PassageExists)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == questions.ErrCreateFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.CreateFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func GetListeningScripts(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	readData, err := questions.GetListenScripts()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", readData)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func UpdateListeningScripts(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	listenReq := questions.PassageInfo{}

	err := json.NewDecoder(r.Body).Decode(&listenReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = questions.UpdateListen(listenReq)
	if err != nil {

		if err.Error() == questions.ErrWordLimit {
			jsonData.Set("error", questions.WordLimit)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == questions.ErrUpdateFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.UpdateFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func DeleteListeningScript(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	listenReq := questions.PassageInfo{}
	err := json.NewDecoder(r.Body).Decode(&listenReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = questions.DeleteListening(listenReq.ID)
	if err != nil {
		if err.Error() == questions.ErrDeleteFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.DeleteFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}
	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

// Purpose: To Add, Retrieve, Update, Delete Grammar Questions

func GetGrammarScripts(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	grammar, err := questions.GetGrammar()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", grammar)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func AddGrammarScript(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	tknData, _ := utils.UserTokenData(token)
	jsonData := simplejson.New()

	grammar := questions.Grammar{}

	err := json.NewDecoder(r.Body).Decode(&grammar)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = questions.AddGrammar(grammar, tknData)
	if err != nil {

		if err.Error() == questions.ErrWordLimit {
			jsonData.Set("error", questions.WordLimit)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == questions.ErrCreateFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.CreateFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func UpdateGrammarScript(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	grammar := questions.Grammar{}

	err := json.NewDecoder(r.Body).Decode(&grammar)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = questions.UpdateGrammar(grammar)
	if err != nil {

		if err.Error() == questions.ErrWordLimit {
			jsonData.Set("error", questions.WordLimit)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == questions.ErrUpdateFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.UpdateFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func DeleteGrammarScript(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	grammar := questions.Grammar{}

	err := json.NewDecoder(r.Body).Decode(&grammar)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = questions.DeleteGrammar(grammar.ID)
	if err != nil {

		if err.Error() == questions.ErrDeleteFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.DeleteFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

// Purpose: To Add, Retrieve, Update, Delete Vocabulary Questions

func AddVocabularyScript(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	vocabulary := questions.Vocalbulary{}

	err := json.NewDecoder(r.Body).Decode(&vocabulary)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = questions.AddVocabulary(vocabulary)
	if err != nil {

		if err.Error() == questions.ErrWordLimit {
			jsonData.Set("error", questions.WordLimit)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == questions.ErrPassageExists {
			jsonData.Set("error", questions.PassageExists)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == questions.ErrCreateFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.CreateFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func GetVocabularyScripts(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	vocabulary, err := questions.GetVocalbulary()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", vocabulary)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func UpdateVocabularyScript(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	vocabulary := questions.Vocalbulary{}

	err := json.NewDecoder(r.Body).Decode(&vocabulary)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = questions.UpdateVocalbulary(vocabulary)
	if err != nil {

		if err.Error() == questions.ErrWordLimit {
			jsonData.Set("error", questions.WordLimit)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		} else if err.Error() == questions.ErrUpdateFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.UpdateFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func DeleteVocabularyScript(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	vocabulary := questions.Vocalbulary{}

	err := json.NewDecoder(r.Body).Decode(&vocabulary)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = questions.DeleteVocabulary(vocabulary.QuestionID)
	if err != nil {

		if err.Error() == questions.ErrDeleteFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.DeleteFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func ExportUsers(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	userReq := reports.EntityUserReq{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	users, err := activity.GetUsersInfo(*userReq.EntityId, *userReq.EntityLevel, userReq.EntityDivId)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	filePath, err := activity.ExportEntityUsersReport(userReq.EntityId, users)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	//Copy the relevant headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	w.Write(fileBytes)
}

func GetCertificationInfo(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	response, total, err := activity.CertificationInfo()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("count", total)
	jsonData.Set("data", response)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func GetMostActiveUsers(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	response, err := activity.MostActiveUsersAsPerDocs()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", response)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func MapIptoEntity(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	entityReq := activity.EntityToIp{}
	err := json.NewDecoder(r.Body).Decode(&entityReq)
	if err != nil {

		if err.Error() == questions.ErrIpAddressExists {
			jsonData.Set("error", activity.IpAddressExists)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = entityReq.MapEntityToIps()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func GetIPAddress(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	response, err := activity.GetEntityIPAddresses()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", response)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func UpdateIPAddress(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	entityIpReq := activity.EntityIpReq{}
	err := json.NewDecoder(r.Body).Decode(&entityIpReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = entityIpReq.UpdateEntityIpAddress()
	if err != nil {

		if err.Error() == questions.ErrUpdateFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.UpdateFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func DeleteIPAddress(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	entityIpReq := activity.EntityIpReq{}
	err := json.NewDecoder(r.Body).Decode(&entityIpReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = entityIpReq.DeleteEntityIpAddress()
	if err != nil {

		if err.Error() == questions.ErrCreateFailure {
			log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
			jsonData.Set("error", questions.DeleteFailure)
			resp.SendResponse(w, resp.StatusInvalidRequestBody, jsonData, &ctx)
			return
		}

		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", resp.SuccessResp)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func GetUserActivtyMetrics(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	userMetrics, loggedInUsers, err := activity.GetUserMetrics()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("userActivtyMetrics", userMetrics)
	jsonData.Set("loggedInUsers", loggedInUsers)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func GetPaymentInfo(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	succesfull, failed, pending, other, err := activity.GetPaymentsInfo()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("succesfull", succesfull)
	jsonData.Set("failed", failed)
	jsonData.Set("pending", pending)
	jsonData.Set("other", other)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func ExportCertificationInfo(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	certification, _, err := activity.CertificationInfo()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	filePath, err := activity.ExportCertificationReport(certification)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		// TODO: log exception
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	// Copy the relevant headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	w.Write(fileBytes)
}

func ExportLoggedInUsers(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	_, loggedInUsers, err := activity.GetUserMetrics()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	filePath, err := activity.LoggedInUsersReport(loggedInUsers)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	// Copy the relevant headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	w.Write(fileBytes)
}

func GetRetailUsers(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	response, err := activity.RetailUsersInfo()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", response)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func ExportRetailUsers(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	retailUsers, err := activity.RetailUsersInfo()
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	filePath, err := activity.RetailUsersReport(retailUsers)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	// Copy the relevant headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	w.Write(fileBytes)
}

func ExportSuccessFullPaymentReport(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	filePath, err := activity.ExportPaymentInfo("success")
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		// TODO: log exception
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	// Copy the relevant headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	w.Write(fileBytes)
}

func ExportFailedPaymentReport(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	filePath, err := activity.ExportPaymentInfo("failure")
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	// Copy the relevant headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	w.Write(fileBytes)
}
func ExportPendingPaymentReport(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	filePath, err := activity.ExportPaymentInfo("pending")
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	// Copy the relevant headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	w.Write(fileBytes)
}

func TransactionDetails(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	orderId := activity.OrderIdReq{}

	err := json.NewDecoder(r.Body).Decode(&orderId)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	txnDetails, err := activity.GetTxnDetails(orderId)
	if err != nil {
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", txnDetails)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func InvoiceDetails(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	orderId := activity.OrderIdReq{}

	err := json.NewDecoder(r.Body).Decode(&orderId)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	txnDetails, err := activity.GetTxnDetails(orderId)
	if err != nil {
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	orderDetails, err := activity.DownloadInvoiceData(txnDetails)
	if err != nil {
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", orderDetails)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {

	ctx, log := logger.GetContextLogger(r.Context())

	token, _ := r.Cookie("token")
	_, _ = utils.UserTokenData(token)
	jsonData := simplejson.New()

	userReq := user.UserDetailReq{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	response, err := reports.SearchUsers(userReq.UserEmail)
	if err != nil {
		// TODO: log exception
		log.Error("Error: ", zap.Error(err), logger.LogUserId(ctx))
		jsonData.Set("error", resp.InternalServerError)
		resp.SendResponse(w, http.StatusInternalServerError, jsonData, &ctx)
		return
	}

	jsonData.Set("data", response)
	resp.SendResponse(w, http.StatusOK, jsonData, &ctx)
}

func uploadUsers(userReq user.CreateUserReq, data dataframe.DataFrame, adminEmail string) {
	failedRecords := user.UserCreationFailStruct{}
	errs := []error{}
	for i := range data.Records() {
		if i == data.Nrow() {
			break
		}
		userEmail := data.Col("User_Email(*)").Records()[i]
		userName := data.Col("User_Name(*)").Records()[i]
		userMob := ""

		colNames := data.Names()
		for j := range colNames {
			if colNames[j] == "User_Mobile" {
				userMob = data.Col("User_Mobile").Records()[i]
			}
		}

		if userEmail != "" && userName != "" {
			if user.ValidEmail(userEmail) {
				userReq.UserEmail = userEmail
				userReq.UserName = userName
				userReq.UserMobile = userMob

				err := userReq.Create()
				if err != nil {
					errs = append(errs, err)
					errorRecord := user.BulkUploadFailResp{userEmail, err.Error()}
					failedRecords.FailResp = append(failedRecords.FailResp, errorRecord)
				}
			} else {
				errs = append(errs, errors.New(user.ErrInvalidEmail))
				errorRecord := user.BulkUploadFailResp{userEmail, user.ErrInvalidEmail}
				failedRecords.FailResp = append(failedRecords.FailResp, errorRecord)
			}
		} else {
			errs = append(errs, errors.New(user.ErrFailed))
			errorRecord := user.BulkUploadFailResp{userEmail, user.ErrFailed}
			failedRecords.FailResp = append(failedRecords.FailResp, errorRecord)
		}
	}

	if len(failedRecords.FailResp) != 0 {

		htmlTemplate, _ := user.UserCreationErrorTemplate(failedRecords.FailResp)
		_ = email.SendEmail("Lanquill Bulk User Creation Failed", []string{adminEmail}, htmlTemplate)
	}
}
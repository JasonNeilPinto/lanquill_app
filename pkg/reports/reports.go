package reports

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/Lanquill/Forge/pkg/db"
	"github.com/Lanquill/Forge/pkg/utils"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func getUsersQuery(entityLevel int) string {
	return `SELECT User_ID, User_Name, User_Email, Last_Login, User_Type FROM sententia.user 
	WHERE Level_` + strconv.Itoa(entityLevel) + `=? ` + utils.GetLowerLevelQuery(entityLevel)
}

func getRootEntities() ([]EntityResp, error) {

	entityResp := []EntityResp{}

	stmt, err := db.MySqlDB.Prepare(SELECT_ROOT_ENTITIES)
	if err != nil {
		log.Println(err)
		return entityResp, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
		return entityResp, err
	}
	defer rows.Close()

	for rows.Next() {
		eachEntityResp := EntityResp{}
		if err := rows.Scan(
			&eachEntityResp.EntityId,
			&eachEntityResp.EntityName,
			&eachEntityResp.EntityTypeId,
			&eachEntityResp.EntityTypeName,
			&eachEntityResp.EntityLevel,
			&eachEntityResp.LevelVerbose,
			&eachEntityResp.EntityHierarchy,
			&eachEntityResp.Users); err != nil {
			log.Println(err)
			return nil, err
		}

		entityResp = append(entityResp, eachEntityResp)

	}

	return entityResp, nil

}

func getSubEntities(entityId int64, entityDivId int64, entityLevel int) ([]EntityResp, error) {

	entityResp := []EntityResp{}

	stmtStr := ""

	if _, ok := utils.HigherLevelTypes[entityLevel]; ok {
		stmtStr = SELECT_HIGHER_ENTITIES
	} else if _, ok := utils.MiddleLevelTypes[entityLevel]; ok {
		stmtStr = SELECT_LOWER_ENTITIES
	} else if _, ok := utils.LowerLevelTypes[entityLevel]; ok {
		stmtStr = SELECT_LOWER_DIV_ENTITIES
		entityId = entityDivId
	} else {
		return entityResp, errors.New(utils.ErrInvalidLevel)
	}

	stmt, err := db.MySqlDB.Prepare(stmtStr)
	if err != nil {
		log.Println(err)
		return entityResp, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(entityId)
	if err != nil {
		log.Println(err)
		return entityResp, err
	}
	defer rows.Close()

	for rows.Next() {
		eachEntityResp := EntityResp{}
		if err := rows.Scan(
			&eachEntityResp.EntityId,
			&eachEntityResp.EntityDivId,
			&eachEntityResp.EntityName,
			&eachEntityResp.EntityTypeId,
			&eachEntityResp.EntityTypeName,
			&eachEntityResp.EntityLevel,
			&eachEntityResp.LevelVerbose,
			&eachEntityResp.EntityHierarchy,
			&eachEntityResp.Users); err != nil {
			log.Println(err)
			return nil, err
		}

		entityResp = append(entityResp, eachEntityResp)

	}

	if _, ok := utils.HigherLevelTypes[entityLevel]; ok {
		stmt, err := db.MySqlDB.Prepare(SELECT_LOWER_ENTITIES)
		if err != nil {
			log.Println(err)
			return entityResp, err
		}

		defer stmt.Close()
		rows, err := stmt.Query(entityId)
		if err != nil {
			log.Println(err)
			return entityResp, err
		}
		defer rows.Close()

		for rows.Next() {
			eachLowerEntityResp := EntityResp{}
			if err := rows.Scan(
				&eachLowerEntityResp.EntityId,
				&eachLowerEntityResp.EntityDivId,
				&eachLowerEntityResp.EntityName,
				&eachLowerEntityResp.EntityTypeId,
				&eachLowerEntityResp.EntityTypeName,
				&eachLowerEntityResp.EntityLevel,
				&eachLowerEntityResp.LevelVerbose,
				&eachLowerEntityResp.EntityHierarchy,
				&eachLowerEntityResp.Users); err != nil {
				log.Println(err)
				return nil, err
			}

			entityResp = append(entityResp, eachLowerEntityResp)

		}
	}

	return entityResp, nil

}

func (entityReq *EntityUserReq) GetEntities() ([]EntityResp, error) {

	entityResp := []EntityResp{}

	if entityReq.EntityId == nil || *entityReq.EntityId == 0 {
		return getRootEntities()
	} else {
		if entityReq.EntityLevel != nil {
			return getSubEntities(*entityReq.EntityId, entityReq.EntityDivId, *entityReq.EntityLevel)
		}
		// Else empty object will be returned
	}

	return entityResp, nil
}

func getRootUsers() ([]UsersResp, error) {
	usersResp := []UsersResp{}

	stmt, err := db.MySqlDB.Prepare(SELECT_ROOT_USERS)
	if err != nil {
		log.Println(err)
		return usersResp, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
		return usersResp, err
	}
	defer rows.Close()

	for rows.Next() {
		usersDB := UsersDB{}
		if err := rows.Scan(
			&usersDB.Id,
			&usersDB.Name,
			&usersDB.Email,
			&usersDB.Login); err != nil {
			log.Println(err)
			return nil, err
		}

		usersResp = append(usersResp, UsersResp{
			usersDB.Id,
			usersDB.Name,
			usersDB.Email,
			usersDB.Login.String,
			true,
		})

	}

	return usersResp, nil
}

func getSubUsers(entityId int64, entityLevel int, entityDivId int64) ([]UsersResp, error) {

	usersResp := []UsersResp{}

	if _, ok := utils.LowerLevelTypes[entityLevel]; ok {
		entityId = entityDivId
	}

	stmt, err := db.MySqlDB.Prepare(getUsersQuery(entityLevel))
	if err != nil {
		log.Println(err)
		return usersResp, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(entityId)
	if err != nil {
		log.Println(err)
		return usersResp, err
	}
	defer rows.Close()

	for rows.Next() {
		usersDB := UsersDB{}
		if err := rows.Scan(
			&usersDB.Id,
			&usersDB.Name,
			&usersDB.Email,
			&usersDB.Login,
			&usersDB.UserType); err != nil {
			log.Println(err)
			return nil, err
		}

		isAdmin := false
		if _, ok := utils.UserTypeMap[usersDB.UserType]; ok {
			isAdmin = true
		} else {
			isAdmin = false
		}

		usersResp = append(usersResp, UsersResp{
			usersDB.Id,
			usersDB.Name,
			usersDB.Email,
			usersDB.Login.String,
			isAdmin,
		})
	}

	return usersResp, nil

}

func (userReq *EntityUserReq) GetUsers() ([]UsersResp, error) {

	usersResp := []UsersResp{}

	if userReq.EntityId == nil || *userReq.EntityId == 0 {
		return getRootUsers()
	} else {
		if userReq.EntityLevel != nil {
			// Return error if level verbose incorrect
			if _, ok := utils.LevelMap[*userReq.EntityLevel]; !ok {
				return usersResp, errors.New(utils.ErrInvalidLevel)
			}

			return getSubUsers(*userReq.EntityId, *userReq.EntityLevel, userReq.EntityDivId)
		}
		// Else empty object will be returned
	}

	return usersResp, nil
}

func (entityReq *EntityReportReq) GetEntityReport() (EntityReportResp, error) {

	entityReportResp := EntityReportResp{}

	if entityReq.EntityId != nil && entityReq.LevelVerbose != nil {
		// Return error if level verbose incorrect
		if _, ok := utils.EntityLevel[*entityReq.LevelVerbose]; !ok {
			return entityReportResp, errors.New(utils.ErrInvalidLevel)
		}

		// For lower level entities - consider entityDivId
		if _, ok := utils.LowerLevelTypes[*entityReq.EntityLevel]; ok {
			entityReq.EntityId = entityReq.EntityDivId
		}

		userCount, err := getEntityUserCount(*entityReq.EntityId, *entityReq.LevelVerbose)
		if err != nil {
			log.Println("ERROR: ", err)
			return entityReportResp, err
		}

		certCount, err := getEntityCertTakenCount(*entityReq.EntityId, *entityReq.LevelVerbose)
		if err != nil {
			log.Println("ERROR: ", err)
			return entityReportResp, err
		}

		generalReport, err := getGeneralReport("entity", *entityReq.EntityId, *entityReq.LevelVerbose)
		if err != nil {
			log.Println("ERROR: ", err)
			return entityReportResp, err
		}

		entityReportResp.TotalUsers = userCount.TotalUsers
		entityReportResp.ActiveUsers = userCount.ActiveUsers
		entityReportResp.CertTaken = certCount
		entityReportResp.GeneralReport = generalReport
	} else {
		userCount, err := getOverallUserCount()
		if err != nil {
			log.Println("ERROR: ", err)
			return entityReportResp, err
		}

		certCount, err := getOverallCertTakenCount()
		if err != nil {
			log.Println("ERROR: ", err)
			return entityReportResp, err
		}

		generalReport, err := getGeneralReport("entity", 1, "Level_8")
		if err != nil {
			log.Println("ERROR: ", err)
			return entityReportResp, err
		}

		entityReportResp.TotalUsers = userCount.TotalUsers
		entityReportResp.ActiveUsers = userCount.ActiveUsers
		entityReportResp.CertTaken = certCount
		entityReportResp.GeneralReport = generalReport
	}
	return entityReportResp, nil
}

func (userReq *UserReportReq) GetUserReport() (UserReportResp, error) {
	userReportResp := UserReportResp{}
	if userReq.UserId != nil {
		userBaseInfo, err := getUserBaseInfo(*userReq.UserId)
		if err != nil {
			log.Println("ERROR: ", err)
			return userReportResp, err
		}

		generalReport, err := getGeneralReport("user", *userReq.UserId, "Level_8")
		if err != nil {
			log.Println("ERROR: ", err)
			return userReportResp, err
		}

		generalReport.KylScore = nil
		userReportResp.LastLogin = userBaseInfo.LastLogin
		userReportResp.SelfAssessmentLevel = userBaseInfo.SelfAssessmentLevel
		userReportResp.Docs = userBaseInfo.Docs
		userReportResp.GeneralReport = generalReport
	}

	return userReportResp, nil
}

func (entityReq *EntityReportReq) ExportEntityReport() (string, error) {
	var filePath = ""
	if entityReq.EntityId != nil {
		// Return error if level verbose incorrect
		if _, ok := utils.EntityLevel[*entityReq.LevelVerbose]; !ok {
			return "", errors.New(utils.ErrInvalidLevel)
		}

		if *entityReq.EntityLevel <= 7 && *entityReq.EntityLevel >= 5 {
			filePath = "%s/excel_report/entity/higher/%d/report.xlsx"
		} else if *entityReq.EntityLevel <= 4 && *entityReq.EntityLevel >= 1 {
			filePath = "%s/excel_report/entity/lower/%d/report.xlsx"
		}

		filePath = fmt.Sprintf(filePath, utils.LQ_DB_PATH, *entityReq.EntityId)
	} else {
		return "", errors.New(utils.ErrInvalidRequest)
	}
	return filePath, nil
}

func (userReq *UserReportReq) ExportUserReport() (string, error) {
	var fileName = ""
	if userReq.UserId != nil {
		userProfileData, err := getUserProfileInfo(*userReq.UserId)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		generalReport, err := getGeneralReport("user", *userReq.UserId, "Level_8")
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		file, err := generateExcelData(userProfileData, *generalReport)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		filePath := "%s/excel_report/users/%d"
		filePath = fmt.Sprintf(filePath, utils.LQ_DB_PATH, *userReq.UserId)

		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		fileName = filePath + "/user_report.xlsx"
		if err = file.SaveAs(fileName); err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

	} else {
		return "", errors.New(utils.ErrInvalidRequest)
	}

	return fileName, nil
}

func getEntityUserCount(entityId int64, entityLevel string) (*UserCountResp, error) {
	entityUserCountResp := UserCountResp{}
	entityUserCountQuery := entityUserCount
	entityUserCountQuery = fmt.Sprintf(entityUserCountQuery, entityLevel)
	stmt, err := db.MySqlDB.Prepare(entityUserCountQuery)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(entityId).Scan(
		&entityUserCountResp.TotalUsers,
		&entityUserCountResp.ActiveUsers,
	)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}
	return &entityUserCountResp, nil
}

func getEntityCertTakenCount(entityId int64, entityLevel string) (int64, error) {
	coll := db.MongoClient.Database("eklaas").Collection("sententia_certification_exam_results")
	arr := []string{"failed", "pass"}
	match := bson.M{"overall_pass_status": bson.M{"$in": arr}, "is_sso_user": "NO", strings.ToLower(entityLevel): entityId}
	count, err := coll.CountDocuments(db.CTX, match)
	if err != nil {
		log.Println("ERROR: ", err)
		return 0, err
	}
	return count, nil
}

func getOverallUserCount() (*UserCountResp, error) {
	entityUserCountResp := UserCountResp{}
	stmt, err := db.MySqlDB.Prepare(overallUserCount)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow().Scan(
		&entityUserCountResp.TotalUsers,
		&entityUserCountResp.ActiveUsers,
	)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}
	return &entityUserCountResp, nil
}

func getOverallCertTakenCount() (int64, error) {
	coll := db.MongoClient.Database("eklaas").Collection("sententia_certification_exam_results")
	arr := []string{"failed", "pass"}
	match := bson.M{"overall_pass_status": bson.M{"$in": arr}, "is_sso_user": "NO"}
	count, err := coll.CountDocuments(db.CTX, match)
	if err != nil {
		log.Println("ERROR: ", err)
		return 0, err
	}
	return count, nil
}

func getUserBaseInfo(userId int64) (*UserReportResp, error) {
	userReportResp := UserReportResp{}
	stmt, err := db.MySqlDB.Prepare(userBaseInfo)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}
	defer stmt.Close()

	lastLogin := sql.NullString{}
	err = stmt.QueryRow(userId).Scan(
		&lastLogin,
		&userReportResp.SelfAssessmentLevel,
		&userReportResp.Docs,
	)
	if err != nil && err != sql.ErrNoRows {
		log.Println("ERROR: ", err)
		return nil, err
	} else if err == sql.ErrNoRows {
		userReportResp.LastLogin = ""
		userReportResp.SelfAssessmentLevel = 0
		userReportResp.Docs = 0
	}

	if lastLogin.Valid {
		userReportResp.LastLogin = lastLogin.String
	} else {
		userReportResp.LastLogin = ""
	}
	return &userReportResp, nil
}

func getGeneralReport(entityType string, id int64, entityLevel string) (*TestResp, error) {
	generalReportResp := GeneralReportResp{}

	if entityType == "entity" {
		stmt, err := db.MySqlDB.Prepare(getEntityReport)
		if err != nil {
			log.Println("ERROR: ", err)
			return nil, err
		}
		defer stmt.Close()

		err = stmt.QueryRow(id, entityLevel).Scan(
			&generalReportResp.LAvgScore,
			&generalReportResp.SAvgScore,
			&generalReportResp.RAvgScore,
			&generalReportResp.WAvgScore,
			&generalReportResp.GAvgScore,
			&generalReportResp.VAvgScore,
			&generalReportResp.LTotalCount,
			&generalReportResp.STotalCount,
			&generalReportResp.RTotalCount,
			&generalReportResp.WTotalCount,
			&generalReportResp.GTotalCount,
			&generalReportResp.VTotalCount,
			&generalReportResp.L1Attempted,
			&generalReportResp.L2Attempted,
			&generalReportResp.L3Attempted,
			&generalReportResp.L4Attempted,
			&generalReportResp.L5Attempted,
			&generalReportResp.L6Attempted,
			&generalReportResp.L7Attempted,
			&generalReportResp.L1Completed,
			&generalReportResp.L2Completed,
			&generalReportResp.L3Completed,
			&generalReportResp.L4Completed,
			&generalReportResp.L5Completed,
			&generalReportResp.L6Completed,
			&generalReportResp.L7Completed,
			&generalReportResp.SuggestedLevel1,
			&generalReportResp.SuggestedLevel2,
			&generalReportResp.SuggestedLevel3,
			&generalReportResp.SuggestedLevel4,
			&generalReportResp.SuggestedLevel5,
			&generalReportResp.SuggestedLevel6,
			&generalReportResp.SuggestedLevel7,
			&generalReportResp.SugLevel1Diff,
			&generalReportResp.SugLevel2Diff,
			&generalReportResp.SugLevel3Diff,
			&generalReportResp.SugLevel4Diff,
			&generalReportResp.SugLevel5Diff,
			&generalReportResp.SugLevel6Diff,
			&generalReportResp.SugLevel7Diff,
			&generalReportResp.Aoi,
		)
		if err != nil && err != sql.ErrNoRows {
			log.Println("ERROR: ", err)
			return nil, err
		}
	} else if entityType == "user" {
		stmt, err := db.MySqlDB.Prepare(getUserReport)
		if err != nil {
			log.Println("ERROR: ", err)
			return nil, err
		}
		defer stmt.Close()

		err = stmt.QueryRow(id).Scan(
			&generalReportResp.LAvgScore,
			&generalReportResp.SAvgScore,
			&generalReportResp.RAvgScore,
			&generalReportResp.WAvgScore,
			&generalReportResp.GAvgScore,
			&generalReportResp.VAvgScore,
			&generalReportResp.LTotalCount,
			&generalReportResp.STotalCount,
			&generalReportResp.RTotalCount,
			&generalReportResp.WTotalCount,
			&generalReportResp.GTotalCount,
			&generalReportResp.VTotalCount,
			&generalReportResp.L1Attempted,
			&generalReportResp.L2Attempted,
			&generalReportResp.L3Attempted,
			&generalReportResp.L4Attempted,
			&generalReportResp.L5Attempted,
			&generalReportResp.L6Attempted,
			&generalReportResp.L7Attempted,
			&generalReportResp.L1Completed,
			&generalReportResp.L2Completed,
			&generalReportResp.L3Completed,
			&generalReportResp.L4Completed,
			&generalReportResp.L5Completed,
			&generalReportResp.L6Completed,
			&generalReportResp.L7Completed,
			&generalReportResp.SuggestedLevel,
			&generalReportResp.Aoi,
		)
		if err != nil && err != sql.ErrNoRows {
			log.Println("ERROR: ", err)
			return nil, err
		}
	}

	if generalReportResp != (GeneralReportResp{}) {
		kylData := KylData{}
		if entityType == "entity" {

			// Level 1
			kylData.Level1.Count = generalReportResp.SuggestedLevel1.Int64
			kylData.Level1.Score = int64(math.Abs(float64(generalReportResp.SugLevel1Diff.Int64)))
			kylData.Level1.Type = getKylType(generalReportResp.SugLevel1Diff.Int64)

			// Level 2
			kylData.Level2.Count = generalReportResp.SuggestedLevel2.Int64
			kylData.Level2.Score = int64(math.Abs(float64(generalReportResp.SugLevel2Diff.Int64)))
			kylData.Level2.Type = getKylType(generalReportResp.SugLevel2Diff.Int64)

			// Level 3
			kylData.Level3.Count = generalReportResp.SuggestedLevel3.Int64
			kylData.Level3.Score = int64(math.Abs(float64(generalReportResp.SugLevel3Diff.Int64)))
			kylData.Level3.Type = getKylType(generalReportResp.SugLevel3Diff.Int64)

			// Level 4
			kylData.Level4.Count = generalReportResp.SuggestedLevel4.Int64
			kylData.Level4.Score = int64(math.Abs(float64(generalReportResp.SugLevel4Diff.Int64)))
			kylData.Level4.Type = getKylType(generalReportResp.SugLevel4Diff.Int64)

			// Level 5
			kylData.Level5.Count = generalReportResp.SuggestedLevel5.Int64
			kylData.Level5.Score = int64(math.Abs(float64(generalReportResp.SugLevel5Diff.Int64)))
			kylData.Level5.Type = getKylType(generalReportResp.SugLevel5Diff.Int64)

			// Level 6
			kylData.Level6.Count = generalReportResp.SuggestedLevel6.Int64
			kylData.Level6.Score = int64(math.Abs(float64(generalReportResp.SugLevel6Diff.Int64)))
			kylData.Level6.Type = getKylType(generalReportResp.SugLevel6Diff.Int64)

			// Level 7
			kylData.Level7.Count = generalReportResp.SuggestedLevel7.Int64
			kylData.Level7.Score = int64(math.Abs(float64(generalReportResp.SugLevel7Diff.Int64)))
			kylData.Level7.Type = getKylType(generalReportResp.SugLevel7Diff.Int64)
		}
		testResp := TestResp{}

		// Listening
		testResp.AverageScore.Listening.Average = generalReportResp.LAvgScore.Int64
		testResp.AverageScore.Listening.Tests = generalReportResp.LTotalCount.Int64

		// Speaking
		testResp.AverageScore.Speaking.Average = generalReportResp.SAvgScore.Int64
		testResp.AverageScore.Speaking.Tests = generalReportResp.STotalCount.Int64

		// Reading
		testResp.AverageScore.Reading.Average = generalReportResp.RAvgScore.Int64
		testResp.AverageScore.Reading.Tests = generalReportResp.RTotalCount.Int64

		// Writing
		testResp.AverageScore.Writing.Average = generalReportResp.WAvgScore.Int64
		testResp.AverageScore.Writing.Tests = generalReportResp.WTotalCount.Int64

		// Grammar
		testResp.AverageScore.Grammar.Average = generalReportResp.GAvgScore.Int64
		testResp.AverageScore.Grammar.Tests = generalReportResp.GTotalCount.Int64

		// Vocabulary
		testResp.AverageScore.Vocabulary.Average = generalReportResp.VAvgScore.Int64
		testResp.AverageScore.Vocabulary.Tests = generalReportResp.VTotalCount.Int64

		// Level 1
		testResp.CertResults.Level1.Attempted = generalReportResp.L1Attempted.Int64
		testResp.CertResults.Level1.Completed = generalReportResp.L1Completed.Int64

		// Level 2
		testResp.CertResults.Level2.Attempted = generalReportResp.L2Attempted.Int64
		testResp.CertResults.Level2.Completed = generalReportResp.L2Completed.Int64

		// Level 3
		testResp.CertResults.Level3.Attempted = generalReportResp.L3Attempted.Int64
		testResp.CertResults.Level3.Completed = generalReportResp.L3Completed.Int64

		// Level 4
		testResp.CertResults.Level4.Attempted = generalReportResp.L4Attempted.Int64
		testResp.CertResults.Level4.Completed = generalReportResp.L4Completed.Int64

		// Level 5
		testResp.CertResults.Level5.Attempted = generalReportResp.L5Attempted.Int64
		testResp.CertResults.Level5.Completed = generalReportResp.L5Completed.Int64

		// Level 6
		testResp.CertResults.Level6.Attempted = generalReportResp.L6Attempted.Int64
		testResp.CertResults.Level6.Completed = generalReportResp.L6Completed.Int64

		// Level 7
		testResp.CertResults.Level7.Attempted = generalReportResp.L7Attempted.Int64
		testResp.CertResults.Level7.Completed = generalReportResp.L7Completed.Int64

		if kylData != (KylData{}) {
			testResp.KylScore = &kylData
			if !generalReportResp.Aoi.Valid || generalReportResp.Aoi.String == "" {
				testResp.Aoi = []string{}
			} else {
				testResp.Aoi = strings.Split(generalReportResp.Aoi.String, ",")
			}

		} else {
			testResp.SuggestedLevel = generalReportResp.SuggestedLevel.Int64
			if !generalReportResp.Aoi.Valid || generalReportResp.Aoi.String == "" {
				testResp.Aoi = []string{}
			} else {
				testResp.Aoi = strings.Split(generalReportResp.Aoi.String, ",")
			}
		}
		return &testResp, nil
	} else {
		return &TestResp{}, nil
	}
}

func getKylType(value int64) string {
	if value == 0 {
		return "none"
	} else if value > 0 {
		return "inc"
	} else {
		return "dec"
	}
}

func getUserProfileInfo(userId int64) (UsersResp, error) {
	userResp := UsersResp{}
	stmt, err := db.MySqlDB.Prepare(userProfileInfo)
	if err != nil {
		log.Println("ERROR: ", err)
		return userResp, err
	}
	defer stmt.Close()

	lastLogin := sql.NullString{}
	err = stmt.QueryRow(userId).Scan(
		&userResp.Id,
		&userResp.Name,
		&userResp.Email,
		&lastLogin,
	)
	if err != nil && err != sql.ErrNoRows {
		log.Println("ERROR: ", err)
		return userResp, err
	} else if err == sql.ErrNoRows {
		userResp.Id = 0
		userResp.Name = ""
		userResp.Email = ""
		userResp.Login = ""
	}

	if lastLogin.Valid {
		userResp.Login = lastLogin.String
	} else {
		userResp.Login = ""
	}

	return userResp, nil
}

func generateExcelData(userProfileData UsersResp, generalReport TestResp) (*excelize.File, error) {
	file := excelize.NewFile()

	sheetName := "User Report"
	file.SetSheetName("Sheet1", sheetName)

	data := [][]interface{}{
		{"Email", "Name", "Last Login", "Average Score", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, "Certification Results", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, "Self-assessment Level", "Areas of Improvement"},
		{nil, nil, nil, "Listening", nil, "Speaking", nil, "Reading", nil, "Writing", nil, "Grammar", nil, "Vocabulary", nil, "Level 1", nil, "Level 2", nil, "Level 3", nil, "Level 4", nil, "Level 5", nil, "Level 6", nil, "Level 7", nil},
		{nil, nil, nil, "Tests Taken", "Average", "Tests Taken", "Average", "Tests Taken", "Average", "Tests Taken", "Average", "Tests Taken", "Average", "Tests Taken", "Average", "Attempted", "Completed", "Attempted", "Completed", "Attempted", "Completed", "Attempted", "Completed", "Attempted", "Completed", "Attempted", "Completed", "Attempted", "Completed"},
		{userProfileData.Email, userProfileData.Name, userProfileData.Login, generalReport.AverageScore.Listening.Tests, generalReport.AverageScore.Listening.Average, generalReport.AverageScore.Speaking.Tests, generalReport.AverageScore.Speaking.Average, generalReport.AverageScore.Reading.Tests, generalReport.AverageScore.Reading.Average, generalReport.AverageScore.Writing.Tests, generalReport.AverageScore.Writing.Average, generalReport.AverageScore.Grammar.Tests, generalReport.AverageScore.Grammar.Average, generalReport.AverageScore.Vocabulary.Tests, generalReport.AverageScore.Vocabulary.Average, generalReport.CertResults.Level1.Attempted, generalReport.CertResults.Level1.Completed, generalReport.CertResults.Level2.Attempted, generalReport.CertResults.Level2.Completed, generalReport.CertResults.Level3.Attempted, generalReport.CertResults.Level3.Completed, generalReport.CertResults.Level4.Attempted, generalReport.CertResults.Level4.Completed, generalReport.CertResults.Level5.Attempted, generalReport.CertResults.Level5.Completed, generalReport.CertResults.Level6.Attempted, generalReport.CertResults.Level6.Completed, generalReport.CertResults.Level7.Attempted, generalReport.CertResults.Level7.Completed, generalReport.SuggestedLevel, strings.Join(generalReport.Aoi, ", ")},
	}

	for i, row := range data {
		startCell, err := excelize.JoinCellName("A", i+1)
		if err != nil {
			log.Println("ERROR: ", err)
			return file, err
		}
		if err := file.SetSheetRow(sheetName, startCell, &row); err != nil {
			log.Println("ERROR: ", err)
			return file, err
		}
	}

	mergeCellRanges := [][]string{{"A1", "A3"}, {"B1", "B3"}, {"C1", "C3"}, {"D1", "O1"}, {"P1", "AC1"}, {"AD1", "AD3"}, {"AE1", "AE3"}, {"D2", "E2"}, {"F2", "G2"}, {"H2", "I2"}, {"J2", "K2"}, {"L2", "M2"}, {"N2", "O2"}, {"P2", "Q2"}, {"R2", "S2"}, {"T2", "U2"}, {"V2", "W2"}, {"X2", "Y2"}, {"Z2", "AA2"}, {"AB2", "AC2"}}
	for _, ranges := range mergeCellRanges {
		if err := file.MergeCell(sheetName, ranges[0], ranges[1]); err != nil {
			log.Println("ERROR: ", err)
			return file, err
		}
	}

	style, err := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Font:      &excelize.Font{Bold: true},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	for _, cell := range []string{"A1", "B1", "C1", "D1", "P1", "AD1", "AE1", "D2", "F2", "H2", "J2", "L2", "N2", "P2", "R2", "T2", "V2", "X2", "Z2", "AB2", "D3", "E3", "F3", "G3", "H3", "I3", "J3", "K3", "L3", "M3", "N3", "O3", "P3", "Q3", "R3", "S3", "T3", "U3", "V3", "W3", "X3", "Y3", "Z3", "AA3", "AB3", "AC3"} {
		if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
			log.Println("ERROR: ", err)
			return file, err
		}
	}

	if err = file.SetColWidth(sheetName, "A", "A", 25); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "B", "C", 21); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "D", "AC", 16); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "AD", "AE", 23); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetRowHeight(sheetName, 4, 38); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetRowHeight(sheetName, 1, 28); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetRowHeight(sheetName, 2, 24); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetRowHeight(sheetName, 3, 20); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	styleWrap, err := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Vertical: "center", WrapText: true},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}
	if err = file.SetCellStyle(sheetName, "AE4", "AE4", styleWrap); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	styleBorder, err := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetCellStyle(sheetName, "A4", "AD4", styleBorder); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	return file, nil
}

func (entityReq *EntityReportReq) ExportEntityUserReport() (string, error) {
	var fileName = ""
	if entityReq.EntityId != nil || entityReq.EntityDivId != nil {

		if _, ok := utils.LowerLevelTypes[*entityReq.EntityLevel]; ok {
			entityReq.EntityId = entityReq.EntityDivId
		}

		generalReport, err := getEntityUsers(*entityReq.EntityId, *entityReq.EntityLevel)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		entityId := int64(0)
		if entityReq.EntityDivId != nil {
			entityId = *entityReq.EntityDivId
		}

		file, err := generateEntityUserExcelData(generalReport)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		filePath := "%s/excel_report/users/%d"
		filePath = fmt.Sprintf(filePath, utils.LQ_DB_PATH, entityId)

		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		fileName = filePath + "/entity_user_report.xlsx"
		if err = file.SaveAs(fileName); err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

	} else {
		return "", errors.New(utils.ErrInvalidRequest)
	}

	return fileName, nil
}

func getEntityUsers(entityId int64, entityLevel int) ([]UsersEntityResp, error) {

	usersResp := []UsersEntityResp{}

	usersQuery := ""

	if _, ok := utils.LowerLevelTypes[entityLevel]; ok {
		usersQuery = strings.Replace(getUserLowerEntityList, "%s", fmt.Sprint(entityLevel), 2)
	} else {
		usersQuery = strings.Replace(getUserHigherEntityList, "%s", fmt.Sprint(entityLevel), 2)
	}

	stmt, err := db.MySqlDB.Prepare(usersQuery)
	if err != nil {
		log.Println(err)
		return usersResp, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(entityId)
	if err != nil {
		log.Println(err)
		return usersResp, err
	}
	defer rows.Close()

	for rows.Next() {
		usersDB := UsersEntityDB{}
		if err := rows.Scan(
			&usersDB.Id,
			&usersDB.Name,
			&usersDB.Email,
			&usersDB.Login,
			&usersDB.UserType,
			&usersDB.UserStatus,
			&usersDB.EntityName,
			&usersDB.DateCreated,
			&usersDB.RenewalDate,
			&usersDB.IsSSOUser); err != nil {
			log.Println(err)
			return nil, err
		}

		isAdmin := "No"
		if _, ok := utils.UserTypeMap[usersDB.UserType]; ok {
			isAdmin = "Yes"
		} else {
			isAdmin = "No"
		}

		usersResp = append(usersResp, UsersEntityResp{
			usersDB.Id,
			usersDB.Name,
			usersDB.Email,
			usersDB.Login.String,
			isAdmin,
			usersDB.UserStatus,
			usersDB.EntityName,
			usersDB.DateCreated,
			usersDB.RenewalDate,
			usersDB.IsSSOUser,
		})
	}

	return usersResp, nil

}

func generateEntityUserExcelData(userProfileData []UsersEntityResp) (*excelize.File, error) {
	file := excelize.NewFile()

	sheetName := "User Report"
	file.SetSheetName("Sheet1", sheetName)

	headers := map[string]string{
		"A1": "Name",
		"B1": "Email",
		"C1": "Last Login",
		"D1": "Entity Name",
		"E1": "Is Admin",
		"F1": "User Status",
		"G1": "SSO User",
		"H1": "Created Date",
		"I1": "Renewal Date",
	}

	for cell, value := range headers {
		file.SetCellValue(sheetName, cell, value)
	}

	for i, eachUser := range userProfileData {

		row := []interface{}{
			eachUser.Name,
			eachUser.Email,
			eachUser.Login,
			eachUser.EntityName,
			eachUser.IsAdmin,
			eachUser.UserStatus,
			eachUser.IsSSOUser,
			eachUser.DateCreated,
			eachUser.RenewalDate}

		startCell, err := excelize.JoinCellName("A", i+2)
		if err != nil {
			log.Println("ERROR : ", err)
			return file, err
		}
		if err := file.SetSheetRow(sheetName, startCell, &row); err != nil {
			log.Println("ERROR : ", err)
			return file, err
		}
	}

	style, err := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Vertical: "center", WrapText: true},
		Font:      &excelize.Font{Bold: true},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	for _, cell := range []string{"A1", "B1", "C1", "D1", "E1", "F1", "G1", "H1", "I1"} {
		if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
			log.Println("ERROR: ", err)
			return file, err
		}
	}

	if err = file.SetColWidth(sheetName, "A", "G", 35); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "C", "C", 25); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "E", "G", 15); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "H", "I", 25); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetRowHeight(sheetName, 1, 20); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	styleBorder, err := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "left", WrapText: true},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetCellStyle(sheetName, "A2", "I"+fmt.Sprint(len(userProfileData)+1), styleBorder); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	return file, nil
}

func SearchUsers(userEmail string) (SearchUserResp, error) {

	usersResp := SearchUserResp{}

	var login sql.NullString
	var entityName sql.NullString
	var lowerEntityName sql.NullString
	userType := 0

	err := db.MySqlDB.QueryRow(getSearchUser, userEmail).Scan(&usersResp.Id,
		&usersResp.Name,
		&usersResp.Email,
		&login,
		&userType,
		&usersResp.UserStatus,
		&entityName,
		&lowerEntityName,
		&usersResp.DateCreated,
		&usersResp.RenewalDate,
		&usersResp.IsSSOUser)
	if err != nil && err != sql.ErrNoRows {
		log.Println("ERROR: ", err)

	}

	isAdmin := false
	if _, ok := utils.UserTypeMap[userType]; ok {
		isAdmin = true
	} else {
		isAdmin = false
	}

	usersResp.IsAdmin = isAdmin

	if login.Valid {
		usersResp.Login = login.String
	} else {
		usersResp.Login = ""
	}

	if lowerEntityName.Valid {
		usersResp.EntityName = lowerEntityName.String
	} else if entityName.Valid {
		usersResp.EntityName = entityName.String
	} else {
		usersResp.EntityName = ""
	}

	return usersResp, nil
}

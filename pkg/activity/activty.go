package activity

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/Lanquill/Forge/pkg/db"
	"github.com/Lanquill/Forge/pkg/questions"
	"github.com/Lanquill/Forge/pkg/utils"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsersInfo(entityId int64, entityLevel int, entityDivId int64) ([]UsersResp, error) {

	userResp := []UsersResp{}
	var UserTypeId int
	var login sql.NullString
	var mobileNumber sql.NullString
	var level1 sql.NullInt64
	var level2 sql.NullInt64
	var level3 sql.NullInt64
	var level4 sql.NullInt64
	var level5 sql.NullInt64
	var level6 sql.NullInt64
	var level7 sql.NullInt64

	stmt, err := db.MySqlDB.Prepare(SELECT_USER)
	if err != nil {
		log.Println(err)
		return userResp, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(entityId, entityId, entityId, entityId, entityId, entityId)
	if err != nil {
		log.Println(err)
		return userResp, err
	}
	defer rows.Close()

	enitityDetailsResp := EntityDetailsTemp{}

	if entityDivId != 0 {
		stmt, err := db.MySqlDB.Prepare(userLowerEntityInfo)
		if err != nil {
			log.Println("ERROR: ", err)
			return nil, err
		}
		defer stmt.Close()

		err = stmt.QueryRow(entityDivId).Scan(
			&enitityDetailsResp.EntityName,
			&enitityDetailsResp.EntityId,
		)
		if err != nil && err != sql.ErrNoRows {
			log.Println("ERROR : ", err)
			return nil, err
		}
	} else {
		stmt, err := db.MySqlDB.Prepare(userHigherEntityInfo)
		if err != nil {
			log.Println("ERROR: ", err)
			return nil, err
		}
		defer stmt.Close()

		err = stmt.QueryRow(entityId).Scan(
			&enitityDetailsResp.EntityName,
			&enitityDetailsResp.EntityId,
		)
		if err != nil && err != sql.ErrNoRows {
			log.Println("ERROR : ", err)
			return nil, err
		}
	}

	if entityDivId != 0 {
		err = SetEntityInfoRedis(entityDivId, enitityDetailsResp)
		if err != nil {
			log.Println("ERROR: ", err)
			return nil, err
		}
	} else {
		err = SetEntityInfoRedis(entityId, enitityDetailsResp)
		if err != nil {
			log.Println("ERROR: ", err)
			return nil, err
		}
	}

	for rows.Next() {

		eachUserResp := UsersResp{}
		if err := rows.Scan(
			&eachUserResp.Id,
			&eachUserResp.Name,
			&eachUserResp.Email,
			&level7,
			&level6,
			&level5,
			&level4,
			&level3,
			&level2,
			&level1,
			&login,
			&mobileNumber,
			&eachUserResp.RenewalDate,
			&UserTypeId,
			&eachUserResp.UserType); err != nil {
			log.Println(err)
			return userResp, err
		}

		if login.Valid {
			eachUserResp.Login = login.String
		} else {
			eachUserResp.Login = ""
		}

		if mobileNumber.Valid {
			eachUserResp.MobileNumber = mobileNumber.String
		} else {
			eachUserResp.MobileNumber = ""
		}

		var entityId = int64(0)

		// TODO: For UserType: 1 --> This can be handled in a better way
		if UserTypeId == 1 {
			if level1.Valid {
				entityId = level1.Int64
			} else if level2.Valid {
				entityId = level2.Int64
			} else if level3.Valid {
				entityId = level3.Int64
			} else if level4.Valid {
				entityId = level4.Int64
			} else if level5.Valid {
				entityId = level5.Int64
			} else if level6.Valid {
				entityId = level6.Int64
			}
		} else {
			levelField := UserTypeLevel[int(UserTypeId)]
			switch levelField {
			case "Level7":
				entityId = level7.Int64
			case "Level6":
				entityId = level6.Int64
			case "Level5":
				entityId = level5.Int64
			case "Level4":
				entityId = level4.Int64
			case "Level3":
				entityId = level3.Int64
			case "Level2":
				entityId = level2.Int64
			case "Level1":
				entityId = level1.Int64
			}
		}

		entityDetails, _ := GetEntityInfoRedis(entityId)
		eachUserResp.EntityName = entityDetails.EntityName

		userResp = append(userResp, eachUserResp)
	}

	return userResp, nil
}

func generateExcelData(usersData []UsersResp) (*excelize.File, error) {
	file := excelize.NewFile()

	sheetName := "User Report"
	file.SetSheetName("Sheet1", sheetName)

	headings := map[string]string{
		"A1": "UserId",
		"B1": "Name",
		"C1": "Email",
		"D1": "Login Time",
		"E1": "Mobile Number",
		"F1": "Renewal Date",
		"G1": "User Type",
		"H1": "Entity Name",
	}

	for cell, value := range headings {
		file.SetCellValue(sheetName, cell, value)
	}

	for i, eachUser := range usersData {

		data := []interface{}{eachUser.Id,
			eachUser.Name,
			eachUser.Email,
			eachUser.Login,
			eachUser.MobileNumber,
			eachUser.RenewalDate,
			eachUser.UserType,
			eachUser.EntityName}

		startCell, err := excelize.JoinCellName("A", i+2)
		if err != nil {
			log.Println("ERROR : ", err)
			return file, err
		}
		if err := file.SetSheetRow(sheetName, startCell, &data); err != nil {
			log.Println("ERROR : ", err)
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

	for _, cell := range []string{"A1", "B1", "C1"} {
		if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
			log.Println("ERROR: ", err)
			return file, err
		}
	}

	if err = file.SetColWidth(sheetName, "A", "A", 10); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "B", "B", 50); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "C", "C", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "D", "D", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "E", "E", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "F", "F", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "G", "G", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "H", "H", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	return file, nil
}

func ExportEntityUsersReport(entityId *int64, userData []UsersResp) (string, error) {
	var fileName = ""
	if userData != nil {
		file, err := generateExcelData(userData)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		filePath := "%s/excel_report/Entity_Users"
		filePath = fmt.Sprintf(filePath, utils.LQ_DB_PATH)

		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		fileName = filePath + "/EnitityUsers_report_%d.xlsx"
		fileName = fmt.Sprintf(fileName, *entityId)
		if err = file.SaveAs(fileName); err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

	} else {
		return "", errors.New(questions.ErrFileCreationFailure)
	}

	return fileName, nil
}

func CertificationInfo() ([]CertificationUsersResp, CountCertification, error) {
	final := []CertificationUsersResp{}

	totalCertification, err := getTotalCertifications()
	if err != nil && err != sql.ErrNoRows {
		log.Println("ERROR: ", err)
		return final, totalCertification, err
	}

	latestCertification, err := getUsersCompletedCertification()
	if err != nil && err != sql.ErrNoRows {
		log.Println("ERROR: ", err)
		return final, totalCertification, err
	}

	final, err = getUserData(latestCertification)
	if err != nil && err != sql.ErrNoRows {
		log.Println("ERROR: ", err)
		return final, totalCertification, err
	}

	return final, totalCertification, nil

}

func getTotalCertifications() (CountCertification, error) {
	totalCertification := CountCertification{}
	coll := db.MongoClient.Database("eklaas").Collection("sententia_certification_exam_results")

	valuesTomatch := []string{"failed", "pass"}
	countMap := make(map[int]int64)

	for i := 1; i <= 7; i++ {
		filter := bson.M{"certification_level": i, "overall_pass_status": bson.M{"$in": valuesTomatch}}

		count, err := coll.CountDocuments(db.CTX, filter)
		if err != nil {
			log.Println("ERROR: ", err)
			return totalCertification, nil
		}
		countMap[i] = count

	}

	totalCertification.Level_1 = countMap[1]
	totalCertification.Level_2 = countMap[2]
	totalCertification.Level_3 = countMap[3]
	totalCertification.Level_4 = countMap[4]
	totalCertification.Level_5 = countMap[5]
	totalCertification.Level_6 = countMap[6]
	totalCertification.Level_7 = countMap[7]

	return totalCertification, nil

}

func getUsersCompletedCertification() ([]LatestCertfication, error) {

	coll := db.MongoClient.Database("eklaas").Collection("sententia_certification_exam_results")

	latestcertification := []LatestCertfication{}
	valuesTomatch := []string{"failed", "pass"}

	filter := bson.M{"overall_pass_status": bson.M{"$in": valuesTomatch}}

	findOptions := options.Find()
	findOptions.SetProjection(bson.M{
		"overall_pass_status": 1,
		"user_id":             1,
		"certification_level": 1,
		"exam_start_date":     1,
		"is_sso_user":         1,
	})
	findOptions.SetSort(bson.D{
		{
			Key:   "exam_start_date",
			Value: -1,
		}})
	findOptions.SetLimit(100)

	cursor, err := coll.Find(db.CTX, filter, findOptions)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(db.CTX, &latestcertification); err != nil {
		log.Println("ERROR: ", err)
		return latestcertification, err
	}

	return latestcertification, nil

}

func getUserData(latestCertification []LatestCertfication) ([]CertificationUsersResp, error) {
	certificationUsers := []CertificationUsersResp{}

	for _, eachUser := range latestCertification {

		certificationUser := CertificationUsersResp{}
		var level6, level5 *int64

		if eachUser.SSOUser == "NO" {
			err := db.MySqlDB.QueryRow(SELECT_USER_DETAILS, eachUser.UserId).Scan(&certificationUser.Name, &certificationUser.Email, &level6, &level5)
			if err != nil && err != sql.ErrNoRows {
				log.Println("ERROR: ", err)
				return certificationUsers, err
			}

			query := "SELECT Entity_Name from sententia.entity WHERE Entity_ID = ?"
			if level5 != nil {
				err := db.MySqlDB.QueryRow(query, level5).Scan(&certificationUser.EntityName)
				if err != nil && err != sql.ErrNoRows {
					log.Println("ERROR: ", err)
					return certificationUsers, err
				}
			} else {
				err := db.MySqlDB.QueryRow(query, level6).Scan(&certificationUser.EntityName)
				if err != nil && err != sql.ErrNoRows {
					log.Println("ERROR: ", err)
					return certificationUsers, err
				}
			}

		} else {
			err := db.MySqlDB.QueryRow(SELECT_SSOUSER_DETAILS, eachUser.UserId).Scan(&certificationUser.Name, &certificationUser.Email, &certificationUser.EntityName)
			if err != nil && err != sql.ErrNoRows {
				log.Println("ERROR: ", err)
				return certificationUsers, err
			}
		}

		certificationUser.ExamStartDate = eachUser.ExamStartDate.Format("02-Jan-2006")
		certificationUser.CertificationLevel = eachUser.CertificationLevel
		certificationUser.OverallPassStatus = eachUser.OverallPassStatus

		certificationUsers = append(certificationUsers, certificationUser)
	}
	return certificationUsers, nil
}

func MostActiveUsersAsPerDocs() ([]MostActiveUsersResp, error) {

	mostActiveUsers := []MostActiveUsersResp{}
	userDetails := []UserDetails{}

	coll := db.MongoClient.Database("stats").Collection("analysis_summary_report")

	project := bson.M{"User_ID": 1, "SSOUser": 1}
	order := bson.M{"last_modified": -1}

	cursor, err := coll.Find(db.CTX, bson.M{}, options.Find().SetProjection(project).SetSort(order))
	if err != nil {
		log.Println(err)
	}

	if err = cursor.All(db.CTX, &userDetails); err != nil {
		log.Println("ERROR: ", err)

	}

	inputSlice := userDetails

	uniqueElements := make(map[string]UserDetails)
	var result []UserDetails

	for _, item := range inputSlice {
		key := fmt.Sprintf("%d-%v", item.UserId, item.SSOUser)

		if _, ok := uniqueElements[key]; !ok {
			uniqueElements[key] = item
			if len(result) != 100 {
				result = append(result, item)
			} else {
				break
			}
		}
	}

	for _, user := range result {

		mostActiveUser := MostActiveUsersResp{}

		count, err := coll.CountDocuments(db.CTX, bson.M{"User_ID": user.UserId, "SSOUser": user.SSOUser})
		if err != nil {
			log.Println(err)
		}
		mostActiveUser.DocCount = count
		var level6, level5 *int64

		if !user.SSOUser {
			err := db.MySqlDB.QueryRow(SELECT_USER_DETAILS, user.UserId).Scan(&mostActiveUser.UserName, &mostActiveUser.UserEmail)
			if err != nil && err != sql.ErrNoRows {
				log.Println("ERROR: ", err)

			}

			query := "SELECT Entity_Name from sententia.entity WHERE Entity_ID = ?"
			if level5 != nil {
				err := db.MySqlDB.QueryRow(query, level5).Scan(&mostActiveUser.EntityName)
				if err != nil && err != sql.ErrNoRows {
					log.Println("ERROR: ", err)
					return mostActiveUsers, err
				}
			} else {
				err := db.MySqlDB.QueryRow(query, level6).Scan(&mostActiveUser.EntityName)
				if err != nil && err != sql.ErrNoRows {
					log.Println("ERROR: ", err)
					return mostActiveUsers, err
				}
			}

		} else {
			err := db.MySqlDB.QueryRow(SELECT_SSOUSER_DETAILS, user.UserId).Scan(&mostActiveUser.UserName, &mostActiveUser.UserEmail, &mostActiveUser.EntityName)
			if err != nil && err != sql.ErrNoRows {
				log.Println("ERROR: ", err)

			}
		}
		mostActiveUsers = append(mostActiveUsers, mostActiveUser)

	}

	return mostActiveUsers, nil
}

func (entityReq EntityToIp) MapEntityToIps() error {

	seperator := func() string {
		if entityReq.IsRange {
			return "-"
		}
		return ","
	}()

	query := "SELECT IP_Address FROM Entity_Ip_Address WHERE IP_Address IN (" + placeholders(len(entityReq.Ip)) + ")"

	var ipdata []interface{}
	for _, ip := range entityReq.Ip {
		ipdata = append(ipdata, ip)
	}

	var ip string

	err := db.MySqlDB.QueryRow(query, ipdata...).Scan(&ip)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	if ip != "" {

		return errors.New(questions.ErrIpAddressExists)
	}

	for _, ip := range entityReq.Ip {

		query = "INSERT INTO Entity_Ip_Address (Entity_ID, IP_Address, IP_Range_Enable) VALUES (?, ?, ?)"
		_, err = db.MySqlDB.Exec(query, entityReq.EntityId, ip, entityReq.IsRange)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
	}

	stmt, err := db.MySqlDB.Prepare(updateIpAddressInEntity)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(
		seperator,
		entityReq.IPEnabled,
		entityReq.EntityId,
	)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	_, err = row.RowsAffected()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	return nil

}

func GetEntityIPAddresses() ([]GetIPAddress, error) {
	var entityIPAddress []GetIPAddress

	//Prepare the query to get the entity IP addresses
	stmt, err := db.MySqlDB.Prepare(getIPAddress)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}

	for rows.Next() {
		eachIpAddress := GetIPAddress{}

		if err := rows.Scan(
			&eachIpAddress.EipID,
			&eachIpAddress.EntityID,
			&eachIpAddress.EntityName,
			&eachIpAddress.IPAddress,
			&eachIpAddress.EntityIPAddress,
			&eachIpAddress.EnableIPLogin,
			&eachIpAddress.IsRange,
		); err != nil {
			log.Println(err)
			return nil, err
		}
		entityIPAddress = append(entityIPAddress, eachIpAddress)
	}

	return entityIPAddress, nil

}

func (entityReq EntityIpReq) UpdateEntityIpAddress() error {

	seperator := func() string {
		if entityReq.IsRange {
			return "-"
		}
		return ","
	}()

	stmt, err := db.MySqlDB.Prepare(updateEntityIpAddress)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	result, err := stmt.Exec(
		entityReq.IPAddress,
		entityReq.IsRange,
		entityReq.EipID,
	)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	if rows == 0 {
		err := errors.New(questions.ErrUpdateFailure)
		return err
	}

	stmt, err = db.MySqlDB.Prepare(updateIpAddressInEntity)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(
		seperator,
		entityReq.IPEnabled,
		entityReq.EntityID,
	)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	_, err = row.RowsAffected()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	return nil

}

func (entityReq EntityIpReq) DeleteEntityIpAddress() error {

	seperator := func() string {
		if entityReq.IsRange {
			return "-"
		}
		return ","
	}()

	stmt, err := db.MySqlDB.Prepare(deleteIpAddress)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	result, err := stmt.Exec(
		entityReq.EipID,
	)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	if rows == 0 {
		err := errors.New(questions.ErrDeleteFailure)
		return err
	}

	stmt, err = db.MySqlDB.Prepare(updateIpAddressInEntity)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(
		seperator,
		entityReq.IPEnabled,
		entityReq.EntityID,
	)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	_, err = row.RowsAffected()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	return nil

}

func placeholders(n int) string {
	return "?" + strings.Repeat(", ?", n-1)
}

func getLoggedInUsers() ([]LoggedInUsers, error) {

	loggedUsers := []LoggedInUsers{}

	stmt, err := db.MySqlDB.Prepare(SELECT_LOGGED_IN_USERS)
	if err != nil {
		log.Println("ERROR: ", err)
		return loggedUsers, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("ERROR: ", err)
		return loggedUsers, err
	}
	defer rows.Close()

	var level5, level6 *int64
	for rows.Next() {
		loggedUser := LoggedInUsers{}
		var lastLogin time.Time
		if err := rows.Scan(
			&loggedUser.UserName,
			&loggedUser.UserEmail,
			&lastLogin,
			&level6,
			&level5); err != nil {
			log.Println(err)
			return loggedUsers, err
		}

		query := "SELECT Entity_Name from sententia.entity WHERE Entity_ID = ?"
		if level5 != nil {
			err := db.MySqlDB.QueryRow(query, level5).Scan(&loggedUser.EntityName)
			if err != nil && err != sql.ErrNoRows {
				log.Println("ERROR: ", err)
				return loggedUsers, err
			}
		} else {
			err := db.MySqlDB.QueryRow(query, level6).Scan(&loggedUser.EntityName)
			if err != nil && err != sql.ErrNoRows {
				log.Println("ERROR: ", err)
				return loggedUsers, err
			}
		}

		loggedUser.LastLogin = lastLogin.Format("02-Jan-2006")
		loggedUsers = append(loggedUsers, loggedUser)
	}

	return loggedUsers, nil

}

func getDocumentAnalysisCount() (*int64, error) {

	coll := db.MongoClient.Database("sententia").Collection("Analysis_Report")
	doc_count, err := coll.CountDocuments(db.CTX, bson.M{})
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}

	return &doc_count, nil
}

func getRegAndRetailUsers() (*int64, *int64, error) {
	var regUsers, retailUsers *int64

	user_registered_query := "SELECT COUNT(*) FROM sententia.user;"
	rows, err := db.MySqlDB.Query(user_registered_query)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&regUsers); err != nil {
			log.Println("ERROR: ", err)
			return nil, nil, err
		}
	}

	retail_users_query := "SELECT COUNT(*) FROM sententia.user WHERE User_Type = 1 AND Level_7 is NULL AND Level_6 is NULL AND Level_5 is NULL;"
	rows, err = db.MySqlDB.Query(retail_users_query)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&retailUsers); err != nil {
			log.Println("ERROR: ", err)
			return nil, nil, err
		}
	}

	return regUsers, retailUsers, nil
}

func GetUserMetrics() (UserActivityMetrics, []LoggedInUsers, error) {

	userMetrics := UserActivityMetrics{}

	loggedInUsers, err := getLoggedInUsers()
	if err != nil {
		log.Println("ERROR: ", err)
		return userMetrics, loggedInUsers, err
	}

	reg_users, retail_users, err := getRegAndRetailUsers()
	if err != nil {
		log.Println("ERROR: ", err)
		return userMetrics, loggedInUsers, err
	}

	docAnalysisCount, err := getDocumentAnalysisCount()
	if err != nil {
		log.Println("ERROR: ", err)
		return userMetrics, loggedInUsers, err
	}

	userMetrics.RegUsers = *reg_users
	userMetrics.RetailUsers = *retail_users
	userMetrics.DocsAnalyzed = *docAnalysisCount

	return userMetrics, loggedInUsers, err

}

func GetPaymentsInfo() ([]PaymentsResp, []PaymentsResp, []PaymentsResp, []PaymentsResp, error) {
	successfull := []PaymentsResp{}
	failed := []PaymentsResp{}
	pending := []PaymentsResp{}
	other := []PaymentsResp{}

	stmt, err := db.MySqlDB.Prepare(PAYMENT_DETAILS)
	if err != nil {
		log.Println("ERROR: ", err)
		return successfull, failed, pending, other, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("ERROR: ", err)
		return successfull, failed, pending, other, err
	}
	defer rows.Close()

	for rows.Next() {
		eachPayment := PaymentsResp{}
		var status sql.NullString
		if err := rows.Scan(
			&eachPayment.Email,
			&eachPayment.OrderID,
			&eachPayment.TransactionStartDate,
			&eachPayment.PaymentType,
			&eachPayment.Amount,
			&status); err != nil {
			log.Println("ERROR: ", err)
			return successfull, failed, pending, other, err
		}

		eachPayment.PaymentStatus = status.String

		if eachPayment.PaymentStatus == "PAYMENT_SUCCESS" || eachPayment.PaymentStatus == "SUCCESS" {
			successfull = append(successfull, eachPayment)

		} else if eachPayment.PaymentStatus == "PAYMENT_ERROR" || eachPayment.PaymentStatus == "FAILED" {
			failed = append(failed, eachPayment)

		} else if eachPayment.PaymentStatus == "PENDING" || eachPayment.PaymentStatus == "PAYMENT_PENDING" {
			pending = append(pending, eachPayment)

		} else {
			other = append(other, eachPayment)
		}
	}

	return successfull, failed, pending, other, nil
}

func generateCertificationExcelData(certificationInfo []CertificationUsersResp) (*excelize.File, error) {
	file := excelize.NewFile()

	sheetName := "Certification Report"
	file.SetSheetName("Sheet1", sheetName)

	headings := map[string]string{
		"A1": "Name",
		"B1": "Email",
		"C1": "Certification Level",
		"D1": "Exam Start Date",
		"E1": "Overall Status",
		"F1": "Entity Name",
	}

	for cell, value := range headings {
		file.SetCellValue(sheetName, cell, value)
	}

	for i, eachUser := range certificationInfo {

		data := []interface{}{
			eachUser.Name,
			eachUser.Email,
			eachUser.CertificationLevel,
			eachUser.ExamStartDate,
			eachUser.OverallPassStatus,
			eachUser.EntityName,
		}

		startCell, err := excelize.JoinCellName("A", i+2)
		if err != nil {
			log.Println("ERROR : ", err)
			return file, err
		}
		if err := file.SetSheetRow(sheetName, startCell, &data); err != nil {
			log.Println("ERROR : ", err)
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

	for _, cell := range []string{"A1", "B1", "C1", "D1", "E1", "F1"} {
		if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
			log.Println("ERROR: ", err)
			return file, err
		}
	}

	if err = file.SetColWidth(sheetName, "A", "A", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "B", "B", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "C", "C", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "D", "D", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "E", "E", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "F", "F", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	return file, nil
}

func ExportCertificationReport(certificationInfo []CertificationUsersResp) (string, error) {
	var fileName = ""
	time := time.Now().Format("02-01-2002")

	if certificationInfo != nil {
		file, err := generateCertificationExcelData(certificationInfo)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		filePath := "%s/excel_report/CertificationReport"
		filePath = fmt.Sprintf(filePath, utils.LQ_DB_PATH)

		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		fileName = filePath + "/certification_report_%s.xlsx"
		fileName = fmt.Sprintf(fileName, time)
		if err = file.SaveAs(fileName); err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

	} else {
		return "", errors.New(questions.ErrFileCreationFailure)

	}

	return fileName, nil
}

func RetailUsersInfo() ([]RetailUsers, error) {
	retailUsers := []RetailUsers{}

	stmt, err := db.MySqlDB.Prepare(RETAIL_USERS)
	if err != nil {
		log.Println("ERROR: ", err)
		return retailUsers, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("ERROR: ", err)
		return retailUsers, err
	}
	defer rows.Close()

	lastLogin := sql.NullString{}
	for rows.Next() {
		retailUser := RetailUsers{}

		if err := rows.Scan(
			&retailUser.UserName,
			&retailUser.UserEmail,
			&retailUser.DateCreated,
			&lastLogin); err != nil {

			log.Println("ERROR: ", err)
			return retailUsers, err
		}

		if err != nil && err != sql.ErrNoRows {

			log.Println("ERROR: ", err)
			return retailUsers, err

		} else if err == sql.ErrNoRows {
			retailUser.UserName = ""
			retailUser.UserEmail = ""
			retailUser.DateCreated = time.Time{}
			retailUser.LastLogin = ""
		}

		if lastLogin.Valid {
			retailUser.LastLogin = lastLogin.String
		} else {
			retailUser.LastLogin = ""
		}

		amount_query := "SELECT amount FROM sententia.payment_details WHERE email = ? ORDER BY txn_end_datetime DESC"
		err := db.MySqlDB.QueryRow(amount_query, retailUser.UserEmail).Scan(&retailUser.Amount)
		if err == sql.ErrNoRows {
			retailUser.Amount = 0.0
		} else if err != nil {
			log.Println("ERROR :", err)
			return retailUsers, err
		}

		retailUsers = append(retailUsers, retailUser)
	}
	return retailUsers, nil
}

func LoggedInUsersReport(loggedInUsers []LoggedInUsers) (string, error) {
	var fileName = ""
	time := time.Now().Format("02-01-2002")

	if loggedInUsers != nil {
		file, err := generateLoginInfoExcelData(loggedInUsers)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		filePath := "%s/excel_report/Login_User_Report"
		filePath = fmt.Sprintf(filePath, utils.LQ_DB_PATH)

		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		fileName = filePath + "/Login_Users_report_%s.xlsx"
		fileName = fmt.Sprintf(fileName, time)
		if err = file.SaveAs(fileName); err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

	} else {
		return "", errors.New(questions.ErrFileCreationFailure)
	}

	return fileName, nil
}

func generateLoginInfoExcelData(loggedInUsers []LoggedInUsers) (*excelize.File, error) {
	file := excelize.NewFile()

	sheetName := "Login Report"
	file.SetSheetName("Sheet1", sheetName)

	headings := map[string]string{
		"A1": "Name",
		"B1": "Email",
		"C1": "Last Login",
		"D1": "Entity Name",
	}

	for cell, value := range headings {
		file.SetCellValue(sheetName, cell, value)
	}

	for i, eachUser := range loggedInUsers {

		data := []interface{}{
			eachUser.UserName,
			eachUser.UserEmail,
			eachUser.LastLogin,
			eachUser.EntityName,
		}

		startCell, err := excelize.JoinCellName("A", i+2)
		if err != nil {
			log.Println("ERROR : ", err)
			return file, err
		}
		if err := file.SetSheetRow(sheetName, startCell, &data); err != nil {
			log.Println("ERROR : ", err)
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

	for _, cell := range []string{"A1", "B1", "C1", "D1"} {
		if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
			log.Println("ERROR: ", err)
			return file, err
		}
	}

	if err = file.SetColWidth(sheetName, "A", "A", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "B", "B", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "C", "C", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "D", "D", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	return file, nil
}

func RetailUsersReport(retailUsers []RetailUsers) (string, error) {
	var fileName = ""
	time := time.Now().Format("02-01-2002")

	if retailUsers != nil {
		file, err := generateRetailUsersExcelData(retailUsers)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		filePath := "%s/excel_report/Retail_Users_Report"
		filePath = fmt.Sprintf(filePath, utils.LQ_DB_PATH)

		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		fileName = filePath + "/Retail_Users_report_%s.xlsx"
		fileName = fmt.Sprintf(fileName, time)
		if err = file.SaveAs(fileName); err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

	} else {
		return "", errors.New(questions.ErrFileCreationFailure)

	}

	return fileName, nil
}

func generateRetailUsersExcelData(retailUsers []RetailUsers) (*excelize.File, error) {
	file := excelize.NewFile()

	sheetName := "Retail Users Report"
	file.SetSheetName("Sheet1", sheetName)

	headings := map[string]string{
		"A1": "Name",
		"B1": "Email",
		"C1": "Registration Date",
		"D1": "Last Login",
		"E1": "Amount",
	}

	for cell, value := range headings {
		file.SetCellValue(sheetName, cell, value)
	}

	for i, eachUser := range retailUsers {

		data := []interface{}{
			eachUser.UserName,
			eachUser.UserEmail,
			eachUser.DateCreated,
			eachUser.LastLogin,
			eachUser.Amount,
		}

		startCell, err := excelize.JoinCellName("A", i+2)
		if err != nil {
			log.Println("ERROR : ", err)
			return file, err
		}
		if err := file.SetSheetRow(sheetName, startCell, &data); err != nil {
			log.Println("ERROR : ", err)
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

	for _, cell := range []string{"A1", "B1", "C1", "D1"} {
		if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
			log.Println("ERROR: ", err)
			return file, err
		}
	}

	if err = file.SetColWidth(sheetName, "A", "A", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "B", "B", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "C", "C", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "D", "D", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "E", "E", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	return file, nil
}

func ExportPaymentInfo(identifier string) (string, error) {

	succesfull, failed, pending, _, err := GetPaymentsInfo()
	if err != nil {
		// TODO: log exception
		log.Println(err)
		return "", err
	}

	switch identifier {
	case "success":
		fileName, err := SuccessfullPaymentReport(succesfull)
		if err != nil {
			// TODO: log exception
			log.Println(err)
			return "", err
		}
		return fileName, nil

	case "failure":
		fileName, err := FailedPaymentReport(failed)
		if err != nil {
			// TODO: log exception
			log.Println(err)
			return "", err
		}

		return fileName, nil

	case "pending":
		fileName, err := PendingPaymentReport(pending)
		if err != nil {
			// TODO: log exception
			log.Println(err)
			return "", err
		}

		return fileName, nil
	}
	return "", nil
}

func SuccessfullPaymentReport(successfull []PaymentsResp) (string, error) {
	var fileName = ""
	time := time.Now().Format("02-01-2002")

	if successfull != nil {
		file, err := generateSuccessPaymentExcelData(successfull)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		filePath := "%s/excel_report/Payments_Report/Successful_Payment_info"
		filePath = fmt.Sprintf(filePath, utils.LQ_DB_PATH)

		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		fileName = filePath + "/Succesfull_Payment_report_%s.xlsx"
		fileName = fmt.Sprintf(fileName, time)
		if err = file.SaveAs(fileName); err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

	} else {
		return "", errors.New(questions.ErrFileCreationFailure)

	}

	return fileName, nil
}

func generateSuccessPaymentExcelData(succesfull []PaymentsResp) (*excelize.File, error) {
	file := excelize.NewFile()

	sheetName := "Successfull Payment Report"
	file.SetSheetName("Sheet1", sheetName)

	headings := map[string]string{
		"A1": "Email",
		"B1": "Order ID",
		"C1": "Transaction Date",
		"D1": "Payment Type",
		"E1": "Amount",
		"F1": "Payment Status",
	}

	for cell, value := range headings {
		file.SetCellValue(sheetName, cell, value)
	}

	for i, eachUser := range succesfull {

		data := []interface{}{
			eachUser.Email,
			eachUser.OrderID,
			eachUser.TransactionStartDate,
			eachUser.PaymentType,
			eachUser.Amount,
			eachUser.PaymentStatus,
		}

		startCell, err := excelize.JoinCellName("A", i+2)
		if err != nil {
			log.Println("ERROR : ", err)
			return file, err
		}
		if err := file.SetSheetRow(sheetName, startCell, &data); err != nil {
			log.Println("ERROR : ", err)
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

	for _, cell := range []string{"A1", "B1", "C1", "D1"} {
		if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
			log.Println("ERROR: ", err)
			return file, err
		}
	}

	if err = file.SetColWidth(sheetName, "A", "A", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "B", "B", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "C", "C", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "D", "D", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "E", "E", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "F", "F", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	return file, nil
}

func FailedPaymentReport(failed []PaymentsResp) (string, error) {
	var fileName = ""
	time := time.Now().Format("02-01-2002")

	if failed != nil {
		file, err := generateFailedPaymentExcelData(failed)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		filePath := "%s/excel_report/Payments_Report/Failed_Payment_info"
		filePath = fmt.Sprintf(filePath, utils.LQ_DB_PATH)

		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		fileName = filePath + "/Failed_Payment_report_%s.xlsx"
		fileName = fmt.Sprintf(fileName, time)
		if err = file.SaveAs(fileName); err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

	} else {
		return "", errors.New(questions.ErrFileCreationFailure)
	}

	return fileName, nil
}

func generateFailedPaymentExcelData(failed []PaymentsResp) (*excelize.File, error) {
	file := excelize.NewFile()

	sheetName := "Failure Payment Report"
	file.SetSheetName("Sheet1", sheetName)

	headings := map[string]string{
		"A1": "Email",
		"B1": "Order ID",
		"C1": "Transaction Date",
		"D1": "Payment Type",
		"E1": "Amount",
		"F1": "Payment Status",
	}

	for cell, value := range headings {
		file.SetCellValue(sheetName, cell, value)
	}

	for i, eachUser := range failed {

		data := []interface{}{
			eachUser.Email,
			eachUser.OrderID,
			eachUser.TransactionStartDate,
			eachUser.PaymentType,
			eachUser.Amount,
			eachUser.PaymentStatus,
		}

		startCell, err := excelize.JoinCellName("A", i+2)
		if err != nil {
			log.Println("ERROR : ", err)
			return file, err
		}
		if err := file.SetSheetRow(sheetName, startCell, &data); err != nil {
			log.Println("ERROR : ", err)
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

	for _, cell := range []string{"A1", "B1", "C1", "D1"} {
		if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
			log.Println("ERROR: ", err)
			return file, err
		}
	}

	if err = file.SetColWidth(sheetName, "A", "A", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "B", "B", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "C", "C", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "D", "D", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "E", "E", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "F", "F", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	return file, nil
}

func PendingPaymentReport(pending []PaymentsResp) (string, error) {
	var fileName = ""
	time := time.Now().Format("02-01-2002")

	if pending != nil {
		file, err := generatePendingPaymentExcelData(pending)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		filePath := "%s/excel_report/Payments_Report/Pending_Payment_info"
		filePath = fmt.Sprintf(filePath, utils.LQ_DB_PATH)

		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

		fileName = filePath + "/Pending_Payment_report_%s.xlsx"
		fileName = fmt.Sprintf(fileName, time)
		if err = file.SaveAs(fileName); err != nil {
			log.Println("ERROR: ", err)
			return "", err
		}

	} else {
		return "", errors.New(questions.ErrFileCreationFailure)
	}

	return fileName, nil
}

func generatePendingPaymentExcelData(pending []PaymentsResp) (*excelize.File, error) {
	file := excelize.NewFile()

	sheetName := "Pending Payment Report"
	file.SetSheetName("Sheet1", sheetName)

	headings := map[string]string{
		"A1": "Email",
		"B1": "Order ID",
		"C1": "Transaction Date",
		"D1": "Payment Type",
		"E1": "Amount",
		"F1": "Payment Status",
	}

	for cell, value := range headings {
		file.SetCellValue(sheetName, cell, value)
	}

	for i, eachUser := range pending {

		data := []interface{}{
			eachUser.Email,
			eachUser.OrderID,
			eachUser.TransactionStartDate,
			eachUser.PaymentType,
			eachUser.Amount,
			eachUser.PaymentStatus,
		}

		startCell, err := excelize.JoinCellName("A", i+2)
		if err != nil {
			log.Println("ERROR : ", err)
			return file, err
		}
		if err := file.SetSheetRow(sheetName, startCell, &data); err != nil {
			log.Println("ERROR : ", err)
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

	for _, cell := range []string{"A1", "B1", "C1", "D1"} {
		if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
			log.Println("ERROR: ", err)
			return file, err
		}
	}

	if err = file.SetColWidth(sheetName, "A", "A", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "B", "B", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "C", "C", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "D", "D", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "E", "E", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	if err = file.SetColWidth(sheetName, "F", "F", 30); err != nil {
		log.Println("ERROR: ", err)
		return file, err
	}

	return file, nil
}

func GetTxnDetails(orderId OrderIdReq) (TxnDetails, error) {
	txnDetails := TxnDetails{}

	var callBackResp sql.NullString
	var statusResp sql.NullString
	var invoiceNumber sql.NullString
	var countryCode sql.NullString
	var status sql.NullString
	var txnEndDate sql.NullTime

	if orderId.OrderId != "" {

		if err := db.MySqlDB.QueryRow(TXN_DETAILS, orderId.OrderId).Scan(&txnDetails.Email,
			&txnDetails.OrderID,
			&txnDetails.TransactionStartDate,
			&txnEndDate,
			&txnDetails.PaymentType,
			&txnDetails.Amount,
			&status,
			&txnDetails.Description,
			&callBackResp,
			&statusResp,
			&txnDetails.OrderDetails,
			&txnDetails.GST,
			&invoiceNumber,
			&countryCode); err != nil {
			log.Println("ERROR: ", err)
			return txnDetails, err
		}
	}

	if callBackResp.Valid {
		txnDetails.CallBackResp = callBackResp.String
	} else {
		txnDetails.CallBackResp = ""
	}

	if statusResp.Valid {
		txnDetails.StatusResp = statusResp.String
	} else {
		txnDetails.StatusResp = ""
	}

	if invoiceNumber.Valid {
		txnDetails.InvoiceNumber = invoiceNumber.String
	} else {
		txnDetails.InvoiceNumber = ""
	}

	if txnEndDate.Valid {
		txnDetails.TransactionEndDate = txnEndDate.Time
	} else {
		txnDetails.TransactionEndDate = time.Time{}
	}

	if status.Valid {
		txnDetails.PaymentStatus = status.String
	} else {
		txnDetails.PaymentStatus = ""
	}

	if countryCode.Valid {
		txnDetails.CountryCode = countryCode.String
	} else {
		txnDetails.CountryCode = ""
	}

	return txnDetails, nil

}

func DownloadInvoiceData(txnDetails TxnDetails) (OrderDetails, error) {
	orderDetails := OrderDetails{}
	statusRespMapper := StatusResp{}

	// Get the BankName from StatusResponse
	if txnDetails.StatusResp != "" {
		err := json.Unmarshal([]byte(txnDetails.StatusResp), &statusRespMapper)
		if err != nil {
			log.Println("ERROR:", err)
			return orderDetails, err
		}
	}

	// Get the BankName from StatusResponse
	if txnDetails.StatusResp != "" {
		err := json.Unmarshal([]byte(txnDetails.StatusResp), &statusRespMapper)
		if err != nil {
			log.Println("ERROR:", err)
			return orderDetails, err
		}
	}

	// Get the Order Details
	err := json.Unmarshal([]byte(txnDetails.OrderDetails), &orderDetails)
	if err != nil {
		log.Println("ERROR:", err)
		return orderDetails, err
	}

	//To get amount from Database
	var amount, amountDiscount float64

	pricingQuery := "SELECT INR_Price, INR_Discount FROM sententia.pricing_plans"
	if err := db.MySqlDB.QueryRow(pricingQuery).Scan(&amount, &amountDiscount); err != nil {
		log.Println("Error :", err)
		return orderDetails, err
	}

	twelveMonthPlansDiscount := (amount * amountDiscount) / 100
	orderDetails.TaxableAmount = amount - twelveMonthPlansDiscount

	gstAmount := func(gst float64, amount float64) float64 {
		gstPercent := gst / 100
		gstAmount := math.Round(amount * gstPercent)
		return gstAmount
	}(18, orderDetails.TaxableAmount)

	orderDetails.GSTAmount = gstAmount

	//Total in words
	if _, ok := CountryCode[txnDetails.CountryCode]; ok {
		orderDetails.TotalInWords = "Indian Rupees " + utils.IntegerToEnUs(int(orderDetails.Total)) + " only"
		orderDetails.CurrencySym = "₹"
	} else {
		orderDetails.TotalInWords = "US Dollars " + utils.IntegerToEnUs(int(orderDetails.Total)) + " only"
		orderDetails.CurrencySym = "$"
	}

	orderDetails.Email = txnDetails.Email
	orderDetails.OrderID = txnDetails.OrderID
	orderDetails.InvoiceNumber = txnDetails.InvoiceNumber
	orderDetails.BankName = statusRespMapper.Data.PaymentInstrument.Type
	orderDetails.TransactionDate = txnDetails.TransactionStartDate

	return orderDetails, nil
}

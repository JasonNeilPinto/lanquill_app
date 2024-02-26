package user

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"github.com/Lanquill/Forge/pkg/db"
	"github.com/Lanquill/Forge/pkg/queue"
	"github.com/Lanquill/Forge/pkg/utils"
)

func (userReq *ResetPassReq) ResetPassword() error {
	stmt, err := db.MySqlDB.Prepare(updateUserPassword)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		utils.EncryptSHA1(userReq.Password),
		userReq.UserId,
	)

	return err
}

// This returns whether the email id exists except for the given Id.
//
// Example if User_ID is 10 and email is id10@email.com,
// then the query will get the count of the email used by any other user
// other than User_ID 10.
//
// Return "true" if the count is greater than 0, else return "false"
func checkEmailExceptId(userId int64, email string) (bool, error) {
	var emailIdCount = 0
	err := db.MySqlDB.QueryRow(countEmailExceptId, email, userId).Scan(&emailIdCount)
	if err != nil {
		log.Println(err)
		return true, err
	}
	return emailIdCount > 0, nil
}

func GetLicenseDetails(entityId int64) (EntityLicenseResp, error) {
	licenseDetails := EntityLicenseResp{}
	err := db.MySqlDB.QueryRow(getLicenseDetailsByLicenseType, entityId).Scan(&licenseDetails.NumberOfLicences,
		&licenseDetails.LicenceTypeId,
		&licenseDetails.LicenceConsumed)
	if err != nil {
		log.Println(err)
		return licenseDetails, err
	}
	return licenseDetails, nil
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ErrorCount(list []error) map[string]int {

	errorFrequency := make(map[string]int)

	for _, item := range list {
		_, exist := errorFrequency[item.Error()]

		if exist {
			errorFrequency[item.Error()] += 1
		} else {
			errorFrequency[item.Error()] = 1
		}
	}
	return errorFrequency
}

func (userReq UpdateUserReq) Update() error {

	emailExists, err := checkEmailExceptId(userReq.UserId, userReq.Email)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	// Return error if email exists
	if emailExists {
		return errors.New(ErrEmailExists)
	}

	// Else update the user name and email id
	stmt, err := db.MySqlDB.Prepare(updateUserInfo)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		userReq.Name,
		userReq.Email,
		userReq.UserId,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	var existingUserType = 0
	err = db.MySqlDB.QueryRow(getUserTypeId, userReq.UserId).Scan(&existingUserType)
	if err != nil {
		log.Println(err)
		return err
	}

	_, existingAdmin := utils.UserTypeMap[existingUserType]

	// If Is Admin - ticked
	if userReq.IsAdmin && !existingAdmin {
		var entityType = ""
		var entityLevel = 0
		var userTypeId = 0

		// Get Entity Type of the user
		err := db.MySqlDB.QueryRow(getEntityTypeName, userReq.UserId, userReq.UserId).Scan(&entityType, &entityLevel)
		if err != nil {
			log.Println(err)
			return err
		}

		// Get Admin UserTypeId to mark the user as admin for the respective entity
		userType := utils.UserEntityTypeMap[entityType]
		err = db.MySqlDB.QueryRow(getUserTypeForAdmin, userType, entityLevel).Scan(&userTypeId)
		if err != nil {
			log.Println(err)
			return err
		}

		// Update user as admin
		stmt, err := db.MySqlDB.Prepare(updateUserType)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(
			userTypeId,
			userReq.UserId,
		)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
	} else if !userReq.IsAdmin && existingAdmin {
		// Remove user from admin role (Mark as individual user)
		stmt, err := db.MySqlDB.Prepare(updateUserType)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(
			1,
			userReq.UserId,
		)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
	}
	// TODO: Update user email and name in MongoDB
	return err
}

func (userReq UserIdReq) Delete() error {

	stmt, err := db.MySqlDB.Prepare(deleteUser)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		userReq.UserId,
	)

	return err
}

func (userReq CreateUserReq) Create() error {

	var email = ""
	err := db.MySqlDB.QueryRow(getUserEmail, userReq.UserEmail, userReq.UserMobile).Scan(&email)
	if err != nil && err != sql.ErrNoRows {
		log.Println("ERROR: ", err)
		return err
	}

	// If the user is for lower entity
	if userReq.EntityDivId != 0 {
		if email != "" {
			var userHigherLevel = ""
			userEntityLevel := UserEntityResp{}
			err = db.MySqlDB.QueryRow(checkIfUserExistsForHigherLevel, userReq.EntityId, userReq.EntityId, userReq.UserEmail).Scan(&userHigherLevel)
			if err != nil && err != sql.ErrNoRows {
				log.Println("ERROR: ", err)
				return err
			}

			if userHigherLevel != "" {

				// Get entity information of the user
				err = db.MySqlDB.QueryRow(userLowerEntityLevelDetails, userReq.EntityDivId).Scan(&userEntityLevel.EntityId,
					&userEntityLevel.ParentEntityDivId,
					&userEntityLevel.EntityDivTypeId,
					&userEntityLevel.EntityDivTypeLevel,
					&userEntityLevel.ParentEntityId,
					&userEntityLevel.EntityTypeId,
					&userEntityLevel.EntityTypeLevel)
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				updateNewUserDetailsQuery := updateUserDetails
				// Update the user table level from EntityDivTypeLevel to EntityTypeLevel
				for i := userEntityLevel.EntityDivTypeLevel; i <= userEntityLevel.EntityTypeLevel; i++ {
					if i == userEntityLevel.EntityTypeLevel {
						updateNewUserDetailsQuery = fmt.Sprintf(updateNewUserDetailsQuery, "Level_"+strconv.Itoa(int(i))+" = ?%s")
					} else {
						updateNewUserDetailsQuery = fmt.Sprintf(updateNewUserDetailsQuery, "Level_"+strconv.Itoa(int(i))+" = ?, %s")
					}
				}

				if userEntityLevel.EntityDivTypeLevel > 1 {
					for i := 1; i < int(userEntityLevel.EntityDivTypeLevel); i++ {
						if i == (int(userEntityLevel.EntityDivTypeLevel) - 1) {
							updateNewUserDetailsQuery = fmt.Sprintf(updateNewUserDetailsQuery, ", Level_"+strconv.Itoa(i)+" = ? ")
						} else {
							updateNewUserDetailsQuery = fmt.Sprintf(updateNewUserDetailsQuery, ", Level_"+strconv.Itoa(i)+" = ?%s")
						}

					}
				} else {
					updateNewUserDetailsQuery = fmt.Sprintf(updateNewUserDetailsQuery, " ")
				}

				// Fetch entityId values to be updated for the levels in user table
				entityValues := []int64{}
				entityDivId := &userReq.EntityDivId
				entityId := &userReq.EntityId

				lowest := 0
				higher := 0
				for i := userEntityLevel.EntityDivTypeLevel; i <= userEntityLevel.EntityTypeLevel; i++ {
					if i <= 4 {
						if lowest == 0 {
							entityValues = append(entityValues, userReq.EntityDivId)
						} else {
							var parentEntityDivId = int64(0)
							err = db.MySqlDB.QueryRow(getParentEntityDivId, entityDivId).Scan(&parentEntityDivId)
							if err != nil {
								log.Println("ERROR: ", err)
								return err
							}
							entityValues = append(entityValues, parentEntityDivId)
							entityDivId = &parentEntityDivId
						}
						lowest++
					}
					if i > 4 {

						if higher == 0 {
							entityValues = append(entityValues, userReq.EntityId)
						} else {
							var parentEntityId = int64(0)
							err = db.MySqlDB.QueryRow(getParentEntityId, entityId).Scan(&parentEntityId)
							if err != nil {
								log.Println("ERROR: ", err)
								return err
							}
							entityValues = append(entityValues, parentEntityId)
							entityId = &parentEntityId
						}
						higher++
					}
				}
				stmt, err := db.MySqlDB.Prepare(updateNewUserDetailsQuery)
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}
				defer stmt.Close()

				args := make([]interface{}, 0, len(entityValues))
				for _, val := range entityValues {
					args = append(args, val)
				}
				
				if userEntityLevel.EntityDivTypeLevel > 1 {
					for i := 1; i < int(userEntityLevel.EntityDivTypeLevel); i++ {
						args = append(args, sql.NullString{})

					}
				}

				args = append(args, userReq.EntityId)
				args = append(args, userReq.EntityId)
				args = append(args, userReq.UserEmail)

				_, err = stmt.Exec(args...)

				return err
			} else {
				return errors.New(ErrEmailOrMobExists)
			}
		}

		// Get entity details to be updated once the new user is inserted
		userEntityLevel := UserEntityResp{}
		err = db.MySqlDB.QueryRow(userLowerEntityLevelDetails, userReq.EntityDivId).Scan(&userEntityLevel.EntityId,
			&userEntityLevel.ParentEntityDivId,
			&userEntityLevel.EntityDivTypeId,
			&userEntityLevel.EntityDivTypeLevel,
			&userEntityLevel.ParentEntityId,
			&userEntityLevel.EntityTypeId,
			&userEntityLevel.EntityTypeLevel)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		updateNewUserDetailStmt := updateNewUserDetails
		for i := userEntityLevel.EntityDivTypeLevel; i <= userEntityLevel.EntityTypeLevel; i++ {
			if i == userEntityLevel.EntityTypeLevel {
				updateNewUserDetailStmt = fmt.Sprintf(updateNewUserDetailStmt, "Level_"+strconv.Itoa(int(i))+" = ?")
			} else {
				updateNewUserDetailStmt = fmt.Sprintf(updateNewUserDetailStmt, "Level_"+strconv.Itoa(int(i))+" = ?, %s")
			}
		}

		entityValues := []int64{}
		entityDivId := &userReq.EntityDivId
		entityId := &userReq.EntityId

		lowest := 0
		higher := 0
		for i := userEntityLevel.EntityDivTypeLevel; i <= userEntityLevel.EntityTypeLevel; i++ {
			if i <= 4 {
				if lowest == 0 {
					entityValues = append(entityValues, userReq.EntityDivId)
				} else {
					var parentEntityDivId = int64(0)
					err = db.MySqlDB.QueryRow(getParentEntityDivId, entityDivId).Scan(&parentEntityDivId)
					if err != nil {
						log.Println("ERROR: ", err)
						return err
					}
					entityValues = append(entityValues, parentEntityDivId)
					entityDivId = &parentEntityDivId
				}
				lowest++
			}
			if i > 4 {
				if higher == 0 {
					entityValues = append(entityValues, userReq.EntityId)
				} else {
					var parentEntityId = int64(0)
					err = db.MySqlDB.QueryRow(getParentEntityId, entityId).Scan(&parentEntityId)
					if err != nil {
						log.Println("ERROR: ", err)
						return err
					}
					entityValues = append(entityValues, parentEntityId)
					entityId = &parentEntityId
				}
				higher++
			}
		}

		// Fetch entity license details
		entityLicenseDetails := []EntityLicenseResp{}
		stmt, err := db.MySqlDB.Prepare(getEntityLicenseDetails)
		if err != nil {
			log.Println(err)
			return err
		}
		defer stmt.Close()

		rows, err := stmt.Query(userReq.EntityId)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return err
		}
		defer rows.Close()

		for rows.Next() {
			eachEntityLicenseResp := EntityLicenseResp{}
			if err := rows.Scan(
				&eachEntityLicenseResp.NumberOfLicences,
				&eachEntityLicenseResp.LicenceTypeId,
				&eachEntityLicenseResp.LicenceConsumed); err != nil {
				log.Println(err)
				return err
			}

			entityLicenseDetails = append(entityLicenseDetails, eachEntityLicenseResp)

		}

		var parentEntityId = int64(0)
		err = db.MySqlDB.QueryRow(getParentEntityId, userReq.EntityId).Scan(&parentEntityId)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		parentEntityLicenseDetails := []EntityLicenseResp{}
		stmt, err = db.MySqlDB.Prepare(getEntityLicenseDetails)
		if err != nil {
			log.Println(err)
			return err
		}
		defer stmt.Close()

		rows, err = stmt.Query(parentEntityId)
		if err != nil {
			log.Println(err)
			return err
		}
		defer rows.Close()

		for rows.Next() {
			eachParentEntityLicenseResp := EntityLicenseResp{}
			if err := rows.Scan(
				&eachParentEntityLicenseResp.NumberOfLicences,
				&eachParentEntityLicenseResp.LicenceTypeId,
				&eachParentEntityLicenseResp.LicenceConsumed); err != nil {
				log.Println(err)
				return err
			}

			parentEntityLicenseDetails = append(parentEntityLicenseDetails, eachParentEntityLicenseResp)

		}

		var isParentOrChild = "child"
		// First check lower eid if availble, deduct from here otherwise from parent
		if entityLicenseDetails != nil {
			for _, entityElement := range entityLicenseDetails {
				if entityElement.LicenceTypeId == 4 {
					if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + 1) {
						for _, parentEntityElement := range parentEntityLicenseDetails {
							if parentEntityElement.LicenceTypeId == 4 {
								if parentEntityElement.NumberOfLicences < (parentEntityElement.LicenceConsumed + 1) {
									return errors.New(ErrNoUserLicense)
								} else {
									isParentOrChild = "parent"
								}
							}
						}
					}
				}
				if entityElement.LicenceTypeId == 1 {
					if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + userReq.DocQuantity) {
						for _, parentEntityElement := range parentEntityLicenseDetails {
							if parentEntityElement.LicenceTypeId == 1 {
								if parentEntityElement.NumberOfLicences < (parentEntityElement.LicenceConsumed + userReq.DocQuantity) {
									return errors.New(ErrNoUserLicense)
								} else {
									isParentOrChild = "parent"
								}
							}
						}
					}
				}
				if entityElement.LicenceTypeId == 5 {
					if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + userReq.CertQuantity) {
						for _, parentEntityElement := range parentEntityLicenseDetails {
							if parentEntityElement.LicenceTypeId == 5 {
								if parentEntityElement.NumberOfLicences < (parentEntityElement.LicenceConsumed + userReq.DocQuantity) {
									return errors.New(ErrNoUserLicense)
								} else {
									isParentOrChild = "parent"
								}
							}
						}
					}
				}
			}

			entityidforupdate := int64(0)
			if isParentOrChild == "child" {
				entityidforupdate = userReq.EntityId
			} else {
				entityidforupdate = parentEntityId
			}

			updateEntityLicenseCert := updateEntityLicense
			updateEntityLicenseCert = fmt.Sprintf(updateEntityLicenseCert, userReq.CertQuantity)
			stmt, err := db.MySqlDB.Prepare(updateEntityLicenseCert)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}
			defer stmt.Close()

			_, err = stmt.Exec(entityidforupdate, 5)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}

			updateEntityLicenseDoc := updateEntityLicense
			updateEntityLicenseDoc = fmt.Sprintf(updateEntityLicenseDoc, userReq.DocQuantity)
			stmt, err = db.MySqlDB.Prepare(updateEntityLicenseDoc)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}
			defer stmt.Close()

			_, err = stmt.Exec(entityidforupdate, 1)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}

			updateEntityLicenceUser := updateEntityLicense
			updateEntityLicenceUser = fmt.Sprintf(updateEntityLicenceUser, 1)

			// Update license details
			stmt, err = db.MySqlDB.Prepare(updateEntityLicenceUser)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}
			defer stmt.Close()

			_, err = stmt.Exec(entityidforupdate, 4)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}

		} else {
			// Deduct from parent
			var parentEntityId = int64(0)
			err = db.MySqlDB.QueryRow(getParentEntityId, userReq.EntityId).Scan(&parentEntityId)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}

			parentEntityLicenseDetails := []EntityLicenseResp{}
			stmt, err := db.MySqlDB.Prepare(getEntityLicenseDetails)
			if err != nil {
				log.Println(err)
				return err
			}
			defer stmt.Close()

			rows, err := stmt.Query(parentEntityId)
			if err != nil {
				log.Println(err)
				return err
			}
			defer rows.Close()

			for rows.Next() {
				eachParentEntityLicenseResp := EntityLicenseResp{}
				if err := rows.Scan(
					&eachParentEntityLicenseResp.NumberOfLicences,
					&eachParentEntityLicenseResp.LicenceTypeId,
					&eachParentEntityLicenseResp.LicenceConsumed); err != nil {
					log.Println(err)
					return err
				}

				parentEntityLicenseDetails = append(parentEntityLicenseDetails, eachParentEntityLicenseResp)

			}

			if parentEntityLicenseDetails != nil {
				for _, entityElement := range parentEntityLicenseDetails {
					if entityElement.LicenceTypeId == 4 {
						if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + 1) {
							return errors.New(ErrNoUserLicense)
						}
					}
					if entityElement.LicenceTypeId == 1 {
						if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + userReq.DocQuantity) {
							return errors.New(ErrNoUserLicense)
						}
					}
					if entityElement.LicenceTypeId == 5 {
						if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + userReq.CertQuantity) {
							return errors.New(ErrNoUserLicense)
						}
					}
				}
			} else {
				return errors.New(ErrNoUserLicense)
			}
		}

		args := make([]interface{}, 0, len(entityValues))
		for _, val := range entityValues {
			args = append(args, val)
		}

		// Insert user details into user table
		stmt, err = db.MySqlDB.Prepare(insertUserDetails)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(userReq.UserName,
			userReq.UserEmail,
			userReq.UserMobile,
			utils.EncryptSHA1(userReq.Password),
			time.Now(),
			time.Now().AddDate(1, 0, 0),
			"Active",
			1)

		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		args = append(args, userReq.UserEmail)

		//TODO: These logs will be removed once the user creation issues are fixed

		log.Println("Update User Statement", updateNewUserDetailStmt)
		log.Println("Arguments in order", args)

		// Update level in user table with associated entityId
		stmt, err = db.MySqlDB.Prepare(updateNewUserDetailStmt)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(args...)

		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		var userId = ""
		err = db.MySqlDB.QueryRow(getUserId, userReq.UserEmail).Scan(&userId)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		licenseType := make(chan int)
		go func() {
			licenseType <- 1
			licenseType <- 2
			close(licenseType)
		}()

		stmt, err = db.MySqlDB.Prepare(insertUserLicense)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
		defer stmt.Close()

		docValidTill, err := time.Parse("2006-01-02", userReq.DocValidTill)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		certValidTill, err := time.Parse("2006-01-02", userReq.CertValidTill)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		for n := range licenseType {
			if n == 1 {
				// Insert for doc
				_, err = stmt.Exec(userId,
					1,
					userReq.DocQuantity,
					time.Now(),
					docValidTill)

				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}
			}
			if n == 2 {
				// Insert for cert
				_, err = stmt.Exec(userId,
					5,
					userReq.CertQuantity,
					time.Now(),
					certValidTill)

				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}
			}
		}

		// Sending welcome email -> feeding info into msg queue
		sent := sendRegistrationEmailQueue(userReq.UserEmail, userReq.UserName, userReq.Password, 0, userReq.EntityDivId)

		if sent < 0 {
			log.Println("ERROR: ", err)
			return err
		}

		return err
	}

	// If the user is for higher entity
	if userReq.EntityId != 0 {
		if email != "" {
			var userHigherLevel = ""
			userEntityLevel := UserHigherEntityResp{}
			err = db.MySqlDB.QueryRow(checkIfUserExistsForHigherLevel, userReq.EntityId, userReq.EntityId, userReq.UserEmail).Scan(&userHigherLevel)
			if err != nil && err != sql.ErrNoRows {
				log.Println("ERROR: ", err)
				return err
			}

			if userHigherLevel != "" {
				err = db.MySqlDB.QueryRow(userHigherEntityLevelDetails, userReq.EntityId).Scan(&userEntityLevel.EntityTypeId,
					&userEntityLevel.EntityTypeLevel,
					&userEntityLevel.ParentEntityId,
					&userEntityLevel.ParentEntityTypeId,
					&userEntityLevel.ParentEntityTypeLevel)
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				updateUserDetailsQuery := updateUserDetails
				for i := userEntityLevel.EntityTypeLevel; i <= userEntityLevel.ParentEntityTypeLevel; i++ {
					if i == userEntityLevel.ParentEntityTypeLevel {
						updateUserDetailsQuery = fmt.Sprintf(updateUserDetailsQuery, "Level_"+strconv.Itoa(int(i))+" = ?%s")
					} else {
						updateUserDetailsQuery = fmt.Sprintf(updateUserDetailsQuery, "Level_"+strconv.Itoa(int(i))+" = ?, %s")
					}
				}

				if userEntityLevel.EntityTypeLevel > 1 {
					for i := 1; i < int(userEntityLevel.EntityTypeLevel); i++ {
						if i == (int(userEntityLevel.EntityTypeLevel) - 1) {
							updateUserDetailsQuery = fmt.Sprintf(updateUserDetailsQuery, ",Level_"+strconv.Itoa(i)+" = ? ")
						} else {
							updateUserDetailsQuery = fmt.Sprintf(updateUserDetailsQuery, ",Level_"+strconv.Itoa(i)+" = ?%s")
						}

					}
				} else {
					updateUserDetailsQuery = fmt.Sprintf(updateUserDetailsQuery, " ")
				}

				entityValues := []int64{}
				entityDivId := &userReq.EntityDivId
				entityId := &userReq.EntityId

				lowest := 0
				higher := 0
				for i := userEntityLevel.EntityTypeLevel; i <= userEntityLevel.ParentEntityTypeLevel; i++ {
					if i <= 4 {
						if lowest == 0 {
							entityValues = append(entityValues, userReq.EntityDivId)
						} else {
							var parentEntityDivId = int64(0)
							err = db.MySqlDB.QueryRow(getParentEntityDivId, entityDivId).Scan(&parentEntityDivId)
							if err != nil {
								log.Println("ERROR: ", err)
								return err
							}
							entityValues = append(entityValues, parentEntityDivId)
							entityDivId = &parentEntityDivId
						}
						lowest++
					}
					if i > 4 {
						if higher == 0 {
							entityValues = append(entityValues, userReq.EntityId)
						} else {
							var parentEntityId = int64(0)
							err = db.MySqlDB.QueryRow(getParentEntityId, entityId).Scan(&parentEntityId)
							if err != nil {
								log.Println("ERROR: ", err)
								return err
							}
							entityValues = append(entityValues, parentEntityId)
							entityId = &parentEntityId
						}
						higher++
					}
				}

				stmt, err := db.MySqlDB.Prepare(updateUserDetailsQuery)
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}
				defer stmt.Close()

				args := make([]interface{}, 0, len(entityValues))
				for _, val := range entityValues {
					args = append(args, val)
				}

				if userEntityLevel.EntityTypeLevel > 1 {
					for i := 1; i < int(userEntityLevel.EntityTypeLevel); i++ {
						args = append(args, sql.NullString{})

					}
				}

				args = append(args, userReq.EntityId)
				args = append(args, userReq.EntityId)
				args = append(args, userReq.UserEmail)

				_, err = stmt.Exec(args...)

				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}
				return err
			} else {
				return errors.New(ErrEmailOrMobExists)
			}
		}
		userEntityLevel := UserHigherEntityResp{}
		err = db.MySqlDB.QueryRow(userHigherEntityLevelDetails, userReq.EntityId).Scan(&userEntityLevel.EntityTypeId,
			&userEntityLevel.EntityTypeLevel,
			&userEntityLevel.ParentEntityId,
			&userEntityLevel.ParentEntityTypeId,
			&userEntityLevel.ParentEntityTypeLevel)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		updateNewUserDetailStatement := updateNewUserDetails
		for i := userEntityLevel.EntityTypeLevel; i <= userEntityLevel.ParentEntityTypeLevel; i++ {
			if i == userEntityLevel.ParentEntityTypeLevel {
				updateNewUserDetailStatement = fmt.Sprintf(updateNewUserDetailStatement, "Level_"+strconv.Itoa(int(i))+" = ?")
			} else {
				updateNewUserDetailStatement = fmt.Sprintf(updateNewUserDetailStatement, "Level_"+strconv.Itoa(int(i))+" = ?, %s")
			}
		}

		entityValues := []int64{}
		entityDivId := &userReq.EntityDivId
		entityId := &userReq.EntityId

		lowest := 0
		higher := 0
		for i := userEntityLevel.EntityTypeLevel; i <= userEntityLevel.ParentEntityTypeLevel; i++ {
			if i <= 4 {
				if lowest == 0 {
					entityValues = append(entityValues, userReq.EntityDivId)
				} else {
					var parentEntityDivId = int64(0)
					err = db.MySqlDB.QueryRow(getParentEntityDivId, entityDivId).Scan(&parentEntityDivId)
					if err != nil {
						log.Println("ERROR: ", err)
						return err
					}
					entityValues = append(entityValues, parentEntityDivId)
					entityDivId = &parentEntityDivId
				}
				lowest++
			}
			if i > 4 {
				if higher == 0 {
					entityValues = append(entityValues, userReq.EntityId)
				} else {
					var parentEntityId = int64(0)
					err = db.MySqlDB.QueryRow(getParentEntityId, entityId).Scan(&parentEntityId)
					if err != nil {
						log.Println("ERROR: ", err)
						return err
					}
					entityValues = append(entityValues, parentEntityId)
					entityId = &parentEntityId
				}
				higher++
			}
		}

		entityLicenseDetails := []EntityLicenseResp{}
		stmt, err := db.MySqlDB.Prepare(getEntityLicenseDetails)
		if err != nil {
			log.Println(err)
			return err
		}
		defer stmt.Close()

		rows, err := stmt.Query(userReq.EntityId)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return err
		}
		defer rows.Close()

		for rows.Next() {
			eachEntityLicenseResp := EntityLicenseResp{}
			if err := rows.Scan(
				&eachEntityLicenseResp.NumberOfLicences,
				&eachEntityLicenseResp.LicenceTypeId,
				&eachEntityLicenseResp.LicenceConsumed); err != nil {
				log.Println(err)
				return err
			}

			entityLicenseDetails = append(entityLicenseDetails, eachEntityLicenseResp)

		}

		var parentEntityId = int64(0)
		err = db.MySqlDB.QueryRow(getParentEntityId, userReq.EntityId).Scan(&parentEntityId)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		parentEntityLicenseDetails := []EntityLicenseResp{}
		stmt, err = db.MySqlDB.Prepare(getEntityLicenseDetails)
		if err != nil {
			log.Println(err)
			return err
		}
		defer stmt.Close()

		rows, err = stmt.Query(parentEntityId)
		if err != nil {
			log.Println(err)
			return err
		}
		defer rows.Close()

		for rows.Next() {
			eachParentEntityLicenseResp := EntityLicenseResp{}
			if err := rows.Scan(
				&eachParentEntityLicenseResp.NumberOfLicences,
				&eachParentEntityLicenseResp.LicenceTypeId,
				&eachParentEntityLicenseResp.LicenceConsumed); err != nil {
				log.Println(err)
				return err
			}

			parentEntityLicenseDetails = append(parentEntityLicenseDetails, eachParentEntityLicenseResp)

		}

		var isParentOrChild = "child"
		if entityLicenseDetails != nil {
			for _, entityElement := range entityLicenseDetails {
				if entityElement.LicenceTypeId == 4 {
					if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + 1) {
						for _, parentEntityElement := range parentEntityLicenseDetails {
							if parentEntityElement.LicenceTypeId == 4 {
								if parentEntityElement.NumberOfLicences < (parentEntityElement.LicenceConsumed + 1) {
									return errors.New(ErrNoUserLicense)
								} else {
									isParentOrChild = "parent"
								}
							}
						}
					}
				}
				if entityElement.LicenceTypeId == 1 {
					if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + userReq.DocQuantity) {
						for _, parentEntityElement := range parentEntityLicenseDetails {
							if parentEntityElement.LicenceTypeId == 1 {
								if parentEntityElement.NumberOfLicences < (parentEntityElement.LicenceConsumed + userReq.DocQuantity) {
									return errors.New(ErrNoUserLicense)
								} else {
									isParentOrChild = "parent"
								}
							}
						}
					}
				}
				if entityElement.LicenceTypeId == 5 {
					if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + userReq.CertQuantity) {
						for _, parentEntityElement := range parentEntityLicenseDetails {
							if parentEntityElement.LicenceTypeId == 5 {
								if parentEntityElement.NumberOfLicences < (parentEntityElement.LicenceConsumed + userReq.DocQuantity) {
									return errors.New(ErrNoUserLicense)
								} else {
									isParentOrChild = "parent"
								}
							}
						}
					}
				}
			}

			entityidforupdate := int64(0)
			if isParentOrChild == "child" {
				entityidforupdate = userReq.EntityId
			} else {
				entityidforupdate = parentEntityId
			}

			updateEntityLicenseCert := updateEntityLicense
			updateEntityLicenseCert = fmt.Sprintf(updateEntityLicenseCert, userReq.CertQuantity)
			stmt, err := db.MySqlDB.Prepare(updateEntityLicenseCert)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}
			defer stmt.Close()

			_, err = stmt.Exec(entityidforupdate, 5)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}

			updateEntityLicenseDoc := updateEntityLicense
			updateEntityLicenseDoc = fmt.Sprintf(updateEntityLicenseDoc, userReq.DocQuantity)
			stmt, err = db.MySqlDB.Prepare(updateEntityLicenseDoc)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}
			defer stmt.Close()

			_, err = stmt.Exec(entityidforupdate, 1)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}

			updateNewUserEntityLicense := updateEntityLicense
			updateNewUserEntityLicense = fmt.Sprintf(updateNewUserEntityLicense, 1)
			stmt, err = db.MySqlDB.Prepare(updateNewUserEntityLicense)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}
			defer stmt.Close()

			_, err = stmt.Exec(entityidforupdate, 4)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}

		} else {
			var parentEntityId = int64(0)
			err = db.MySqlDB.QueryRow(getParentEntityId, userReq.EntityId).Scan(&parentEntityId)
			if err != nil {
				log.Println("ERROR: ", err)
				return err
			}

			parentEntityLicenseDetails := []EntityLicenseResp{}
			stmt, err := db.MySqlDB.Prepare(getEntityLicenseDetails)
			if err != nil {
				log.Println(err)
				return err
			}
			defer stmt.Close()

			rows, err := stmt.Query(parentEntityId)
			if err != nil {
				log.Println(err)
				return err
			}
			defer rows.Close()

			for rows.Next() {
				eachParentEntityLicenseResp := EntityLicenseResp{}
				if err := rows.Scan(
					&eachParentEntityLicenseResp.NumberOfLicences,
					&eachParentEntityLicenseResp.LicenceTypeId,
					&eachParentEntityLicenseResp.LicenceConsumed); err != nil {
					log.Println(err)
					return err
				}

				parentEntityLicenseDetails = append(parentEntityLicenseDetails, eachParentEntityLicenseResp)

			}

			if parentEntityLicenseDetails != nil {
				for _, entityElement := range parentEntityLicenseDetails {
					if entityElement.LicenceTypeId == 4 {
						if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + 1) {
							return errors.New(ErrNoUserLicense)
						}
					}
					if entityElement.LicenceTypeId == 1 {
						if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + userReq.DocQuantity) {
							return errors.New(ErrNoUserLicense)
						}
					}
					if entityElement.LicenceTypeId == 5 {
						if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + userReq.CertQuantity) {
							return errors.New(ErrNoUserLicense)
						}
					}
				}
			} else {
				return errors.New(ErrNoUserLicense)
			}
		}

		args := make([]interface{}, 0, len(entityValues))
		for _, val := range entityValues {
			args = append(args, val)
		}

		stmt, err = db.MySqlDB.Prepare(insertUserDetails)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(userReq.UserName,
			userReq.UserEmail,
			userReq.UserMobile,
			utils.EncryptSHA1(userReq.Password),
			time.Now(),
			time.Now().AddDate(1, 0, 0),
			"Active",
			1)

		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		args = append(args, userReq.UserEmail)

		//TODO: These logs will be removed once the user creation issues are fixed

		log.Println("Update User Statement", updateNewUserDetailStatement)
		log.Println("Arguments in order", args)

		// Update level in user table with associated entityId
		stmt, err = db.MySqlDB.Prepare(updateNewUserDetailStatement)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(args...)

		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		var userId = ""
		err = db.MySqlDB.QueryRow(getUserId, userReq.UserEmail).Scan(&userId)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		licenseType := make(chan int)
		go func() {
			licenseType <- 1
			licenseType <- 2
			close(licenseType)
		}()

		stmt, err = db.MySqlDB.Prepare(insertUserLicense)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
		defer stmt.Close()

		docValidTill, err := time.Parse("2006-01-02", userReq.DocValidTill)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		certValidTill, err := time.Parse("2006-01-02", userReq.CertValidTill)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		for n := range licenseType {
			if n == 1 {
				// Insert for doc
				_, err = stmt.Exec(userId,
					1,
					userReq.DocQuantity,
					time.Now(),
					docValidTill)

				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}
			}
			if n == 2 {
				// Insert for cert
				_, err = stmt.Exec(userId,
					5,
					userReq.CertQuantity,
					time.Now(),
					certValidTill)

				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}
			}
		}

		// Sending welcome email -> feeding info into msg queue
		sent := sendRegistrationEmailQueue(userReq.UserEmail, userReq.UserName, userReq.Password, userReq.EntityId, 0)

		if sent < 0 {
			log.Println("ERROR: ", err)
			return err
		}

		return err
	} else { // If the user is not related to any entity

		if email != "" {
			return errors.New(ErrEmailOrMobExists)
		}
		stmt, err := db.MySqlDB.Prepare(insertUserDetails)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(userReq.UserName,
			userReq.UserEmail,
			userReq.UserMobile,
			utils.EncryptSHA1(userReq.Password),
			time.Now(),
			time.Now().AddDate(1, 0, 0),
			"Active",
			1)

		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		var userId = ""
		err = db.MySqlDB.QueryRow(getUserId, userReq.UserEmail).Scan(&userId)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		licenseType := make(chan int)
		go func() {
			licenseType <- 1
			licenseType <- 2
			close(licenseType)
		}()

		stmt, err = db.MySqlDB.Prepare(insertUserLicense)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
		defer stmt.Close()

		docValidTill, err := time.Parse("2006-01-02", userReq.DocValidTill)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		certValidTill, err := time.Parse("2006-01-02", userReq.CertValidTill)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		for n := range licenseType {
			if n == 1 {
				// Insert for doc
				_, err = stmt.Exec(userId,
					1,
					userReq.DocQuantity,
					time.Now(),
					docValidTill)

				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}
			}
			if n == 2 {
				// Insert for cert
				_, err = stmt.Exec(userId,
					5,
					userReq.CertQuantity,
					time.Now(),
					certValidTill)

				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}
			}
		}

		// Sending welcome email -> feeding info into msg queue
		sent := sendRegistrationEmailQueue(userReq.UserEmail, userReq.UserName, userReq.Password, 0, 0)
		if sent < 0 {
			log.Println("ERROR: ", err)
			return err
		}
		return err

	}
}

func sendRegistrationEmailQueue(email string, userName string, password string, entityId int64, entityDivId int64) int {
	userData := make(map[string]interface{})
	userData["name"] = userName
	userData["email"] = email
	userData["pass"] = password
	userData["entityId"] = entityId
	userData["entityDivId"] = entityDivId
	return queue.ServiceMQ("user_registration_welcome_queue", userData)
}

func (changeExpiryDateReq *ChangeExpiryDateReq) Update() error {

	updateStmt := ""
	id := int64(0)

	if changeExpiryDateReq.RequestType == "user" {
		updateStmt = changeSingleUserExpiryDate
		id = *changeExpiryDateReq.UserId
	} else if changeExpiryDateReq.RequestType == "entity" {
		if strings.ToLower(changeExpiryDateReq.EntityType) == "university" ||
			strings.ToLower(changeExpiryDateReq.EntityType) == "corporate_global" ||
			strings.ToLower(changeExpiryDateReq.EntityType) == "school_chain" {
			updateStmt = changeHigherEntityExpiryDate
		}

		if strings.ToLower(changeExpiryDateReq.EntityType) == "college" ||
			strings.ToLower(changeExpiryDateReq.EntityType) == "corporate_country" ||
			strings.ToLower(changeExpiryDateReq.EntityType) == "school" {
			updateStmt = changeLowerEntityExpiryDate
		}
		id = *changeExpiryDateReq.EntityId
	} else {
		return errors.New(ErrFailed)
	}

	expiryDate, err := time.Parse("2006-01-02", changeExpiryDateReq.ExpiryDate)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	stmt, err := db.MySqlDB.Prepare(updateStmt)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		expiryDate,
		id,
	)

	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	updateLicence, err := db.MySqlDB.Prepare(updateEntityLicenceValidity)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	_, err = updateLicence.Exec(
		expiryDate,
		id,
	)

	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	return err
}

func GetUserDetails(userEmail string) (UserDetailsResp, error) {
	userDetails := UserDetailsResp{}
	err := db.MySqlDB.QueryRow(getUserDetail, userEmail).Scan(&userDetails.UserId,
		&userDetails.UserEmail,
		&userDetails.UserName,
		&userDetails.PasswordAttempts,
		&userDetails.UserStatus,
		&userDetails.UserType)
	if err != nil {
		log.Println(err)
		return userDetails, err
	}
	return userDetails, nil
}

func (entityDetailsReq *EntityDetailsReq) GetEntityDetails() ([]EntityDetailsResp, error) {
	entityDetails := []EntityDetailsResp{}

	if entityDetailsReq.RequestType == "" {
		if entityDetailsReq.UserEmail != "" {
			stmt, err := db.MySqlDB.Prepare(getEntityDetailsByEmail)
			if err != nil {
				log.Println(err)
				return entityDetails, err
			}
			defer stmt.Close()

			rows, err := stmt.Query(entityDetailsReq.UserEmail)
			if err != nil {
				log.Println(err)
				return entityDetails, err
			}
			defer rows.Close()

			for rows.Next() {
				eachEntityDetails := EntityDetailsResp{}
				if err := rows.Scan(&eachEntityDetails.EntityName,
					&eachEntityDetails.EntityId,
					&eachEntityDetails.EntityType,
					&eachEntityDetails.EntityInfo); err != nil {
					log.Println(err)
					return entityDetails, err
				}
				entityDetails = append(entityDetails, eachEntityDetails)

			}
		} else {
			stmt, err := db.MySqlDB.Prepare(getEntityDetailsByEntityId)
			if err != nil {
				log.Println(err)
				return entityDetails, err
			}
			defer stmt.Close()

			if entityDetailsReq.UserType == 2 {
				entityDetailsReq.EntityId = nil
			}

			rows, err := stmt.Query(entityDetailsReq.EntityId, entityDetailsReq.EntityId)
			if err != nil {
				log.Println(err)
				return entityDetails, err
			}
			defer rows.Close()

			for rows.Next() {
				eachEntityDetails := EntityDetailsResp{}
				if err := rows.Scan(&eachEntityDetails.EntityName,
					&eachEntityDetails.EntityId,
					&eachEntityDetails.EntityType,
					&eachEntityDetails.EntityInfo); err != nil {
					log.Println(err)
					return entityDetails, err
				}

				entityDetails = append(entityDetails, eachEntityDetails)

			}
		}
	}

	if entityDetailsReq.RequestType == "user" {
		eachEntityDetails := EntityDetailsResp{}
		err := db.MySqlDB.QueryRow(getUserDetailsByEmail, entityDetailsReq.UserEmail).Scan(&eachEntityDetails.UserId,
			&eachEntityDetails.UserName,
			&eachEntityDetails.ExpiryDate)

		eachEntityDetails.NumOfUsers = int64(1)

		if err != nil {
			log.Println("ERROR: ", err)
			return entityDetails, err
		}
		entityDetails = append(entityDetails, eachEntityDetails)
	}

	if entityDetailsReq.RequestType == "entity" {
		eachEntityDetails := EntityDetailsResp{}
		updateStmt := ""
		if strings.ToLower(entityDetailsReq.EntityType) == "university" ||
			strings.ToLower(entityDetailsReq.EntityType) == "corporate_global" ||
			strings.ToLower(entityDetailsReq.EntityType) == "school_chain" ||
			strings.ToLower(entityDetailsReq.EntityType) == "college" {
			updateStmt = getHigherEntityDetails
		}

		if strings.ToLower(entityDetailsReq.EntityType) == "department" ||
			strings.ToLower(entityDetailsReq.EntityType) == "corporate_country" ||
			strings.ToLower(entityDetailsReq.EntityType) == "school" {
			updateStmt = getLowerEntityDetails
		}

		err := db.MySqlDB.QueryRow(updateStmt, entityDetailsReq.EntityId, entityDetailsReq.EntityId).Scan(&eachEntityDetails.ExpiryDate,
			&eachEntityDetails.NumOfUsers)

		if err != nil {
			log.Println("ERROR: ", err)
			return entityDetails, err
		}
		entityDetails = append(entityDetails, eachEntityDetails)
	}

	return entityDetails, nil
}

func UserCreationErrorTemplate(bulkUploadFailResp []BulkUploadFailResp) (string, error) {
	upload_err_template := `
    <strong>Bulk Upload Failed Users</strong>
    <br>
    <br>
    <table>
		<tr>
			<th style="border: 1px solid #dddddd; font-weight: bold;">User Email</th>
			<th style="border: 1px solid #dddddd; font-weight: bold;">Error</th>
		</tr>
		{{range .BulkUploadFailResp }}
			<tr>
				<td style="border: 1px solid #dddddd">{{.UserEmail}}</td>
				<td style="border: 1px solid #dddddd">{{.Error}}</td>
			</tr>
		{{end}}
    </table>
    `

	renderData := struct {
		BulkUploadFailResp []BulkUploadFailResp
	}{
		BulkUploadFailResp: bulkUploadFailResp,
	}

	tmpl, err := template.New("upload_err").Parse(upload_err_template)
	if err != nil {
		log.Println("ERROR :", err)
		return "", err
	}

	var htmlBuffer bytes.Buffer
	err = tmpl.Execute(&htmlBuffer, renderData)
	if err != nil {
		log.Println("ERROR :", err)
		return "", err
	}

	return htmlBuffer.String(), nil
}
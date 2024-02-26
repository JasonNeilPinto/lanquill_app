package entity

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Lanquill/Forge/pkg/db"
	"github.com/Lanquill/Forge/pkg/utils"
)

func (entityReq *UpdateEntityReq) UpdateHigherEntity(file multipart.File, filename string) error {

	var fileNameDB *string = nil
	logoEnabled := "no"
	if file != nil {
		f, err := os.OpenFile(utils.LOGOPATH+"/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
			return err
		}
		defer f.Close()
		io.Copy(f, file)

		tempPath := "/static/logo/" + filename
		fileNameDB = &tempPath
		logoEnabled = "yes"
	}

	stmt, err := db.MySqlDB.Prepare(updateHigherEntityInfo)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(
		entityReq.Name,
		&fileNameDB,
		logoEnabled,
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

	return err
}

func (entityReq *UpdateEntityReq) UpdateLowerEntity() error {

	stmt, err := db.MySqlDB.Prepare(updateLowerEntityInfo)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	defer stmt.Close()

	row, err := stmt.Exec(
		entityReq.Name,
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

	return err
}

func (entityReq *CreateEntityReq) CreateEntity() (int64, error) {

	insertedEntityId := int64(0)
	if entityReq.EntityTypeName != "" {
		entityReq.EntityTypeName = strings.Replace(entityReq.EntityTypeName, " ", "_", -1)
	}

	// For higher entities
	if strings.ToLower(entityReq.EntityTypeName) == "university" ||
		strings.ToLower(entityReq.EntityTypeName) == "college" ||
		strings.ToLower(entityReq.EntityTypeName) == "corporate_global" ||
		strings.ToLower(entityReq.EntityTypeName) == "school_chain" ||
		strings.ToLower(entityReq.EntityTypeName) == "corporate_country" ||
		strings.ToLower(entityReq.EntityTypeName) == "school" {

		// If parent entity is available
		if entityReq.EntityId != 0 {

			var entityName = ""
			err := db.MySqlDB.QueryRow(getEntityName, entityReq.EntityName, entityReq.EntityId).Scan(&entityName)
			if err != nil && err != sql.ErrNoRows {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			if entityName != "" {
				return insertedEntityId, errors.New(ErrEntityExists)
			}

			var entityTypeId = ""
			err = db.MySqlDB.QueryRow(getEntityTypeId, entityReq.EntityTypeName).Scan(&entityTypeId)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			var parentEntityTypeId = ""
			err = db.MySqlDB.QueryRow(getParentEntityTypeId, entityReq.EntityId).Scan(&parentEntityTypeId)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			parentEntityLicenseDetails := []EntityLicenseResp{}
			stmt, err := db.MySqlDB.Prepare(getParentEntityLicenseDetails)
			if err != nil {
				log.Println(err)
				return insertedEntityId, err
			}
			defer stmt.Close()

			rows, err := stmt.Query(entityReq.EntityId)
			if err != nil && err != sql.ErrNoRows {
				log.Println(err)
				return insertedEntityId, err
			}
			defer rows.Close()

			for rows.Next() {
				eachParentEntityLicenseResp := EntityLicenseResp{}
				if err := rows.Scan(
					&eachParentEntityLicenseResp.NumberOfLicences,
					&eachParentEntityLicenseResp.LicenceTypeId,
					&eachParentEntityLicenseResp.LicenceConsumed); err != nil {
					log.Println(err)
					return insertedEntityId, err
				}

				parentEntityLicenseDetails = append(parentEntityLicenseDetails, eachParentEntityLicenseResp)

			}

			if parentEntityLicenseDetails != nil {
				for _, entityElement := range parentEntityLicenseDetails {
					if entityElement.LicenceTypeId == 4 {
						if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + entityReq.NumOfUsers) {
							return insertedEntityId, errors.New(ErrNoUserLicense)
						}
					}
					if entityElement.LicenceTypeId == 1 {
						if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + entityReq.NumOfUsers*10) {
							return insertedEntityId, errors.New(ErrNoUserLicense)
						}
					}
					if entityElement.LicenceTypeId == 5 {
						if entityElement.NumberOfLicences < (entityElement.LicenceConsumed + entityReq.NumOfUsers) {
							return insertedEntityId, errors.New(ErrNoUserLicense)
						}
					}
				}
			} else {
				return insertedEntityId, errors.New(ErrNoUserLicense)
			}

			stmt, err = db.MySqlDB.Prepare(insertEntityDetails)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			defer stmt.Close()

			entity, err := stmt.Exec(entityTypeId,
				entityReq.EntityId,
				parentEntityTypeId,
				entityReq.EntityName,
				entityReq.ContactName,
				entityReq.ContactEmail,
				entityReq.ContactMob,
				time.Now(),
				time.Now(),
				entityReq.LogoPath)

			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			entityId, err := entity.LastInsertId()
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			insertedEntityId = entityId

			if entityReq.LogoPath != "" {
				var fileNameDB *string = nil
				tempPath := "/static/logo/" + strconv.FormatInt(insertedEntityId, 10)
				fileNameDB = &tempPath
				logoEnabled := "yes"
				stmt, err = db.MySqlDB.Prepare(updateLogoPath)
				if err != nil {
					log.Println("ERROR: ", err)
					return insertedEntityId, err
				}
				defer stmt.Close()

				_, err = stmt.Exec(&fileNameDB, logoEnabled, entityId)
				if err != nil {
					log.Println("ERROR: ", err)
					return insertedEntityId, err
				}
			}

			licenseType := make(chan int)
			go func() {
				licenseType <- 1
				licenseType <- 4
				licenseType <- 5
				close(licenseType)
			}()

			stmt, err = db.MySqlDB.Prepare(insertEntityLicenseDetails)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			defer stmt.Close()
			for n := range licenseType {
				if n == 1 {
					_, err = stmt.Exec(entityId,
						n,
						entityReq.NumOfUsers*10,
						time.Now(),
						time.Now().AddDate(1, 0, 0))

					if err != nil {
						log.Println("ERROR: ", err)
						return insertedEntityId, err
					}
				}
				if n == 4 {
					_, err = stmt.Exec(entityId,
						n,
						entityReq.NumOfUsers,
						time.Now(),
						time.Now().AddDate(1, 0, 0))

					if err != nil {
						log.Println("ERROR: ", err)
						return insertedEntityId, err
					}
				}
				if n == 5 {
					_, err = stmt.Exec(entityId,
						n,
						entityReq.NumOfUsers,
						time.Now(),
						time.Now().AddDate(1, 0, 0))

					if err != nil {
						log.Println("ERROR: ", err)
						return insertedEntityId, err
					}
				}
			}

			var parentEntityId = ""
			err = db.MySqlDB.QueryRow(getParentEntityId, entityReq.EntityId).Scan(&parentEntityId)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			var licenseCount = 0
			for n := range licenseType {
				if n == 1 {
					licenseCount = int(entityReq.NumOfUsers) * 10
				}
				if n == 4 {
					licenseCount = int(entityReq.NumOfUsers)
				}
				if n == 5 {
					licenseCount = int(entityReq.NumOfUsers)
				}

				updateEntityLicenseCert := fmt.Sprintf(updateEntityLicense, licenseCount)

				func() {
					stmt, err := db.MySqlDB.Prepare(updateEntityLicenseCert)
					if err != nil {
						log.Println("ERROR: ", err)
						return
					}
					defer stmt.Close()

					_, err = stmt.Exec(n, parentEntityId)
					if err != nil {
						log.Println("ERROR: ", err)
						return
					}
				}()
			}
			return insertedEntityId, err
		} else { // If parent entity is not available
			var entityName = ""
			err := db.MySqlDB.QueryRow(getExistingEntityName, entityReq.EntityName).Scan(&entityName)
			if err != nil && err != sql.ErrNoRows {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			if entityName != "" {
				return insertedEntityId, errors.New(ErrEntityExists)
			}

			var entityTypeId = ""
			err = db.MySqlDB.QueryRow(getEntityTypeIdDirect, entityReq.EntityTypeName).Scan(&entityTypeId)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			parentEntityTypeId := 1

			stmt, err := db.MySqlDB.Prepare(insertEntityDetails)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			defer stmt.Close()

			entity, err := stmt.Exec(entityTypeId,
				0,
				parentEntityTypeId,
				entityReq.EntityName,
				entityReq.ContactName,
				entityReq.ContactEmail,
				entityReq.ContactMob,
				time.Now(),
				time.Now(),
				entityReq.LogoPath)

			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			entityId, err := entity.LastInsertId()
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			insertedEntityId = entityId
			stmt, err = db.MySqlDB.Prepare(updateParentEntityId)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			defer stmt.Close()

			_, err = stmt.Exec(entityId, entityId)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			if entityReq.LogoPath != "" {
				var fileNameDB *string = nil
				tempPath := "/static/logo/" + strconv.FormatInt(insertedEntityId, 10)
				fileNameDB = &tempPath
				logoEnabled := "yes"
				stmt, err = db.MySqlDB.Prepare(updateLogoPath)
				if err != nil {
					log.Println("ERROR: ", err)
					return insertedEntityId, err
				}
				defer stmt.Close()

				_, err = stmt.Exec(&fileNameDB, logoEnabled, entityId)
				if err != nil {
					log.Println("ERROR: ", err)
					return insertedEntityId, err
				}
			}

			licenseType := make(chan int)
			go func() {
				licenseType <- 1
				licenseType <- 4
				licenseType <- 5
				close(licenseType)
			}()

			stmt, err = db.MySqlDB.Prepare(insertEntityLicenseDetails)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			defer stmt.Close()
			for n := range licenseType {
				if n == 1 {
					_, err = stmt.Exec(entityId,
						n,
						entityReq.NumOfUsers*10,
						time.Now(),
						time.Now().AddDate(1, 0, 0))

					if err != nil {
						log.Println("ERROR: ", err)
						return insertedEntityId, err
					}
				}
				if n == 4 {
					_, err = stmt.Exec(entityId,
						n,
						entityReq.NumOfUsers,
						time.Now(),
						time.Now().AddDate(1, 0, 0))

					if err != nil {
						log.Println("ERROR: ", err)
						return insertedEntityId, err
					}
				}
				if n == 5 {
					_, err = stmt.Exec(entityId,
						n,
						entityReq.NumOfUsers,
						time.Now(),
						time.Now().AddDate(1, 0, 0))

					if err != nil {
						log.Println("ERROR: ", err)
						return insertedEntityId, err
					}
				}
			}
			return insertedEntityId, err
		}
	}

	// For lower entities
	if strings.ToLower(entityReq.EntityTypeName) == "department" ||
		strings.ToLower(entityReq.EntityTypeName) == "corporate_locality" ||
		strings.ToLower(entityReq.EntityTypeName) == "business_unit" ||
		strings.ToLower(entityReq.EntityTypeName) == "degree" ||
		strings.ToLower(entityReq.EntityTypeName) == "semester" ||
		strings.ToLower(entityReq.EntityTypeName) == "section" ||
		strings.ToLower(entityReq.EntityTypeName) == "grade" ||
		strings.ToLower(entityReq.EntityTypeName) == "corporate_city" ||
		strings.ToLower(entityReq.EntityTypeName) == "division" {

		// If parent entity is available
		if entityReq.EntityDivId != 0 {
			var entityDivName = ""
			err := db.MySqlDB.QueryRow(getEntityDivisionName, entityReq.EntityName, entityReq.EntityId, entityReq.EntityDivId).Scan(&entityDivName)
			if err != nil && err != sql.ErrNoRows {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			if entityDivName != "" {
				return insertedEntityId, errors.New(ErrEntityExists)
			}

			var entityTypeId = ""
			err = db.MySqlDB.QueryRow(getParentEntityTypeId, entityReq.EntityId).Scan(&entityTypeId)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			var entityDivTypeId = ""
			err = db.MySqlDB.QueryRow(getEntityDivisionTypeId, entityReq.EntityTypeName, entityTypeId).Scan(&entityDivTypeId)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			stmt, err := db.MySqlDB.Prepare(insertEntityDivisionDetails)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			defer stmt.Close()

			entityDiv, err := stmt.Exec(entityReq.EntityId,
				entityReq.EntityName,
				entityReq.EntityDivId,
				entityDivTypeId,
				entityTypeId,
				entityReq.EntityTypeName,
				entityReq.ContactName,
				entityReq.ContactEmail,
				entityReq.ContactMob)

			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			insertedEntityId, err = entityDiv.LastInsertId()
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			return insertedEntityId, err
		} else { // If parent entity is not available
			var entityDivName = ""
			err := db.MySqlDB.QueryRow(getParentEntityDivisionName, entityReq.EntityName, entityReq.EntityId).Scan(&entityDivName)
			if err != nil && err != sql.ErrNoRows {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			if entityDivName != "" {
				return insertedEntityId, errors.New(ErrEntityExists)
			}

			var entityTypeId = ""
			err = db.MySqlDB.QueryRow(getParentEntityTypeId, entityReq.EntityId).Scan(&entityTypeId)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			var entityDivTypeId = ""
			err = db.MySqlDB.QueryRow(getEntityDivisionTypeId, entityReq.EntityTypeName, entityTypeId).Scan(&entityDivTypeId)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			stmt, err := db.MySqlDB.Prepare(insertEntityDivisionDetails)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			defer stmt.Close()

			entity, err := stmt.Exec(entityReq.EntityId,
				entityReq.EntityName,
				0,
				entityDivTypeId,
				entityTypeId,
				entityReq.EntityTypeName,
				entityReq.ContactName,
				entityReq.ContactEmail,
				entityReq.ContactMob)

			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			entityDivId, err := entity.LastInsertId()
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}

			// Update entityDivId of the inserted record as parentEntityDivId
			stmt, err = db.MySqlDB.Prepare(updateParentEntityDivId)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			defer stmt.Close()

			_, err = stmt.Exec(entityDivId, entityDivId)
			if err != nil {
				log.Println("ERROR: ", err)
				return insertedEntityId, err
			}
			return insertedEntityId, err
		}
	}
	return insertedEntityId, errors.New("")
}

func (entityReq *DeleteEntityReq) DeleteLowerEntity() error {

	stmt, err := db.MySqlDB.Prepare(deleteLowerEntity)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(
		entityReq.EntityDivId,
		entityReq.EntityDivId,
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
	return err
}

func (entityReq *DeleteEntityReq) DeleteHigherEntity() error {

	stmt, err := db.MySqlDB.Prepare(deleteHigherEntity)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(
		entityReq.EntityId,
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

	stmt, err = db.MySqlDB.Prepare(deleteLowerEntityWithParent)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	row, err = stmt.Exec(
		entityReq.EntityId)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	_, err = row.RowsAffected()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	return err
}

func (entityReq *DeleteEntityReq) DeleteEntityUser() error {

	deleteUsers := deleteEntityUsers
	deleteUsers = fmt.Sprintf(deleteUsers, "Level_"+strconv.Itoa(int(entityReq.EntityLevel)))
	stmt, err := db.MySqlDB.Prepare(deleteUsers)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	// If its a lower entity then consider entity div id
	if _, ok := utils.LowerLevelTypes[entityReq.EntityLevel]; ok {
		entityReq.EntityId = entityReq.EntityDivId
	}

	row, err := stmt.Exec(
		entityReq.EntityId)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	_, err = row.RowsAffected()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	return err
}

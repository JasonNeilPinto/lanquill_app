package user

var updateUserPassword = `UPDATE 
        sententia.user
    SET
        Password = ?, User_Status='Active', Password_Attempts=0
    WHERE
        User_ID = ?`

var countEmailExceptId = `SELECT 
        COUNT(User_ID) FROM sententia.user
    WHERE
        User_Email = ? AND User_ID != ? AND User_Status != 'Deleted'`

var updateUserInfo = `UPDATE 
        sententia.user
    SET
        User_Name = ?, User_Email = ?
    WHERE
        User_ID = ?`

var deleteUser = `UPDATE 
        sententia.user
    SET
        User_Status = "Deleted", User_Email = CONCAT(User_Email,"_Deleted"),
        Last_Modified = CURRENT_TIMESTAMP()
    WHERE
        User_ID = ?`

var getUserEmail = `SELECT User_Email FROM sententia.user WHERE (User_Email = ? OR (User_Mobile = ? AND User_Mobile != '')) AND User_Status != 'Deleted'`

var getUserId = `SELECT User_ID FROM sententia.user WHERE User_Email = ? AND User_Status != 'Deleted'`

var getUserTypeId = `SELECT User_Type FROM sententia.user WHERE User_ID = ?`

var checkIfUserExistsForHigherLevel = `SELECT User_Email FROM sententia.user WHERE (Level_6 = ? or Level_5 = ?) AND User_Email = ? AND User_Status != 'Deleted'`

var userLowerEntityLevelDetails = `SELECT ed.Entity_ID, ed.Parent_Entity_Div_ID, ed.Entity_Div_Type_ID, edt.Entity_Div_Type_Level,
ea.Parent_Entity_ID, eb.Entity_Type_ID, et.Entity_Type_Level
FROM sententia.entity_division ed
JOIN sententia.entity_division_type edt ON ed.Entity_Div_Type_ID = edt.Entity_Div_Type_ID
LEFT JOIN sententia.entity ea ON ea.Entity_ID = ed.Entity_ID
LEFT JOIN sententia.entity eb ON eb.Entity_ID = ea.Parent_Entity_ID
JOIN sententia.entity_type et ON et.Entity_Type_ID = eb.Entity_Type_ID
WHERE ed.Entity_Div_ID = ?`

var updateUserDetails = `UPDATE sententia.user SET %s WHERE (Level_6 = ? OR Level_5 = ?) AND User_Email = ?`

var getParentEntityDivId = `SELECT Parent_Entity_Div_ID FROM sententia.entity_division WHERE Entity_Div_ID = ?`

var getParentEntityId = `SELECT Parent_Entity_ID FROM sententia.entity WHERE Entity_ID = ?`

var insertUserDetails = `INSERT INTO sententia.user (User_Name, User_Email, User_Mobile, Password, Date_Created, Renewal_Date, User_Status, User_Type) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

var updateNewUserDetails = `UPDATE sententia.user SET %s WHERE User_Email = ?`

var getEntityLicenseDetails = `SELECT Number_of_Licences, Licence_Type_ID, Licence_Consumed FROM sententia.entity_licence WHERE Entity_ID = ?`

var updateEntityLicense = `UPDATE sententia.entity_licence SET Licence_Consumed = Licence_Consumed + %d WHERE Entity_ID=? AND Licence_Type_ID = ?`

var insertUserLicense = `INSERT INTO sententia.user_licence (User_ID, Licence_Type_ID, Number_of_Licences, Date_Purchased, Valid_Till) VALUES (?, ?, ?, ?, ?)`

var userHigherEntityLevelDetails = `SELECT e.Entity_Type_ID, et.Entity_Type_Level, e.Parent_Entity_ID, 
pe.Entity_Type_ID AS Parent_Entity_Type_ID, pet.Entity_Type_Level AS Parent_Entity_Type_Level
FROM sententia.entity e
JOIN sententia.entity_type et ON et.Entity_Type_ID = e.Entity_Type_ID
LEFT JOIN sententia.entity pe ON pe.Entity_ID = e.Parent_Entity_ID
JOIN sententia.entity_type pet ON pet.Entity_Type_ID = pe.Entity_Type_ID
WHERE e.Entity_ID = ?`

var getLicenseDetailsByLicenseType = `SELECT Number_of_Licences, Licence_Type_ID, Licence_Consumed FROM sententia.entity_licence WHERE Entity_ID = ? AND Licence_Type_ID = 4`

var changeSingleUserExpiryDate = `UPDATE sententia.user SET Renewal_date = ? WHERE User_ID = ?`

var changeHigherEntityExpiryDate = `UPDATE sententia.user SET Renewal_date = ? WHERE Level_6 = ?`

var changeLowerEntityExpiryDate = `UPDATE sententia.user SET Renewal_date = ? WHERE Level_5 = ?`

var updateEntityLicenceValidity = `UPDATE sententia.entity_licence SET Valid_Till = ? WHERE Entity_ID = ?`

var getUserDetail = `SELECT User_ID,User_Email,User_Name,Password_Attempts,User_Status, User_Type FROM sententia.user WHERE User_Email=?`

var getEntityDetailsByEmail = `SELECT e.Entity_Name, e.Entity_ID, et.Entity_Type_Name, CONCAT(e.Entity_ID, "//", et.Entity_Type_Name) AS Entity_Info FROM sententia.entity e
JOIN sententia.entity_type et ON e.Entity_Type_ID = et.Entity_Type_ID
WHERE e.Parent_Entity_ID=
(SELECT CASE WHEN Level_6 IS NOT NULL THEN Level_6 ELSE Level_5 END
FROM sententia.user 
WHERE Level_4 IS NULL AND Level_3 IS NULL AND Level_2 IS NULL and Level_1 IS NULL AND User_Email=?)`

var getEntityDetailsByEntityId = `SELECT e.Entity_Name, e.Entity_ID, et.Entity_Type_Name, CONCAT(e.Entity_ID, "//", et.Entity_Type_Name) AS Entity_Info FROM sententia.entity e
JOIN sententia.entity_type et ON e.Entity_Type_ID = et.Entity_Type_ID
WHERE ? IS NULL OR e.Entity_ID = ?`

var getUserDetailsByEmail = `SELECT User_ID, User_Name, Renewal_date FROM sententia.user WHERE User_Email = ?`

var getHigherEntityDetails = `SELECT Renewal_Date, COUNT(*) FROM sententia.user WHERE Level_6 = ? OR Level_5 = ? GROUP BY Renewal_Date LIMIT 1`

var getLowerEntityDetails = `SELECT Renewal_Date, COUNT(*) FROM sententia.user WHERE Level_4 = ? OR Level_3 = ? GROUP BY Renewal_Date LIMIT 1`

var getEntityTypeName = `SELECT et.Entity_Type_Name AS 'Entity_Type',
et.Entity_Type_Level AS 'Entity_Level'
FROM sententia.entity e
JOIN sententia.entity_type et ON e.Entity_Type_ID = et.Entity_Type_ID
WHERE e.Entity_ID IN (
	SELECT CASE 
			WHEN Level_1 IS NOT NULL
				THEN Level_1
			WHEN Level_2 IS NOT NULL
				THEN Level_2
			WHEN Level_3 IS NOT NULL
				THEN Level_3
			WHEN Level_4 IS NOT NULL
				THEN Level_4
			WHEN Level_5 IS NOT NULL
				THEN Level_5
			WHEN Level_6 IS NOT NULL
				THEN Level_6
			WHEN Level_7 IS NOT NULL
				THEN Level_7
			END AS LEVEL
	FROM sententia.user
	WHERE User_ID = ?
	)

UNION

SELECT edt.Entity_Div_Type_Name AS 'Entity_Type',
edt.Entity_Div_Type_Level AS 'Entity_Level'
FROM sententia.entity_division ed
JOIN sententia.entity_division_type edt ON ed.Entity_Div_Type_ID = edt.Entity_Div_Type_ID
WHERE ed.Entity_Div_ID IN (
	SELECT CASE 
			WHEN Level_1 IS NOT NULL
				THEN Level_1
			WHEN Level_2 IS NOT NULL
				THEN Level_2
			WHEN Level_3 IS NOT NULL
				THEN Level_3
			WHEN Level_4 IS NOT NULL
				THEN Level_4
			WHEN Level_5 IS NOT NULL
				THEN Level_5
			WHEN Level_6 IS NOT NULL
				THEN Level_6
			WHEN Level_7 IS NOT NULL
				THEN Level_7
			END AS LEVEL
	FROM sententia.user
	WHERE User_ID = ?
	)
`

var getUserTypeForAdmin = `SELECT User_Type_ID FROM sententia.user_type WHERE User_Type = ? AND User_Type_Level = ?`

var updateUserType = `UPDATE sententia.user SET User_Type = ? WHERE User_ID = ?`

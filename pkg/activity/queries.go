package activity

var SELECT_USER = `SELECT 
					u.User_ID, 
					u.User_Name, 
					u.User_Email,
					u.Level_7,
					u.Level_6,
					u.Level_5,
					u.Level_4,
					u.Level_3,
					u.Level_2,
					u.Level_1, 
					u.Last_Login, 
					u.User_Mobile, 
					u.Renewal_Date,
					u.User_Type, 
					ut.User_Type
				FROM
					sententia.user u INNER JOIN sententia.user_type ut ON u.User_Type = ut.User_Type_ID
				WHERE 
					(Level_6= ? or Level_5= ? or Level_4 =? or Level_3 =? or Level_2 =? or Level_1= ?) 
				AND  
					User_Status = 'Active'`

var userHigherEntityInfo = `SELECT
							e.Entity_Name,
							e.Entity_ID
							FROM sententia.entity e
							JOIN sententia.entity_type et ON e.Entity_Type_ID = et.Entity_Type_ID
							WHERE e.Entity_ID = ?`

var userLowerEntityInfo = `SELECT 
							ed.Entity_Div_Name AS 'Entity_Name',
							ed.Entity_Div_ID AS 'Entity_ID'
							FROM sententia.entity_division ed
							JOIN sententia.entity_division_type edt ON ed.Entity_Div_Type_ID = edt.Entity_Div_Type_ID
							WHERE ed.Entity_Div_ID = ?
							`

var SELECT_USER_DETAILS = `SELECT 
							User_Name, User_Email, Level_6, Level_5 
						FROM 
							sententia.user
						WHERE 
							User_ID = ?`

var SELECT_SSOUSER_DETAILS = `SELECT  
								sso_User_Name, sso_User_Email, sso_Institute_Name 
							FROM 
								sso_user 
							WHERE 
								sso_User_ID= ?`

var SELECT_LOGGED_IN_USERS = `SELECT 
								User_Name, User_Email, Last_Login, Level_6, Level_5 
							FROM 
								sententia.user 
							ORDER BY Last_Login DESC LIMIT 100`

var PAYMENT_DETAILS = `SELECT 
						email, order_id, txn_start_datetime, title, amount, status 
					FROM 
						payment_details 
					ORDER BY 
						txn_end_datetime DESC`

var RETAIL_USERS = `SELECT  
						User_Name, 
						User_Email, 
						Date_Created, 
						Last_Login 
					FROM sententia.user 
					WHERE User_Type = 1 
					AND Level_7 is NULL 
					AND Level_6 is NULL 
					AND Level_5 is NULL 
					AND Level_4 is NULL 
					AND Level_3 is NULL 
					AND Level_2 is NULL
					AND Level_1 is NULL
					AND User_Status = 'Active'
					ORDER BY Date_Created DESC LIMIT 100;`

var TXN_DETAILS = `SELECT
						email,
						order_id,
						txn_start_datetime,
						txn_end_datetime,
						title,
						amount,
						status,
						description,
						callback_response,
						status_response,
						order_det,
						gst,
						invoice_no,
						country_code
					FROM 
						payment_details WHERE order_id= ? `

var getIPAddress = `
				SELECT eip.Eip_ID, e.Entity_ID, e.Entity_Name, eip.IP_Address, e.IP_Address, e.Enable_IP_Login, eip.IP_Range_Enable
				FROM sententia.Entity_Ip_Address eip JOIN sententia.entity e 
				ON eip.Entity_ID = e.Entity_ID`

var updateEntityIpAddress = `
UPDATE sententia.Entity_Ip_Address 
SET IP_Address = ?, IP_Range_Enable = ?
WHERE Eip_ID = ?
`
var deleteIpAddress = `
DELETE FROM sententia.Entity_Ip_Address
WHERE Eip_ID = ?
`

// var updateIpAddressInEntity = `
// UPDATE sententia.entity e
// SET IP_Address = (
//         SELECT GROUP_CONCAT(IP_Address SEPARATOR ', ') AS IP_Address
//         FROM sententia.Entity_Ip_Address eip
//         WHERE eip.Entity_ID = e.Entity_ID
//         GROUP BY Entity_ID
//     ),
//     Enable_IP_Login = CASE
//                          WHEN (
//                                  SELECT GROUP_CONCAT(IP_Address SEPARATOR ', ') AS IP_Address
//                                  FROM sententia.Entity_Ip_Address eip
//                                  WHERE eip.Entity_ID = e.Entity_ID
//                                  GROUP BY Entity_ID
//                              ) IS NULL THEN 'No'
//                          ELSE 'True'
//                      END
// WHERE e.Entity_ID = ?
// `

var updateIpAddressInEntity = `
UPDATE sententia.entity e
SET IP_Address = (
        SELECT GROUP_CONCAT(IP_Address SEPARATOR '? ') AS IP_Address
        FROM sententia.Entity_Ip_Address eip
        WHERE eip.Entity_ID = e.Entity_ID
        GROUP BY Entity_ID
    ),
    Enable_IP_Login = ?
WHERE e.Entity_ID = ?
`

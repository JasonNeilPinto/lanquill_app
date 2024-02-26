package reports

var SELECT_ROOT_ENTITIES = `SELECT
								a.Entity_ID,
								a.Entity_Name,
								a.Entity_Type_ID,
								b.Entity_Type_Name,
								b.Entity_Type_Level,
								b.Level_Verbose,
								b.entity_hierarchy_id,
								COUNT(c.User_ID) as Users
							FROM
								sententia.entity a
								LEFT JOIN sententia.entity_type b ON a.Entity_Type_ID = b.Entity_Type_ID
								AND a.Parent_Entity_Type_ID = b.Parent_Entity_Type_ID
								LEFT JOIN user as c ON (
									(c.Level_7 = a.Entity_ID
									OR c.Level_6 = a.Entity_ID
									OR c.Level_5 = a.Entity_ID)
									AND c.User_Status != 'Deleted'
								)
							WHERE
								a.Parent_Entity_Type_ID = 1 AND a.Entity_Status != 'Deleted'
							GROUP BY
								a.Entity_ID
							ORDER BY
								a.Entity_ID`

var SELECT_HIGHER_ENTITIES = `SELECT
								a.Entity_ID,
								null AS Entity_Div_ID,
								a.Entity_Name,
								a.Entity_Type_ID,
								b.Entity_Type_Name,
								b.Entity_Type_Level,
								b.Level_Verbose,
								b.entity_hierarchy_id,
								COUNT(c.User_ID) as Users
							FROM
								sententia.entity a
								LEFT JOIN sententia.entity_type b ON a.Entity_Type_ID = b.Entity_Type_ID
								AND a.Parent_Entity_Type_ID = b.Parent_Entity_Type_ID
								LEFT JOIN user as c ON (
									(c.Level_7 = a.Entity_ID
									OR c.Level_6 = a.Entity_ID
									OR c.Level_5 = a.Entity_ID)
									AND c.User_Status != 'Deleted'
								)
							WHERE
								Parent_Entity_ID = ?
								AND Entity_ID != Parent_Entity_ID
								AND a.Entity_Status != 'Deleted'
							GROUP BY
								a.Entity_ID
							ORDER BY
								a.Entity_ID`

var SELECT_LOWER_ENTITIES = `SELECT
								a.Entity_ID,
								a.Entity_Div_ID,
								a.Entity_Div_Name,
								a.Entity_Div_Type_ID,
								b.Entity_Div_Type_Name,
								b.Entity_Div_Type_Level,
								b.Level_Verbose,
								b.entity_hierarchy_id,
								COUNT(c.User_ID) as Users
								FROM
								sententia.entity_division a
									LEFT JOIN sententia.entity_division_type b ON a.Entity_Div_Type_ID = b.Entity_Div_Type_ID
									AND a.Entity_Type_ID = b.Entity_Type_ID
									LEFT JOIN user as c ON (
										(c.Level_4 = a.Entity_Div_ID
										OR c.Level_3 = a.Entity_Div_ID
										OR c.Level_2 = a.Entity_Div_ID)
										AND c.User_Status != 'Deleted'
									)
								WHERE
									Entity_ID = ?
									AND Entity_Div_ID=Parent_Entity_Div_ID
									AND a.Entity_Status != 'Deleted'
									GROUP BY
									a.Entity_Div_ID
									ORDER BY
									a.Entity_Div_ID`

var SELECT_LOWER_DIV_ENTITIES = `SELECT
									a.Entity_ID,
									a.Entity_Div_ID,
									a.Entity_Div_Name,
									a.Entity_Div_Type_ID,
									b.Entity_Div_Type_Name,
									b.Entity_Div_Type_Level,
									b.Level_Verbose,
									b.entity_hierarchy_id,
									COUNT(c.User_ID) as Users
								FROM
									sententia.entity_division a
									LEFT JOIN sententia.entity_division_type b ON a.Entity_Div_Type_ID = b.Entity_Div_Type_ID
									AND a.Entity_Type_ID = b.Entity_Type_ID
									LEFT JOIN user as c ON (
										(c.Level_4 = a.Entity_Div_ID
										OR c.Level_3 = a.Entity_Div_ID
										OR c.Level_2 = a.Entity_Div_ID)
										AND c.User_Status != 'Deleted'
									)
								WHERE
									Parent_Entity_Div_ID = ?
									AND Entity_Div_ID != Parent_Entity_Div_ID
									AND a.Entity_Status != 'Deleted'
								GROUP BY
									a.Entity_Div_ID
								ORDER BY
									a.Entity_Div_ID`

var SELECT_ROOT_USERS = `SELECT
								User_ID,
								User_Name,
								User_Email,
								Last_Login
							FROM
								sententia.user
							WHERE
								Level_7 IS NULL
								AND Level_6 IS NULL
								AND Level_5 IS NULL
								AND Level_4 IS NULL
								AND Level_3 IS NULL
								AND Level_2 IS NULL
								AND Level_1 IS NULL
								AND User_Status != 'Deleted'`

var overallUserCount = `SELECT COUNT(User_ID) AS Total_Users,
COUNT(Last_Login) AS Active_Users
FROM sententia.user
WHERE User_Status != 'Deleted'`

var entityUserCount = `SELECT COUNT(User_ID) AS Total_Users,
COUNT(Last_Login) AS Active_Users
FROM sententia.user
WHERE %s = ? AND User_Status != 'Deleted'`

var getEntityReport = `SELECT LAvgScore,
SAvgScore,
RAvgScore,
WAvgScore,
GAvgScore,
VAvgScore,
LTotalCount,
STotalCount,
RTotalCount,
WTotalCount,
GTotalCount,
VTotalCount,
L1Attempted,
L2Attempted,
L3Attempted,
L4Attempted,
L5Attempted,
L6Attempted,
L7Attempted,
L1Completed,
L2Completed,
L3Completed,
L4Completed,
L5Completed,
L6Completed,
L7Completed,
SuggestedLevel1,
SuggestedLevel2,
SuggestedLevel3,
SuggestedLevel4,
SuggestedLevel5,
SuggestedLevel6,
SuggestedLevel7,
SugLevel1Diff,
SugLevel2Diff,
SugLevel3Diff,
SugLevel4Diff,
SugLevel5Diff,
SugLevel6Diff,
SugLevel7Diff,
AoI
FROM Reports.Entity_Report
WHERE Entity_ID = ?
AND Entity_Level = ?`

var getUserReport = `SELECT LAvgScore,
SAvgScore,
RAvgScore,
WAvgScore,
GAvgScore,
VAvgScore,
LTotalCount,
STotalCount,
RTotalCount,
WTotalCount,
GTotalCount,
VTotalCount,
L1Attempted,
L2Attempted,
L3Attempted,
L4Attempted,
L5Attempted,
L6Attempted,
L7Attempted,
L1Completed,
L2Completed,
L3Completed,
L4Completed,
L5Completed,
L6Completed,
L7Completed,
SuggestedLevel,
AoI
FROM Reports.User_Report
WHERE User_ID = ?`

var userBaseInfo = `SELECT b.Last_Login,
a.SuggestedLevel,
a.WTotalCount
FROM Reports.User_Report a
LEFT JOIN sententia.user b ON a.User_ID = b.User_ID
WHERE a.User_ID = ?`

var userProfileInfo = `SELECT User_ID, User_Name, User_Email, Last_Login FROM sententia.user WHERE User_ID = ?`

var getUserHigherEntityList = `SELECT u.User_ID,
							u.User_Name,
							u.User_Email,
							u.Last_Login,
							u.User_Type,
							u.User_Status,
							e.Entity_Name,
							u.Date_Created,
							u.Renewal_Date,
							CASE 
								WHEN u.Referral_Id != ''
									THEN 'Yes'
								ELSE 'No'
								END AS 'IsSSOUser'
							FROM sententia.user u
							JOIN sententia.entity e ON e.Entity_ID = u.Level_%s
							WHERE u.Level_%s = ? AND u.User_Status != 'Deleted'`

var getUserLowerEntityList = `SELECT ud.User_ID,
								ud.User_Name,
								ud.User_Email,
								ud.Last_Login,
								ud.User_Type,
								ud.User_Status,
								ed.Entity_Div_Name AS 'Entity_Name',
								ud.Date_Created,
								ud.Renewal_Date,
								CASE 
									WHEN ud.Referral_Id != ''
										THEN 'Yes'
									ELSE 'No'
									END AS 'IsSSOUser'
								FROM sententia.user ud
								JOIN sententia.entity_division ed ON ed.Entity_Div_ID = ud.Level_%s
								WHERE ud.Level_%s = ? AND ud.User_Status != 'Deleted'`

var getSearchUser = `SELECT U.User_Id,
						U.User_Name,
						U.User_Email,
						U.Last_Login,
						U.User_Type,
						U.User_Status,
						e.Entity_Name,
						et.Entity_Name AS Lower_Entity_Name,
						U.Date_Created,
						U.Renewal_Date,
						CASE 
							WHEN U.Referral_Id != ''
								THEN 'Yes'
							ELSE 'No'
							END AS 'IsSSOUser'
						FROM sententia.user U
						LEFT JOIN sententia.entity e ON e.Entity_ID = U.Level_6
						LEFT JOIN sententia.entity et ON et.Entity_ID = U.Level_5
						WHERE U.User_Email = ?
`

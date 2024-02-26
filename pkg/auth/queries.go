package auth

var SELECT_HIGHER_ENTITY_LEVEL = `SELECT et.Entity_Type_Level
FROM sententia.entity e
JOIN sententia.entity_type et ON e.Entity_Type_ID = et.Entity_Type_ID
	AND e.Parent_Entity_Type_ID = et.Parent_Entity_Type_ID
WHERE e.Entity_ID = ?`

var SELECT_LOWER_ENTITY_LEVEL = `SELECT et.Entity_Div_Type_Level
FROM sententia.entity_division e
JOIN sententia.entity_division_type et ON e.Entity_Div_Type_ID = et.Entity_Div_Type_ID
	AND e.Entity_Type_ID = et.Entity_Type_ID
WHERE e.Entity_Div_ID = ?
`

var SELECT_USER_TYPE_LEVEL = `SELECT User_Type_Level FROM sententia.user_type WHERE User_Type_ID = ?`

var SELECT_USER_LEVEL = `SELECT ut.User_Type_Level FROM sententia.user u JOIN sententia.user_type ut ON u.User_Type = ut.User_Type_ID
WHERE u.User_ID = ?`

var SELECT_USER_ENTITY_DETAILS = `SELECT User_Type, Level_6, Level_5, Level_4, Level_3, Level_2, Level_1 FROM sententia.user WHERE User_ID = ?`

var SELECT_USER_DETAILS = `SELECT COUNT(User_ID) FROM sententia.user WHERE User_Email=? %s`

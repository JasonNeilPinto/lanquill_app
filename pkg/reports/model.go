package reports

import (
	"database/sql"
)

type EntityUserReq struct {
	EntityId    *int64 `json:"entityId"`
	EntityDivId int64  `json:"entityDivId"`
	EntityLevel *int   `json:"entityLevel"`
}

type EntityReportReq struct {
	EntityId     *int64  `json:"entityId"`
	EntityDivId  *int64  `json:"entityDivId"`
	LevelVerbose *string `json:"levelVerbose"`
	EntityLevel  *int    `json:"entityLevel"`
}

type EntityResp struct {
	EntityId        int64  `json:"entityId"`
	EntityDivId     *int64 `json:"entityDivId"`
	EntityName      string `json:"entityName"`
	EntityTypeId    int    `json:"entityTypeId"`
	EntityTypeName  string `json:"entityTypeName"`
	EntityLevel     int    `json:"entityLevel"`
	LevelVerbose    string `json:"levelVerbose"`
	EntityHierarchy int    `json:"entityHierarchy"`
	Users           int64  `json:"users"`
}

type UserReportReq struct {
	UserId *int64 `json:"userId"`
}

type UsersDB struct {
	Id       int64          `db:"User_ID"`
	Name     string         `db:"User_Name"`
	Email    string         `db:"User_Email"`
	Login    sql.NullString `db:"Last_Login"`
	UserType int            `db:"User_Type"`
}

type UsersResp struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Login   string `json:"login"`
	IsAdmin bool   `json:"isAdmin"`
}

type UserCountResp struct {
	TotalUsers  int64 `json:"totalUsers"`
	ActiveUsers int64 `json:"activeUsers"`
}

type KylLevelInfo struct {
	Count int64  `json:"count"`
	Score int64  `json:"score"`
	Type  string `json:"type"`
}

type KylData struct {
	Level1 KylLevelInfo `json:"1"`
	Level2 KylLevelInfo `json:"2"`
	Level3 KylLevelInfo `json:"3"`
	Level4 KylLevelInfo `json:"4"`
	Level5 KylLevelInfo `json:"5"`
	Level6 KylLevelInfo `json:"6"`
	Level7 KylLevelInfo `json:"7"`
}

type Skills struct {
	Listening  TestResults `json:"listening"`
	Speaking   TestResults `json:"speaking"`
	Reading    TestResults `json:"reading"`
	Writing    TestResults `json:"writing"`
	Grammar    TestResults `json:"grammar"`
	Vocabulary TestResults `json:"vocabulary"`
}

type TestLevels struct {
	Level1 TestStatus `json:"1"`
	Level2 TestStatus `json:"2"`
	Level3 TestStatus `json:"3"`
	Level4 TestStatus `json:"4"`
	Level5 TestStatus `json:"5"`
	Level6 TestStatus `json:"6"`
	Level7 TestStatus `json:"7"`
}

type TestResults struct {
	Average int64 `json:"average"`
	Tests   int64 `json:"tests"`
}

type TestStatus struct {
	Attempted int64 `json:"attempted"`
	Completed int64 `json:"completed"`
}

type TestResp struct {
	AverageScore   Skills     `json:"avgScore"`
	CertResults    TestLevels `json:"certResults"`
	KylScore       *KylData   `json:"kylScore,omitempty"`
	Aoi            []string   `json:"AoI"`
	SuggestedLevel int64      `json:"suggestedLevel,omitempty"`
}

type EntityReportResp struct {
	TotalUsers    int64     `json:"totalUsers"`
	ActiveUsers   int64     `json:"activeUsers"`
	CertTaken     int64     `json:"certTaken"`
	GeneralReport *TestResp `json:"reportData"`
}

type UserReportResp struct {
	LastLogin           string    `json:"lastLogin"`
	SelfAssessmentLevel int64     `json:"selfAssessmentLevel"`
	Docs                int64     `json:"docs"`
	GeneralReport       *TestResp `json:"reportData"`
}

type GeneralReportResp struct {
	LAvgScore       sql.NullInt64
	SAvgScore       sql.NullInt64
	RAvgScore       sql.NullInt64
	WAvgScore       sql.NullInt64
	GAvgScore       sql.NullInt64
	VAvgScore       sql.NullInt64
	LTotalCount     sql.NullInt64
	STotalCount     sql.NullInt64
	RTotalCount     sql.NullInt64
	WTotalCount     sql.NullInt64
	GTotalCount     sql.NullInt64
	VTotalCount     sql.NullInt64
	L1Attempted     sql.NullInt64
	L2Attempted     sql.NullInt64
	L3Attempted     sql.NullInt64
	L4Attempted     sql.NullInt64
	L5Attempted     sql.NullInt64
	L6Attempted     sql.NullInt64
	L7Attempted     sql.NullInt64
	L1Completed     sql.NullInt64
	L2Completed     sql.NullInt64
	L3Completed     sql.NullInt64
	L4Completed     sql.NullInt64
	L5Completed     sql.NullInt64
	L6Completed     sql.NullInt64
	L7Completed     sql.NullInt64
	SuggestedLevel  sql.NullInt64
	SuggestedLevel1 sql.NullInt64
	SuggestedLevel2 sql.NullInt64
	SuggestedLevel3 sql.NullInt64
	SuggestedLevel4 sql.NullInt64
	SuggestedLevel5 sql.NullInt64
	SuggestedLevel6 sql.NullInt64
	SuggestedLevel7 sql.NullInt64
	SugLevel1Diff   sql.NullInt64
	SugLevel2Diff   sql.NullInt64
	SugLevel3Diff   sql.NullInt64
	SugLevel4Diff   sql.NullInt64
	SugLevel5Diff   sql.NullInt64
	SugLevel6Diff   sql.NullInt64
	SugLevel7Diff   sql.NullInt64
	Aoi             sql.NullString
}

type UsersEntityDB struct {
	Id          int64          `db:"User_ID"`
	Name        string         `db:"User_Name"`
	Email       string         `db:"User_Email"`
	Login       sql.NullString `db:"Last_Login"`
	UserType    int            `db:"User_Type"`
	UserStatus  string         `db:"User_Status"`
	EntityName  string         `db:"Entity_Name"`
	DateCreated string         `db:"Date_Created"`
	RenewalDate string         `db:"Renewal_Date"`
	IsSSOUser   string         `db:"IsSSOUser"`
}

type UsersEntityResp struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Login       string `json:"login"`
	IsAdmin     string `json:"isAdmin"`
	UserStatus  string `json:"userStatus"`
	EntityName  string `json:"entityName"`
	DateCreated string `json:"dateCreated"`
	RenewalDate string `json:"renewalDate"`
	IsSSOUser   string `json:"isSSOUser"`
}

type SearchUserResp struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Login       string `json:"login"`
	IsAdmin     bool   `json:"isAdmin"`
	UserStatus  string `json:"userStatus"`
	EntityName  string `json:"entityName"`
	DateCreated string `json:"dateCreated"`
	RenewalDate string `json:"renewalDate"`
	IsSSOUser   string `json:"isSSOUser"`
}

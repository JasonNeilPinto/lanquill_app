package auth

import "github.com/dgrijalva/jwt-go"

type JWTData struct {
	UserEmail string `json:"User_Email"`
	UserID    int    `json:"User_ID"`
	UserName  string `json:"User_Name"`
	UserType  int    `json:"User_Type"`
	Level7    *int64 `json:"Level_7"`
	Level6    *int64 `json:"Level_6"`
	Level5    *int64 `json:"Level_5"`
	Level4    *int64 `json:"Level_4"`
	Level3    *int64 `json:"Level_3"`
	Level2    *int64 `json:"Level_2"`
	Level1    *int64 `json:"Level_1"`
	SSOUser   bool   `json:"SSO_User"`
	jwt.StandardClaims
}

type EntityDetail struct {
	EntityId    *int64 `json:"entityId"`
	EntityLevel *int   `json:"entityLevel"`
}

type Entity struct {
	EntityId     *int64  `json:"entityId"`
	EntityDivId  *int64  `json:"entityDivId"`
	EntityLevel  *int    `json:"entityLevel"`
	LevelVerbose *string `json:"levelVerbose"`
}

type User struct {
	UserId int64 `json:"userId"`
}

type ResetReq struct {
	UserId    int64  `json:"userId"`
	Password  string `json:"password"`
	UserEmail string `json:"userEmail"`
}

type UserEntityResp struct {
	UserType int
	Level6   *int64
	Level5   *int64
	Level4   *int64
	Level3   *int64
	Level2   *int64
	Level1   *int64
}

type ChangeExpiryDateReq struct {
	UserEmail   string `json:"userEmail"`
	RequestType string `json:"requestType"`
	EntityId    *int64 `json:"entityId"`
}

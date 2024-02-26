package user

import "time"

type UserIdReq struct {
	UserId int64 `json:"userId"`
}

type ResetPassReq struct {
	UserId   int64  `json:"userId"`
	Password string `json:"password"`
}

type UpdateUserReq struct {
	UserId  int64  `json:"userId"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"isAdmin"`
}

type CreateUserReq struct {
	UserName      string `json:"userName"`
	Password      string `json:"password"`
	UserMobile    string `json:"userMob"`
	UserEmail     string `json:"userEmail"`
	EntityDivId   int64  `json:"entityDivId"`
	EntityId      int64  `json:"entityId"`
	DocQuantity   int64  `json:"docQuantity"`
	CertQuantity  int64  `json:"certQuantity"`
	DocValidTill  string `json:"docValidTill"`
	CertValidTill string `json:"certValidTill"`
}

type UserEntityResp struct {
	EntityId           int64 `json:"entityId"`
	ParentEntityDivId  int64 `json:"parentEntityDivId"`
	EntityDivTypeId    int64 `json:"entityDivTypeId"`
	EntityDivTypeLevel int64 `json:"entityDivTypeLevel"`
	ParentEntityId     int64 `json:"parentEntityId"`
	EntityTypeId       int64 `json:"entityTypeId"`
	EntityTypeLevel    int64 `json:"entityTypeLevel"`
}

type EntityLicenseResp struct {
	NumberOfLicences int64 `json:"numberOfLicences"`
	LicenceTypeId    int64 `json:"licenceTypeId"`
	LicenceConsumed  int64 `json:"licenceConsumed"`
}

type UserHigherEntityResp struct {
	EntityTypeId          int64 `json:"entityTypeId"`
	EntityTypeLevel       int64 `json:"EntityTypeLevel"`
	ParentEntityId        int64 `json:"ParentEntityId"`
	ParentEntityTypeId    int64 `json:"parentEntityTypeId"`
	ParentEntityTypeLevel int64 `json:"parentEntityTypeLevel"`
}

type ChangeExpiryDateReq struct {
	RequestType string `json:"requestType"`
	ExpiryDate  string `json:"expiryDate"`
	UserId      *int64 `json:"userId"`
	EntityId    *int64 `json:"entityId"`
	EntityType  string `json:"entityType"`
}

type UserDetailReq struct {
	UserEmail string `json:"userEmail"`
}

type UserDetailsResp struct {
	UserId           int64  `json:"userId"`
	UserEmail        string `json:"userEmail"`
	UserName         string `json:"userName"`
	PasswordAttempts int64  `json:"passwordAttempts"`
	UserStatus       string `json:"userStatus"`
	UserType         string `json:"userType"`
}

type EntityDetailsReq struct {
	UserEmail   string `json:"userEmail"`
	RequestType string `json:"requestType"`
	EntityId    *int64 `json:"entityId"`
	UserType    int    `json:"userType"`
	EntityType  string `json:"entityType"`
}

type EntityDetailsResp struct {
	UserId     *int64     `json:"userId"`
	UserName   string     `json:"userName"`
	NumOfUsers int64      `json:"numOfUsers"`
	ExpiryDate *time.Time `json:"expiryDate"`
	EntityId   *int64     `json:"entityId"`
	EntityName string     `json:"entityName"`
	EntityType string     `json:"entityType"`
	EntityInfo string     `json:"entityInfo"`
}

type UserCreationFailStruct struct {
	FailResp []BulkUploadFailResp
}

type BulkUploadFailResp struct {
	UserEmail string `json:"userEmail"`
	Error     string `json:"error"`
}
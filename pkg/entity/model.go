package entity

type UpdateEntityReq struct {
	EntityId    int64  `json:"entityId"`
	EntityDivId int64  `json:"entityDivId"`
	EntityLevel int    `json:"entityLevel"`
	Name        string `json:"name"`
}

type CreateEntityReq struct {
	EntityId       int64  `json:"entityId"`
	EntityDivId    int64  `json:"entityDivId"`
	EntityName     string `json:"entityName"`
	EntityTypeName string `json:"entityTypeName"`
	NumOfUsers     int64  `json:"numOfUsers"`
	ContactName    string `json:"contactName"`
	ContactEmail   string `json:"contactEmail"`
	ContactMob     string `json:"contactMob"`
	LogoPath       string `json:"logoPath"`
}

type EntityLicenseResp struct {
	NumberOfLicences int64 `json:"numberOfLicences"`
	LicenceTypeId    int64 `json:"licenceTypeId"`
	LicenceConsumed  int64 `json:"licenceConsumed"`
}

type DeleteEntityReq struct {
	EntityId    int64 `json:"entityId"`
	EntityDivId int64 `json:"entityDivId"`
	EntityLevel int   `json:"entityLevel"`
}

package entity

import "github.com/Lanquill/Forge/pkg/resp"

const (
	ErrEntityExists = "entity exists"
)

const (
	ErrNoUserLicense = "no user license"
)

var EntityExists = resp.RespMessage{
	Title:       "Entity name already exists",
	Description: " ",
}

var EntityInfoUpdateSuccessful = resp.RespMessage{
	Title:       "Entity info updated successfully!",
	Description: "",
}

var EntityDeleteSuccessful = resp.RespMessage{
	Title:       "Entity deleted successfully!",
	Description: "",
}

var NoUserLicense = resp.RespMessage{
	Title:       "Please buy additional user licences. You do not have those many user licences",
	Description: " ",
}

var EntityCreateSuccessful = resp.RespMessage{
	Title:       "Entity created successfully!",
	Description: "",
}

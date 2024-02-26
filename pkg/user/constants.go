package user

import "github.com/Lanquill/Forge/pkg/resp"

const (
	ErrEmailExists = "email exists"
)

const (
	ErrEmailOrMobExists = "email or mobile exists"
)

const (
	ErrNoUserLicense = "no user license"
)

const (
	ErrInvalidEmail = "invalid email"
)

const (
	ErrFailed = "operation failed"
)

var EmailExists = resp.RespMessage{
	Title:       "Email Id exists",
	Description: " ",
}

var EmailOrMobExists = resp.RespMessage{
	Title:       "User Email or Mobile already registered",
	Description: " ",
}

var InvalidEmail = resp.RespMessage{
	Title:       "Invalid Email",
	Description: " ",
}

var NoUserLicense = resp.RespMessage{
	Title:       "Please buy additional user licences. You do not have those many user licences",
	Description: " ",
}

var ExceedsMaxUser = resp.RespMessage{
	Title:       "Please upload max 1000 users at a time",
	Description: " ",
}

var UserInfoUpdateSuccessful = resp.RespMessage{
	Title:       "User info updated successfully!",
	Description: "",
}

var UserDeleteSuccessful = resp.RespMessage{
	Title:       "User deleted successfully!",
	Description: "",
}

var UserCreateSuccessful = resp.RespMessage{
	Title:       "User created successfully!",
	Description: "",
}

var UserBulkCreateFailed = resp.RespMessage{
	Title:       "Invalid file",
	Description: "",
}

var UserBulkCreateResponse = resp.RespMessage{
	Title:       "Status - %d Success | %d Duplicate | %d not uploaded successfully | %d not valid email",
	Description: "",
}

var ExpiryDateUpdateSuccessful = resp.RespMessage{
	Title:       "Expiry date changed!",
	Description: "",
}

var UserBulkCreateProcessing = resp.RespMessage{
	Title:       "We are processing your request, Users will be created in few minutes",
	Description: "",
}
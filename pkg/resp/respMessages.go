package resp

type RespMessage struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"desc"  bson:"desc"`
}

var InternalErrorMsg = "There was error while processing your request. Please try again later."

var InternalServerError = RespMessage{
	"Something went wrong!",
	InternalErrorMsg,
}

var FileTooLarge = RespMessage{
	"Uploaded file is too large!",
	" ",
}

var InvalidLogoFileType = RespMessage{
	"Invalid file type!",
	"Only .jpg, .png and .webp are accepted.",
}

var NoCookie = RespMessage{
	"Request unauthorized",
	"Authorization cookie not found",
}

var InvalidToken = RespMessage{
	"Request unauthorized",
	"Invalid access token",
}

var InvalidReqBody = RespMessage{
	"Invalid request",
	"The content you are requesting is invalid. Please try again.",
}

var InvalidEmail = RespMessage{
	"Invalid email",
	"The email you are requesting is invalid. Please try again.",
}

var SuccessResp = RespMessage{
	"Success",
	"Success",
}

var ResetPasswordResp = RespMessage{
	"New password updated successfully!",
	"",
}

const (
	StatusInvalidRequestBody = 461
)

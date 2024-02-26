package questions

import "github.com/Lanquill/Forge/pkg/resp"

const (
	ErrWordLimit = "exceeded the maximum words limit "
)

const (
	ErrPassageExists = "passage with the specified title/text already exists"
)

const (
	ErrCreateFailure = "creation of document failed"
)

const (
	ErrUpdateFailure = "no documents were updated"
)

const (
	ErrDeleteFailure = "no documents were deleted"
)

const (
	ErrNoDocumentsExits = "the document was not found"
)

const (
	ErrIpAddressExists = "ip address already exists"
)

const (
	ErrFileCreationFailure = "file creation failed"
)

var WordLimit = resp.RespMessage{
	Title:       "Exceeded the maximum words limit",
	Description: "",
}

var PassageExists = resp.RespMessage{
	Title:       "Passage already exists",
	Description: " ",
}

var CreateFailure = resp.RespMessage{
	Title:       "Document not Created",
	Description: "No Document was created",
}

var UpdateFailure = resp.RespMessage{
	Title:       "Update failure",
	Description: "No Documents were updated",
}

var DeleteFailure = resp.RespMessage{
	Title:       "Delete failure",
	Description: "No Documents were Deleted",
}

var NoDocumentsExits = resp.RespMessage{
	Title:       "Document not found",
	Description: " ",
}

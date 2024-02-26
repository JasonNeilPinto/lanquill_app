package activity

import "github.com/Lanquill/Forge/pkg/resp"

var CountryCode = map[string]string{
	"IN":    "IN",
	"India": "India",
}

var UserTypeLevel = map[int]string{
	4:  "Level6",
	5:  "Level6",
	6:  "Level6",
	7:  "Level5",
	8:  "Level5",
	9:  "Level5",
	10: "Level4",
	11: "Level4",
	12: "Level4",
	13: "Level3",
	14: "Level3",
	15: "Level3",
	16: "Level2",
	17: "Level2",
	18: "Level2",
	19: "Level1",
	20: "Level1",
}
var UserTypeLevelMap = map[int]int{
	4:  6,
	5:  6,
	6:  6,
	7:  5,
	8:  5,
	9:  5,
	10: 4,
	11: 4,
	12: 4,
	13: 3,
	14: 3,
	15: 3,
	16: 2,
	17: 2,
	18: 2,
	19: 1,
	20: 1,
}

var IpAddressExists = resp.RespMessage{
	Title:       "Ip Address Exists",
	Description: "The provided IP address already exists",
}

var FileCreationFailed = resp.RespMessage{
	Title:       "File creation failed",
	Description: "The creation of the report failed",
}

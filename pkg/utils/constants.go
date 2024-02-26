package utils

import "os"

var STATICPATH = os.Getenv("LQ_STATIC_PATH")
var LQ_DB_PATH = os.Getenv("LQ_DB_PATH")
var LOGOPATH = STATICPATH + "/logo"

var HigherLevelTypes = map[int]int{
	7: 7,
	6: 6,
}
var MiddleLevelTypes = map[int]int{
	5: 5,
}
var LowerLevelTypes = map[int]int{
	4: 4,
	3: 3,
	2: 2,
}

var LevelWithLogo = map[int]int{
	7: 7,
	6: 6,
	5: 5,
}

var LevelWithoutLogo = map[int]int{
	4: 4,
	3: 3,
	2: 2,
}

var LevelMap = map[int]string{
	8: "Level_8",
	7: "Level_7",
	6: "Level_6",
	5: "Level_5",
	4: "Level_4",
	3: "Level_3",
	2: "Level_2",
	1: "Level_1",
}

var EntityLevel = map[string]string{
	"Level_8": "Level_8",
	"Level_7": "Level_7",
	"Level_6": "Level_6",
	"Level_5": "Level_5",
	"Level_4": "Level_4",
	"Level_3": "Level_3",
	"Level_2": "Level_2",
	"Level_1": "Level_1",
}

var AcceptedEntityLogoType = map[string]string{
	"image/jpeg": ".jpg",
	"image/webp": ".webp",
	"image/png":  ".png",
}

var UserTypeMap = map[int]int{
	2:  2,
	3:  3,
	4:  4,
	5:  5,
	6:  6,
	7:  7,
	8:  8,
	9:  9,
	10: 10,
	11: 11,
	12: 12,
	13: 13,
	14: 14,
	15: 15,
	16: 16,
	17: 17,
	18: 18,
	19: 19,
	20: 20,
}

var UserEntityTypeMap = map[string]string{
	"University":         "University_Admin",
	"Corporate_Global":   "Corporate_Global_Admin",
	"School_Chain":       "SchoolChain_Admin",
	"College":            "College_Admin",
	"Corporate_Country":  "Corporate_Country_Admin",
	"School":             "School_Admin",
	"Department":         "Department_Admin",
	"Degree":             "Degree_Admin",
	"Semester":           "Semester_Admin",
	"Section":            "Section_Admin",
	"Grade":              "Grade_Admin",
	"Corporate_City":     "Corporate_City_Admin",
	"Corporate_Locality": "Corporate_Locality_Admin",
	"Business_Unit":      "Business_Unit_Admin",
	"Division":           "Division_Admin",
}

const (
	ErrInvalidLevel = "invalid level"
)

const (
	ErrInvalidRequest = "invalid request"
)

var AcceptedAudioType = map[string]string{
	"audio/mpeg": "mp3",
}

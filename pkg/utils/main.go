package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lanquill/Forge/pkg/auth"
)

func UserTokenData(token *http.Cookie) (*auth.JWTData, error) {
	return auth.UnverifiedJwt(token.Value)
}

func GetLowerLevelQuery(entityLevel int) string {
	sub_query := ""
	for i := entityLevel - 1; i >= 1; i-- {
		sub_query += "AND Level_" + strconv.Itoa(i) + " IS NULL AND User_Status != 'Deleted'"
	}

	return sub_query
}

func EncryptSHA1(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func IntegerToEnUs(input int) string {
	var englishMegas = []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion", "octillion", "nonillion", "decillion", "undecillion", "duodecillion", "tredecillion", "quattuordecillion"}
	var englishUnits = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var englishTens = []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	var englishTeens = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

	words := []string{}

	if input < 0 {
		words = append(words, "minus")
		input *= -1
	}

	triplets := integerToTriplets(input)

	if len(triplets) == 0 {
		return "zero"
	}

	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]

		if triplet == 0 {
			continue
		}

		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		if hundreds > 0 {
			words = append(words, englishUnits[hundreds], "hundred")
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, englishUnits[units])
		case 1:
			words = append(words, englishTeens[units])
		default:
			if units > 0 {
				word := fmt.Sprintf("%s-%s", englishTens[tens], englishUnits[units])
				words = append(words, word)
			} else {
				words = append(words, englishTens[tens])
			}
		}

	tripletEnd:
		if mega := englishMegas[idx]; mega != "" {
			words = append(words, mega)
		}
	}

	return strings.Join(words, " ")
}

func integerToTriplets(number int) []int {
	triplets := []int{}

	for number > 0 {
		triplets = append(triplets, number%1000)
		number = number / 1000
	}

	return triplets
}

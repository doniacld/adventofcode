package passport

import (
	"fmt"
	"regexp"
	"strconv"
)

// Validator is a function validating an entry
type Validator func(string) bool

const (
	// mandatory fields of the passport
	birthYear      = "byr"
	issueYear      = "iyr"
	eyeColor       = "ecl"
	expirationYear = "eyr"
	height         = "hgt"
	hairColor      = "hcl"
	passportId     = "pid"
)

var mandatoryFields = map[string]Validator{
	birthYear:      yearRule(1920, 2002),
	issueYear:      yearRule(2010, 2020),
	expirationYear: yearRule(2020, 2030),
	eyeColor:       regexpRule("amb|blu|brn|gry|grn|hzl|oth"),
	height:         heightRule,
	hairColor:      regexpRule("#[0-9a-f]{6}"),
	passportId:     regexpRule("[0-9]{9}"),
}

func yearRule(min, max int) func(string) bool {
	return func(s string) bool {
		value, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return value >= min && value <= max
	}
}

func heightRule(s string) bool {
	r := regexp.MustCompile("([0-9]{2,3})(cm|in)")
	if !r.MatchString(s) {
		return false
	}

	values := r.FindStringSubmatch(s)

	height, err := strconv.Atoi(values[1])
	if err != nil {
		fmt.Println(err)
		return false
	}

	unit := values[2]
	switch unit {
	case "cm":
		return height >= 150 && height <= 193
	case "in":
		return height >= 59 && height <= 76
	default:
		return false
	}
}

func regexpRule(reg string) func(string) bool {
	return func(s string) bool {
		r := "^" + reg + "$"
		res, err := regexp.MatchString(r, s)
		if err != nil {
			return false
		}
		return res
	}
}

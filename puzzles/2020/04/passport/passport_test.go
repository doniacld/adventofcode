package passport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassport_IsValid(t *testing.T) {
	tt := []struct {
		description string
		passport    Passport
		expected    bool
	}{
		{"valid passport", Passport{"pid": "087499704", "hgt": "74in", "ecl": "grn", "iyr": "2012",
			"eyr": "2030", "byr": "1980", "hcl": "#623a2f"}, true},
		{"invalid height", Passport{"eyr": "1972", "cid": "100",
			"hcl": "#18171d", "ecl": "amb", "hgt": "170", "pid": "186cm", "iyr": "2018", "byr": "1926"}, false},
		{"invalid expiration year", Passport{"iyr": "2019", "hcl": "#602927", "eyr": "1967", "hgt": "170cm",
			"ecl": "grn", "pid": "012533040", "byr": "1946"}, false},
		{"invalid hair color", Passport{"hcl": "dab227", "iyr": "2012",
			"ecl": "brn", "hgt": "182cm", "pid": "021572410", "eyr": "2020", "byr": "1992", "cid": "277"}, false},
		{"invalid birth year", Passport{"hgt": "59cm", "ecl": "zzz",
			"eyr": "2038", "hcl": "74454a", "iyr": "2023",
			"pid": "3556412378", "byr": "2007"}, false},
			{"invalid number of fields - less", Passport{"hgt": "155cm"}, false},
		{"invalid number of fields - higher", Passport{"pid": "087499704", "hgt": "74in", "ecl": "grn", "iyr": "2012",
			"eyr": "2030", "byr": "1980", "hcl": "#623a2f", "hct": "toto"}, true},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.passport.IsValid())
		})
	}

}

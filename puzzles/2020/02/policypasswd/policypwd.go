package policypasswd

import (
	"strconv"
	"strings"
)

// policy password based on letter occurrence holds all the information about the validity of the given password
type policy struct {
	firstParam  int
	secondParam int
	letter      string
	password    string
}

// RetrieveValidPasswordsOccurrencePolicy returns the valid passwords using the occurrence policy
func RetrieveValidPasswordsOccurrencePolicy(inputs []string) (int, error) {
	var counter int
	for _, in := range inputs {
		password, err := extractPolicy(in)
		if err != nil {
			return 0, err
		}
		if password.isPolicyOccurrenceValid() {
			counter++
		}
	}
	return counter, nil
}

// RetrieveValidPasswordsPositionPolicy returns the valid passwords using the position policy
func RetrieveValidPasswordsPositionPolicy(inputs []string) (int, error) {
	var counter int
	for _, in := range inputs {
		password, err := extractPolicy(in)
		if err != nil {
			return 0, err
		}
		if password.isPositionPolicyValid() {
			counter++
		}
	}
	return counter, nil
}

// Extract the policy and the password from the given string
// e.g.: "1-3 a: abcde"
// first param = 1; second param = 3; letter = a; password = abcde
func extractPolicy(input string) (policyPwd policy, err error) {
	policyAndPwdString := strings.Split(input, " ")
	policyPwd.password = policyAndPwdString[2]

	letter := strings.Split(policyAndPwdString[1], "")
	policyPwd.letter = letter[0]

	limits := strings.Split(policyAndPwdString[0], "-")
	policyPwd.firstParam, err = strconv.Atoi(limits[0])
	if err != nil {
		return policy{}, err
	}
	policyPwd.secondParam, err = strconv.Atoi(limits[1])
	if err != nil {
		return policy{}, err
	}
	return policyPwd, nil
}

// Check the validity of the password
// the given letter should have its occurrence in the given range
// assume that the min is always smaller than max
func (p *policy) isPolicyOccurrenceValid() bool {
	// early exit if the pwd does not contain the letter at all
	if !strings.Contains(p.password, p.letter) {
		return false
	}

	pwdArr := strings.Split(p.password, "")

	var counter int
	for _, c := range pwdArr {
		if c == p.letter {
			counter++
			// early exit if the counter is beyond the given max
			if counter > p.secondParam {
				return false
			}
		}
	}
	// finally check if the counter is min
	if counter < p.firstParam {
		return false
	}
	return true
}

// Check the validity of the password
// the given letter should be whereas at the first param position or the second one
func (p *policy) isPositionPolicyValid() bool {
	// early exit if the pwd does not contain the letter at all
	if !strings.Contains(p.password, p.letter) {
		return false
	}

	pwdArr := strings.Split(p.password, "")
	sub := []string{pwdArr[p.firstParam-1], pwdArr[p.secondParam-1]}

	counter := strings.Count(strings.Join(sub, ""), p.letter)
	if counter != 1 {
		return false
	}
	return true
}

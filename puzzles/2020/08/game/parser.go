package game

import (
	"regexp"
	"strconv"
)

var (
	reg = regexp.MustCompile("^([a-z]{3}) ([+|-][0-9]+)")
)

func (ops *Operations) ParseLine(line string) error {
	r := reg.FindStringSubmatch(line)

	v, err := strconv.Atoi(r[2])
	if err != nil {
		return err
	}

	o := NewOperation(r[1], v)
	*ops = append(*ops, o)
	return nil
}

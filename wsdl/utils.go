package wsdl

import (
	"strconv"
)

var epsilon = 1e-9

func atof(s string) *float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil
	}
	return &f
}

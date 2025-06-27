package helper

import (
	"strconv"
	"tarmac/logger"
)

var epsilon = 1e-9

func atof(s string) *float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		logger.Log.Log("Cannot convert", s, "to float64")
		return nil
	}
	return &f
}

package helper

import (
	"log"
	"strconv"
)

var epsilon = 1e-9

func atof(s string) *float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println("Cannot convert", s, "to float64")
		return nil
	}
	return &f
}

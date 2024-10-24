package utils

// File: utils/utils.go

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package utils

import "fmt"

func CheckErr(err error) bool {
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	return true
}

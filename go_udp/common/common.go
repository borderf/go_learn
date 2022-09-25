package common

import (
	"fmt"
	"os"
)

var Delemeter = []byte{54, 36, 35, 35, 34}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

package main

import (
	"os"
	"strconv"
)

// PathExists ...
func PathExists(path string) (bool, error) {
	var err error
	if _, err = os.Stat(path); err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ParseBool str to bool
func ParseBool(str string) bool {
	b, _ := strconv.ParseBool(str)
	return b
}

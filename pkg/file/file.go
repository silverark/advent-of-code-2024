package file

import (
	"os"
	"strings"
)

func GetFile(filename string) []string {
	dat := GetData(filename)
	return strings.Split(dat, "\n")
}

func GetData(filename string) string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic("Unable to load file: " + err.Error())
	}
	return string(dat)
}

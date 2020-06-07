package kiyaui

import (
	"io/ioutil"
	"os"
)

func ensureFileExists(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_ = ioutil.WriteFile(filename, []byte("[]"), 0600)
	}
}

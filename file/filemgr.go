package file

import (
	"os"
)

func FileNotExist(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err)
}

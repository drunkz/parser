package prep

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	SCRIPT_ROOT = filepath.Join(path, "../example")
}

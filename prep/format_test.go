package prep

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"
)

func TestFormatScript(t *testing.T) {
	start := time.Now()

	list, err := FormatScript(filepath.Join(SCRIPT_ROOT, "main.txt"))
	if err != nil {
		err.Print()
	} else {
		for i := list.Front(); i != nil; i = i.Next() {
			fmt.Println(i.Value)
		}
	}

	fmt.Println(time.Since(start))

}

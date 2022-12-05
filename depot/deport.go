package depot

import (
	"log"
	"os"
)

var WorkDir string = "work"
var WorkDirPattern string = "unodata-*"

func MkdirTemp() {
	d, _ := os.MkdirTemp("", WorkDirPattern)
	WorkDir = d
	log.Printf("Server working directory '%s'", d)
}

func CleanTemp() {
	log.Printf("Removing directory '%s'", WorkDir)
	os.RemoveAll(WorkDir)
}

package sqlite3

// #include "bundle/sqlite3.h"
import "C"
import "os"

type ProgressReport struct {
	os.Error
	Total			int
	Remaining		int
	Source			string
	Target			string
}

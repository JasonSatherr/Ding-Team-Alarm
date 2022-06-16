package fileTools

import (
	"os"
	"path/filepath"

	"github.com/JasonSatherr/Ding-Team-Alarm/tree/main/Client/src/convenience"
)

//geExecutionPath is a function that will return the path the the program is executing from
func GetExecutionPath() string {
	ex, err := os.Executable()
	convenience.HandlePotErr(err)
	exPath := filepath.Dir(ex)
	return exPath
}

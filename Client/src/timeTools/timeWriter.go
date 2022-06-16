package timeTools

import (
	"fmt"
	"os"
	"time"

	"github.com/JasonSatherr/Ding-Team-Alarm/tree/main/Client/src/fileTools"
)

func PrintCurrentTime() {
	currTime := time.Now()
	formattedTime := currTime.Format(time.UnixDate)
	fmt.Println(formattedTime)
}

//unfortunately involves hard code
func getPathToCSV() string {
	execPath := fileTools.GetExecutionPath()
	pathToOpen := execPath + "/src/data/time-open.csv"
	return pathToOpen
}

func writeTimeToFile(filePath string, timeToRecord time.Time) {

	//first open the file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	//then get the csv writer into a var csvWriter

	//then get the string version of timeToRecord.format(time.UnixDate)
}

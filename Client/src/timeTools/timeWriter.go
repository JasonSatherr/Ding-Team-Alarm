//package timeTools attempts to provide
package timeTools

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/JasonSatherr/Ding-Team-Alarm/tree/main/Client/src/convenience"
)

func PrintCurrentTime() {
	currTime := time.Now()
	formattedTime := currTime.Format(time.UnixDate)
	fmt.Println(formattedTime)
}

//unfortunately involves hard code
func getPathToStorage() string {
	//execPath := fileTools.GetExecutionPath()
	//pathToOpen := execPath + "src\\data\\time-open.csv"
	pathToOpen := "src\\data\\time-open.csv"
	//dangerous because only works on windows?
	fmt.Printf("\nThe file path we are looking for is...%s\n", pathToOpen)
	return pathToOpen
}

func WriteTimeToStorage(filePath string, timeToRecord time.Time) {

	//first open the file
	//file, err := os.Open(filePath)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	convenience.HandlePotErr(err)
	defer file.Close()
	//then get the csv writer into a var csvWriter
	csvWriter := csv.NewWriter(file) //file implements the io.writer interface

	//then get the string version of timeToRecord.format(time.UnixDate)
	timeString := timeToRecord.Format(time.UnixDate)
	fmt.Printf("\n time to write is %s \n", timeString)
	//then write the time
	var record []string
	record = append(record, timeString)
	ok := csvWriter.Write(record)
	if ok != nil {
		panic(ok)
	}
	csvWriter.Flush()
	err = csvWriter.Error()
	convenience.HandlePotErr(err)
}

func WriteCurrentTimeToFile() {
	filePath := getPathToStorage()
	WriteTimeToStorage(filePath, time.Now())
}

func WriteToHi() {
	//first open the file
	file, err := os.Create("src\\data\\hi.txt")
	convenience.HandlePotErr(err)
	defer file.Close()
	//then get the csv writer into a var csvWriter
	var message []byte
	message = append(message, 77, 78)
	file.Write(message)
}

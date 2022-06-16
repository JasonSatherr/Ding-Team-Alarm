package data

import (
	"fmt"
	"time"
)

func printCurrentTime() {
	currTime := time.Now()
	formattedTime := currTime.Format(time.UnixDate)
	fmt.Println(formattedTime)
}

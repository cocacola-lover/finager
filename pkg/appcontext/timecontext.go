package appcontext

import (
	"fmt"
	"time"

	"github.com/peterh/liner"
)

func DisplayTime(date time.Time) {
	fmt.Printf("%d/%d/%d\n", date.Year(), date.Month(), date.Day())
}

func ParseTime(line *liner.State, timeContext *time.Time) error {
	dateStr, err := line.Prompt("Time: ")
	if err != nil {
		return err
	}
	newTime, err := time.Parse("2/1/2006", dateStr)
	if err != nil {
		return err
	}
	*timeContext = newTime
	return nil
}

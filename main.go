package main

import (
	"fmt"
	"time"

	"github.com/blakek/get-time-cli/internal/options"
	"github.com/blakek/get-time-cli/internal/timesheet"
)

func printLongSummary(timesheet *timesheet.Timesheet, options *options.Options) {
	fmt.Printf("Current:   %0.2f\n", timesheet.TimeWorked.Hours())
	fmt.Printf("Remaining: %0.2f", timesheet.TimeRemaining.Hours())

	if timesheet.IsCompleted {
		fmt.Printf(" (%v)", timesheet.CompletionTime.Format(time.Kitchen))
	}

	fmt.Printf("\n")
}

func printShortSummary(timesheet *timesheet.Timesheet, options *options.Options) {
	fmt.Printf("%0.2f hr", timesheet.TimeWorked.Hours())

	if timesheet.IsCompleted {
		fmt.Printf(" (@%v)", timesheet.CompletionTime.Format(time.Kitchen))
	}

	fmt.Printf("\n")
}

func main() {
	opts := options.FetchOptions()
	// filePath := options.TimesheetPath("2021-01-28.md", opts)
	// filePath := "/Users/bknight/Desktop/scratchpad/times-txt/2021/2021-08/2021-08-03.md"
	filePath := "/Users/bknight/Desktop/scratchpad/times-txt/2021-08-25.md"

	timesheet := timesheet.ParseFile(filePath, opts)

	printLongSummary(timesheet, opts)
	fmt.Println("---")
	printShortSummary(timesheet, opts)
}

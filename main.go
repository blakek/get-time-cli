package main

import (
	"fmt"
	"time"

	"github.com/blakek/get-time-cli/internal/options"
	"github.com/blakek/get-time-cli/internal/timesheet"
)

func getCompletionTime(hoursRemaining float32, options *options.Options) time.Time {
	// `time.Duration` doesn't accept floats, so the duration is converted to
	// minutes instead of fractions of hours
	minutesRemaining := time.Minute * time.Duration(hoursRemaining*60)
	granularity := 60 * time.Minute / time.Duration(options.Granularity)
	return time.Now().Add(minutesRemaining).Round(granularity)
}

func getHoursRemaining(hoursWorked float32, options *options.Options) float32 {
	hoursRemaining := options.HoursPerDay - hoursWorked

	if hoursRemaining > 0 {
		return hoursRemaining
	}

	return 0
}

func printLongSummary(hoursWorked float32, options *options.Options) {
	hoursRemaining := getHoursRemaining(hoursWorked, options)
	completionTime := getCompletionTime(hoursRemaining, options)

	fmt.Printf("Current:   %0.2f\n", hoursWorked)
	fmt.Printf("Remaining: %0.2f", hoursRemaining)

	if hoursRemaining > 0 {
		fmt.Printf(" (%v)", completionTime.Format(time.Kitchen))
	}

	fmt.Printf("\n")
}

func printShortSummary(hoursWorked float32, options *options.Options) {
	hoursRemaining := getHoursRemaining(hoursWorked, options)
	completionTime := getCompletionTime(hoursRemaining, options)

	fmt.Printf("%0.2f hr", hoursWorked)

	if hoursRemaining > 0 {
		fmt.Printf(" (@%v)", completionTime.Format(time.Kitchen))
	}

	fmt.Printf("\n")
}

func main() {
	opts := options.FetchOptions()
	// filePath := options.TimesheetPath("2021-01-28.md", opts)
	// filePath := "/Users/bknight/Desktop/scratchpad/times-txt/2021/2021-08/2021-08-03.md"
	filePath := "/Users/bknight/Desktop/scratchpad/times-txt/2021-08-25.md"

	timeEntries := timesheet.ParseFile(filePath, opts)
	hoursWorked := timesheet.GetHoursWorked(timeEntries)

	printLongSummary(hoursWorked, opts)
	fmt.Println("---")
	printShortSummary(hoursWorked, opts)
}

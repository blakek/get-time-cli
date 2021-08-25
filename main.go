package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/blakek/get-time-cli/internal/options"
	"github.com/blakek/get-time-cli/internal/timesheet"
)

func getSeparatorLength(timesheet *timesheet.Timesheet) int {
	const (
		maxLength = 64
		padding   = 2
	)

	length := 10

	for _, entry := range timesheet.Entries {
		nameLength := len(entry.Name)

		if nameLength > length-padding && nameLength < maxLength {
			length = nameLength + padding
		}
	}

	return length
}

func printLongSummary(timesheet *timesheet.Timesheet, options *options.Options) {
	const (
		maxNoteLength = 64
	)

	fmt.Printf("Current:   %0.2f\n", timesheet.TimeWorked.Hours())
	fmt.Printf("Remaining: %0.2f", timesheet.TimeRemaining.Hours())

	if !timesheet.IsCompleted {
		fmt.Printf(" (%v)", timesheet.CompletionTime.Format(time.Kitchen))
	}

	fmt.Printf("\n")

	separatorLength := getSeparatorLength(timesheet)

	if len(timesheet.Entries) > 0 {
		fmt.Printf("\n")
	}

	for _, entry := range timesheet.Entries {
		var noteText string
		separator := strings.Repeat(".", separatorLength-len(entry.Name))

		// Ensure text fits in max area
		if len(entry.Notes) > 0 {
			noteText = "# " + entry.Notes
			if len(noteText) > maxNoteLength {
				noteText = noteText[0:maxNoteLength-1] + "…"
			}
		}

		fmt.Printf("- %s%s%.2f  %s\n", entry.Name, separator, entry.Time, noteText)
	}
}

func printShortSummary(timesheet *timesheet.Timesheet, options *options.Options) {
	fmt.Printf("%0.2f hr", timesheet.TimeWorked.Hours())

	if !timesheet.IsCompleted {
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

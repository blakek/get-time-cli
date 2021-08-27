package summary

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/blakek/get-time-cli/internal/options"
	"github.com/blakek/get-time-cli/internal/timesheet"
	"golang.org/x/term"
)

var maxOutputWidth = getOutputWidth()

func PrintLongSummary(timesheet *timesheet.Timesheet, options *options.Options) {
	const (
		maxNameWidth = 32
		maxTimeWidth = 6
	)

	separatorLength := getSeparatorLength(timesheet, maxNameWidth)
	maxNoteLength := maxOutputWidth - maxTimeWidth - separatorLength

	fmt.Printf("Current:   %0.2f\n", timesheet.TimeWorked.Hours())

	if !timesheet.IsCompleted {
		fmt.Printf("Remaining: %0.2f", timesheet.TimeRemaining.Hours())
		fmt.Printf(" (%v)", timesheet.CompletionTime.Format(time.Kitchen))
		fmt.Printf("\n")
	}

	if len(timesheet.Entries) > 0 {
		fmt.Printf("\n")
	}

	for _, entry := range timesheet.Entries {
		var noteText string
		separator := strings.Repeat(".", separatorLength-len(entry.Name))

		// Ensure text fits in max area
		if len(entry.Notes) > 0 {
			noteText = "# " + entry.Notes
			additionalLength := 4
			if len(noteText)+additionalLength > maxNoteLength {
				noteText = noteText[0:maxNoteLength-additionalLength+1] + "…"
			}
		}

		fmt.Printf("- %s%s%.2f  %s\n", entry.Name, separator, entry.Time, noteText)
	}
}

func PrintShortSummary(timesheet *timesheet.Timesheet, options *options.Options) {
	fmt.Printf("%0.2f hr", timesheet.TimeWorked.Hours())

	if !timesheet.IsCompleted {
		fmt.Printf(" (@%v)", timesheet.CompletionTime.Format(time.Kitchen))
	}

	fmt.Printf("\n")
}

func getOutputWidth() int {
	const (
		maximumWidth = 320
		minimumWidth = 32
	)

	stdoutFd := int(os.Stdout.Fd())

	if term.IsTerminal(stdoutFd) {
		width, _, _ := term.GetSize(stdoutFd)

		if width >= minimumWidth {
			return width
		}
	}

	return maximumWidth
}

func getSeparatorLength(timesheet *timesheet.Timesheet, maxLength int) int {
	const (
		padding = 2
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

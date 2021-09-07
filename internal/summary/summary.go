package summary

import (
	"fmt"
	"os"
	"time"

	"github.com/DataDrake/flair"
	"github.com/blakek/get-time-cli/internal/options"
	"github.com/blakek/get-time-cli/internal/timesheet"
	"golang.org/x/term"
)

var (
	stdoutFd   = int(os.Stdout.Fd())
	isTerminal = term.IsTerminal(stdoutFd)
)

func PrintLongSummary(timesheet *timesheet.Timesheet, options *options.Options) {

	fmt.Printf("%0.2f hours worked\n", timesheet.TimeWorked.Hours())

	if !timesheet.IsCompleted {
		fmt.Printf(
			"Done at %v (%0.2f remaining)\n",
			timesheet.CompletionTime.Format(time.Kitchen),
			timesheet.TimeRemaining.Hours(),
		)
	}

	if len(timesheet.Entries) > 0 {
		fmt.Printf("\n")
	}

	for _, entry := range timesheet.Entries {
		sectionTitle := fmt.Sprintf("%s: %.2f", entry.Name, entry.Time)

		fmt.Printf("%s\n", formatTitleText(sectionTitle))

		if entry.Notes != "" {
			fmt.Print(formatNoteText(entry.Notes))
		}

		fmt.Println()
	}
}

func PrintShortSummary(timesheet *timesheet.Timesheet, options *options.Options) {
	fmt.Printf("%0.2f hr", timesheet.TimeWorked.Hours())

	if !timesheet.IsCompleted {
		fmt.Printf(" (@%v)", timesheet.CompletionTime.Format(time.Kitchen))
	}

	fmt.Printf("\n")
}

func formatNoteText(noteText string) string {
	return fmt.Sprintf("    %s\n", maybeFormatText(noteText, flair.Dim))
}

func formatTitleText(title string) string {
	return maybeFormatText(title, flair.Bold)
}

func maybeFormatText(text string, formatFunc func(string) string) string {
	if isTerminal {
		return formatFunc(text)
	}

	return text
}

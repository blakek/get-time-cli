package summary

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/DataDrake/flair"
	"github.com/blakek/get-time-cli/internal/options"
	"github.com/blakek/get-time-cli/internal/timesheet"
	"golang.org/x/term"
)

const (
	LongSummary    string = "long"
	ShortSummary   string = "short"
	UnicodeSummary string = "unicode"
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

func PrintUnicodeArtSummary(timesheet *timesheet.Timesheet, options *options.Options) {
	symbols := []string{"⣿", "⣷", "⣶", "⣦", "⣤", "⣄", "⣀", "⡀"}

	symbolIndex := int(math.Min(
		float64(len(symbols)-1),
		math.Max(0, math.Round(float64(timesheet.TimeRemaining.Hours()))),
	))

	fmt.Print(symbols[symbolIndex])

	if !timesheet.IsCompleted {
		fmt.Printf(" (%v)", timesheet.CompletionTime.Format(time.Kitchen))
	}

	fmt.Printf("\n")
}

func PrintSummary(summaryType string, timesheet *timesheet.Timesheet, options *options.Options) error {
	switch summaryType {
	case LongSummary:
		PrintLongSummary(timesheet, options)

	case ShortSummary:
		PrintShortSummary(timesheet, options)

	case UnicodeSummary:
		PrintUnicodeArtSummary(timesheet, options)

	default:
		return fmt.Errorf("summary: unknown summary type \"%s\"", summaryType)
	}

	return nil
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

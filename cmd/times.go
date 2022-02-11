package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/blakek/get-time-cli/internal/options"
	"github.com/blakek/get-time-cli/internal/summary"
	"github.com/blakek/get-time-cli/internal/timesheet"
	"github.com/spf13/cobra"
)

var summaryType string

// timesCmd represents the times command
var timesCmd = &cobra.Command{
	Use:   "times",
	Short: "Get times for a timesheet file",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		dayOffset := 0

		if len(args) > 0 {
			parsedDayOffset, err := strconv.ParseInt(args[0], 10, 0)

			if err != nil {
				errorMessage := fmt.Sprintf("invalid day offset: \"%s\"", args[0])
				log.Fatal(errorMessage)
			}

			dayOffset = int(parsedDayOffset)
		}

		opts := options.FetchOptions()
		fileDate := time.Now().AddDate(0, 0, dayOffset)
		timesheetPath := timesheet.GetPathForDate(fileDate, opts)
		timesheet := timesheet.ParseFile(timesheetPath, opts)

		err := summary.PrintSummary(summaryType, timesheet, opts)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(timesCmd)

	timesCmd.Flags().StringVarP(&summaryType, "summary-type", "s", summary.LongSummary, "What kind of summary to show")
}

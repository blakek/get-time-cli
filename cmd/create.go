package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/blakek/get-time-cli/internal/options"
	"github.com/blakek/get-time-cli/internal/timesheet"
	"github.com/spf13/cobra"
)

var (
	startTime uint
	endTime   uint
)

var createCmd = &cobra.Command{
	Use:     "create [day offset]",
	Aliases: []string{"new"},
	Short:   "Create a new timesheet file",
	Args:    cobra.RangeArgs(0, 1),
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

		_, err := os.Stat(timesheetPath)

		if !os.IsNotExist(err) {
			log.Fatal("A timesheet already exists at " + timesheetPath)
		}

		timesheetFile, err := os.Create(timesheetPath)

		if err != nil {
			log.Fatal(err)
		}

		defer timesheetFile.Close()

		template := timesheet.Create(startTime, endTime, opts)

		_, err = timesheetFile.WriteString(template)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().UintVarP(&endTime, "end-time", "e", 15, "Sets the ending time for a new timesheet")
	createCmd.Flags().UintVarP(&startTime, "start-time", "s", 7, "Sets the beginning time for a new timesheet")
}

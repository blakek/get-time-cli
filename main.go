package main

import (
	"github.com/blakek/get-time-cli/internal/options"
	"github.com/blakek/get-time-cli/internal/summary"
	"github.com/blakek/get-time-cli/internal/timesheet"
)

func main() {
	opts := options.FetchOptions()
	// filePath := options.TimesheetPath("2021-01-28.md", opts)
	filePath := options.TimesheetPath("2021-01-29.md", opts)
	// filePath := "/Users/bknight/Desktop/scratchpad/times-txt/2021/2021-08/2021-08-03.md"
	// filePath := "/Users/bknight/Desktop/scratchpad/times-txt/2021-09-07.md"

	timesheet := timesheet.ParseFile(filePath, opts)

	summary.PrintLongSummary(timesheet, opts)
}

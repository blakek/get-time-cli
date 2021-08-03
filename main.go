package main

import (
	"fmt"

	"github.com/blakek/get-time-cli/internal/options"
	"github.com/blakek/get-time-cli/internal/timesheet"
)

func main() {
	opts := options.FetchOptions()
	filePath := options.TimesheetPath("2021-01-28.md", opts)
	times := timesheet.ParseFile(filePath, opts)

	for _, entry := range times {
		fmt.Println(entry)
	}
}

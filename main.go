package main

import (
	"fmt"

	"github.com/blakek/get-time-cli/internal/options"
	"github.com/blakek/get-time-cli/internal/timesheet"
)

func main() {
	options := &options.Options{
		Granularity: 4,
	}

	template := timesheet.Create(7, 15, options)

	fmt.Println(template)
}

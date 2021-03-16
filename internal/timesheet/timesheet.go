package timesheet

import (
	"fmt"
	"log"
	"strings"

	"github.com/blakek/get-time-cli/internal/options"
)

func Create(startHour uint, endHour uint, options *options.Options) string {
	if startHour > 23 || endHour > 23 {
		log.Fatal("start and end times must be a valid hour number")
	}

	if endHour <= startHour {
		log.Fatal("end time must be greater than start time")
	}

	var template strings.Builder

	emptyTimeEntry := strings.Repeat("-", options.Granularity)
	timeEntryColumn := fmt.Sprintf(" %s |", emptyTimeEntry)
	timeEntries := strings.Repeat(timeEntryColumn, int(endHour-startHour+1))

	// create header row
	template.WriteString(fmt.Sprintf("| %-*s |", options.Granularity, "Name"))

	for hour := startHour; hour <= endHour; hour++ {
		template.WriteString(fmt.Sprintf(" %-*d |", options.Granularity, hour))
	}

	template.WriteString("\n")

	// create header separator row
	template.WriteString(fmt.Sprintf("|%s", timeEntryColumn))
	template.WriteString(timeEntries)
	template.WriteString("\n")

	// create empty time entry row
	template.WriteString(fmt.Sprintf("| %-*s |", options.Granularity, "-"))
	template.WriteString(timeEntries)
	template.WriteString("\n")

	return template.String()
}

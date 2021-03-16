package main

import (
	"fmt"
	"strings"
)

const (
	// How many sections to break each hour into. Default is 4 for 15 minutes (i.e. 60 / 15).
	granularity = 4
	// The character to use when there isn't a time entry
	timeNoEntry = "-"
)

func createTemplate(startHour uint, endHour uint) string {
	if startHour < 0 || startHour > 23 || endHour > 23 {
		panic("start and end times must be a valid hour number")
	}

	if endHour <= startHour {
		panic("end time must be greater than start time")
	}

	var template strings.Builder

	emptyTimeEntry := strings.Repeat(timeNoEntry, granularity)
	timeEntryColumn := fmt.Sprintf(" %s |", emptyTimeEntry)
	timeEntries := strings.Repeat(timeEntryColumn, int(endHour-startHour+1))

	// create header row
	template.WriteString(fmt.Sprintf("| %-*s |", granularity, "Name"))

	for hour := startHour; hour <= endHour; hour++ {
		template.WriteString(fmt.Sprintf(" %-*d |", granularity, hour))
	}

	template.WriteString("\n")

	// create header separator row
	template.WriteString(fmt.Sprintf("|%s", timeEntryColumn))
	template.WriteString(timeEntries)
	template.WriteString("\n")

	// create empty time entry row
	template.WriteString(fmt.Sprintf("| %-*s |", granularity, "-"))
	template.WriteString(timeEntries)
	template.WriteString("\n")

	return template.String()
}

func main() {
	template := createTemplate(8, 15)
	fmt.Println(template)
}

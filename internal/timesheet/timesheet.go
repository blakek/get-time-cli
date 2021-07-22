package timesheet

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/blakek/get-time-cli/internal/options"
)

type TimeEntry struct {
	Name string
	Time float32
}

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

func ParseFile(filePath string, options *options.Options) []TimeEntry {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	timeEntries := []TimeEntry{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		timeEntry := TimeEntry{}

		line := scanner.Text()
		lineParts := strings.SplitN(line, "|", 3)

		if len(lineParts) != 3 {
			continue
		}

		fmt.Printf("%v\n", lineParts)

		timeEntry.Name = strings.TrimSpace(lineParts[1])
		timeEntry.Time = float32(strings.Count(lineParts[2], "x")) / float32(options.Granularity)
		timeEntries = append(timeEntries, timeEntry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return timeEntries
}

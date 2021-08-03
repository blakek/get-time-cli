package timesheet

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/blakek/get-time-cli/internal/options"
)

type TimeEntry struct {
	Name       string
	noteNumber string
	Notes      string
	Time       float32
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

func ParseFile(filePath string, options *options.Options) []*TimeEntry {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	noteNumberRegex := regexp.MustCompile(`\[\^(\d+)\](?::\s*(.*))?`)
	skipEntryRegex := regexp.MustCompile(`^\s*(?: -+|Name)\s*$`)

	timeEntries := []*TimeEntry{}
	hasNotes := false
	hasTimeEntries := false

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.SplitN(line, "|", 3)
		timeEntry := TimeEntry{}

		isTimeEntryLine := len(lineParts) == 3 && !skipEntryRegex.MatchString(lineParts[1])
		noteNumberMatches := noteNumberRegex.FindStringSubmatch(line)
		hasNoteNumber := len(noteNumberMatches) > 0
		isNoteLine := !isTimeEntryLine && len(noteNumberMatches) > 2

		// Add note text to matching entry
		if isNoteLine {
			noteNumber := noteNumberMatches[1]
			noteText := noteNumberMatches[2]

			for _, entry := range timeEntries {
				if entry.noteNumber == noteNumber {
					entry.Notes = noteText
					break
				}
			}
		}

		if !isTimeEntryLine {
			// Stop reading the file if all time entries and notes have been read
			if hasNotes && hasTimeEntries {
				return timeEntries
			}

			continue
		}

		timeEntry.Name = strings.TrimSpace(lineParts[1])
		timeEntry.Time = float32(strings.Count(lineParts[2], "x")) / float32(options.Granularity)

		if hasNoteNumber {
			timeEntry.noteNumber = noteNumberMatches[1]
		}

		timeEntries = append(timeEntries, &timeEntry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return timeEntries
}

package timesheet

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/blakek/get-time-cli/internal/options"
)

type TimesheetEntry struct {
	Name       string
	noteNumber string
	Notes      string
	Time       float32
}

type Timesheet struct {
	Entries        []*TimesheetEntry
	CompletionTime time.Time
	IsCompleted    bool
	TimeRemaining  time.Duration
	TimeWorked     time.Duration
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

func ParseFile(filePath string, options *options.Options) *Timesheet {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	noteNumberRegex := regexp.MustCompile(`\[\^(\d+)\](?::\s*(.*))?`)
	skipEntryRegex := regexp.MustCompile(`^\s*(?: -+|Name)\s*$`)

	timeEntries := []*TimesheetEntry{}
	hasNotes := false
	hasTimeEntries := false

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.SplitN(line, "|", 3)
		timeEntry := TimesheetEntry{}

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
				}
			}
		}

		if !isTimeEntryLine {
			// Stop reading the file if all time entries and notes have been read
			if hasNotes && hasTimeEntries {
				return getTimesheetFromEntries(timeEntries, options)
			}

			continue
		}

		timeEntry.Name = strings.TrimSpace(lineParts[1])
		timeEntry.Time = float32(strings.Count(lineParts[2], "x")) / float32(options.Granularity)

		if hasNoteNumber {
			timeEntry.noteNumber = noteNumberMatches[1]
			timeEntry.Name = strings.TrimSuffix(timeEntry.Name, " [^"+timeEntry.noteNumber+"]")
		}

		timeEntries = append(timeEntries, &timeEntry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return getTimesheetFromEntries(timeEntries, options)
}

func GetCompletionTime(timeRemaining time.Duration, options *options.Options) time.Time {
	granularity := time.Hour / time.Duration(options.Granularity)
	return time.Now().Add(timeRemaining).Round(granularity)
}

func GetPath(timesheetName string, options *options.Options) string {
	return path.Join(options.TimesheetDirectory, timesheetName)
}

func GetPathForDate(date time.Time, options *options.Options) string {
	return GetPath(
		date.Format(options.TimesheetNameFormat),
		options,
	)
}

func GetTimeRemaining(timeWorked time.Duration, options *options.Options) time.Duration {
	hoursPerDay := getDurationFromHours(options.HoursPerDay)
	timeRemaining := hoursPerDay - timeWorked

	if timeRemaining > 0 {
		return timeRemaining
	}

	return 0
}

func GetTimeWorked(times []*TimesheetEntry) time.Duration {
	var hoursWorked float32 = 0

	for _, timeEntry := range times {
		hoursWorked += timeEntry.Time
	}

	return getDurationFromHours(hoursWorked)
}

// This is a utility to convert a float of hours to a `time.Duration`.
// `time.Duration` doesn't accept floats, so the duration needs to be converted
// to minutes instead of fractions of hours
func getDurationFromHours(hours float32) time.Duration {
	return time.Duration(hours*60) * time.Minute
}

func getTimesheetFromEntries(entries []*TimesheetEntry, options *options.Options) *Timesheet {
	timeWorked := GetTimeWorked(entries)
	timeRemaining := GetTimeRemaining(timeWorked, options)
	completionTime := GetCompletionTime(timeRemaining, options)

	return &Timesheet{
		CompletionTime: completionTime,
		Entries:        entries,
		IsCompleted:    timeRemaining <= 0,
		TimeWorked:     timeWorked,
		TimeRemaining:  timeRemaining,
	}
}

package options

import (
	"os"
	"path"
)

type Options struct {
	// How many sections to break each hour into. Default is 4 for 15 minutes
	// (i.e. 60 / 15).
	Granularity int
	// How many hours are worked per day (for estimating ending time)
	HoursPerDay float32
	// Location where timesheets are stored
	TimesheetDirectory string
	// Time format string for naming timesheets
	TimesheetNameFormat string
}

func FetchOptions() *Options {
	configDir := ensureConfigDir()

	defaultOptions := &Options{
		Granularity:         4,
		HoursPerDay:         8,
		TimesheetDirectory:  configDir,
		TimesheetNameFormat: "2006-01-02.md",
	}

	return defaultOptions
}

func ensureConfigDir() string {
	configDir := getConfigDir()
	_, error := os.Stat(configDir)

	if os.IsNotExist(error) {
		os.Mkdir(configDir, 0775)
	}

	return configDir
}

func getConfigDir() string {
	userHomeDir, _ := os.UserHomeDir()
	configDir := path.Join(userHomeDir, ".get-time")
	return configDir
}

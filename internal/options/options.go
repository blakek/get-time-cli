package options

import "path"

type Options struct {
	// How many sections to break each hour into. Default is 4 for 15 minutes
	// (i.e. 60 / 15).
	Granularity int
	// Location where timesheets are stored
	TimesheetDirectory string
}

func TimesheetPath(timesheetName string, options *Options) string {
	return path.Join(options.TimesheetDirectory, timesheetName)
}

func FetchOptions() *Options {
	defaultOptions := &Options{
		Granularity:        4,
		TimesheetDirectory: "times-test",
	}

	return defaultOptions
}

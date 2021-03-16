package options

type Options struct {
	// How many sections to break each hour into. Default is 4 for 15 minutes
	// (i.e. 60 / 15).
	Granularity int
}

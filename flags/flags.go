package flags

import (
	flag "github.com/spf13/pflag"
)

var (
	LogFile = flag.StringP("log", "l", "/logs/User.log", "Log Location")
	Verbose = flag.BoolP("verbose", "v", false, "Outputs the log to the screen")
)

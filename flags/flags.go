package flags

import (
	flag "github.com/spf13/pflag"
)

var (
	DatabaseFile = flag.StringP("database", "d", "/config/User.db", "Database Location")
	LogFile      = flag.StringP("log", "l", "/logs/User.log", "Log Location")
	Verbose      = flag.BoolP("verbose", "v", false, "Outputs the log to the screen")
)

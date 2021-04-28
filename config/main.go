package config

type cliConfig struct {
	Debug   bool
	Verbose bool
}

var CliConfig = cliConfig{false, false}

func SetCliConfig(debug bool, verbose bool) {
	CliConfig = cliConfig{debug, verbose}
}

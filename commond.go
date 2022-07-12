package commond

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Debug flag parsing
var Debug = os.Getenv("DEBUGFLAGS") != ""

// Your custom Config struct needs a Flags() method. It could be quick returning an array of Flag.
type Config interface {
	Flags() IFlags
}

// Advanced: custom IFlags to bypass BaseConfig
type IFlags interface {
	GetFlags() *flag.FlagSet
}

// App wraps your customized Config and provides Run(config) function.
type App struct {
	Config Config
}

// HasConfigFile if your config type implements ConfigFile method, reads config file before flags (toml-then-flag)
// For flag-then-toml parse, call ReadConfigFile yourself (no ConfigFile method)
type HasConfigFile interface {
	ConfigFile() string
}

// New is alias for Parse
var New = Parse

// Run app with your config
func (a App) Run(mainfn func(c Config) error) error {
	if mainfn == nil {
		panic("no main function given")
	}
	return mainfn(a.Config)
}

// LogFlag is the log flag int for the log package
var LogFlag = func() int {
	if Debug {
		return log.Lshortfile
	}
	return 0
}()

// Parse a config into App. Given config should be 'sane default values'.
// See HasConfigFile and IFlags for advanced customization possibilities
func Parse(config Config, cliargs ...string) (*App, error) {
	if cliargs != nil && len(cliargs) != 1 {
		return nil, fmt.Errorf("too many args")
	}
	if cliargs == nil {
		cliargs = os.Args[1:]
	}
	log.SetFlags(LogFlag)
	if config == nil {
		return nil, fmt.Errorf("no config no app")
	}
	fs := config.Flags().GetFlags()
	err := fs.Parse(cliargs)
	_, hasConfigFile := config.(HasConfigFile)
	if hasConfigFile {
		configFile := config.(HasConfigFile).ConfigFile()
		if err := ReadConfigFile(configFile, config, true); err != nil {
			if os.IsNotExist(err) {
				err = WriteConfig(configFile, config)
			}
			return nil, err
		}
	}
	return &App{Config: config}, err
}

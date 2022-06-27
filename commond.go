package commond

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Config interface {
	Flags() Flags
}

type Flag struct {
	Name        string
	Shortname   string `json:"omitempty"`
	Description string
	Ptr         any
	Default     any
}
type Flags []Flag

func (f Flags) GetFlags() *flag.FlagSet {
	errorHandling := flag.ExitOnError
	cmdname := ""
	flags := flag.NewFlagSet(cmdname, errorHandling)
	if len(f) == 0 {
		panic("zero length flags")
	}
	if Debug {
		log.Printf("got %d flags: %T", len(f), f[0].Default)
		log.Printf("got %d flags: %T", len(f), f[1].Default)
	}
	for i, flaginfo := range f {
		switch flaginfo.Default.(type) {
		case string:
			if Debug {
				log.Println("is string", flaginfo)
			}
			flags.StringVar(flaginfo.Ptr.(*string), flaginfo.Name, flaginfo.Default.(string), flaginfo.Description)
		case int:
			if Debug {
				log.Println("is int", flaginfo)
			}
			flags.IntVar(flaginfo.Ptr.(*int), flaginfo.Name, flaginfo.Default.(int), flaginfo.Description)
		case bool:
			if Debug {
				log.Println("is bool", flaginfo)
			}
			flags.BoolVar(flaginfo.Ptr.(*bool), flaginfo.Name, flaginfo.Default.(bool), flaginfo.Description)
		case time.Duration:
			if Debug {
				log.Println("is time.Duration", flaginfo)
			}
			flags.DurationVar(flaginfo.Ptr.(*time.Duration), flaginfo.Name, flaginfo.Default.(time.Duration), flaginfo.Description)
		default:
			panic(fmt.Sprintf("unhandled flag #%d type: %T", i, flaginfo.Default))
		}
	}

	return flags
}

var Debug = os.Getenv("DEBUGFLAGS") != ""
var LogFlag = func() int {
	if Debug {
		return log.Lshortfile
	}
	return 0
}()

type App struct {
	Config Config
}

func (a App) Run(mainfn func(c Config) error) error {
	if mainfn == nil {
		panic("no main function given")
	}
	return mainfn(a.Config)
}

var New = Parse

func Parse(config Config, cliargs ...[]string) (*App, error) {
	if cliargs != nil && len(cliargs) != 1 {
		return nil, fmt.Errorf("too many args")
	}
	if cliargs == nil {
		cliargs = [][]string{os.Args[1:]}
	}
	log.SetFlags(LogFlag)
	if config == nil {
		config = DefaultBaseConfig()
	}
	fs := config.Flags().GetFlags()
	err := fs.Parse(cliargs[0])
	return &App{Config: config}, err
}

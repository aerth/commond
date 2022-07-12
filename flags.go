package commond

import (
	"flag"
	"fmt"
	"log"
	"time"
)

// Flag represents everything needed to produce a command line flag.
type Flag struct {
	Name        string
	Shortname   string `json:"omitempty"`
	Description string
	Ptr         any
	Default     any
}

// Flags is the default IFlags implementation
type Flags []Flag

// GetFlags creates a CLI flagset from a typed config.
func (f Flags) GetFlags() *flag.FlagSet {
	errorHandling := flag.ExitOnError
	cmdname := ""
	flags := flag.NewFlagSet(cmdname, errorHandling)
	if len(f) == 0 {
		panic("zero length flags")
	}
	for i, flaginfo := range f {
		// TODO implement more types?
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

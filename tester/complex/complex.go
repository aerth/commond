package main

import (
	"fmt"
	"log"
	"os"
	"time"

	commond "example.com/m"
)

type MyConfig struct {
	Motd     string
	Duration time.Duration
}

func (m *MyConfig) Flags() commond.Flags {
	return commond.Flags{commond.Flag{Name: "motd", Default: m.Motd, Description: "message of the day, change me!", Ptr: &m.Motd}}
}

func main() {
	conf := &MyConfig{
		Motd: "Hello",
	}

	// end advanced preflag
	app, err := commond.New(conf) // flag args if preflag
	if err != nil {
		log.Fatalln("couldnt parse commond:", err)
	}

	app.Run(func(c commond.Config) error {
		return mainfn(c.(*MyConfig))
	})
}

func mainfn(config *MyConfig) error {
	fmt.Fprintf(os.Stderr, "%s\n", config.Motd)
	return nil
}

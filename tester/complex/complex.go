package main

import (
	"fmt"
	"log"
	"os"

	commond "github.com/aerth/commond"
)

const configFile = "/tmp/fake.toml"

type myConfig struct {
	Motd string
	Foo  string
}

func (m myConfig) ConfigFile() string {
	log.Println("reading config file:", configFile)
	return configFile
}

func (m *myConfig) Flags() commond.IFlags {
	return commond.Flags{commond.Flag{Name: "motd", Default: m.Motd, Description: "message of the day, change me!", Ptr: &m.Motd}}
}

func main() {
	// end advanced preflag
	app, err := commond.New(&myConfig{
		Motd: "Hello Default",
		Foo:  "edit /tmp/fake.toml right now!",
	}) // flag args if preflag
	if err != nil {
		log.Fatalln("couldnt parse commond:", err)
	}

	// small wrapper
	app.Run(func(c commond.Config) error {
		return realmain(c.(*myConfig))
	})
}

// this is the real main func
func realmain(config *myConfig) error {
	fmt.Fprintf(os.Stderr, "%s\n", config.Motd)
	fmt.Fprintf(os.Stderr, "%s\n", config.Foo)
	return nil
}

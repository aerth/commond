package main

import (
	"fmt"
	"log"
	"os"

	commond "example.com/m"
)

func main() {
	app, err := commond.New(nil)
	if err != nil {
		log.Fatalln("couldnt parse commond:", err)
	}
	app.Run(mainfn)
}

func mainfn(config commond.Config) error {
	isDebug := config.(*commond.BaseConfig).Debug
	fmt.Fprintf(os.Stderr, "works with no config no flags, debug=%v\n", isDebug)
	return nil
}

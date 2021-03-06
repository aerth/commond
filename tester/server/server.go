package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	commond "github.com/aerth/commond"
)

type MyConfig struct {
	Addr    string
	Timeout time.Duration
	Motd    string
	conf    string
}

func (m MyConfig) ConfigFile() string {
	log.Println("reading config file:", m.conf)
	return m.conf
}

func (m *MyConfig) Flags() commond.IFlags {
	return commond.Flags{
		commond.Flag{Name: "addr", Default: m.Addr, Description: "address to listen", Ptr: &m.Addr},
		commond.Flag{Name: "timeout", Default: m.Timeout, Description: "timeout for http connectiong", Ptr: &m.Timeout},
		commond.Flag{Name: "motd", Default: m.Motd, Description: "hello message", Ptr: &m.Motd},
		commond.Flag{Name: "config", Default: "/tmp/fake.toml", Description: "path to config file", Ptr: &m.conf},
	}
}

func main() {
	conf := &MyConfig{
		Addr:    "127.0.0.1:8080",
		Timeout: time.Second * 3,
		Motd:    "unset motd flag",
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

type Server struct {
	config *MyConfig
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", s.config.Motd)
}
func newServer(config *MyConfig) http.Handler {
	return Server{
		config: config,
	}
}
func mainfn(config *MyConfig) error {
	server := newServer(config)
	go log.Println("listening:", config.Addr)
	// return http.ListenAndServe(config.Addr, server)
	httpserver := &http.Server{
		Addr:              config.Addr,
		Handler:           server,
		ReadTimeout:       config.Timeout,
		WriteTimeout:      config.Timeout,
		ReadHeaderTimeout: config.Timeout,
		IdleTimeout:       config.Timeout,
		ErrorLog:          log.New(os.Stderr, "http>", log.Lshortfile|log.LstdFlags),
	}
	return httpserver.ListenAndServe()
}

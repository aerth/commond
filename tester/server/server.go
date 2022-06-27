package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	commond "example.com/m"
)

type MyConfig struct {
	Addr    string
	Timeout time.Duration
	Motd    string
}

func (m *MyConfig) Flags() commond.Flags {
	return commond.Flags{
		commond.Flag{Name: "addr", Default: m.Addr, Description: "address to listen", Ptr: &m.Addr},
		commond.Flag{Name: "timeout", Default: m.Timeout, Description: "timeout for http connectiong", Ptr: &m.Timeout},
		commond.Flag{Name: "motd", Default: m.Motd, Description: "hello message", Ptr: &m.Motd},
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
	return http.ListenAndServe(config.Addr, server)
}

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
}

func (m *MyConfig) Flags() commond.Flags {
	return commond.Flags{
		commond.Flag{Name: "addr", Default: m.Addr, Description: "address to listen", Ptr: &m.Addr},
		commond.Flag{Name: "timeout", Default: m.Timeout, Description: "timeout for http connectiong", Ptr: &m.Timeout},
	}
}

func main() {
	conf := &MyConfig{
		Addr:    "127.0.0.1:8080",
		Timeout: time.Second * 3,
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

type Server struct{}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi\n")
}
func newServer(config *MyConfig) http.Handler {
	return Server{
		//
	}
}
func mainfn(config *MyConfig) error {
	server := newServer(config)
	return http.ListenAndServe(config.Addr, server)
}

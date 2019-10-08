package main

import (
	"flag"

	server "github.com/andreylm/guru/pkg/server/v1"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "Assigning the port")
	flag.Parse()
}

func main() {
	srv := server.NewServer(port)
	srv.Start()
}

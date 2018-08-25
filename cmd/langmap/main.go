package main

import (
	"flag"
	"log"

	"github.com/qwwqe/langmap"
)

func main() {
	configPath := flag.String("f", "config.json", "configuration file")
	flag.Parse()

	e := langmap.Engine{Config: &langmap.Config{}}
	e.Config.FromFile(*configPath)

	err := e.Run()
	if err != nil {
		log.Fatal(err)
	}
}

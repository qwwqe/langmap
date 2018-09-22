package main

import (
	"flag"
	"log"

	"github.com/qwwqe/langmap"
)

var opts = struct {
	CreateForeignKeys bool
	CreateIndexes     bool
	CreateTables      bool
	Config            string
}{
	CreateForeignKeys: false,
	CreateIndexes:     false,
	CreateTables:      false,
	Config:            "config.json",
}

func main() {
	flag.BoolVar(&opts.CreateForeignKeys, "create-foreign-keys", opts.CreateForeignKeys, "Create foreign keys on start up")
	flag.BoolVar(&opts.CreateIndexes, "create-indexes", opts.CreateIndexes, "Create table indexes on start up")
	flag.BoolVar(&opts.CreateTables, "create-tables", opts.CreateTables, "Create tables on start up")
	flag.StringVar(&opts.Config, "config", opts.Config, "Set configuration file to use")
	flag.Parse()

	e := langmap.Engine{Config: &langmap.Config{}}
	e.Config.FromFile(opts.Config)

	err := e.Run(
		opts.CreateTables,
		opts.CreateIndexes,
		opts.CreateForeignKeys,
	)
	if err != nil {
		log.Fatal(err)
	}
}

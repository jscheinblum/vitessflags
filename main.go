package main

import (
	"flag"
	"os"
	"vitess.io/vitess/go/vt/vitessdriver"
)

func main() {
	flag.Parse()

	db, err := vitessdriver.Open("vtgate.example.com", "@master")
	if err != nil {
		os.Exit(1)
	}
	err = db.Close()
	if err != nil {
		os.Exit(1)
	}
}

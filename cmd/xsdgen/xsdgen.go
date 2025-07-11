package main

import (
	"log"
	"os"

	"github.com/llgoer/go-xml/xsdgen"
)

func main() {
	log.SetFlags(0)
	var cfg xsdgen.Config
	cfg.Option(xsdgen.DefaultOptions...)
	cfg.Option(xsdgen.LogOutput(log.New(os.Stderr, "", 0)))

	if err := cfg.GenCLI(os.Args[1:]...); err != nil {
		log.Fatal(err)
	}
}

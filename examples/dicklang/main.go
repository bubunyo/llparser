package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	parser "github.com/romshark/llparser"
)

var flagFilePath = flag.String(
	"src",
	"./dicks.txt",
	"source file path",
)
var flagPrintParseTree = flag.Bool(
	"ptree",
	false,
	"prints the parse-tree only",
)

func main() {
	flag.Parse()
	mod, err := Parse(*flagFilePath)
	if err != nil {
		log.Fatal("ERR: ", err)
	}

	if *flagPrintParseTree {
		// Print the parse-tree only
		_, err := parser.PrintFragment(mod.Frag, os.Stdout, "  ")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// Print all parsed dicks
	fmt.Printf("%d dicks parsed:\n", len(mod.Dicks))
	for ix, dick := range mod.Dicks {
		fmt.Printf(" %d: %s\n", ix, string(dick.Frag.Src()))
	}
}

package main

import (
	"flag"

	"github.com/fogleman/mol/mol"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		return
	}
	molecule, err := mol.ParseFile(args[0])
	if err != nil {
		panic(err)
	}
	molecule.Render(args[0] + ".png")
}

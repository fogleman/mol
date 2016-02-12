package main

import (
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/fogleman/mol/mol"
)

func main() {
	flag.Parse()
	args := flag.Args()
	for _, arg := range args {
		fmt.Println(arg)
		molecule, err := mol.ParseFile(arg)
		if err != nil {
			panic(err)
		}
		molecule.Render(outputFilename(arg))
	}
}

func outputFilename(x string) string {
	_, file := path.Split(x)
	i := strings.LastIndex(file, ".")
	if i >= 0 {
		file = file[:i]
	}
	return file + ".png"
}

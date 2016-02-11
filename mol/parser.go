package mol

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ParseFile(path string) (Molecule, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Molecule{}, err
	}
	return ParseString(string(data)), nil
}

func ParseString(data string) Molecule {
	lines := strings.Split(data, "\n")
	fields := strings.Fields(lines[3])
	natoms := parseInt(fields[0])
	nbonds := parseInt(fields[1])
	atoms := make([]Atom, natoms)
	for i := 0; i < natoms; i++ {
		fields := strings.Fields(lines[i+4])
		x := parseFloat(fields[0])
		y := parseFloat(fields[1])
		z := parseFloat(fields[2])
		symbol := fields[3]
		atoms[i] = Atom{x, y, z, symbol}
	}
	bonds := make([]Bond, nbonds)
	for i := 0; i < nbonds; i++ {
		fields := strings.Fields(lines[i+4+natoms])
		a := parseInt(fields[0]) - 1
		b := parseInt(fields[1]) - 1
		t := parseInt(fields[2])
		bonds[i] = Bond{a, b, t}
	}
	return Molecule{atoms, bonds}
}

func parseInt(x string) int {
	value, _ := strconv.ParseInt(x, 0, 0)
	return int(value)
}

func parseFloat(x string) float64 {
	value, _ := strconv.ParseFloat(x, 64)
	return value
}

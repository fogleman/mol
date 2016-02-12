package mol

import "github.com/fogleman/ln/ln"

func Render(path string, molecule *Molecule) {
	scene := ln.Scene{}

	nodes := make([]ln.Vector, len(molecule.Atoms))
	for i, atom := range molecule.Atoms {
		nodes[i] = ln.Vector{atom.X, atom.Y, atom.Z}
	}

	eye := CameraPosition(nodes)
	center := CameraCenter(nodes)
	fov := CameraFOV(nodes, eye, center)
	up := ln.Vector{0, 0, 1}

	for i, node := range nodes {
		atom := molecule.Atoms[i]
		radius := float64(AtomicRadii[atom.Symbol]) / 100
		scene.Add(ln.NewOutlineSphere(eye, up, node, radius*0.5))
	}

	for _, bond := range molecule.Bonds {
		v0 := nodes[bond.I]
		v1 := nodes[bond.J]
		r := float64(bond.Type) / 16
		scene.Add(ln.NewTransformedOutlineCylinder(eye, up, v0, v1, r))
	}

	width := 1024.0
	height := 1024.0
	paths := scene.Render(eye, center, up, width, height, fov, 0.1, 100, 0.01)
	paths.WriteToPNG(path, width, height)
}

var AtomicRadii = map[string]int{
	"H":  53,
	"He": 31,
	"Li": 167,
	"Be": 112,
	"B":  87,
	"C":  67,
	"N":  56,
	"O":  48,
	"F":  42,
	"Ne": 38,
	"Na": 190,
	"Mg": 145,
	"Al": 118,
	"Si": 111,
	"P":  98,
	"S":  88,
	"Cl": 79,
	"Ar": 71,
	"K":  243,
	"Ca": 194,
	"Sc": 184,
	"Ti": 176,
	"V":  171,
	"Cr": 166,
	"Mn": 161,
	"Fe": 156,
	"Co": 152,
	"Ni": 149,
	"Cu": 145,
	"Zn": 142,
	"Ga": 136,
	"Ge": 125,
	"As": 114,
	"Se": 103,
	"Br": 94,
	"Kr": 88,
	"Rb": 265,
	"Sr": 219,
	"Y":  212,
	"Zr": 206,
	"Nb": 198,
	"Mo": 190,
	"Tc": 183,
	"Ru": 178,
	"Rh": 173,
	"Pd": 169,
	"Ag": 165,
	"Cd": 161,
	"In": 156,
	"Sn": 145,
	"Sb": 133,
	"Te": 123,
	"I":  115,
	"Xe": 108,
	"Cs": 298,
	"Ba": 253,
	"Pr": 247,
	"Nd": 206,
	"Pm": 205,
	"Sm": 238,
	"Eu": 231,
	"Gd": 233,
	"Tb": 225,
	"Dy": 228,
	"Er": 226,
	"Tm": 222,
	"Yb": 222,
	"Lu": 217,
	"Hf": 208,
	"Ta": 200,
	"W":  193,
	"Re": 188,
	"Os": 185,
	"Ir": 180,
	"Pt": 177,
	"Au": 174,
	"Hg": 171,
	"Tl": 156,
	"Pb": 154,
	"Bi": 143,
	"Po": 135,
	"Rn": 120,
}

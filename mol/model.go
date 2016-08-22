package mol

import "github.com/fogleman/ln/ln"

type Atom struct {
	X, Y, Z float64
	Symbol  string
}

type Bond struct {
	I, J int
	Type int
}

type Molecule struct {
	Atoms []Atom
	Bonds []Bond
}

type Sphere struct {
	Center Vector
	Radius float64
	Symbol string
}

type Cylinder struct {
	A, B   Vector
	Radius float64
	Type   int
}

func (m *Molecule) Solids() ([]Sphere, []Cylinder) {
	spheres := make([]Sphere, len(m.Atoms))
	cylinders := make([]Cylinder, len(m.Bonds))

	for i, atom := range m.Atoms {
		center := Vector{atom.X, atom.Y, atom.Z}
		radius := float64(AtomicRadii[atom.Symbol]) / 100
		spheres[i] = Sphere{center, radius, atom.Symbol}
	}

	for i, bond := range m.Bonds {
		s0 := spheres[bond.I]
		s1 := spheres[bond.J]
		radius := float64(bond.Type) / 16
		cylinders[i] = Cylinder{s0.Center, s1.Center, radius, bond.Type}
	}

	return spheres, cylinders
}

func (m *Molecule) Camera() Camera {
	points := make([]Vector, len(m.Atoms))
	for i, atom := range m.Atoms {
		points[i] = Vector{atom.X, atom.Y, atom.Z}
	}
	return MakeCamera(points)
}

func (m *Molecule) Paths(width, height float64) ln.Paths {
	scene := ln.Scene{}

	camera := m.Camera()
	eye := camera.Eye.ln()
	center := camera.Center.ln()
	up := camera.Up.ln()
	fovy := camera.Fovy

	spheres, cylinders := m.Solids()

	for _, s := range spheres {
		scene.Add(ln.NewOutlineSphere(eye, up, s.Center.ln(), s.Radius*0.5))
	}

	for _, c := range cylinders {
		scene.Add(ln.NewTransformedOutlineCylinder(eye, up, c.A.ln(), c.B.ln(), c.Radius))
	}

	return scene.Render(eye, center, up, width, height, fovy, 0.1, 100, 0.01)
}

func (m *Molecule) Render(path string, width, height float64) {
	paths := m.Paths(width, height)
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

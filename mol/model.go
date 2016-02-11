package mol

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

func (m *Molecule) Render(path string) {
	Render(path, m)
}

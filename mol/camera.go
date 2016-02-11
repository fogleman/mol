package mol

import "github.com/fogleman/ln/ln"

func CameraCenter(nodes []ln.Vector) ln.Vector {
	box := ln.BoxForVectors(nodes)
	return box.Center()
}

func CameraPosition(nodes []ln.Vector) ln.Vector {
	var result ln.Vector
	best := 1e9
	for i := 0; i < 1000; i++ {
		eye := ln.RandomUnitVector().MulScalar(20)
		score := cameraScore(nodes, eye)
		if score < best {
			best = score
			result = eye
		}
	}
	return result
}

func cameraScore(nodes []ln.Vector, eye ln.Vector) float64 {
	var result float64
	for _, n1 := range nodes {
		d1 := n1.Sub(eye).Normalize()
		for _, n2 := range nodes {
			d2 := n2.Sub(eye).Normalize()
			a := d1.Dot(d2)
			result += a * a
		}
	}
	return result
}

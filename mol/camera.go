package mol

import (
	"math"

	"github.com/fogleman/ln/ln"
)

func CameraFOV(points []ln.Vector, eye, center ln.Vector) float64 {
	var result float64
	c := center.Sub(eye).Normalize()
	for _, point := range points {
		d := point.Sub(eye).Normalize()
		a := math.Acos(d.Dot(c))
		result = math.Max(result, a)
	}
	return ln.Degrees(result * 2 * 1.2)
}

func CameraPosition(points []ln.Vector) ln.Vector {
	var result ln.Vector
	best := 1e9
	for i := 0; i < 1000; i++ {
		eye := ln.RandomUnitVector().MulScalar(50)
		score := cameraScore(points, eye)
		if score < best {
			best = score
			result = eye
		}
	}
	return result
}

func cameraScore(points []ln.Vector, eye ln.Vector) float64 {
	var result float64
	for _, p1 := range points {
		d1 := p1.Sub(eye).Normalize()
		for _, p2 := range points {
			d2 := p2.Sub(eye).Normalize()
			a := d1.Dot(d2)
			result += a * a
		}
	}
	return result
}

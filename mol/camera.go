package mol

import "math"

type Camera struct {
	Eye    Vector
	Center Vector
	Up     Vector
	Fovy   float64
}

func MakeCamera(points []Vector) Camera {
	up := Vector{0, 0, 1}

	var center Vector
	for _, point := range points {
		center = center.Add(point)
	}
	center = center.DivScalar(float64(len(points)))

	var eye Vector
	best := 1e9
	for i := 0; i < 1000; i++ {
		v := RandomUnitVector().MulScalar(50)
		score := cameraScore(points, v)
		if score < best {
			best = score
			eye = v
		}
	}

	var fovy float64
	c := center.Sub(eye).Normalize()
	for _, point := range points {
		d := point.Sub(eye).Normalize()
		a := Degrees(math.Acos(d.Dot(c)) * 2 * 1.2)
		fovy = math.Max(fovy, a)
	}

	return Camera{eye, center, up, fovy}
}

func cameraScore(points []Vector, eye Vector) float64 {
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

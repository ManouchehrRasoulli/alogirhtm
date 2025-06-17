/*
package kata

import (
	"math"
)

type Point struct {
	x, y float64
}

// Distance from point P to segment AB
func pointToSegmentDistance(px, py, x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1

	if dx == 0 && dy == 0 {
		return math.Hypot(px-x1, py-y1)
	}

	t := ((px-x1)*dx + (py-y1)*dy) / (dx*dx + dy*dy)
	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}

	closestX := x1 + t*dx
	closestY := y1 + t*dy
	return math.Hypot(px-closestX, py-closestY)
}

// Check if a point is inside a polygon using ray casting
func pointInPolygon(p Point, polygon []Point) bool {
	inside := false
	n := len(polygon)
	j := n - 1
	for i := 0; i < n; i++ {
		xi, yi := polygon[i].x, polygon[i].y
		xj, yj := polygon[j].x, polygon[j].y
		if (yi > p.y) != (yj > p.y) &&
			p.x < (xj-xi)*(p.y-yi)/(yj-yi)+xi {
			inside = !inside
		}
		j = i
	}
	return inside
}

func FitBiggestCircle(raw []float64) []float64 {
	polygon := make([]Point, len(raw)/2)
	for i := 0; i < len(raw); i += 2 {
		polygon[i/2] = Point{raw[i], raw[i+1]}
	}

	// Bounding box
	minX, maxX := polygon[0].x, polygon[0].x
	minY, maxY := polygon[0].y, polygon[0].y
	for _, p := range polygon {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	// Grid sampling resolution
	step := 0.003
	best := Point{0, 0}
	maxRadius := -1.0

	for x := minX; x <= maxX; x += step {
		for y := minY; y <= maxY; y += step {
			p := Point{x, y}
			if !pointInPolygon(p, polygon) {
				continue
			}
			minDist := math.MaxFloat64
			for i := 0; i < len(polygon); i++ {
				a := polygon[i]
				b := polygon[(i+1)%len(polygon)]
				d := pointToSegmentDistance(p.x, p.y, a.x, a.y, b.x, b.y)
				if d < minDist {
					minDist = d
				}
			}
			if minDist > maxRadius {
				maxRadius = minDist
				best = p
			}
		}
	}

	return []float64{best.x, best.y, maxRadius}
}
*/

package kata

import (
	"math"
)

type Point struct {
	x float64
	y float64
}

type Line struct {
	p1 Point
	p2 Point
	A  float64
	B  float64
	C  float64
}

type Case struct {
	a1 []Point
	a2 []Point
	l  Line
}

func Slope(p1, p2 Point) float64 {
	return (p2.y - p1.y) / (p2.x - p1.x)
}

func LineFromPoints(p1, p2 Point) Line {
	m := Slope(p1, p2)

	A := m
	B := -1.0
	C := p1.y - m*p1.x

	return Line{A: A, B: B, C: C, p1: p1, p2: p2}
}

func AngleBisector(A, B, C Point) (float64, float64) {
	// Calculate the distances (lengths of sides)
	AB := math.Sqrt(math.Pow(B.x-A.x, 2) + math.Pow(B.y-A.y, 2))
	AC := math.Sqrt(math.Pow(C.x-A.x, 2) + math.Pow(C.y-A.y, 2))

	// Weighted average of the points B and C based on side lengths
	x := (AB*C.x + AC*B.x) / (AB + AC)
	y := (AB*C.y + AC*B.y) / (AB + AC)

	// Calculate the slope of the bisector line (using A and the weighted point)
	slope := Slope(A, Point{x, y})

	if slope == math.Inf(1) || slope == math.Inf(-1) {
		slope = 0
	}

	// Calculate y-intercept
	intercept := A.y - slope*A.x

	return slope, intercept
}

func Intersection(m1, c1, m2, c2 float64) (float64, float64) {
	// Check if the slopes are equal (lines are parallel)
	if m1 == m2 {
		return 0, 0
	}

	// Solve for x
	x := (c2 - c1) / (m1 - m2)

	// Solve for y using either of the line equations
	y := m1*x + c1

	return x, y
}

func DistanceFromLine(x, y float64, l Line) float64 {
	if l.A == math.Inf(1) || l.A == math.Inf(-1) {
		return math.Abs(l.p1.x - x)
	}

	p := Point{x: x, y: y}

	distance := math.Abs(l.A*p.x+l.B*p.y+l.C) / math.Sqrt(l.A*l.A+l.B*l.B)

	return distance
}

func IsSolution(AB, BC, CD, DA Line, x, y, r float64) bool {

	lines := []Line{AB, BC, CD, DA}

	count := 0
	isIntercept := false
	for _, l := range lines {
		if d := DistanceFromLine(x, y, l); math.Abs(r-d) <= 1e-12 { // tangent to line
			count++
		} else if r > d { // intercepts line
			isIntercept = true
		}
	}
	return count == 4 || (count >= 3 && !isIntercept)
}

func FitBiggestCircle(corners []float64) []float64 {
	cx, cy, r := 0.0, 0.0, -1.0
	tmpX, tmpY, tmpR := 0.0, 0.0, -1.0
	m1, c1, m2, c2 := 0.0, 0.0, 0.0, 0.0

	A, B, C, D := Point{x: corners[0], y: corners[1]},
		Point{x: corners[2], y: corners[3]},
		Point{x: corners[4], y: corners[5]},
		Point{x: corners[6], y: corners[7]}

	AB := LineFromPoints(A, B)
	BC := LineFromPoints(B, C)
	CD := LineFromPoints(C, D)
	DA := LineFromPoints(D, A)

	case1 := Case{
		a1: []Point{A, B, D},
		a2: []Point{B, A, C},
		l:  AB,
	}

	case2 := Case{
		a1: []Point{B, A, C},
		a2: []Point{C, B, D},
		l:  BC,
	}

	case3 := Case{
		a1: []Point{C, B, D},
		a2: []Point{D, A, C},
		l:  CD,
	}

	case4 := Case{
		a1: []Point{D, A, C},
		a2: []Point{A, B, D},
		l:  DA,
	}

	cases := []Case{
		case1,
		case2,
		case3,
		case4,
	}

	for _, c := range cases {
		m1, c1 = AngleBisector(c.a1[0], c.a1[1], c.a1[2])
		m2, c2 = AngleBisector(c.a2[0], c.a2[1], c.a2[2])

		tmpX, tmpY = Intersection(m1, c1, m2, c2)

		tmpR = DistanceFromLine(tmpX, tmpY, c.l)

		if IsSolution(AB, BC, CD, DA, tmpX, tmpY, tmpR) && tmpR > r {
			cx, cy, r = tmpX, tmpY, tmpR
		}
	}

	return []float64{cx, cy, r}
}

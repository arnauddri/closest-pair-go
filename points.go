package closest_pair

import (
	"math"
	"strconv"
)

type point struct {
	x uint32
	y uint32
}

type Point interface {
	toString() string
}

func makePoint(x uint32, y uint32) *point {
	A := new(point)
	A.x = x
	A.y = y

	return A
}

func (A *point) toString() string {
	return "(" + string(A.x) + "," + string(A.y) + ")"
}

type pair struct {
	point1   point
	point2   point
	distance float64
}

type Pair interface {
	calcDistance() float64
	toString() string
}

func makePair(A point, B point) *pair {
	P := new(pair)
	P.point1 = A
	P.point2 = B
	P.distance = calcDistance(A, B)

	return P
}

func calcDistance(A point, B point) float64 {
	xdist := A.x - B.x
	ydist := A.y - B.y

	return math.Sqrt(float64(xdist*xdist + ydist*ydist))
}

func (P *pair) toString() string {
	return P.point1.toString() + "-" + P.point2.toString() + "-" + strconv.FormatFloat(P.distance, 'f', -1, 64)
}

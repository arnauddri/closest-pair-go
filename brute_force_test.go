package closest_pair

import (
	"fmt"
	"testing"
)

func TestBruteForce(t *testing.T) {
	Points := make([]point, 10)
	for i := 0; i < 10; i++ {
		a := makePoint(uint32(i), uint32(i*i))

		Points[i] = *a
	}

	A := makePoint(uint32(0), uint32(0))
	B := makePoint(uint32(1), uint32(1))
	P := bruteForce(Points)
	if P.point1 != *A || P.point2 != *B {
		fmt.Println(P.point1, *A, *B)
		t.Error()
	}
}

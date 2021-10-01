package keeper

import (
	"testing"
	"time"
)

func TestCalcD(t *testing.T) {
	A := 100.0
	weight := 1.0
	x := 10.0
	y := 10.0
	now := time.Now()
	D := solveD(A, x, y, weight, weight)
	elapsed := time.Now().Sub(now)
	t.Error(elapsed, D, invarLeft(A, D, x, y), invarRight(A, D, x, y, weight, weight))

	newx := x + 1
	D = 20
	newy := loopCalcNewY(A, D, newx, y, weight, weight)
	t.Error(D, newy, invarLeft(A, D, newx, newy), invarRight(A, D, newx, newy, weight, weight))
}

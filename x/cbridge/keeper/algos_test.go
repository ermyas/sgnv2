package keeper

import (
	"testing"
	"time"
)

func TestCalcD(t *testing.T) {
	A := 2000.0
	m := 1.25
	n := 2 - m
	x := 839571.827415
	y := 777552.260981
	now := time.Now()
	D := solveD(A, x, y, m, n)
	elapsed := time.Now().Sub(now)
	t.Error(elapsed, D, invarLeft(A, D, x, y), invarRight(A, D, x, y, m, n))

	newx := 4505472.157971
	newy := loopCalcNewY(A, D, newx, y, m, n)
	t.Error(D, newy, invarLeft(A, D, newx, newy), invarRight(A, D, newx, newy, m, n))
}

func TestLoopNewY(t *testing.T) {
	A := 2000.0
	m := 1.0
	n := 2 - m
	// x := 839571.827415
	y := 777552.260981
	D := 1.6171276365928054e+06
	newx := 4505472.157971
	t.Error(loopCalcNewY(A, D, newx, y, m, n), solveY(A, D, newx, y))
}

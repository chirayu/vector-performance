package vector

import (
	"math/rand"
	"testing"
)

var seed int64 = 0xCA0541
var nCells int = 1e5

func rows88(nCells int) (map[int]*Cell88, map[int]*Cell88, map[int]*Cell88) {
	one := make(map[int]*Cell88, nCells)
	two := make(map[int]*Cell88, nCells)
	res := make(map[int]*Cell88, nCells)

	rand := rand.New(rand.NewSource(seed))

	for i := 0; i < nCells; i++ {
		one[i] = &Cell88{value: rand.Float64()}
		two[i] = &Cell88{value: rand.Float64()}
	}

	return one, two, res
}

func rows8(nCells int) (map[int]*Cell8, map[int]*Cell8, map[int]*Cell8) {
	one := make(map[int]*Cell8, nCells)
	two := make(map[int]*Cell8, nCells)
	res := make(map[int]*Cell8, nCells)

	rand := rand.New(rand.NewSource(seed))

	for i := 0; i < nCells; i++ {
		one[i] = &Cell8{value: rand.Float64()}
		two[i] = &Cell8{value: rand.Float64()}
	}

	return one, two, res
}

func makeVectors(nCells int) ([]float64, []float64, []float64) {
	one := make([]float64, nCells, nCells)
	two := make([]float64, nCells, nCells)
	res := make([]float64, nCells, nCells)

	rand := rand.New(rand.NewSource(seed))

	for i := 0; i < nCells; i++ {
		one[i] = rand.Float64()
		two[i] = rand.Float64()
	}

	return one, two, res
}

func BenchmarkCell88(b *testing.B) {
	one, two, res := rows88(nCells)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Mul88(one, two, res, nCells)
	}
}

func BenchmarkCell8(b *testing.B) {
	one, two, res := rows8(nCells)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Mul8(one, two, res, nCells)
	}
}

func BenchmarkVector(b *testing.B) {
	one, two, res := makeVectors(nCells)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		MulVector(one, two, res, nCells)
	}
}

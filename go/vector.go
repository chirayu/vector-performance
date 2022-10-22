package vector

type Cell88 struct {
	padding [80]byte // just to simulate what would be real stuff
	value   float64
}

type Cell8 struct {
	value float64
}

func MulVector(one, two, res []float64, nCells int) {
	for i := 0; i < nCells; i++ {
		res[i] = one[i] * two[i]
	}
}

func Mul8(one, two, res map[int]*Cell8, nCells int) {
	for i := 0; i < nCells; i++ {
		res[i] = &Cell8{value: one[i].value * two[i].value}
	}
}

func Mul88(one, two, res map[int]*Cell88, nCells int) {
	for i := 0; i < nCells; i++ {
		res[i] = &Cell88{value: one[i].value * two[i].value}
	}
}

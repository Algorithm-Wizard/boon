package math

type Vectori struct {
	X int
	Y int
}

type Vector struct {
	X float32
	Y float32
}

func Absi(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

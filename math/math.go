package math

type Vector struct {
	X float32
	Y float32
}

func Abs(val float32) float32 {
	if val < 0 {
		return -val
	}
	return val
}

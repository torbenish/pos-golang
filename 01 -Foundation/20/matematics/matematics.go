package matematics

func Sum[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Car struct {
	Brand string
}

func (c Car) Walk() string {
	return "Car walks"
}
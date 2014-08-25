package lemon

type Shaper interface {
	Area() float32
}
type Square struct {
	Side float32
}

func (sq *Square) Area() float32 {
	return sq.Side * sq.Side
}

package shape

type Circle struct {
	radius int
}

func NewCircle(radius int) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}

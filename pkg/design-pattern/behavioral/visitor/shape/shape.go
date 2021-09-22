package shape

type Shape interface {
	GetType() string
	Accept(Visitor)
}
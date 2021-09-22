package main

import (
	"fmt"
	"github.com/jjmengze/mygo/pkg/design-pattern/behavioral/visitor/shape"
)

func main() {
	square := shape.NewSquare(2)
	circle := shape.NewCircle(3)
	rectangle := shape.NewRectangle(2, 3)

	areaCalculator := &shape.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	//middleCoordinates := &middleCoordinates{}
	//square.accept(middleCoordinates)
	//circle.accept(middleCoordinates)
	//rectangle.accept(middleCoordinates)
}

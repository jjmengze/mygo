package shape

import (
	"fmt"
)

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) VisitForSquare(square *Square) {
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) VisitForCircle(circle *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) VisitForRectangle(rectangle *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}


package main

import (
	"fmt"
	"github.com/jjmengze/mygo/pkg/design-pattern/creational/factory"
)

func main() {
	var calculationFactory factory.CalculationFactory
	calculationFactory = factory.PlusCalculationFactory{}
	calculation := calculationFactory.Create()
	calculation.SetCounted(1)
	calculation.SetCounted(2)
	fmt.Println(calculation.Result())

	calculationFactory = factory.MinusCalculationFactory{}
	calculation.SetCounted(2)
	calculation.SetCounted(1)
	fmt.Println(calculation.Result())
}

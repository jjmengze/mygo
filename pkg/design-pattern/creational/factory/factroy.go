package factory

import "github.com/jjmengze/mygo/pkg/design-pattern/creational/factory/calculation"

//CalculationFactory 定義如何建立 Calculation
type CalculationFactory interface {
	Create() calculation.Calculation
}

//PlusCalculationFactory 是實作 CalculationFactory 用來生產 PlusCalculation 的類別
type PlusCalculationFactory struct{}

func (PlusCalculationFactory) Create() calculation.Calculation {
	return &calculation.PlusCalculation{
		Base: &calculation.Base{},
	}
}

//MinusCalculationFactory 是實作 CalculationFactory 用來生產 MinusCalculation 的類別
type MinusCalculationFactory struct{}

func (MinusCalculationFactory) Create() calculation.Calculation {
	return &calculation.MinusCalculation{
		Base: &calculation.Base{},
	}
}

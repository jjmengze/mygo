package calculation

//MinusCalculation 實作 Calculation  interface 並且注重 reset 如何實作 Minus
type MinusCalculation struct {
	*Base
}

//Result 获取结果
func (o MinusCalculation) Result() int {
	return o.beCounted - o.counted
}

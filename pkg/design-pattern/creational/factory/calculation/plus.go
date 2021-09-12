package calculation



//PlusCalculation 實作 Calculation  interface 並且注重 reset 如何實作 Plus
type PlusCalculation struct {
	*Base
}

//Result 获取结果
func (o PlusCalculation) Result() int {
	return o.beCounted + o.counted
}

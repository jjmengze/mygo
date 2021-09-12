package calculation

//Calculation 是一個被封裝的物件行為
type Calculation interface {
	SetBeCounted(int)
	SetCounted(int)
	Result() int
}

//Base 是實作 Calculation interface 的基礎類別，通常會抽出重複的邏輯做為基礎類別，再讓其他類別嵌入 。
type Base struct {
	beCounted int
	counted   int
}

//SetBeCounted 設定 BeCounted
func (o *Base) SetBeCounted(beCounted int) {
	o.beCounted = beCounted
}

//SetCounted 設定 Counted
func (o *Base) SetCounted(counted int) {
	o.counted = counted
}

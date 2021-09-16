#  抽象工廠模式 abstract factory Pattern

抽象實體化物件的過程，需要各自實作不同類型的工廠。

這邊舉兩個例子，第一個例子是傢俱開發商。

我們身為一個家具開發商，負責開發各種不同的傢俱。例如椅子、沙發、桌子。

家具也會有很多不同的風格例如 維也納風、現代風以及藝術風等等。


## 工廠
如果我們用工廠模式會怎麼設計？

```go
//ChairFactory 定義如何建立 chair
type ChairFactory interface {
	Create() chair.Chair
}

//ModernChairFactory 是實作 ChairFactory 用來生產 ModernChair 的類別
type ModernChairFactory struct{}

func (*ModernChairFactory) Create() chair.Chair {
	return &chair.ModernChair{}
}
```

看起來我們用工廠模式順利的完成現代風椅子的工廠，現在還有椅子、桌子等等。
所以我們還要再新增現代風的椅子、桌子等等的工廠

## 抽象工廠
抽象工廠會把`所有需要建立的家具`都寫到抽象工廠的 interface ，實作該interface的工廠必須要能夠生廠所有家具。
```go
type FurnishingFactory interface {
	CreateChair() chair.Chair
	CreateDesk() desk.Desk
	CreateSofa() sofa.Sofa
}
```
例如這個工廠主要生產所有現代風的家具我們可能會這樣設計。
```go
type ModernFurnishingFactory struct {}

var _ FurnishingFactory = &ModernFurnishingFactory{}

func (a ModernFurnishingFactory) CreateChair() chair.Chair {
	return &chair.ModernChair{}
}

func (a ModernFurnishingFactory) CreateDesk() desk.Desk {
	return &desk.ModernDesk{}
}

func (a ModernFurnishingFactory) CreateSofa() sofa.Sofa {
	return &sofa.ModernSofa{}
}
```

工廠需要把所有需要生產的家具，依照特定的風格撰寫生產的方式。以上範例為工廠生產現代風的桌子、椅子、沙發的生產方式。


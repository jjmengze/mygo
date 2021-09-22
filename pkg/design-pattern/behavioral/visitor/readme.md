# 訪問者模式 Visitor Pattern 

在資料結構穩定(甚至固定)時才適合使用，如果任意使用反而會造成(資料結構)擴充困難的結果

>訪問者模式 Visitor Pattern 使你可以再不改變各元素之類別的前提之下，定義作用於這些元素的新操作。
> -- Design Pattern by GoF [大話設計模式 p.434]


## 優點
1. 訪問者模式容易增加新的操作，不會影響元件類別
2. 特定的操作(邏輯行為)可以包裝在訪問者裡

## 缺點
1. 透過訪問者去調用實際元件類別的方法，某種程度上元件類別需要暴露資訊給訪問者
2. 如果有更複雜的訪問者，需要儲存元件狀態的話，元件的狀態資訊就需要放在訪問者加深了複雜度


## example 

假設我們是一個具有不同形狀結構的library的維護者，例如：

- Square
- Circle
- Triangle

每個形狀結構都實現了通用行形狀介面(interface)，一旦有人開始使用這個形狀library後，我們會收到大量新功能請求。

例如以下一個非常簡單的範例：
xxxx同事要求新增一個 GetArea function 到library中。

xxxx同事要求新增一個 GetMiddleCoordinates function 到library中。

有很多方法可以解決這個問題

1. 第一個方法是將 GetArea GetMiddleCoordinates function 直接添加到形狀介面(interface)中，然後在每個形狀結構中實作。這是一個聽起來滿合理的解決方案，但實際上是需要付出相應的代價。
   通常我們作為library的維護者，`可能`不想在每次有新的要求時冒險破壞已經測試過且沒問題的程式碼。這種方法會入侵到你既有的程式碼中，造成一定程度的風險。
   
2. 第二種方法為讓需要該功能的使用者自己實作想要的行為，例如 GetArea GetMiddleCoordinates function 。但如果我們的結構有私有的變數或私有方法會第三方團隊可能無法取得某些變數或是方法。

3. 第三種為本篇要介紹的`訪問者模式 Visitor Pattern`可以在最小的修改原始的程式碼上進行擴增。

```go
type visitor interface {
   visitForSquare(square)
   visitForCircle(circle)
   visitForTriangle(triangle)
}
```

訪問者需要實作各種圖形的擴增功能例如擴增`areaCalculator`功能。
```go
type areaCalculator struct {
    area int
}

func (a *areaCalculator) visitForSquare(s *square) {
    // Calculate area for square.
    // Then assign in to the area instance variable.
    fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
    fmt.Println("Calculating area for circle")
}
func (a *areaCalculator) visitForrectangle(s *rectangle) {
    fmt.Println("Calculating area for rectangle")
}
```
有了擴增功能的實作者後，我們需要把擴增功能與原有的物件綁定再一起，如下所示我們可以看到 shape 接受一個 visitor 的 interface 的物件。

如此一來原有的物件就跟訪問者 （visitor interface）有了掛勾，原有的物件能夠透過訪問者（visitor interface）的 function 把自己傳出去，讓訪問者可以取得物件的部分資訊 。

```go

type shape interface {
    getType() string
    accept(visitor)
}

type rectangle struct {
    l int
    b int
}

func (t *rectangle) accept(v visitor) {
    v.visitForrectangle(t)
}

func (t *rectangle) getType() string {
    return "rectangle"
}
```

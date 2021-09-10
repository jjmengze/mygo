# 中介者模式-Mediator Pattern
定義一個 Mediator interface 用來封裝物件的互動方式。

Mediator 設計模式主要是藉由 Mediator interface 作為中介者，避免物件間相互直接引用從而降低物件之間的耦合程度，從而我們能夠獨立地改變這些物件不會影響其他物件。

## 優點
減少物件之間的依賴，把原有的一對多的依賴變成了一對一的依賴。

以範例來看 Intercity Bus 與 Shuttle Bus 只依賴 mediator (station Manager)，物件之間減少依賴，同時會降低了物件的耦合程度，物件修改時就不會影響其他物件。

>Intercity Bus 或 Shuttle Bus 時都需要知道對方目前是否到站，若是對方已經到站需要等待對方駛離才能進站。（若是兩者之間沒有中介人的話，就會需要互相引用才能知道彼此的狀態）

若是透過  mediator (station Manager) ，Intercity Bus 或 Shuttle Bus 進站時把資訊告訴 station Manager 再由 station Manager 告訴 Bus 車輛目前是否可以進入。

## 缺點

中介者會隨著時間膨脹，原本Ｎ個物件之間相互依賴，現在變成物件依賴中介者，隨著業務物邏輯的增長物件也會跟著增加伴隨而來的就是中介者的膨脹。

以範例來看目前只有 Intercity Bus 與 Shuttle Bus。

如果之後新增 Sightseeing Bus 、  Night-Time Bus 以及 Open-Top Bus 的話，對於 Bus 來說只要實踐自己的進出站邏輯就好。

但對於 mediator (station Manager)來說就需要管理 Bus 進站的安排，舉例來說 Sightseeing Bus 的權限較高需要需要優先被調度，或是等待最久的車輛可以被優先調度。

這些業務邏輯就落在mediator (station Manager)身上了。
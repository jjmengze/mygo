# 責任練模式 chain of responsibility pattern

使責任鍊上的物件都有機會處理請求，避免請求的發送者和接收者與責任鏈上的物件產生的耦合關係。


例如流程當中有有多個 if-else 狀況，可以考慮用責任鍊重構。
## example

## gin example

```go
// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc


type Context struct {
    ...
    handlers HandlersChain
	index    int8
}


func (c *Context) Next() {
    c.index++
    for c.index < int8(len(c.handlers)) {
        c.handlers[c.index](c)
        c.index++
    }
}
```
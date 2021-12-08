# 責任練模式 chain of responsibility pattern

使責任鍊上的物件都有機會處理請求，避免請求的發送者和接收者與責任鏈上的物件產生的耦合關係。


例如流程當中有有多個 if-else 狀況，可以考慮用責任鍊重構。
## example
例如我們設定任務需要經過一系列的工作審核，最終才能交付到客戶手中

例如我們需要對產品進行包裝（Packaging ），對產品進行組裝（Assembly），對產品需要的原物料進行收集（Gathering materials）。
順序會是
1. Gathering materials
2. Assembly
3. Packaging


我們要打造一個玩具，我是這樣定義玩具
```go
type Task struct {
    Name              string
    MaterialCollected bool
    AssemblyExecuted  bool
    PackagingExecuted bool
}

task := &basic.Task{Name: "truck_toy"}
```
把玩具需送進`chain of responsibility pattern`的每個階段進行組裝與審核。
```go
    //建立包裝工作
    packaging := &basic.Packaging{}

    //建立組裝工作
	assembly := &basic.Assembly{}
	
	
    //建立原物料收集工作
	material := &basic.Material{}
	//設定原物料的下一關為組裝
	material.SetNext(assembly)
	//設定組裝完的下一關為包裝
	assembly.SetNext(packaging)
	
	//最後把任務送進原物料收集的工作中(入口處)
    material.Execute(task)
```

原物料部門（入口處）會收到組裝玩具的任務，對玩具進行修修改改。在這個部門若出現任何問題就`不會送到下一個部門`而是直接把任務丟棄。


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

## kubernetes example
```go
func NewMultiplexerPluginWrapper(plugins ...Plugin) *MultiplexerPluginWrapper {
	return &MultiplexerPluginWrapper{
		plugins: plugins,
	}
}

func NewAuthorLoggerPluginWrapper(plugin Plugin) *AuthorLoggerPluginWrapper {
    return &AuthorLoggerPluginWrapper{
        plugin: plugin,
    }
}
```
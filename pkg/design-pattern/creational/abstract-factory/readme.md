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



第二個例子為 kubernetes 裡面常見的範例 SharedInformerFactory 主宰了如何建立 kubernetes 各項 資源的 informer 如何生成，例如 app 的 informer 、Batch 的 informer 等等。
```go
// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	Admissionregistration() admissionregistration.Interface
	Internal() apiserverinternal.Interface
	Apps() apps.Interface
	Autoscaling() autoscaling.Interface
	Batch() batch.Interface
	Certificates() certificates.Interface
	Coordination() coordination.Interface
	Core() core.Interface
	Discovery() discovery.Interface
	Events() events.Interface
	Extensions() extensions.Interface
	Flowcontrol() flowcontrol.Interface
	Networking() networking.Interface
	Node() node.Interface
	Policy() policy.Interface
	Rbac() rbac.Interface
	Scheduling() scheduling.Interface
	Storage() storage.Interface
}

type sharedInformerFactory struct {
    client           kubernetes.Interface
    namespace        string
    tweakListOptions internalinterfaces.TweakListOptionsFunc
    lock             sync.Mutex
    defaultResync    time.Duration
    customResync     map[reflect.Type]time.Duration
    
    informers map[reflect.Type]cache.SharedIndexInformer
    // startedInformers is used for tracking which informers have been started.
    // This allows Start() to be called multiple times safely.
    startedInformers map[reflect.Type]bool
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(client kubernetes.Interface, defaultResync time.Duration, options ...SharedInformerOption) SharedInformerFactory {
    factory := &sharedInformerFactory{
    client:           client,
    namespace:        v1.NamespaceAll,
    defaultResync:    defaultResync,
    informers:        make(map[reflect.Type]cache.SharedIndexInformer),
    startedInformers: make(map[reflect.Type]bool),
    customResync:     make(map[reflect.Type]time.Duration),
    }
    
    // Apply all options
    for _, opt := range options {
    factory = opt(factory)
    }
    
return factory
}

func (f *sharedInformerFactory) Apps() apps.Interface {
    return apps.New(f, f.namespace, f.tweakListOptions)
}

```
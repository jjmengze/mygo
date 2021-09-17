package chain_of_responsibility

type Section interface {
	Execute(*Task)
	SetNext(Section)
}

type Task struct {
	Name              string
	MaterialCollected bool
	AssemblyExecuted  bool
	PackagingExecuted bool
}

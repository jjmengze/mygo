package basic

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

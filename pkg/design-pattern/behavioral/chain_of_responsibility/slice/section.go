package slice

type Section interface {
	Execute(*Task)
	//SetNext(Section)
}

type SectionFilterChain struct {
	filters []Section
}

// AddFilter ...
func (c *SectionFilterChain) AddFilter(filter ...Section) {
	c.filters = append(c.filters, filter...)
}

// Execute 执行
func (c *SectionFilterChain) Execute(task *Task) bool {
	for _, filter := range c.filters {
		filter.Execute(task)
	}
	return false
}

type Task struct {
	Name              string
	MaterialCollected bool
	AssemblyExecuted  bool
	PackagingExecuted bool
}

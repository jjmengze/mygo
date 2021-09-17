package slice

import (
	"fmt"
)

type Packaging struct {
	next Section
}

func (p *Packaging) Execute(t *Task) {
	if t.PackagingExecuted {
		fmt.Println("Packaging already done")
		//p.next.Execute(t)
		return
	}
	fmt.Println("Packaging Section doing Packaging")
}


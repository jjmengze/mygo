package basic

import (
	"fmt"
)

type Assembly struct {
	next Section
}

func (a *Assembly) Execute(t *Task) {
	if t.AssemblyExecuted {
		fmt.Println("Assembly already done")
		a.next.Execute(t)
		return
	}
	fmt.Println("Assembly Section assembling...")
	t.AssemblyExecuted = true
	a.next.Execute(t)
}

func (a *Assembly) SetNext(next Section) {
	a.next = next
}

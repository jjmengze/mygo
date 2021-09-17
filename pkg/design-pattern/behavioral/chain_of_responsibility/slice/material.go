package slice

import (
	"fmt"
)

type Material struct {
	next Section
}

func (m *Material) Execute(t *Task) {
	if t.MaterialCollected {
		fmt.Println("Material already collected")
		//m.next.Execute(t)
		return
	}
	fmt.Println("Material section gathering materials")
	t.MaterialCollected = true
	//m.next.Execute(t)
}


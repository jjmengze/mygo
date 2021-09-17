package main

import (
	"github.com/jjmengze/mygo/pkg/design-pattern/behavioral/chain_of_responsibility/slice"
)

func main() {
	sc := slice.SectionFilterChain{}

	packaging := &slice.Packaging{}

	// set next for assembly section
	assembly := &slice.Assembly{}

	material := &slice.Material{}

	//sc.AddFilter([]slice.Section{packaging, assembly, material}...)
	sc.AddFilter(packaging, assembly, material)

	task := &slice.Task{Name: "truck_toy"}
	sc.Execute(task)
}

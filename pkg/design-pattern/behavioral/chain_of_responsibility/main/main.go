package main

import "github.com/jjmengze/mygo/pkg/design-pattern/behavioral/chain_of_responsibility"

func main() {
	packaging := &chain_of_responsibility.Packaging{}

	// set next for assembly section
	assembly := &chain_of_responsibility.Assembly{}
	assembly.SetNext(packaging)

	material := &chain_of_responsibility.Material{}
	material.SetNext(assembly)

	task := &chain_of_responsibility.Task{Name: "truck_toy"}
	material.Execute(task)
}

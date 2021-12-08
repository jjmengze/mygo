package main

import (
	"github.com/jjmengze/mygo/pkg/design-pattern/behavioral/chain_of_responsibility/basic"
)

func main() {
	packaging := &basic.Packaging{}

	assembly := &basic.Assembly{}
	assembly.SetNext(packaging)

	material := &basic.Material{}
	material.SetNext(assembly)

	task := &basic.Task{Name: "truck_toy"}
	material.Execute(task)
}

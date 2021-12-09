package main

import "github.com/jjmengze/mygo/pkg/design-pattern/structural/composite"

func main() {
	file1 := &composite.file{name: "File1"}
	file2 := &composite.file{name: "File2"}
	file3 := &composite.file{name: "File3"}

	folder1 := &composite.folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &composite.folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}

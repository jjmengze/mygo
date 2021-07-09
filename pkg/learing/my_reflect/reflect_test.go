package my_reflect

import "testing"

func Test_inspectMap(t *testing.T) {
	inspectMap(map[uint32]uint32{
		1: 2,
		3: 4,
	})
	inspectMap(map[string]int{
		"KeyA": 2,
		"Key":  4,
	})
}

func Test_inspectSliceArray(t *testing.T) {
	inspectSliceArray([]int{1, 2, 3})
	inspectSliceArray([3]int{4, 5, 6})
}

type Info struct {
	note string
}

type User struct {
	Name    string
	Age     int
	Married bool
	note    *Info
}

func (u User) Add(a, b int) int {
	return a + b
}

func (u *User) SetAge(a int) {
	u.Age = a
}

func Test_inspectStruct(t *testing.T) {
	u := User{
		Name:    "dj",
		Age:     18,
		Married: true,
		note:    &Info{"nothing"},
	}

	inspectStruct(u, 1, 2)
}

func Add(a, b int) int {
	return a + b
}

func Greeting(name string) string {
	return "hello " + name
}

func Test_inspectFunc(t *testing.T) {
	inspectFunc("Add", Add)
	inspectFunc("Greeting", Greeting)
}

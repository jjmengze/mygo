package my_reflect

import (
	"fmt"
	"reflect"
)

func inspectMap(m interface{}) {
	v := reflect.ValueOf(m)
	for _, k := range v.MapKeys() {
		field := v.MapIndex(k)

		fmt.Printf("%v => %v\n", k.Interface(), field.Interface())
	}
}

func inspectSliceArray(sa interface{}) {
	v := reflect.ValueOf(sa)

	fmt.Printf("%c", '[')
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		fmt.Printf("%v ", elem.Interface())
	}
	fmt.Printf("%c\n", ']')
}

func inspectStruct(u interface{}, args ...interface{}) {
	v := reflect.ValueOf(u)

	argsV := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		argsV = append(argsV, reflect.ValueOf(arg))
		fmt.Printf("args: %v \n", reflect.ValueOf(arg).Interface())
	}

	for i := 0; i < v.NumMethod(); i++ {
		m := v.Method(i)
		rest := m.Call(argsV)
		fmt.Println(m)
		for _, ret := range rest {
			fmt.Println(ret.Interface())
		}
	}
	fmt.Println("")
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Printf("field:%d type:%s value:%d\n", i, field.Type().Name(), field.Int())

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fmt.Printf("field:%d type:%s value:%d\n", i, field.Type().Name(), field.Uint())

		case reflect.Bool:
			fmt.Printf("field:%d type:%s value:%t\n", i, field.Type().Name(), field.Bool())

		case reflect.String:
			fmt.Printf("field:%d type:%s value:%q\n", i, field.Type().Name(), field.String())

		default:
			fmt.Printf("field:%d unhandled kind:%s\n", i, field.Kind())
		}
	}
}

func inspectFunc(name string, f interface{}) {
	t := reflect.TypeOf(f)
	fmt.Println(name, "input:")
	for i := 0; i < t.NumIn(); i++ {
		t := t.In(i)
		fmt.Print(t.Name())
		fmt.Print(" ")
	}
	fmt.Println()

	fmt.Println("output:")
	for i := 0; i < t.NumOut(); i++ {
		t := t.Out(i)
		fmt.Print(t.Name())
		fmt.Print(" ")
	}
	fmt.Println("\n===========")
}

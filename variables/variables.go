package main

import (
	"fmt"
	"reflect"
)

/*var (
	name, course string
	module       float64
)*/

var (
	name   = "Nigel"
	course = "Docker Deep Dive"
	module = 3.2
)

func main() {
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))

	ptr := &module

	fmt.Println("Memory address of *module* variable is", ptr, "and the value of *module* is", *ptr)

	a := 10.0000000000
	b := 3

	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is of type", reflect.TypeOf(b))

	c := int(a) + b

	fmt.Println("\nC has value:", c, "and is of type:", reflect.TypeOf(c))

	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourse(&course)

	fmt.Println("\nYou are now watching course", course)

	const myLastName = "MASSON"

	fmt.Println("\nMy last name is", myLastName)
}

func changeCourse(course *string) string {
	*course = "First Look: Native Docker Clustering"

	fmt.Println("\nTrying to change your course to", *course)

	return *course
}

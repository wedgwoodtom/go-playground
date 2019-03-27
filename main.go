package main

import "fmt"

func main() {
	fmt.Println("Hi")

	listFun()
	mapFun()
}

func listFun() {
	// slice
	list := make([]string, 3)
	list[0] = "A"
	list[1] = "B"
	list[2] = "C"

	// panic
	//listFun[3] = "C"

	list = append(list, "D")
	fmt.Println(list)

	// A, B
	sublist := list[0:2]
	fmt.Println(sublist)
}

func mapFun() {
	dict := make(map[string]string)
	dict["A"] = "First"
	dict["B"] = "Second"
	dict["C"] = "Third"

	fmt.Println(dict)

	fmt.Println(len(dict))
	delete(dict, "A")
	delete(dict, "Z")
	fmt.Println(dict)

	value, present := dict["A"]
	fmt.Println("present ", present)
	fmt.Println("value ", value)

	allAtOnce := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("allAtOnce:", allAtOnce)

}

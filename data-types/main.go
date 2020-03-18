package main

import "fmt"

var vname1, vname2, vname3 int

// var vname1, vname2, vname3 string = "Joe", "Iris", "Wally"
const constantName = 55

// you can assign type of constants if it's necessary
const Pi float32 = 3.1415926

var a int8

var b int32

var c complex64 = 5 + 5i

//output: (5+5i)
// fmt.Printf("Value is: %v", c)

// Arrays
// var arr [n]type
//in [n]type, n is the length of the array, type is the type of its elements
var arr [10]int // an array of type [10]int

func main() {
	// c := a + b
	arr[0] = 42 // array is 0-based
	arr[1] = 13 // assign value to element
	arr[9] = 96
	last := arr[9]
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	m := make(map[string]string) // returns none zero value
	k := new(map[int]int) // returns pointer to map
	fmt.Println(vname1, vname2, vname3, constantName, Pi)
	fmt.Println(easyArray, arr, last, m, k)
}

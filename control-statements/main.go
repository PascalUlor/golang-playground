package main

import "fmt"

type testInt func(int) bool // define a function type of variable

func isOdd(integer int) bool {
    return integer%2 != 0
}

func isEven(integer int) bool {
    return integer%2 == 0
}

// pass the function `f` as an argument (a call back) to another function
func filter(slice []int, f testInt) []int {
    var result []int
    for _, value := range slice {
        if f(value) {
            result = append(result, value)
        }
    }
    return result
}

var slice = []int{1, 2, 3, 4, 5, 7}

// main fnction
func main() {
	integer := 6
switch integer {
case 4:
    fmt.Println("integer <= 4")
    fallthrough
case 5:
    fmt.Println("integer <= 5")
    fallthrough
case 6:
	fmt.Println("integer <= 6")
	integer++
    fallthrough
case 7:
	fmt.Println("integer <= 7")
	integer++
    fallthrough
case 8:
	fmt.Println("integer <= 8")
	integer++
    fallthrough
default:
    fmt.Println("default case")
}
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}


  odd := filter(slice, isOdd)
  even := filter(slice, isEven)

  fmt.Println("slice = ", slice)
    fmt.Println("Odd elements of slice are: ", odd)
    fmt.Println("Even elements of slice are: ", even)
}
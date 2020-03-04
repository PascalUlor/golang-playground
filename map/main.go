package main

import "fmt"


func main()  {
	dcComics := map[string]string {
		"supeman": "clark kent",
		"batman": "bruce wayne",
		"wonder woman": "Diana Prince",
	}
	dcComics["aquaman"] = "Authur Curry"
	dcComics["the-flash"] = "Barry Allen"
	printComics(dcComics)
}

func printComics(c map[string]string)  {
	for key, value := range(c) {
		fmt.Println("super hero name", key, "actual name", value)
	}
}
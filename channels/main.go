package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	c <- "connecting to site"

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down!")
		c <- link
		return
	}
	c <- link
	fmt.Println(link, "is ok!")
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)


type logWriter struct{}
func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// first pass
	// fmt.Println(res)

	//second pass
	// bs := make([]byte, 100000)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	//third pass
	lw := logWriter{}
	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (int, error)  {
	fmt.Println(string(bs))
	fmt.Println("This is the response in bytes", len(bs))
	return len(bs), nil
}

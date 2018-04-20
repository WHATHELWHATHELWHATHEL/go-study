package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		bytes, error := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error in reading the response %s %v \n", url, error)
		}
		fmt.Printf("%s", bytes)
	}
}

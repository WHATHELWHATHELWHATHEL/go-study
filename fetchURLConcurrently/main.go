package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	patchRequestStartTime := time.Now()
	// Key part, this create a channel which pass string types of value
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// Here we consume the value sent through the channel
		fmt.Println(<-ch)
	}
	fmt.Printf("total time: %.3f\n", time.Since(patchRequestStartTime).Seconds())
}

// Here if you review the path result,

// All these ch are used to passed the string
// result to the main function to handle the problem
func fetch(url string, ch chan<- string) {
	startTime := time.Now()
	response, netWorkError := http.Get(url)
	if netWorkError != nil {
		// This is the mechanism call channel
		ch <- fmt.Sprint(netWorkError)
		return
	}
	// make a useless write operation to write the staff
	numberOfBytes, IOError := io.Copy(ioutil.Discard, response.Body)
	// The response stream is useless, close it
	response.Body.Close()
	if IOError != nil {
		ch <- fmt.Sprintf("error in counting bytes url %s with error %v", url, IOError)
		return
	}
	requestDuration := time.Since(startTime).Seconds()
	ch <- fmt.Sprintf("%.3f %7d %s", requestDuration, numberOfBytes, url)
}

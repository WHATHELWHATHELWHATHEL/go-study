package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create a map the map[string]int indicate
	// it use string as key and int as value
	// the building make then create the map instance
	// by the designated map type
	counts := make(map[string]int)
	// Like the Java Scanner
	input := bufio.NewScanner(os.Stdin)
	// scan the staff and input
	for input.Scan() {
		counts[input.Text()] += 1
		// loop for map will return two results,
		// first as key, second as value
		for line, n := range counts {
			if n > 1 {
				// Similar to C printf, use different placeholder
				// to decide the position of params
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

package main;

import (
	"fmt"
	"os"
);

// func main() {
// 	var result, seperator string;
// 	// Last elements in the declaration decide the type of variable
// 	for i := 1; i < len(os.Args); i += 1 {
// 		result += seperator + os.Args[i];
// 		seperator = " ";
// 	}
// 	fmt.Println(result);
// }
func main() {
	var result, seperator string;
	// Last elements in the declaration decide the type of variable
	// For string type, it has a default value like " "
	for _, arg := range os.Args[2:] {
		// Get the CLI arguments from the CLI console
		result += seperator + arg;
		seperator = " ";
	}
	fmt.Println(result);
}

package main

import (
  "fmt"
  "os"
  "strings"
  "io/ioutil"
)

func main() {
  filename := os.Args[1];
  dataCounts := make(map[string]int);
  fileBytes, fileError := ioutil.ReadFile(filename)
  if (fileError != nil) {
    fmt.Printf("Error occur in readin file %s", filename)
  } else {
    lines := strings.Split(string(fileBytes), "\n")
    for index, line := range lines {
      fmt.Printf("processing line %d\n", index)
      dataCounts[line] += 1
    }
    printCounts(dataCounts);
  }

}

func printCounts (dataCounts map[string]int) {
  for dataName, dataAppearCount := range dataCounts {
    fmt.Printf("data %s appear in %d times\n", dataName, dataAppearCount)
  }
}

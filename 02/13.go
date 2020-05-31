package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	col1 := "col1.txt"
	col2 := "col2.txt"

	c1, err := os.Open(col1)
	if err != nil {
		os.Exit(1)
	}

	c2, err := os.Open(col2)
	if err != nil {
		os.Exit(1)
	}

	reader1 := bufio.NewReaderSize(c1, 4096*1024)
	reader2 := bufio.NewReaderSize(c2, 4096*1024)
	merged, err := os.Create("merged.txt")
	if err != nil {
		os.Exit(1)
	}
	writer := bufio.NewWriter(merged)
	for {
		first, _, err := reader1.ReadLine()
		second, _, err := reader2.ReadLine()
		if err != nil {
			os.Exit(1)
		}
		line := strings.Join([]string{string(first), string(second)}, "\t")
		writer.WriteString(line + "\n")
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	filename := "popular-names.txt"
	f, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	reader := bufio.NewReaderSize(f, 4096*1024)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			os.Exit(1)
		}

		line = strings.ReplaceAll(line, "\t", " ")
		fmt.Print(line)
	}
}

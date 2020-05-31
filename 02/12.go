package main

import (
	"bufio"
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

	reader := bufio.NewReaderSize(f, 4096*1024)
	fileCol1 := "col1.txt"
	fileCol2 := "col2.txt"
	f1, err := os.Create(fileCol1)
	if err != nil {
		os.Exit(1)
	}
	defer f1.Close()

	f2, err := os.Create(fileCol2)
	if err != nil {
		os.Exit(1)
	}
	defer f2.Close()

	f1Writer := bufio.NewWriter(f1)
	f2Writer := bufio.NewWriter(f2)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			os.Exit(1)
		}
		words := strings.Fields(line)
		f1Writer.WriteString(words[0] + "\n")
		f2Writer.WriteString(words[1] + "\n")
	}
}

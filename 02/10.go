package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func countLineNum(r *bufio.Reader) int {
	var lineNum int
	for {
		_, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			os.Exit(1)
		}
		lineNum++
	}
}

func main() {
	filename := "popular-names.txt"
	f, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	reader := bufio.NewReaderSize(f, 4096*1024)
	lineNum := countLineNum(reader)
	fmt.Println(lineNum)
}

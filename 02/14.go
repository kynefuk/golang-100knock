package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := "popular-names.txt"
	var n int
	flag.IntVar(&n, "n", 1, "line num")
	flag.Parse()

	f, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}

	reader := bufio.NewReaderSize(f, 4096*1024)
	for {
		if n == 0 {
			break
		}
		n--
		line, _, err := reader.ReadLine()
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(string(line))
	}
}

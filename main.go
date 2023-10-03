package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./import_checklist.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	// var line string
	var lineCount int64

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("line => %v", line)
		if err != nil {
			panic(err)
		}
		if lineCount == 0 {
			lineCount++
			continue
		}
		lineCount++
		//use read line to csv column, delimiter by '|'
		r := csv.NewReader(strings.NewReader(string(line)))
		r.Comma = '|'
		records, err := r.ReadAll()
		if err != nil {
			panic(err)
		}

		//col 1
		for _, row := range records {
			fmt.Println(row[0])
		}
		//col 2
		for _, row := range records {
			fmt.Println(row[1])
		}
		//col 3
		for _, row := range records {
			fmt.Println(row[2])
		}
	}
	fmt.Println("done")
}

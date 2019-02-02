package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		panic("ERROR: no path argument given")
	}
	args := os.Args[1:]
	filepath := args[0]
	file, err := os.Open(filepath)

	if err != nil {
		panic("a problem has occured: file not found!")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Title(line)
		fmt.Println(line)
	}

}

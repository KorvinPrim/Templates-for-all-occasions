package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	for _, file := range files {
		open_f, err := os.Open(file)
		if err != nil {
			fmt.Printf("File %v no found!", file)
		} else {
			input := bufio.NewScanner(open_f)
			for input.Scan() {
				str := input.Text()
				fmt.Println(str)
			}
		}

	}

}

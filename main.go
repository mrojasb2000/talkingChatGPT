package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	for true {
		fmt.Print("What is you name ? > ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("hello %s", line)
	}
}

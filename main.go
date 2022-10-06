package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(wd)
}

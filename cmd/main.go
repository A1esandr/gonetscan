package main

import (
	"fmt"
	"github.com/A1esandr/gonetscan"
)

func main() {
	err := gonetscan.NewScanner().Scan("127.0.0.1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Finish")
}

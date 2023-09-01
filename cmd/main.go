package main

import (
	"fmt"
	"github.com/A1esandr/gonetscan"
)

func main() {
	ports := make([]int, 0, 10000)
	for i := 1; i < 100; i++ {
		ports = append(ports, i)
	}
	gonetscan.NewScanner().Scan("127.0.0.1", ports)
	fmt.Println("Finish")
}

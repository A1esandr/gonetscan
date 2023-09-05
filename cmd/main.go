package main

import (
	"fmt"
	"strconv"

	"github.com/A1esandr/gonetscan"
)

func main() {
	ports := make([]int, 0, 10000)
	for i := 1; i < 1000; i++ {
		ports = append(ports, i)
	}
	address := "127.0.0.1"
	result := gonetscan.NewScanner().Scan(address, ports)
	if len(result.Open) > 0 {
		fmt.Println("Open")
	}
	for _, v := range result.Open {
		fmt.Println(address, ":", strconv.Itoa(v))
	}
	fmt.Println("Finish")
}

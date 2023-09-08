package main

import (
	"flag"
	"fmt"

	"github.com/A1esandr/gonetscan"
)

var address string

func init() {
	flag.StringVar(&address, "address", "127.0.0.1", "address to scan")
}

func main() {
	flag.Parse()
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
		fmt.Printf("%s:%d", address, v)
	}
	fmt.Println("Finish")
}

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/A1esandr/gonetscan"
)

var address string

func init() {
	flag.StringVar(&address, "address", "127.0.0.1", "address to scan")
}

func main() {
	flag.Parse()
	ports := make([]int, 0, 10000)
	fmt.Println("Scan", int(math.Pow(2, 16)), "ports")
	for i := 1; i < int(math.Pow(2, 16)); i++ {
		ports = append(ports, i)
	}
	result := gonetscan.NewScanner().Scan(address, ports)
	data := make([]string, 0, len(result.Open)+2)
	if len(result.Open) > 0 {
		fmt.Println("Open")
		data = append(data, "Open")
	}
	for _, v := range result.Open {
		fmt.Printf("%s:%d\n", address, v)
		data = append(data, fmt.Sprintf("%s:%d\n", address, v))
	}
	data = append(data, "Finish")
	write(data)
	fmt.Println("Finish")
}

func write(data []string) {
	file, err := os.Create("result" + time.Now().Format("2023-01-01") + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, line := range data {
		file.WriteString(line + "\n")
	}
}

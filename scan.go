package gonetscan

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

type Result struct {
	Open   []int
	Closed []int
}

type scan struct {
}

type Scanner interface {
	Scan(address string, ports []int) *Result
}

func NewScanner() Scanner {
	return &scan{}
}

func (s *scan) Scan(address string, ports []int) *Result {
	var wg sync.WaitGroup
	for _, port := range ports {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			addr := address + ":" + strconv.Itoa(i)
			c, err := net.Dial("tcp", addr)
			if err != nil {
				fmt.Println(err)
			}
			if c != nil {
				fmt.Println("Success", addr)
				err = c.Close()
				if err != nil {
					fmt.Println(err)
				}
			}
		}(port)
	}
	wg.Wait()
}

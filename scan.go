package gonetscan

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

type scan struct {
}

type Scanner interface {
	Scan(address string, ports []int)
}

func NewScanner() Scanner {
	return &scan{}
}

func (s *scan) Scan(address string, ports []int) {
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

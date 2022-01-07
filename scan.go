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
	Scan(address string) error
}

func NewScanner() Scanner {
	return &scan{}
}

func (s *scan) Scan(address string) error {
	var wg sync.WaitGroup
	for i := 80; i <= 81; i++ {
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
		}(i)
	}
	wg.Wait()

	return nil
}

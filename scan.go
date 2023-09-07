package gonetscan

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"sync"
)

type (
	Result struct {
		Open   []int
		Closed []int
	}
	RawResult struct {
		Result *Result
		mu     sync.Mutex
	}
	Config struct {
		ShowLogs bool
	}
)

type scan struct {
	config Config
}

type Scanner interface {
	Scan(address string, ports []int) *Result
}

func NewScanner() Scanner {
	return &scan{}
}

func (s *scan) WithConfig(config Config) {
	s.config = config
}

func (s *scan) Scan(address string, ports []int) *Result {
	var wg sync.WaitGroup
	result := &RawResult{Result: &Result{}}
	for _, port := range ports {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			addr := address + ":" + strconv.Itoa(i)
			c, err := net.Dial("tcp", addr)
			if err != nil {
				if s.config.ShowLogs {
					fmt.Println(err)
				}
			}
			if c != nil {
				if s.config.ShowLogs {
					fmt.Println("Success", addr)
				}
				result.mu.Lock()
				result.Result.Open = append(result.Result.Open, i)
				result.mu.Unlock()
				err = c.Close()
				if err != nil {
					fmt.Println(err)
				}
			}
		}(port)
	}
	wg.Wait()
	sort.IntSlice(result.Result.Open).Sort()
	return result.Result
}

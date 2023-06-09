package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("www.xiaodi8.com:%d", p)
		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= 500; i++ {
			ports <- i
		}
	}()
	for i := 0; i < 500; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

func scanPortAsync(protocol, hostname string, port int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.Dial(protocol, address)
	if err == nil {
		conn.Close()
		results <- port
	}
}

func main() {
	hostname := "www.xiaodi8.com"
	protocol := "tcp"
	lowerPort := 1
	upperPort := 500

	startTime := time.Now() // 记录程序开始时间

	// Create a wait group and result channel.
	wg := sync.WaitGroup{}
	results := make(chan int)

	// Start workers to scan ports asynchronously.
	for port := lowerPort; port <= upperPort; port++ {
		wg.Add(1)
		go scanPortAsync(protocol, hostname, port, &wg, results)
	}

	// Collect results from workers and close the channel when all done.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read open ports from the result channel.
	openPorts := []int{}
	for port := range results {
		openPorts = append(openPorts, port)
	}

	// Sort and print the open ports.
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	fmt.Printf("程序运行时间：%.2f 秒\n", elapsedTime.Seconds())
}

package portscanner

import (
	"fmt"
	"net"
	"sort"
)

// SlowScan checks a single target one port at a time
func SlowScan(target string) {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf(target+":%d", i)

		conn, err := net.Dial("tcp", address)

		if err != nil {
			continue
		}
		conn.Close()

		fmt.Printf("%d open\n", i)
	}
}

// Scan checks a single target
func Scan(target string) {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, target)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
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

func worker(ports, results chan int, target string) {
	for p := range ports {
		address := fmt.Sprintf(target+":%d", p)

		conn, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0
			continue
		}

		conn.Close()

		results <- p
	}
}

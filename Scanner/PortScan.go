package scanner

import (
	"fmt"
	"sync"
	"time"
	"vidar-scan/basework"

	"github.com/panjf2000/ants/v2"
)

func PortScan(targetIp string, begin_port int, end_port int) []int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	concurrencylimit := 60
	rateLimit := 3 * time.Millisecond

	pool, err := ants.NewPool(concurrencylimit)
	if err != nil {
		fmt.Println("error: %v", err)
	}

	fmt.Println("-----START-----")
	var ports []int
	for port := begin_port; port <= end_port; port++ {
		wg.Add(1)
		p := port

		pool.Submit(func() {
			time.Sleep(rateLimit)
			defer wg.Done()
			if basework.TcpConnect(targetIp, p, 3*time.Second) {
				mu.Lock()
				ports = append(ports, p)
				mu.Unlock()

				fmt.Printf("[OPEN] %d\n", p)
			}
		})

		//if port%500 == 0 {
		//	fmt.Printf("%d\n", port)
		//}
	}

	wg.Wait()
	pool.Release()
	fmt.Println("-----OVER-----")

	return ports
}

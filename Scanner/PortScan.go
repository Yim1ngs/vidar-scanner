package scanner

import (
	"fmt"
	"sync"
	"time"
	"vidar-scan/basework"

	"github.com/panjf2000/ants/v2"
)

func PortScan(targetIp string) {
	var wg sync.WaitGroup

	concurrencylimit := 3000
	// sleeptime := 500 * time.Millisecond

	pool, _ := ants.NewPool(concurrencylimit)

	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		p := port

		pool.Submit(func() {
			defer wg.Done()
			if basework.TcpConnect(targetIp, p, 2*time.Second) {
				fmt.Printf("[OPEN] %d\n", p)
			}
		})
	}

	wg.Wait()
	pool.Release()
}

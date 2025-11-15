package scanner

import (
	"fmt"
	"net/http"
	"sync"
	"time"
	"vidar-scan/basework"
)

func Getscan(url string, filename string) {
	finalpath := basework.UrlConstruct(url, filename)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	var wg sync.WaitGroup

	ratelimit := 2

	limitch := make(chan struct{}, ratelimit)

	fmt.Println("-----START-----")

	for _, url = range finalpath {
		limitch <- struct{}{}

		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			defer func() { <-limitch }()

			basework.SendMessage(client, url)

		}(url)
	}

	wg.Wait()

	fmt.Println("-----OVER-----")

}

package basework

import (
	"bufio"
	"net/url"
	"os"
	"strings"
)

func UrlConstruct(urlStr string, filename string) (<-chan string, error) {
	out := make(chan string, 100) // 缓冲通道

	url1, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	if url1.Scheme == "" {
		url1.Scheme = "http"
	}

	if !strings.HasPrefix(url1.Path, "/") {
		url1.Path = url1.Path + "/"
	}

	urlStr = url1.String()

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	go func() {
		defer file.Close()
		defer close(out)
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			if line != "" {
				out <- urlStr + line
			}
		}
	}()

	return out, nil
}

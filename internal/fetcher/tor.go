package fetcher

import (
	"bufio"
	"net/http"
	"strings"
)

func FetchTorNodes() ([]string, error) {
    resp, err := http.Get("https://check.torproject.org/torbulkexitlist")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var ips []string
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" {
            ips = append(ips, line)
        }
    }
    return ips, scanner.Err()
}

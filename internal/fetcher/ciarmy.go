package fetcher

import (
	"bufio"
	"net/http"
	"strings"
)

func FetchCIArmy() ([]string, error) {
    resp, err := http.Get("https://cinsscore.com/list/ci-badguys.txt")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var ips []string
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" && !strings.HasPrefix(line, "#") {
            ips = append(ips, line)
        }
    }
    return ips, scanner.Err()
}
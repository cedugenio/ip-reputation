package fetcher

import (
	"bufio"
	"net/http"
	"strings"
)

func FetchSpamhausDROP() ([]string, error) {
    resp, err := http.Get("https://www.spamhaus.org/drop/drop.txt")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var ips []string
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, ";") {
            continue
        }
        fields := strings.Fields(line)
        if len(fields) > 0 {
            ips = append(ips, fields[0])
        }
    }
    return ips, scanner.Err()
}

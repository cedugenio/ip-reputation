package fetcher

func FetchAllSources() ([]string, error) {
    var all []string

    sources := []func() ([]string, error){
        FetchFireHOL,
        FetchCIArmy,
        FetchSpamhausDROP,
        FetchTorNodes,
    }

    for _, fetch := range sources {
        ips, err := fetch()
        if err != nil {
            continue // ou logar erro
        }
        all = append(all, ips...)
    }

    return deduplicate(all), nil
}

func deduplicate(ips []string) []string {
    m := make(map[string]bool)
    var result []string
    for _, ip := range ips {
        if !m[ip] {
            m[ip] = true
            result = append(result, ip)
        }
    }
    return result
}
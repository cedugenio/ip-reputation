package main

import (
	"log"
	"time"

	"github.com/cedugenio/ip-reputation/internal/api"
	"github.com/cedugenio/ip-reputation/internal/fetcher"
	"github.com/cedugenio/ip-reputation/internal/store"
)

func main() {
    store.InitRedis()

    go func() {
        for {
            log.Println("[FETCHER] Atualizando IPs maliciosos...")
            ips, err := fetcher.FetchFireHOL()
            if err != nil {
                log.Println("Erro ao buscar FireHOL:", err)
            } else {
                for _, ip := range ips {
                    _ = store.SaveIP(ip, 1, time.Hour*24)
                }
                log.Printf("Adicionados %d IPs\n", len(ips))
            }
            time.Sleep(time.Hour * 6)
        }
    }()

    api.StartServer()
}
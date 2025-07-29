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
        	log.Println("[FETCHER] Coletando dados de múltiplas fontes...")
        	ips, err := fetcher.FetchAllSources()
        	if err != nil {
            	log.Println("Erro na coleta:", err)
        	} else {
            	for _, ip := range ips {
                	_ = store.SaveIP(ip, 1, 24*time.Hour)
            	}
            log.Printf("Total de IPs coletados: %d\n", len(ips))
        }
        time.Sleep(6 * time.Hour)
    }
	}()

    api.StartServer()
}
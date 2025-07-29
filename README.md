# IP Reputation Service (Go)

Um microserviço para verificar reputação de IPs baseado em fontes públicas (ex: FireHOL). Inspirado em Spamhaus.

## 🚀 Como rodar

go run ./cmd/main.go
curl http://localhost:8080/check/1.2.3.4

1. Instale o Redis localmente ou use um container Docker:

```bash
docker run -d -p 6379:6379 redis
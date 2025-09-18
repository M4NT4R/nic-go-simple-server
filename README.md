# nic-go-simple-server
1. `/ping`
   - Запрос: GET http://localhost:8080/ping
   - Ответ: 
     ```json
     { "message": "pong" }

2. `/sum?a=2&b=3`
   - Запрос: GET http://localhost:8080/sum?a=2&b=3
   - Ответ:
     ```json
     { "result": 5 }

## Запуск

```bash
go run main.go

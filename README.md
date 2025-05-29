### Buat file .env
 - Buat file .env(di root), kemudian copy semua isinya dari .env.example
 - Sesuaikan semuanya (nama database dan lain lain)

### Jalankan Swagger
```go
swag init -g cmd/main.go
```

### Jalnkan Program
```go
go run cmd/main.go
```

### Buka Swagger UI
```go
http://localhost:8080/swagger/index.html
```
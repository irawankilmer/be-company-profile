### Buat file .env
 - Buat file .env(di root), kemudian copy semua isinya dari .env.example
 - Sesuaikan semuanya (nama database dan lain lain)

### Jalankan Swagger
```go
swag init -g cmd/main.go
```

### Jalnkan Program
```go
air
```

### Buka Swagger UI
```go
http://localhost:8080/swagger/index.html
```
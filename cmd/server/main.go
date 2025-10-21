package main

import (
	"github.com/duonghuu/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}

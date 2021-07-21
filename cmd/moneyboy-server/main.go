package main

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/server"
)

func main() {
	server := server.New()
	server.Start(":3000")
}

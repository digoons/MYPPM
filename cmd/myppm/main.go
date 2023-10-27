package main

import (
	"fmt"
	"myppm/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	//Use cfg.DatabaseURL and cfg.ServerPort to setup your application
	fmt.Printf("Server is running on port %s\n", cfg.ServerPort)
}

//TEXTEST

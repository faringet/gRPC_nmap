package main

import (
	"fmt"
	"gRPC_nmap/config"
	"gRPC_nmap/server"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("чтение конфигурации: %w", err)
	}
	return server.Start(cfg)
}

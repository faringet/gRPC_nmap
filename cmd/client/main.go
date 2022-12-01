package main

import (
	"fmt"
	"gRPC_nmap/client"
	"gRPC_nmap/config"
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
	return client.SendRequest(cfg)
}

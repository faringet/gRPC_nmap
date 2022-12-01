package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Addr string
}

func New() (Config, error) {
	if len(os.Args) != 2 {
		return Config{}, errors.New("неверное кол-во аргументов")
	}
	addr := os.Args[1]
	if !strings.Contains(addr, ":") {
		return Config{}, fmt.Errorf("неверный адрес: %s", addr)
	}
	return Config{addr}, nil
}

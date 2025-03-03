package main

import (
	"fmt"

	"github.com/HIS-Vita/auth-service/internal/config"
)

func main() {
	cfg := config.MustLoad("config/local.yaml")

	fmt.Println(cfg)
}

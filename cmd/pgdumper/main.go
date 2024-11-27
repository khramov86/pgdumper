package main

import (
	"fmt"
	"pgdumper/internal/config"
)

func main() {
	config := config.GetConfig()
	fmt.Println(config)
}

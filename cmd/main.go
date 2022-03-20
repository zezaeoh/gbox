package main

import (
	"github.com/zezaeoh/gbox/internal/cmd"
	"github.com/zezaeoh/gbox/internal/logger"
)

func main() {
	log := logger.Logger()
	if err := cmd.Execute(); err != nil {
		log.Errorf("error occurred while executing gbox.")
		log.Fatalf("%v", err)
	}
}

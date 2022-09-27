package main

import (
	"log"

	"github.com/rasulov-emirlan/template_microgo/config"
)

func main() {
	cfg, err := config.LoadConfigs()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg)
}

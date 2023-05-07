package main

import (
	"flag"
	"log"

	"VKBot/config"
	"VKBot/internal/vk"
)

var (
	cfgPath = flag.String("config", "", "path to config file")
)

func main() {
	flag.Parse()

	// read config file
	cfg, err := config.ReadConfig(*cfgPath)
	if err != nil {
		log.Fatalf("can't read config config: %v", err)
	}

	// New connection VKApi
	api, err := vk.NewVK(cfg)
	if err != nil {
		log.Fatalf("can't create vk api: %v", err)
	}

	api.RegisterEvents()

	// Run VKBot
	err = api.Run()
	if err != nil {
		log.Fatalf("can't run api: %v", err)
	}
}

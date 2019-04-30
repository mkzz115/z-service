package main

import (
	"os"

	"github.com/mkzz115/z-service.git/api_service/config"
	"github.com/mkzz115/z-service.git/api_service/processor"
	"github.com/mkzz115/z-service.git/common/log"
)

func main() {
	if len(os.Args) != 2 {
		panic("empty config file!")
	}
	configPath := os.Args[1]
	cfg, err := config.FromToml(configPath)
	if err != nil {
		print("init error")
		panic(err.Error())
	}
	log.Init(cfg.Logger.Path)
	// install http service
	s := processor.NewApiServer(cfg)
	if err := s.Init(); err != nil {
		log.Error("server init error[%s]", err.Error())
		os.Exit(-1)
	}
	if err := s.Start(); err != nil {
		log.Error("server start error[%s]", err.Error())
	}
}

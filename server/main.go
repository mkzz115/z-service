package main

import (
	"os"
	"path/filepath"

	"github.com/mkzz115/z-service.git/common/confutil"
	"github.com/mkzz115/z-service.git/common/log"
	"github.com/mkzz115/z-service.git/server/processor"
)

func main() {
	if len(os.Args) != 2 {
		panic("empty config file!")
	}
	configPath := os.Args[1]
	cfg, err := confutil.FromToml(configPath)
	if err != nil {
		print("init error")
		panic(err.Error())
	}

	err = log.Init(filepath.Join(cfg.Logger.Path, os.Args[0]))
	if err != nil {
		print("init error")
		panic(err.Error())
	}
	s := processor.NewThriftServer(cfg)
	if err := s.Init(); err != nil {
		log.Error("server init error[%s]", err.Error())
		os.Exit(-1)
	}
	if err := s.Start(); err != nil {
		log.Error("server start error[%s]", err.Error())
	}
}

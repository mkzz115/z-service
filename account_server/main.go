package main

import (
	"os"
	"path/filepath"

	"github.com/mkzz115/z-service.git/account_server/logic"
	"github.com/mkzz115/z-service.git/account_server/processor"
	"github.com/mkzz115/z-service.git/common/confutil"
	"github.com/mkzz115/z-service.git/common/log"
	"github.com/mkzz115/z-service.git/common/serve"
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

	proc := processor.NewAccountProcess()
	zserve := serve.NewZServer("account_serve", cfg.Account.Addr)
	err = zserve.Init(cfg, logic.InitHandle, proc)
	if err != nil {
		print("zserve.Init error:", err.Error())
		panic(err.Error())
	}
	zserve.Start()
}

package main

import (
	"log"
	"os"
	"time"

	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/server"
)

func main() {
	cfg := config.Get()
	svr := server.New(cfg)

	if err := svr.Start(); err != nil {
		end(err, "failed to start server")
	}
}

func end(err error, message string) {
	log.Fatalf("%+v: %+v", message, err)
	time.Sleep(time.Millisecond * 50)

	os.Exit(1)
}

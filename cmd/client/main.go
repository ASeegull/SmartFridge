package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/ASeegull/SmartFridge/staticServer"
	"github.com/ASeegull/SmartFridge/staticServer/config"
	"github.com/davecheney/errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfgPath := flag.String(
		"config",
		"staticServer/config/files/clientConfig.yaml",
		"Location of config File",
	)

	log.Info(*cfgPath)
	flag.Parse()
	cfg, err := config.GetSettings(*cfgPath)
	if err != nil {
		log.Fatal(errors.Annotate(err, "Server couldn't start"))
	}

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)

	defer func() {
		signal.Stop(ch)
		cancel()
	}()

	go func() {
		select {
		case <-ch:
			cancel()
		case <-ctx.Done():
		}
	}()

	staticServer.Run(cfg, ctx)
}

package main

import (
	config2 "github.com/instinctG/exchanges/internal/config"
	"github.com/instinctG/exchanges/internal/transport"
	"github.com/sirupsen/logrus"
)

func main() {
	config, err := config2.LoadConfig()
	if err != nil {
		logrus.Fatal("Error loading config: ", err)
	}
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logrus.Warn("invalid log level,defaulting to info: ", err)
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

	logrus.Infof("Starting server at %s:%s", config.Host, config.Port)
	logrus.Info("Starting exchanges...")

	httpHandler := transport.NewHandler(config.Host, config.Port)
	if err := httpHandler.Serve(); err != nil {
		logrus.Fatal(err)
	}

}

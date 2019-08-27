package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/ypapax/logrus_conf"
	"github.com/ypapax/players/config"
	"github.com/ypapax/players/parser"
	"os"
)

func main() {
	logrus_conf.Files("players", logrus.TraceLevel)
	var confPath string
	flag.StringVar(&confPath, "conf", "conf.yaml", "path to config file")
	flag.Parse()
	f, err := os.Open(confPath)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	c, err := config.Parse(f)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	pl, err := parser.GetPlayers(c.UrlTemplate, c.Teams, c.Timeout)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	for _, p := range pl {
		fmt.Println(p)
	}
}

// Autogenerated - DO NOT EDIT
package main

import (
	"flag"
	"syscall"
	"os"
	"os/signal"
	"fmt"

	"github.com/THE108/go-skeleton/config"
	"github.com/THE108/go-skeleton/resources"
	"github.com/THE108/go-skeleton/app"
	"github.com/THE108/go-skeleton/log"
	"github.com/THE108/go-skeleton/monitoring"
)

// Variables values set during compile time (see Makefile)
var (
	// Application name
	AppName = batler{ .Project.Name }
	// AppVersion application version
	AppVersion string
	// GoVersion GO version that used to build application
	GoVersion string
	// BuildDate date when application is built
	BuildDate string
	// GitLog log of commit when build app
	GitLog string
)

func main() {
	var filename string
	var showVersionAndExit bool

	flag.StringVar(&filename, "config", "config.json", "config file name")
	flag.BoolVar(&showVersionAndExit, "v", false, "version")
	flag.Parse()

	if showVersionAndExit {
		fmt.Printf("%s version: %s go: %s build: %s git: %s", AppName, AppVersion, GoVersion, BuildDate, GitLog)
		return
	}

	cfg, err := config.Parse(filename, AppVersion, GoVersion, BuildDate, GitLog)
	if err != nil {
		fmt.Println(err)
		return
	}

	logger, err := log.NewLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	mon := monitoring.NewMonitoring()

	res, err := resources.InitResources(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	logger.Info("Start application")
	if err := app.Run(cfg, logger, mon, res); err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for s := range ch {
		switch s {
		case syscall.SIGINT, syscall.SIGTERM:
			signal.Stop(ch)
			app.Shutdown()
			logger.Info("Shutdown application")
			return
		}
	}
}

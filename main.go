package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"net/http"

	"butler{ .Vars.repoPath }/butler{ .Project.Name }/monitoring"
	"butler{ .Vars.repoPath }/butler{ .Project.Name }/app"
	"butler{ .Vars.repoPath }/butler{ .Project.Name }/config"
)

// Variables values set during compile time (see Makefile)
var (
	// Application name
	AppName = "butler{ .Project.Name }"
	// Application description
	AppDescription = "butler{ .Project.Description }"
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
	var filename, pprofListenAddress string
	var showVersionAndExit, doProfile bool

	flag.StringVar(&filename, "config", "config.json", "config file name")
	flag.StringVar(&pprofListenAddress,"pprof", "localhost:6060", "pprof listen address")
	flag.BoolVar(&doProfile, "profile", false, "run with pprof enabled")
	flag.BoolVar(&showVersionAndExit, "v", false, "version")
	flag.Parse()

	if showVersionAndExit {
		fmt.Printf("%s - %s\nversion: %s go: %s build: %s git: %s\n",
			AppName, AppDescription, AppVersion, GoVersion, BuildDate, GitLog)
		return
	}

	cfg, err := config.Parse(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	if doProfile {
		go func() {
			log.Println(http.ListenAndServe(pprofListenAddress, nil))
		}()
	}

	monitoring.InitMonitoring(&cfg.Monitoring)

	application := app.NewApplication(cfg)

	quitCh := make(chan struct{})
	go func() {
		defer close(quitCh)

		log.Println("Start application")
		if err := application.Run(); err != nil {
			fmt.Println(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case s := <-ch:
			switch s {
			case syscall.SIGINT, syscall.SIGTERM:
				signal.Stop(ch)
				application.Shutdown()
				log.Println("Shutdown application")
			}

		case <-quitCh:
			return
		}
	}
}

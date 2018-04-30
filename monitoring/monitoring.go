package monitoring

import (
	"errors"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/cyberdelia/go-metrics-graphite"
	"github.com/rcrowley/go-metrics"
)

type Config struct {
	Server string
	Prefix string
}

func InitMonitoring(cfg *Config) {
	metrics.RegisterDebugGCStats(metrics.DefaultRegistry)
	metrics.RegisterRuntimeMemStats(metrics.DefaultRegistry)
	go metrics.CaptureDebugGCStats(metrics.DefaultRegistry, time.Minute)
	go metrics.CaptureRuntimeMemStats(metrics.DefaultRegistry, time.Minute)

	if err := startGraphite(cfg); err != nil {
		log.Print("Error starting Graphite client", err)

		go metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
	}
}

func startGraphite(cfg *Config) error {
	server := strings.TrimSpace(cfg.Server)
	if server == "" {
		return errors.New("empty graphite server address")
	}

	addr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		return err
	}

	go graphite.Graphite(metrics.DefaultRegistry, time.Minute, cfg.Prefix, addr)

	return nil
}
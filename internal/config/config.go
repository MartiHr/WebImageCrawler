package config

import (
	"flag"
	"time"
)

type Config struct {
	Workers        int
	MaxGoroutines  int
	FollowExternal bool
	Timeout        time.Duration
	StartURLs      []string
	MySQLDSN       string
}

func Load() *Config {
	cfg := &Config{}

	flag.IntVar(&cfg.Workers, "workers", 10, "worker pool size")
	flag.IntVar(&cfg.MaxGoroutines, "max-goroutines", 50, "maximum goroutines")
	flag.BoolVar(&cfg.FollowExternal, "external", false, "follow external links")
	flag.DurationVar(&cfg.Timeout, "timeout", 2*time.Minute, "crawler timeout")

	flag.StringVar(
		&cfg.MySQLDSN,
		"mysql",
		"crawler:crawler@tcp(127.0.0.1:3306)/webcrawler",
		"MySQL DSN",
	)

	flag.Parse()
	cfg.StartURLs = flag.Args()

	return cfg
}

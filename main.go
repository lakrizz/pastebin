package main

import (
	"encoding/gob"
	"log"
	"time"

	"github.com/namsral/flag"
	"github.com/patrickmn/go-cache"
)

var cfg Config

func main() {

	gob.Register(&Paste{}) // we need to manually register our new paste-type ;)

	var (
		config    string
		bind      string
		fqdn      string
		permstore bool
		expiry    time.Duration
	)

	flag.StringVar(&config, "config", "", "config file")
	flag.StringVar(&bind, "bind", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.StringVar(&fqdn, "fqdn", "localhost", "FQDN for public access")
	flag.BoolVar(&permstore, "permanent", false, "Pastes never expire [this overrides any expiry setting]")
	flag.DurationVar(&expiry, "expiry", 5*time.Minute, "expiry time for pastes")
	flag.Parse()

	if expiry.Seconds() < 60 {
		log.Fatalf("expiry of %s is too small", expiry)
	}

	if permstore {
		cfg.expiry = cache.NoExpiration
	} else {
		cfg.expiry = expiry
	}

	// TODO: Abstract the Config and Handlers better
	cfg.fqdn = fqdn
	cfg.permstore = permstore

	NewServer(bind, cfg).ListenAndServe()
}

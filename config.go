package main

import (
	"time"
)

// Config ...
type Config struct {
	permstore bool
	expiry    time.Duration
	fqdn      string
}

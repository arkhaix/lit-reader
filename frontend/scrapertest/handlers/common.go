package handlers

import (
	api "github.com/arkhaix/lit-reader/api/scraper"
	"time"
)

var (
	ScraperClient  api.ScraperClient
	ScraperTimeout time.Duration
)

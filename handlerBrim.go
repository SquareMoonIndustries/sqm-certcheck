package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func init() {
	routes = append(routes, Route{"handleBrim", "GET", "/sqmcertcheck", handlerBrim})
}

func handlerBrim(w http.ResponseWriter, r *http.Request) {
	fmsUrls := []string{
		"https://brimuat.ajg.com/",
		"https://brim.ajg.com/",
		"https://fmsuat.ajg.com/",
		"https://fms.ajg.com",
	}
	results := []struct {
		Url    string
		Expire time.Time
		Error  error
	}{}

	for _, url := range fmsUrls {
		t, err := checkUrl(url)

		results = append(results, struct {
			Url    string
			Expire time.Time
			Error  error
		}{
			Url:    url,
			Expire: t,
			Error:  err,
		})
	}

	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if settings.Debug || true {
		enc.SetIndent("", "\t")
	}
	enc.Encode(results)
}

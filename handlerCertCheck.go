package main

import (
	"crypto/x509"
	"encoding/json"
	"net/http"
	"time"
)

type UrlsData struct {
	Urls []UrlData `json:"urls"`
}

type UrlData struct {
	URL    string    `json:"url"`
	Expire time.Time `json:"expire"`
	Error  string    `json:"error"`
}

func init() {
	routes = append(routes, Route{"handlerCertCheck", "POST", "/", handlerCertCheck})
}

func handlerCertCheck(w http.ResponseWriter, r *http.Request) {
	data := UrlsData{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		logger.Error(err)
		http.Error(w, "Not valid input data", http.StatusInternalServerError)
	}
	for i, v := range data.Urls {
		v.Expire, v.Error = checkUrl(v.URL)
		data.Urls[i] = v
	}

	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if settings.Debug {
		enc.SetIndent("", "\t")
	}
	enc.Encode(data)
}

func checkUrl(url string) (time.Time, string) {
	resp, err := http.Get(url)
	if err != nil {
		return time.Time{}, err.Error()
	}
	defer resp.Body.Close()
	cert, err := x509.ParseCertificate(resp.TLS.PeerCertificates[0].Raw)
	if err != nil {
		return time.Time{}, err.Error()
	}
	return cert.NotAfter, ""
}

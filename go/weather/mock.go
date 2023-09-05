package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"time"
)

var rnd *rand.Rand

func newMockServer() string {
	rnd = rand.New(rand.NewSource(99))
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := fmt.Sprintf(`{
			"current": {
				"condition": {
					"text": "Overcast"
				},
				"is_day": 1,
				"last_updated": "%s",
				"precip_in": 0.0,
				"precip_mm": 0.0,
				"temp_c": %.1f,
				"uv": 5.0,
				"wind_degree": 350,
				"wind_dir": "N",
				"wind_mph": %.1f
			}
		}`, time.Now(), randFloat(), randFloat())
		fmt.Fprintln(w, body)
	}))
	return ts.URL
}

func randFloat() float64 {
	min := 0.0
	max := 30.0
	r := min + rand.Float64()*(max-min)
	return r
}

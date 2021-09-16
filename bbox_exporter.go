// Copyright (C) 2021 Nicolas Lamirault <nicolas.lamirault@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/nlamirault/bbox_exporter/exporter"
	"github.com/nlamirault/bbox_exporter/version"
)

const (
	banner = "bbox_exporter - %s\n"
)

var (
	debug         bool
	vrs           bool
	listenAddress string
	metricsPath   string
	endpoint      string
	password      string
)

func init() {
	// parse flags
	flag.BoolVar(&vrs, "version", false, "print version and exit")
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.StringVar(&listenAddress, "web.listen-address", ":9311", "Address to listen on for web interface and telemetry.")
	flag.StringVar(&metricsPath, "web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	flag.StringVar(&endpoint, "bbox", "https://mabbox.bytel.fr", "Endpoint of Bbox")
	flag.StringVar(&password, "password", "", "The admin password")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(banner, version.Version))
		flag.PrintDefaults()
	}

	flag.Parse()
	if vrs {
		fmt.Printf("%s", version.Version)
		os.Exit(0)
	}

	if len(endpoint) == 0 {
		usageAndExit("bbox endpoint cannot be empty.", 1)
	}
	if len(password) == 0 {
		usageAndExit("bbox password cannot be empty.", 1)
	}
}

func main() {
	exporter, err := exporter.NewExporter(endpoint, password)
	if err != nil {
		log.Errorf("Can't create exporter : %s", err)
		os.Exit(1)
	}
	log.Infoln("Register exporter")
	prometheus.MustRegister(exporter)

	http.Handle(metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>bbox Exporter</title></head>
             <body>
             <h1>bbox Exporter</h1>
             <p><a href='` + metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	log.Infoln("Listening on", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n")
	}
	flag.Usage()
	os.Exit(exitCode)
}

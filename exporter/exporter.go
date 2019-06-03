// Copyright (C) 2019 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/nlamirault/bbox_exporter/bbox"
)

const (
	namespace = "bbox"
)

var (
	up = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Was the last query of BBox successful.",
		nil, nil,
	)
)

// Exporter collects Bbox stats from the given server and exports them using
// the prometheus metrics package.
type Exporter struct {
	Bbox *bbox.Client
}

// NewExporter returns an initialized Exporter.
func NewExporter(endpoint string) (*Exporter, error) {
	log.Infof("Setup BBox exporter using URL: %s", endpoint)
	bboxClient, err := bbox.NewClient(endpoint)
	if err != nil {
		return nil, err
	}
	return &Exporter{
		Bbox: bboxClient,
	}, nil
}

// Describe describes all the metrics ever exported by the Bbox exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
	describeWanMetrics(ch)
}

// Collect the stats from channel and delivers them as Prometheus metrics.
// It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	log.Infof("Bbox exporter starting")
	resp, err := e.Bbox.GetMetrics()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(
			up, prometheus.GaugeValue, 0,
		)
		log.Errorf("Bbox error: %s", err.Error())
		return
	}
	log.Infof("Bbox metrics retrieved")
	storeWanMetrics(ch, resp.Wan.IPStatistics[0])
	storeWanFtthMetric(ch, resp.FtthState)
	ch <- prometheus.MustNewConstMetric(
		up, prometheus.GaugeValue, 1,
	)
	log.Infof("BBox exporter finished")
}

func storeMetric(ch chan<- prometheus.Metric, value float64, desc *prometheus.Desc, labels ...string) {
	ch <- prometheus.MustNewConstMetric(
		desc, prometheus.GaugeValue, value, labels...)
}

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

package exporter

import (
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"

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
	Bbox   *bbox.Client
	logger log.Logger
}

// NewExporter returns an initialized Exporter.
func NewExporter(endpoint string, password string, logger log.Logger) (*Exporter, error) {
	level.Info(logger).Log("msg", "Setup BBox exporter", "endpoint", endpoint)
	bboxClient, err := bbox.NewClient(endpoint, password, logger)
	if err != nil {
		return nil, err
	}
	return &Exporter{
		Bbox:   bboxClient,
		logger: logger,
	}, nil
}

// Describe describes all the metrics ever exported by the Bbox exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
	describeWanMetrics(ch)
	describeLanMetrics(ch)
	describeDeviceMetrics(ch)
	describeDNSMetrics(ch)
	describeIPTVMetrics(ch)
	describeServicesMetrics(ch)
	describeWirelessMetrics(ch)
}

// Collect the stats from channel and delivers them as Prometheus metrics.
// It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	level.Info(e.logger).Log("msg", "Bbox exporter starting")

	if err := e.Bbox.Authenticate(); err != nil {
		ch <- prometheus.MustNewConstMetric(
			up, prometheus.GaugeValue, 0,
		)
		level.Error(e.logger).Log("msg", "Bbox authentication error: %s", err.Error())
		return
	}

	resp, err := e.Bbox.GetMetrics()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(
			up, prometheus.GaugeValue, 0,
		)
		level.Error(e.logger).Log("msg", "Bbox API error: %s", err.Error())
		return
	}
	level.Info(e.logger).Log("msg", "Bbox metrics retrieved")
	storeServicesMetrics(ch, resp.Services)
	storeDeviceMetrics(ch, resp.Device)
	storeDNSMetrics(ch, resp.DNS)
	storeLanMetrics(e.logger, ch, resp.Lan)
	storeWanMetrics(ch, resp.Wan)
	storeWanFtthMetric(ch, resp.FtthState)
	storeWirelessMetrics(ch, resp.Wireless)
	storeIPTVMetrics(ch, resp.IPTV)
	ch <- prometheus.MustNewConstMetric(
		up, prometheus.GaugeValue, 1,
	)
	level.Info(e.logger).Log("msg", "BBox exporter finished")
}

func storeMetric(ch chan<- prometheus.Metric, value float64, desc *prometheus.Desc, labels ...string) {
	ch <- prometheus.MustNewConstMetric(
		desc, prometheus.GaugeValue, value, labels...)
}

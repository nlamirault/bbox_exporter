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
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/nlamirault/bbox_exporter/bbox"
)

var (
	ftthState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_ftth_state"),
		"LinkState of the GEth FTTH port",
		nil, nil,
	)
	txBytesWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_bytes"),
		"TX bytes",
		nil, nil,
	)
	txPacketsWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_packets"),
		"TX packets",
		nil, nil,
	)
	txPacketsErrorsWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_packets_errors"),
		"TX packets in error",
		nil, nil,
	)
	txPacketsDiscardsWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_packets_discards"),
		"TX packets discards",
		nil, nil,
	)
	txLineOccupationWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_line_occupation"),
		"TX line occupation",
		nil, nil,
	)
	txBandwidthWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_bandwidth"),
		"TX bandwith available",
		nil, nil,
	)
	txBandwidthMaxWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_bandwidth_max"),
		"TX maximum bandwith available",
		nil, nil,
	)

	rxBytesWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_bytes"),
		"RX bytes",
		nil, nil,
	)
	rxPacketsWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_packets"),
		"RX packets",
		nil, nil,
	)
	rxPacketsErrorsWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_packets_errors"),
		"RX packets in error",
		nil, nil,
	)
	rxPacketsDiscardsWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_packets_discards"),
		"RX packets discards",
		nil, nil,
	)
	rxLineOccupationWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_line_occupation"),
		"RX line occupation",
		nil, nil,
	)
	rxBandwidthWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_bandwidth"),
		"RX bandwith available",
		nil, nil,
	)
	rxBandwidthMaxWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_bandwidth_max"),
		"RX bandwith available",
		nil, nil,
	)
)

func describeWanMetrics(ch chan<- *prometheus.Desc) {
	ch <- ftthState
	ch <- txBytesWan
	ch <- txPacketsWan
	ch <- txPacketsErrorsWan
	ch <- txPacketsDiscardsWan
	ch <- txLineOccupationWan
	ch <- txBandwidthWan
	ch <- txBandwidthMaxWan
	ch <- rxBytesWan
	ch <- rxPacketsWan
	ch <- rxPacketsErrorsWan
	ch <- rxPacketsDiscardsWan
	ch <- rxLineOccupationWan
	ch <- rxBandwidthWan
	ch <- rxBandwidthMaxWan
}

func storeWanMetrics(ch chan<- prometheus.Metric, metrics bbox.WanIPStatistics) {
	log.Info("Store WAN metrics")
	storeMetric(ch, metrics.WAN.IP.Stats.Tx.Bytes, txBytesWan)
	storeMetric(ch, metrics.WAN.IP.Stats.Tx.Packets, txPacketsWan)
	storeMetric(ch, metrics.WAN.IP.Stats.Tx.Packetserrors, txPacketsErrorsWan)
	storeMetric(ch, metrics.WAN.IP.Stats.Tx.Packetsdiscards, txPacketsDiscardsWan)
	storeMetric(ch, metrics.WAN.IP.Stats.Rx.Bytes, rxBytesWan)
	storeMetric(ch, metrics.WAN.IP.Stats.Rx.Packets, rxPacketsWan)
	storeMetric(ch, metrics.WAN.IP.Stats.Rx.Packetserrors, rxPacketsErrorsWan)
	storeMetric(ch, metrics.WAN.IP.Stats.Rx.Packetsdiscards, rxPacketsDiscardsWan)
}

func storeWanFtthMetric(ch chan<- prometheus.Metric, metric string) {
	fftStateValue := float64(0)
	if strings.ToUpper(metric) == "UP" {
		fftStateValue = float64(1)
	}
	ch <- prometheus.MustNewConstMetric(
		ftthState, prometheus.GaugeValue, fftStateValue,
	)
}

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
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/nlamirault/bbox_exporter/bbox"
)

var (
	txBytesLan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "lan_transmitted_bytes"),
		"TX bytes",
		nil, nil,
	)
	txPacketsLan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "lan_transmitted_packets"),
		"TX packets",
		nil, nil,
	)
	txPacketsErrorsLan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "lan_transmitted_packets_errors"),
		"TX packets in error",
		nil, nil,
	)
	txPacketsDiscardsLan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "lan_transmitted_packets_discards"),
		"TX packets discards",
		nil, nil,
	)

	rxBytesLan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "lan_received_bytes"),
		"RX bytes",
		nil, nil,
	)
	rxPacketsLan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "lan_received_packets"),
		"RX packets",
		nil, nil,
	)
	rxPacketsErrorsLan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "lan_received_packets_errors"),
		"RX packets in error",
		nil, nil,
	)
	rxPacketsDiscardsLan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "lan_received_packets_discards"),
		"RX packets discards",
		nil, nil,
	)
)

func describeLanMetrics(ch chan<- *prometheus.Desc) {
	ch <- txBytesLan
	ch <- txPacketsLan
	ch <- txPacketsErrorsLan
	ch <- txPacketsDiscardsLan
	ch <- rxBytesLan
	ch <- rxPacketsLan
	ch <- rxPacketsErrorsLan
	ch <- rxPacketsDiscardsLan
}

func storeLanMetrics(ch chan<- prometheus.Metric, metrics bbox.LanMetrics) {
	log.Info("Store LAN metrics")
	// storeMetric(ch, metrics.Statistics[0].Lan.Stats.Tx.Bytes, txBytesLan)
	if val, err := strconv.ParseFloat(metrics.Statistics[0].Lan.Stats.Tx.Bytes, 64); err == nil {
		storeMetric(ch, val, txBytesLan)
	}
	storeMetric(ch, metrics.Statistics[0].Lan.Stats.Tx.Packets, txPacketsLan)
	storeMetric(ch, metrics.Statistics[0].Lan.Stats.Tx.Packetserrors, txPacketsErrorsLan)
	storeMetric(ch, metrics.Statistics[0].Lan.Stats.Tx.Packetsdiscards, txPacketsDiscardsLan)
	// storeMetric(ch, metrics.Statistics[0].Lan.Stats.Rx.Bytes, rxBytesLan)
	if val, err := strconv.ParseFloat(metrics.Statistics[0].Lan.Stats.Rx.Bytes, 64); err == nil {
		storeMetric(ch, val, rxBytesLan)
	}
	storeMetric(ch, metrics.Statistics[0].Lan.Stats.Rx.Packets, rxPacketsLan)
	storeMetric(ch, metrics.Statistics[0].Lan.Stats.Rx.Packetserrors, rxPacketsErrorsLan)
	storeMetric(ch, metrics.Statistics[0].Lan.Stats.Rx.Packetsdiscards, rxPacketsDiscardsLan)
}

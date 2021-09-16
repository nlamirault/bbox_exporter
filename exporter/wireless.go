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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/nlamirault/bbox_exporter/bbox"
)

var (
	txBytesWireless = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wireless_transmitted_bytes"),
		"TX bytes",
		[]string{"frequency"}, nil,
	)
	txPacketsWireless = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wireless_transmitted_packets"),
		"TX packets",
		[]string{"frequency"}, nil,
	)
	txPacketsErrorsWireless = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wireless_transmitted_packets_errors"),
		"TX packets in error",
		[]string{"frequency"}, nil,
	)
	txPacketsDiscardsWireless = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wireless_transmitted_packets_discards"),
		"TX packets discards",
		[]string{"frequency"}, nil,
	)

	rxBytesWireless = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wireless_received_bytes"),
		"RX bytes",
		[]string{"frequency"}, nil,
	)
	rxPacketsWireless = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wireless_received_packets"),
		"RX packets",
		[]string{"frequency"}, nil,
	)
	rxPacketsErrorsWireless = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wireless_received_packets_errors"),
		"RX packets in error",
		[]string{"frequency"}, nil,
	)
	rxPacketsDiscardsWireless = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wireless_received_packets_discards"),
		"RX packets discards",
		[]string{"frequency"}, nil,
	)
)

func describeWirelessMetrics(ch chan<- *prometheus.Desc) {
	ch <- hosts
	ch <- txBytesWireless
	ch <- txPacketsWireless
	ch <- txPacketsErrorsWireless
	ch <- txPacketsDiscardsWireless
	ch <- rxBytesWireless
	ch <- rxPacketsWireless
	ch <- rxPacketsErrorsWireless
	ch <- rxPacketsDiscardsWireless
}

func storeWirelessMetrics(ch chan<- prometheus.Metric, metrics bbox.WirelessMetrics) {
	log.Info("Store Wireless metrics")
	// storeMetric(ch, float64(metrics.Wireless5GhzStatistics[0].Wireless.SSID.Stats.Tx.Bytes), txBytesWireless, "5ghz")
	// storeMetric(ch, float64(metrics.Wireless5GhzStatistics[0].Wireless.SSID.Stats.Tx.Packets), txPacketsWireless, "5ghz")
	// storeMetric(ch, float64(metrics.Wireless5GhzStatistics[0].Wireless.SSID.Stats.Tx.Packetserrors), txPacketsErrorsWireless, "5ghz")
	// storeMetric(ch, float64(metrics.Wireless5GhzStatistics[0].Wireless.SSID.Stats.Tx.Packetsdiscards), txPacketsDiscardsWireless, "5ghz")
	// storeMetric(ch, float64(metrics.Wireless5GhzStatistics[0].Wireless.SSID.Stats.Rx.Bytes), rxBytesWireless, "5ghz")
	// storeMetric(ch, float64(metrics.Wireless5GhzStatistics[0].Wireless.SSID.Stats.Rx.Packets), rxPacketsWireless, "5ghz")
	// storeMetric(ch, float64(metrics.Wireless5GhzStatistics[0].Wireless.SSID.Stats.Rx.Packetserrors), rxPacketsErrorsWireless, "5ghz")
	// storeMetric(ch, float64(metrics.Wireless5GhzStatistics[0].Wireless.SSID.Stats.Rx.Packetsdiscards), rxPacketsDiscardsWireless, "5ghz")
	storeMetric(ch, float64(metrics.Wireless24GhzStatistics[0].Wireless.SSID.Stats.Tx.Bytes), txBytesWireless, "24ghz")
	storeMetric(ch, float64(metrics.Wireless24GhzStatistics[0].Wireless.SSID.Stats.Tx.Packets), txPacketsWireless, "24ghz")
	storeMetric(ch, float64(metrics.Wireless24GhzStatistics[0].Wireless.SSID.Stats.Tx.Packetserrors), txPacketsErrorsWireless, "24ghz")
	storeMetric(ch, float64(metrics.Wireless24GhzStatistics[0].Wireless.SSID.Stats.Tx.Packetsdiscards), txPacketsDiscardsWireless, "24ghz")
	storeMetric(ch, float64(metrics.Wireless24GhzStatistics[0].Wireless.SSID.Stats.Rx.Bytes), rxBytesWireless, "24ghz")
	storeMetric(ch, float64(metrics.Wireless24GhzStatistics[0].Wireless.SSID.Stats.Rx.Packets), rxPacketsWireless, "24ghz")
	storeMetric(ch, float64(metrics.Wireless24GhzStatistics[0].Wireless.SSID.Stats.Rx.Packetserrors), rxPacketsErrorsWireless, "24ghz")
	storeMetric(ch, float64(metrics.Wireless24GhzStatistics[0].Wireless.SSID.Stats.Rx.Packetsdiscards), rxPacketsDiscardsWireless, "24ghz")
}

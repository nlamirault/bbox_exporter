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

	"github.com/nlamirault/bbox_exporter/bbox"
)

var (
	hosts = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "lan_connected_devices"),
		"Number of devices connected",
		[]string{"link"}, nil,
	)

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
	ch <- hosts
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
	// storeMetric(ch, float64(len(metrics.Devices[0].Hosts.List)), hosts)
	lanHosts := map[string]int{}
	for _, host := range metrics.Devices[0].Hosts.List {
		// log.Infof("Host: %s, IP: %s %s %s => [%s]", host.Hostname, host.Ipaddress, host.Type, host.Link, host.Active)
		if host.Active == 1 {
			lanHosts[host.Link] = lanHosts[host.Link] + 1
		}
	}
	for link, val := range lanHosts {
		storeMetric(ch, float64(val), hosts, link)
	}
	// log.Infof("%+v", metrics[0].Hosts.List[0])
	storeMetric(ch, float64(metrics.Statistics[0].Lan.Stats.Tx.Bytes), txBytesLan)
	storeMetric(ch, float64(metrics.Statistics[0].Lan.Stats.Tx.Packets), txPacketsLan)
	storeMetric(ch, float64(metrics.Statistics[0].Lan.Stats.Tx.Packetserrors), txPacketsErrorsLan)
	storeMetric(ch, float64(metrics.Statistics[0].Lan.Stats.Tx.Packetsdiscards), txPacketsDiscardsLan)
	storeMetric(ch, float64(metrics.Statistics[0].Lan.Stats.Rx.Bytes), rxBytesLan)
	storeMetric(ch, float64(metrics.Statistics[0].Lan.Stats.Rx.Packets), rxPacketsLan)
	storeMetric(ch, float64(metrics.Statistics[0].Lan.Stats.Rx.Packetserrors), rxPacketsErrorsLan)
	storeMetric(ch, float64(metrics.Statistics[0].Lan.Stats.Rx.Packetsdiscards), rxPacketsDiscardsLan)
}

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
	diagnosticsMinWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_diagnostics_min"),
		"Minimum response Time",
		[]string{"mode"}, nil,
	)
	diagnosticsMaxWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_diagnostics_max"),
		"Maximum response Time",
		[]string{"mode"}, nil,
	)
	diagnosticsAvgWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_diagnostics_avg"),
		"Average response Time",
		[]string{"mode"}, nil,
	)
	diagnosticsNumberOfSuccess = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_diagnostics_success"),
		"Number of sucess",
		[]string{"mode"}, nil,
	)
	diagnosticsNumberOfError = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_diagnostics_error"),
		"Number of error",
		[]string{"mode"}, nil,
	)
	diagnosticsNumberOfTries = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_diagnostics_tries"),
		"Number of tries",
		[]string{"mode"}, nil,
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
	ch <- diagnosticsMinWan
	ch <- diagnosticsMaxWan
	ch <- diagnosticsAvgWan
	ch <- diagnosticsNumberOfSuccess
	ch <- diagnosticsNumberOfError
	ch <- diagnosticsNumberOfTries
}

func storeWanMetrics(ch chan<- prometheus.Metric, metrics bbox.WanMetrics) {
	log.Info("Store WAN metrics")
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.Bytes), txBytesWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.Packets), txPacketsWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.Packetserrors), txPacketsErrorsWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.Packetsdiscards), txPacketsDiscardsWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.Occupation), txLineOccupationWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.Bandwidth), txBandwidthWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.MaxBandwidth), txBandwidthMaxWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.Bytes), rxBytesWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.Packets), rxPacketsWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.Packetserrors), rxPacketsErrorsWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.Packetsdiscards), rxPacketsDiscardsWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.Occupation), rxLineOccupationWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.Bandwidth), rxBandwidthWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.MaxBandwidth), rxBandwidthMaxWan)

	log.Infof("%+v", metrics.DiagnosticsStatistics[0].Diags.Ping[0])
	for _, val := range metrics.DiagnosticsStatistics[0].Diags.DNS {
		if val.Tries > 0 {
			storeMetric(ch, float64(val.Min), diagnosticsMinWan, "dns")
			storeMetric(ch, float64(val.Max), diagnosticsMaxWan, "dns")
			storeMetric(ch, float64(val.Average), diagnosticsAvgWan, "dns")
			break
		}
	}
	for _, val := range metrics.DiagnosticsStatistics[0].Diags.Ping {
		if val.Tries > 0 {
			storeMetric(ch, float64(val.Min), diagnosticsMinWan, "ping")
			storeMetric(ch, float64(val.Max), diagnosticsMaxWan, "ping")
			storeMetric(ch, float64(val.Average), diagnosticsAvgWan, "ping")
			break
		}
	}
	for _, val := range metrics.DiagnosticsStatistics[0].Diags.HTTP {
		if val.Tries > 0 {
			storeMetric(ch, float64(val.Min), diagnosticsMinWan, "http")
			storeMetric(ch, float64(val.Max), diagnosticsMaxWan, "http")
			storeMetric(ch, float64(val.Average), diagnosticsAvgWan, "http")
			break
		}
	}
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

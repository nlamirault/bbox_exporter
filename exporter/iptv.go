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
	iptvChannel = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "iptv_channel"),
		"Name of channel",
		[]string{"name"}, nil,
	)
)

func describeIPTVMetrics(ch chan<- *prometheus.Desc) {
	ch <- iptvChannel
}

func storeIPTVMetrics(ch chan<- prometheus.Metric, metrics bbox.IPTVMetrics) {
	// storeMetric(ch, float64(metrics.Informations[0].IPTV[0].Receipt), iptvChannel, metrics.Informations[0].IPTV[0].Name)
}

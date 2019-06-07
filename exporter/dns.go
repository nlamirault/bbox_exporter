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

var (
	dnsNumberOfQueries = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dns_number_of_queries"),
		"Number of queries",
		nil, nil,
	)
	dnsMin = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dns_min"),
		"Minimun of average dns response time",
		nil, nil,
	)
	dnsMax = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dns_max"),
		"Maximun of average dns response time",
		nil, nil,
	)
	dnsAverage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dns_average"),
		"Average of average dns response time",
		nil, nil,
	)
)

func describeDNSMetrics(ch chan<- *prometheus.Desc) {
	ch <- dnsNumberOfQueries
	ch <- dnsMin
	ch <- dnsMax
	ch <- dnsAverage
}

func storeDNSMetrics(ch chan<- prometheus.Metric, metrics bbox.DNSMetrics) {
	log.Info("Store DNS metrics")
	storeMetric(ch, metrics.Principal[0].DNS.NumberOfQueries, dnsNumberOfQueries)
	storeMetric(ch, metrics.Principal[0].DNS.Min, dnsMin)
	storeMetric(ch, metrics.Principal[0].DNS.Max, dnsMax)
	storeMetric(ch, metrics.Principal[0].DNS.Average, dnsAverage)
}
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
	serviceUp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "service_status"),
		"BBox services status",
		[]string{"name"}, nil,
	)
)

func describeServicesMetrics(ch chan<- *prometheus.Desc) {
	ch <- serviceUp
}

func storeServicesMetrics(ch chan<- prometheus.Metric, metrics bbox.ServicesMetrics) {
	log.Info("Store Services metrics")
	storeMetric(ch, float64(metrics.Informations[0].Services.Firewall.Enable), serviceUp, "firewall")
	storeMetric(ch, float64(metrics.Informations[0].Services.Dyndns.Enable), serviceUp, "dydns")
	storeMetric(ch, float64(metrics.Informations[0].Services.Dhcp.Enable), serviceUp, "dhcp")
	storeMetric(ch, float64(metrics.Informations[0].Services.Nat.Enable), serviceUp, "nat")
	storeMetric(ch, float64(metrics.Informations[0].Services.Gamermode.Enable), serviceUp, "gamermode")
	storeMetric(ch, float64(metrics.Informations[0].Services.Upnp.Igd.Enable), serviceUp, "upnp")
	storeMetric(ch, float64(metrics.Informations[0].Services.Remote.Proxywol.Enable), serviceUp, "remote_proxywol")
	storeMetric(ch, float64(metrics.Informations[0].Services.Remote.Admin.Enable), serviceUp, "remote_admin")
	storeMetric(ch, float64(metrics.Informations[0].Services.Parentalcontrol.Enable), serviceUp, "parentalcontrol")
	storeMetric(ch, float64(metrics.Informations[0].Services.Wifischeduler.Enable), serviceUp, "wifischeduler")
	storeMetric(ch, float64(metrics.Informations[0].Services.Voipscheduler.Enable), serviceUp, "voipscheduler")
	storeMetric(ch, float64(metrics.Informations[0].Services.Notification.Enable), serviceUp, "notification")
	storeMetric(ch, float64(metrics.Informations[0].Services.Hotspot.Enable), serviceUp, "hotspot")
	storeMetric(ch, float64(metrics.Informations[0].Services.Usb.Samba.Enable), serviceUp, "usb_samba")
	storeMetric(ch, float64(metrics.Informations[0].Services.Usb.Printer.Enable), serviceUp, "usb_printer")
	storeMetric(ch, float64(metrics.Informations[0].Services.Usb.Dlna.Enable), serviceUp, "user_dlna")

}

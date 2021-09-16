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
	deviceModelName = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_model_name"),
		"Device model name",
		[]string{"model_name"}, nil,
	)
	deviceUsing = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_fai_usage"),
		"FAI box usage",
		[]string{"using"}, nil,
	)
	deviceStatus = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_status"),
		"Current status",
		nil, nil,
	)
	deviceNumberOfBoots = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_number_of_boots"),
		"Number of boots since last reset to factory default",
		nil, nil,
	)
	deviceTemperature = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_temperature"),
		"Current internal temperature in °C",
		nil, nil,
	)

	deviceMemory = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_memory"),
		"Memory in kB",
		[]string{"type"}, nil,
	)

	deviceCPU = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_cpu"),
		"CPU Total Time",
		[]string{"mode"}, nil,
	)

	deviceProcess = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_process"),
		"Device process",
		[]string{"type"}, nil,
	)
)

func describeDeviceMetrics(ch chan<- *prometheus.Desc) {
	ch <- deviceModelName
	ch <- deviceUsing
	ch <- deviceStatus
	ch <- deviceNumberOfBoots
	ch <- deviceTemperature
	ch <- deviceMemory
	ch <- deviceCPU
	ch <- deviceProcess
}

func storeDeviceMetrics(ch chan<- prometheus.Metric, metrics bbox.DeviceMetrics) {
	storeMetric(ch, 1.0, deviceModelName, metrics.Informations[0].Device.ModelName)
	storeMetric(ch, float64(metrics.Informations[0].Device.Using.IPv4), deviceUsing, "ipv4")
	storeMetric(ch, float64(metrics.Informations[0].Device.Using.IPv6), deviceUsing, "ipv6")
	storeMetric(ch, float64(metrics.Informations[0].Device.Using.FTTH), deviceUsing, "ftth")
	storeMetric(ch, float64(metrics.Informations[0].Device.Using.ADSL), deviceUsing, "adsl")
	storeMetric(ch, float64(metrics.Informations[0].Device.Using.VDSL), deviceUsing, "vdsl")
	storeMetric(ch, metrics.Informations[0].Device.Status, deviceStatus)
	storeMetric(ch, metrics.Informations[0].Device.NumberOfBoots, deviceNumberOfBoots)
	storeMetric(ch, metrics.Informations[0].Device.Temperature.Current, deviceTemperature)
	storeMetric(ch, metrics.Memory[0].Device.Memory.Total, deviceMemory, "total")
	storeMetric(ch, metrics.Memory[0].Device.Memory.Free, deviceMemory, "free")
	storeMetric(ch, metrics.Memory[0].Device.Memory.Cached, deviceMemory, "cached")
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.Total, deviceCPU, "total")
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.User, deviceCPU, "user")
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.Nice, deviceCPU, "nice")
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.System, deviceCPU, "system")
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.IO, deviceCPU, "io")
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.Idle, deviceCPU, "idle")
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.Irq, deviceCPU, "irq")
	storeMetric(ch, metrics.CPU[0].Device.CPU.Process.Created, deviceProcess, "created")
	storeMetric(ch, metrics.CPU[0].Device.CPU.Process.Running, deviceProcess, "running")
	storeMetric(ch, metrics.CPU[0].Device.CPU.Process.Blocked, deviceProcess, "blocked")
}

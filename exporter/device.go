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

	deviceMemoryTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_memory_total"),
		"Total memory in kB",
		nil, nil,
	)
	deviceMemoryFree = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_memory_free"),
		"Free memory in kB",
		nil, nil,
	)
	deviceMemoryCached = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_memory_cached"),
		"Cached memory in kB",
		nil, nil,
	)

	deviceCPUTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_cpu_total"),
		"CPU Total Time",
		nil, nil,
	)
	deviceCPUUser = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_cpu_user"),
		"CPU User Time",
		nil, nil,
	)
	deviceCPUNice = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_cpu_nice"),
		"CPU Nice Time",
		nil, nil,
	)
	deviceCPUSystem = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_cpu_system"),
		"CPU System Time",
		nil, nil,
	)
	deviceCPUIO = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_cpu_io"),
		"CPU IO Time",
		nil, nil,
	)
	deviceCPUIdle = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_cpu_idle"),
		"CPU Idle Time",
		nil, nil,
	)
	deviceCPUIrq = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_cpu_irq"),
		"CPU Irq Time",
		nil, nil,
	)

	deviceProcessCreated = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_process_created"),
		"Number of created processus",
		nil, nil,
	)
	deviceProcessRunning = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_process_running"),
		"Number of running processus",
		nil, nil,
	)
	deviceProcessBlocked = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_process_blocked"),
		"Number of blocked processus",
		nil, nil,
	)

)

func describeDeviceMetrics(ch chan<- *prometheus.Desc) {
	ch <- deviceStatus
	ch <- deviceNumberOfBoots
	ch <- deviceTemperature
	ch <- deviceCPUTotal
	ch <- deviceCPUUser
	ch <- deviceCPUNice
	ch <- deviceCPUSystem
	ch <- deviceCPUIO
	ch <- deviceCPUIdle
	ch <- deviceCPUIrq
}

func storeDeviceMetrics(ch chan<- prometheus.Metric, metrics bbox.DeviceMetrics) {
	log.Info("Store Device metrics")
	storeMetric(ch, metrics.Informations[0].Device.Status, deviceStatus)
	storeMetric(ch, metrics.Informations[0].Device.NumberOfBoots, deviceNumberOfBoots)
	storeMetric(ch, metrics.Informations[0].Device.Temperature.Current, deviceTemperature)
	storeMetric(ch, metrics.Memory[0].Device.Memory.Total, deviceMemoryTotal)
	storeMetric(ch, metrics.Memory[0].Device.Memory.Free, deviceMemoryFree)
	storeMetric(ch, metrics.Memory[0].Device.Memory.Cached, deviceMemoryCached)
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.Total, deviceCPUTotal)
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.User, deviceCPUUser)
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.Nice, deviceCPUNice)
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.System, deviceCPUSystem)
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.IO, deviceCPUIO)
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.Idle, deviceCPUIdle)
	storeMetric(ch, metrics.CPU[0].Device.CPU.Time.Irq, deviceCPUIrq)
	storeMetric(ch, metrics.CPU[0].Device.CPU.Process.Created, deviceProcessCreated)
	storeMetric(ch, metrics.CPU[0].Device.CPU.Process.Running, deviceProcessRunning)
	storeMetric(ch, metrics.CPU[0].Device.CPU.Process.Blocked, deviceProcessBlocked)
}
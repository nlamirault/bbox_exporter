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

package bbox

import (
	"github.com/prometheus/common/log"
)

type DeviceMetrics struct {
	Informations []DeviceInformations `json:"informations"`
	Memory       []DeviceMemory       `json:"device"`
	CPU          []DeviceCPU          `json:"cpu"`
}

type DeviceInformations struct {
	Device struct {
		Now           string  `json:"now"`
		Status        float64 `json:"status"`
		NumberOfBoots float64 `json:"number_of_boots"`
		ModelName     string  `json:"modelname"`
		Temperature   struct {
			Current float64 `json:"current"`
			Status  string  `json:"status"`
		} `json:"temperature"`
		Using struct {
			IPv4 int `json:"ipv4"`
			IPv6 int `json:"ipv6"`
			FTTH int `json:"ftth"`
			ADSL int `json:"adsl"`
			VDSL int `json:"vdsl"`
		} `json:"using"`
	} `json:"device"`
}

type DeviceMemory struct {
	Device struct {
		Memory struct {
			Total  float64 `json:"total"`
			Free   float64 `json:"free"`
			Cached float64 `json:"cached"`
		} `json:"mem"`
	} `json:"device"`
}

type DeviceCPU struct {
	Device struct {
		CPU struct {
			Time struct {
				Total  float64 `json:"total"`
				User   float64 `json:"user"`
				Nice   float64 `json:"nice"`
				System float64 `json:"system"`
				IO     float64 `json:"io"`
				Idle   float64 `json:"idle"`
				Irq    float64 `json:"irq"`
			} `json:"time"`
			Process struct {
				Created float64 `json:"created"`
				Running float64 `json:"running"`
				Blocked float64 `json:"blocked"`
			} `json:"process"`
		} `json:"cpu"`
	} `json:"device"`
}

func (client *Client) getDeviceMetrics() (*DeviceMetrics, error) {
	var deviceStats DeviceMetrics

	informations, err := client.getDeviceInformations()
	if err != nil {
		return nil, err
	}
	deviceStats.Informations = informations

	cpu, err := client.getDeviceCPU()
	if err != nil {
		return nil, err
	}
	deviceStats.CPU = cpu

	memory, err := client.getDeviceMemory()
	if err != nil {
		return nil, err
	}
	deviceStats.Memory = memory

	return &deviceStats, nil
}

// getDeviceInformations returns Bbox information
// See: https://api.bbox.fr/doc/apirouter/#api-Device-GetDevice
func (client *Client) getDeviceInformations() ([]DeviceInformations, error) {
	log.Info("Retrieve device informations")
	var informations []DeviceInformations
	if err := client.apiRequest("/device", &informations); err != nil {
		return nil, err
	}
	return informations, nil
}

// getDeviceCPU returns Bbox CPU information
// See: https://api.bbox.fr/doc/apirouter/#api-Device-GetDeviceCPU
func (client *Client) getDeviceCPU() ([]DeviceCPU, error) {
	log.Info("Retrieve device CPU")
	var cpu []DeviceCPU
	if err := client.apiRequest("/device/cpu", &cpu); err != nil {
		return nil, err
	}
	return cpu, nil
}

// getDeviceMemory returns Bbox Memory information
// See: https://api.bbox.fr/doc/apirouter/#api-Device-GetDeviceMem
func (client *Client) getDeviceMemory() ([]DeviceMemory, error) {
	log.Info("Retrieve device memory")
	var memory []DeviceMemory
	if err := client.apiRequest("/device/mem", &memory); err != nil {
		return nil, err
	}
	return memory, nil
}

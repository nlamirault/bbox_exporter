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

package bbox

import (
	"fmt"

	"github.com/go-kit/kit/log/level"
)

type WirelessMetrics struct {
	Wireless5GhzStatistics  []WirelessStatistics
	Wireless24GhzStatistics []WirelessStatistics
}

// WirelessStatistics represents statistics information of the Bbox WIFI
type WirelessStatistics struct {
	Wireless struct {
		SSID struct {
			ID    interface{} `json:"id"`
			Stats struct {
				Rx struct {
					Packets         flexInt `json:"packets"`
					Bytes           flexInt `json:"bytes"` // See: https://github.com/nlamirault/bbox_exporter/issues/1
					Packetserrors   flexInt `json:"packetserrors"`
					Packetsdiscards flexInt `json:"packetsdiscards"`
				} `json:"rx"`
				Tx struct {
					Packets         flexInt `json:"packets"`
					Bytes           flexInt `json:"bytes"` // See: https://github.com/nlamirault/bbox_exporter/issues/1
					Packetserrors   flexInt `json:"packetserrors"`
					Packetsdiscards flexInt `json:"packetsdiscards"`
				} `json:"tx"`
			} `json:"stats"`
		} `json:"ssid"`
	} `json:"wireless"`
}

func (client *Client) getWirelessMetrics() (*WirelessMetrics, error) {
	var metrics WirelessMetrics

	// wifi5Ghz, err := client.getWirelessStatistics("5")
	// if err != nil {
	// 	return nil, err
	// }
	// metrics.Wireless5GhzStatistics = wifi5Ghz

	wifi24Ghz, err := client.getWirelessStatistics("24")
	if err != nil {
		return nil, err
	}
	metrics.Wireless24GhzStatistics = wifi24Ghz

	return &metrics, nil
}

func (client *Client) getWirelessStatistics(which string) ([]WirelessStatistics, error) {
	level.Info(client.logger).Log("msg", "Retrieve WIFI %sGhz metrics from Bbox", which)

	var metrics []WirelessStatistics
	if err := client.apiRequest(fmt.Sprintf("/wireless/%s/stats", which), &metrics); err != nil {
		return nil, err
	}
	return metrics, nil
}

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

import "github.com/go-kit/kit/log/level"

type IPTVMetrics struct {
	Informations []IPTVInformations `json:"informations"`
	// Diagnostics  []IPTVDiagnostics  `json:"diagnostics"`
}

type IPTVInformations struct {
	IPTV []struct {
		// the IP Address of the multicast
		Address string `json:"address"`
		// the IP Address of receiver
		Ipaddress string `json:"ipaddress"`
		// image name of the logo
		Logo string `json:"logo"`
		// the offset for the logo
		// Logooffset string `json:"logooffset"`
		// The channel name
		Name string `json:"name"`
		// the channel number
		Number int `json:"number"`
		// Defines if the channel is really received or not
		Receipt int `json:"receipt"`
		// Channel Id in the epg
		Epgid int `json:"epgid"`
	} `json:"iptv"`
	Now string `json:"now"`
}

func (client *Client) getIPTVMetrics() (*IPTVMetrics, error) {
	var metrics IPTVMetrics

	informations, err := client.getIPTVInformations()
	if err != nil {
		return nil, err
	}
	metrics.Informations = informations

	return &metrics, nil
}

func (client *Client) getIPTVInformations() ([]IPTVInformations, error) {
	level.Info(client.logger).Log("msg", "Retrieve IP TV informations")
	var iptvInformations []IPTVInformations
	if err := client.apiRequest("/iptv", &iptvInformations); err != nil {
		return nil, err
	}
	return iptvInformations, nil
}

// func (client *Client) getIPTVDiagnostic() ([]IPTVInformations, error) {
// 	level.Info(client.logger).Log("msg", "Retrieve IP TV diagnostic")
// 	if err := client.apiRequest("/iptv/diags", nil); err != nil {
// 		return nil, err
// 	}
// 	return nil, nil
// }

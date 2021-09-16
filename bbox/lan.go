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
	"github.com/go-kit/kit/log/level"
)

type LanMetrics struct {
	IPInformations []LanIPInformations `json:"ip_informations"`
	Devices        []LanDevice         `json:"devices"`
	Statistics     []LanStatistics     `json:"statistics"`
}

type LanDevice struct {
	Hosts struct {
		List []LanHost `json:"list"`
	} `json:"hosts"`
}

type LanHost struct {
	ID         int           `json:"id"`
	Hostname   string        `json:"hostname"`
	Macaddress string        `json:"macaddress"`
	Ipaddress  string        `json:"ipaddress"`
	Type       string        `json:"type"`
	Link       string        `json:"link"`
	Devicetype string        `json:"devicetype"`
	Firstseen  string        `json:"firstseen"`
	Lastseen   interface{}   `json:"lastseen"`
	IP6Address []interface{} `json:"ip6address"`
	Ethernet   struct {
		Physicalport int    `json:"physicalport"`
		Logicalport  int    `json:"logicalport"`
		Speed        int    `json:"speed"`
		Mode         string `json:"mode"`
	} `json:"ethernet"`
	Stb struct {
		Product string `json:"product"`
		Serial  string `json:"serial"`
	} `json:"stb,omitempty"`
	Wireless struct {
		Band       string      `json:"band"`
		Rssi0      interface{} `json:"rssi0"` // String or int ? "rssi0":"-76","rssi1":0,"rssi2":0
		Rssi1      interface{} `json:"rssi1"`
		Rssi2      interface{} `json:"rssi2"`
		Mcs        interface{} `json:"mcs"` // Same string or int ??
		Rate       interface{} `json:"rate"`
		Idle       interface{} `json:"idle"`
		Wexindex   interface{} `json:"wexindex"`
		Starealmac interface{} `json:"starealmac"`
	} `json:"wireless"`
	Plc struct {
		Rxphyrate        string `json:"rxphyrate"`
		Txphyrate        string `json:"txphyrate"`
		Associateddevice int    `json:"associateddevice"`
		Interface        int    `json:"interface"`
		Ethernetspeed    int    `json:"ethernetspeed"`
	} `json:"plc"`
	Lease           flexInt `json:"lease"`
	Active          int     `json:"active"`
	Parentalcontrol struct {
		Enable          int    `json:"enable"`
		Status          string `json:"status"`
		StatusRemaining int    `json:"statusRemaining"`
		StatusUntil     string `json:"statusUntil"`
	} `json:"parentalcontrol"`
	Ping struct {
		Average int `json:"average"`
	} `json:"ping"`
	Scan struct {
		Services []interface{} `json:"services"`
	} `json:"scan"`
}

// LanStatistics represents statistics information of the Bbox LAN
type LanStatistics struct {
	Lan struct {
		Stats struct {
			Rx struct {
				Packets         flexInt `json:"packets"`
				Bytes           flexInt `json:"bytes"`
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
	} `json:"lan"`
}

type LanIPInformations struct {
	Lan struct {
		IP struct {
			State      string        `json:"state"`
			Mtu        int           `json:"mtu"`
			Ipaddress  string        `json:"ipaddress"`
			IP6Enable  int           `json:"ip6enable"`
			IP6State   string        `json:"ip6state"`
			IP6Address []interface{} `json:"ip6address"`
			IP6Prefix  []interface{} `json:"ip6prefix"`
			Netmask    string        `json:"netmask"`
			Mac        string        `json:"mac"`
			Hostname   string        `json:"hostname"`
			Domain     string        `json:"domain"`
			Aliases    string        `json:"aliases"`
		} `json:"ip"`
		Switch struct {
			Ports []struct {
				ID         int    `json:"id"`
				State      string `json:"state"`
				LinkMode   string `json:"link_mode"`
				Blocked    int    `json:"blocked"`
				Flickering int    `json:"flickering"`
			} `json:"ports"`
		} `json:"switch"`
	} `json:"lan"`
}

func (client *Client) getLanMetrics() (*LanMetrics, error) {
	var metrics LanMetrics

	lanStats, err := client.getLanStatistics()
	if err != nil {
		return nil, err
	}
	metrics.Statistics = lanStats

	devices, err := client.getLanDevices()
	if err != nil {
		return nil, err
	}
	metrics.Devices = devices

	return &metrics, nil
}

// returns ip configuration of the Bbox local Network.
// See: https://api.bbox.fr/doc/apirouter/#api-LAN-GetLanIP
func (client *Client) getLanInformations() ([]LanIPInformations, error) {
	level.Info(client.logger).Log("msg", "Retrieve LAN IP informations from Bbox")
	var informations []LanIPInformations
	if err := client.apiRequest("/lan/ip", &informations); err != nil {
		return nil, err
	}
	return informations, nil
}

// getLanDevices returns information on all devices connected to the Bbox.
// See: https://api.bbox.fr/doc/apirouter/#api-LAN-GetHosts
func (client *Client) getLanDevices() ([]LanDevice, error) {
	level.Info(client.logger).Log("msg", "Retrieve LAN devices from Bbox")
	var metrics []LanDevice
	if err := client.apiRequest("/hosts", &metrics); err != nil {
		return nil, err
	}
	return metrics, nil
}

// getLanStatistics returns statistics of the Bbox local Network.
// See: https://api.bbox.fr/doc/apirouter/#api-LAN-GetLanStats
func (client *Client) getLanStatistics() ([]LanStatistics, error) {
	level.Info(client.logger).Log("msg", "Retrieve LAN IP statistics")
	var metrics []LanStatistics
	if err := client.apiRequest("/lan/stats", &metrics); err != nil {
		return nil, err
	}
	return metrics, nil
}

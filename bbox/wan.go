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

type WanMetrics struct {
	IPInformations []WanIPInformations `json:"ip_informations"`
	IPStatistics   []WanIPStatistics   `json:"ip_statistics"`
	FtthStatistics *FtthStatistics     `json:"ftth_statistics"`
}

type WanIPStatistics struct {
	WAN struct {
		IP struct {
			Stats struct {
				Rx struct {
					Packets         float64 `json:"packets"`
					Bytes           string  `json:"bytes"` // See: https://github.com/nlamirault/bbox_exporter/issues/1
					Packetserrors   float64 `json:"packetserrors"`
					Packetsdiscards float64 `json:"packetsdiscards"`
					Occupation      float64 `json:"occupation"`
					Bandwidth       float64 `json:"bandwidth"`
					MaxBandwidth    float64 `json:"maxBandwidth"`
				} `json:"rx"`
				Tx struct {
					Packets         float64 `json:"packets"`
					Bytes           string  `json:"bytes"` // See: https://github.com/nlamirault/bbox_exporter/issues/1
					Packetserrors   float64 `json:"packetserrors"`
					Packetsdiscards float64 `json:"packetsdiscards"`
					Occupation      float64 `json:"occupation"`
					Bandwidth       float64 `json:"bandwidth"`
					MaxBandwidth    float64 `json:"maxBandwidth"`
				} `json:"tx"`
			} `json:"stats"`
		} `json:"ip"`
	} `json:"wan"`
}

type FtthStatistics []struct {
	Ftth Ftth `json:"ftth"`
}

type Ftth struct {
	Wan struct {
		Ftth struct {
			Mode  string `json:"mode"`
			State string `json:"state"`
		} `json:"ftth"`
	} `json:"wan"`
}

type WanIPInformations struct {
	Wan struct {
		Internet struct {
			State int `json:"state"`
		} `json:"internet"`
		Interface struct {
			ID      int `json:"id"`
			Default int `json:"default"`
			State   int `json:"state"`
		} `json:"interface"`
		IP struct {
			Address    string        `json:"address"`
			State      string        `json:"state"`
			Gateway    string        `json:"gateway"`
			Dnsservers string        `json:"dnsservers"`
			Subnet     string        `json:"subnet"`
			IP6State   string        `json:"ip6state"`
			IP6Address []interface{} `json:"ip6address"`
			IP6Prefix  []interface{} `json:"ip6prefix"`
			Mac        string        `json:"mac"`
			Mtu        int           `json:"mtu"`
		} `json:"ip"`
		Link struct {
			State string `json:"state"`
			Type  string `json:"type"`
		} `json:"link"`
	} `json:"wan"`
}

func (client *Client) getWanMetrics() (*WanMetrics, error) {
	var metrics WanMetrics

	wanIPInformations, err := client.getWanInformations()
	if err != nil {
		return nil, err
	}
	metrics.IPInformations = wanIPInformations

	wanIPStats, err := client.getWanStatistics()
	if err != nil {
		return nil, err
	}
	metrics.IPStatistics = wanIPStats

	ftthMetrics, err := client.getWanFtthStatistics()
	if err != nil {
		return nil, err
	}
	metrics.FtthStatistics = ftthMetrics

	return &metrics, nil
}

// getWanInformations returns WAN IP Information
// See: https://api.bbox.fr/doc/apirouter/#api-WAN-GetWANIP
func (client *Client) getWanInformations() ([]WanIPInformations, error) {
	log.Info("Retrieve WAN IP informations from Bbox")
	var informations []WanIPInformations
	if err := client.apiRequest("%s/wan/ip", &informations); err != nil {
		return nil, err
	}
	return informations, nil
}

// getWanStatistics returns WAN IP statistics
// See: https://api.bbox.fr/doc/apirouter/#api-WAN-GetWANIPStats
func (client *Client) getWanStatistics() ([]WanIPStatistics, error) {
	log.Info("Retrieve WAN metrics from Bbox")
	// resp, err := http.Get(fmt.Sprintf("%s/wan/ip/stats", client.URL))
	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// log.Infof("Wan metrics response: %s", body)
	var metrics []WanIPStatistics
	// dec := json.NewDecoder(bytes.NewBuffer(body))
	// if err := dec.Decode(&metrics); err != nil {
	// 	return nil, err
	// }
	if err := client.apiRequest("%s/wan/ip/stats", &metrics); err != nil {
		return nil, err
	}
	return metrics, nil
}

// getWanFtthStatistics returns information about FTTH
// See: https://api.bbox.fr/doc/apirouter/#api-WAN-GetFTTHStats
func (client *Client) getWanFtthStatistics() (*FtthStatistics, error) {
	log.Info("Retrieve WAN metrics from Bbox")
	var metrics FtthStatistics
	if err := client.apiRequest("%s/wan/ftth/stats", &metrics); err != nil {
		return nil, err
	}
	return &metrics, nil
}

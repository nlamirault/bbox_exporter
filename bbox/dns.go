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

import "github.com/prometheus/common/log"

type DNSMetrics struct {
	Principal []DNSAverage `json:"principal"`
}

type DNSAverage struct {
	DNS struct {
		NumberOfQueries float64 `json:""`
		Min             float64 `json:"min"`
		Max             float64 `json:"max"`
		Average         float64 `json:"avg"`
	} `json:"dns"`
}

func (client *Client) getDNSMetrics() (*DNSMetrics, error) {
	var metrics DNSMetrics

	dns, err := client.getDNSAverage()
	if err != nil {
		return nil, err
	}
	metrics.Principal = dns

	return &metrics, nil
}

// getDNSAverage returns information about dns average.
// See: https://api.bbox.fr/doc/apirouter/#api-DNS-GetDNS
func (client *Client) getDNSAverage() ([]DNSAverage, error) {
	log.Info("Retrieve DNS informations")
	var dns []DNSAverage
	if err := client.apiRequest("/dns/stats", &dns); err != nil {
		return nil, err
	}
	return dns, nil
}

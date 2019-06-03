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
	// "bytes"
	// "encoding/json"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "io/ioutil"
	"net/http"
	"net/url"

	"github.com/prometheus/common/log"

	"github.com/nlamirault/bbox_exporter/version"
)

const (
	acceptHeader = "application/json"
	mediaType    = "application/json"

	apiVersion = "/api/v1"
)

var (
	userAgent = fmt.Sprintf("bbox-exporter/%s", version.Version)
)

// Metrics define Bbox Prometheus metrics
type Metrics struct {
	Device    DeviceMetrics   `json:"device"`
	Wan       WanMetrics      `json:"wan"`
	Lan       LanMetrics      `json:"lan"`
	DNS       DNSMetrics      `json:"dns"`
	Services  ServicesMetrics `json:"services"`
	FtthState string          `json:"ftth_state"`
}

type Client struct {
	URL string
}

func NewClient(endpoint string) (*Client, error) {
	url, err := url.Parse(endpoint)
	if err != nil || url.Scheme != "https" {
		return nil, fmt.Errorf("Invalid bbox address: %s", err)
	}
	log.Infof("bbox client creation")
	return &Client{
		URL: fmt.Sprintf("%s%s", url.String(), apiVersion),
	}, nil
}

func (client *Client) setupHeaders(request *http.Request) {
	request.Header.Add("Content-Type", mediaType)
	request.Header.Add("Accept", acceptHeader)
	request.Header.Add("User-Agent", userAgent)
}

// GetMetrics retrieve available metrics for the API Router
func (client *Client) GetMetrics() (*Metrics, error) {
	log.Infof("Get metrics")

	var metrics Metrics

	deviceMetrics, err := client.getDeviceMetrics()
	if err != nil {
		return nil, err
	}
	log.Infof("Device metrics: %#v", deviceMetrics)
	metrics.Device = *deviceMetrics

	wanMetrics, err := client.getWanMetrics()
	if err != nil {
		return nil, err
	}
	log.Infof("WAN metrics: %#v", wanMetrics)
	metrics.Wan = *wanMetrics

	lanMetrics, err := client.getLanMetrics()
	if err != nil {
		return nil, err
	}
	log.Infof("LAN metrics: %#v", lanMetrics)
	metrics.Lan = *lanMetrics

	dnsMetrics, err := client.getDNSMetrics()
	if err != nil {
		return nil, err
	}
	log.Infof("DNS metrics: %#v", dnsMetrics)
	metrics.DNS = *dnsMetrics

	servicesMetrics, err := client.getServicesMetrics()
	if err != nil {
		return nil, err
	}
	log.Infof("Services metrics: %#v", servicesMetrics)
	metrics.Services = *servicesMetrics

	return &metrics, nil
}

func (client *Client) apiRequest(request string, v interface{}) error {
	url := fmt.Sprintf(request, client.URL)
	log.Infof("Bbox API request : %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Infof("Bbox API response: %s", body)
	dec := json.NewDecoder(bytes.NewBuffer(body))
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}

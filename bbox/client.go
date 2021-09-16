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
	// "encoding/json"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/prometheus/common/log"

	"github.com/nlamirault/bbox_exporter/version"
)

const (
	acceptHeader = "application/json"
	mediaType    = "application/json"

	apiVersion = "/api/v1"
)

var (
	application = "bbox-exporter"
	userAgent   = fmt.Sprintf("%s/%s", application, version.Version)
)

// Metrics define Bbox Prometheus metrics
type Metrics struct {
	Device    DeviceMetrics   `json:"device"`
	Wan       WanMetrics      `json:"wan"`
	Lan       LanMetrics      `json:"lan"`
	DNS       DNSMetrics      `json:"dns"`
	Services  ServicesMetrics `json:"services"`
	FtthState string          `json:"ftth_state"`
	Wireless  WirelessMetrics `json:"wireless"`
	IPTV      IPTVMetrics     `json:"iptv"`
}

type Client struct {
	url      string
	cookies  []*http.Cookie
	password string
}

func NewClient(endpoint string, password string) (*Client, error) {
	url, err := url.Parse(endpoint)
	if err != nil || url.Scheme != "https" {
		return nil, fmt.Errorf("Invalid bbox address: %s", err)
	}
	log.Infof("bbox client creation")
	return &Client{
		url:      fmt.Sprintf("%s%s", url.String(), apiVersion),
		password: password,
	}, nil
}

// func (client *Client) setupHeaders(request *http.Request) {
// 	request.Header.Add("Content-Type", mediaType)
// 	request.Header.Add("X-Requested-By", application)
// 	request.Header.Add("Accept", acceptHeader)
// 	request.Header.Add("User-Agent", userAgent)
// }

// GetMetrics retrieve available metrics for the API Router
func (client *Client) GetMetrics() (*Metrics, error) {
	log.Infof("Get metrics")

	var metrics Metrics

	deviceMetrics, err := client.getDeviceMetrics()
	if err != nil {
		return nil, fmt.Errorf("Device metrics : %s", err)
	}
	log.Infof("Device metrics: %#v", deviceMetrics)
	metrics.Device = *deviceMetrics

	servicesMetrics, err := client.getServicesMetrics()
	if err != nil {
		return nil, fmt.Errorf("Services metrics: %s", err)
	}
	log.Infof("Services metrics: %#v", servicesMetrics)
	metrics.Services = *servicesMetrics

	wanMetrics, err := client.getWanMetrics()
	if err != nil {
		return nil, fmt.Errorf("WAN metrics: %s", err)
	}
	log.Infof("WAN metrics: %#v", wanMetrics)
	metrics.Wan = *wanMetrics

	lanMetrics, err := client.getLanMetrics()
	if err != nil {
		return nil, fmt.Errorf("LAN metrics: %s", err)
	}
	log.Infof("LAN metrics: %#v", lanMetrics)
	metrics.Lan = *lanMetrics

	wirelessMetrics, err := client.getWirelessMetrics()
	if err != nil {
		return nil, fmt.Errorf("Wireless metrics: %s", err)
	}
	log.Infof("WIFI metrics: %#v", wirelessMetrics)
	metrics.Wireless = *wirelessMetrics

	dnsMetrics, err := client.getDNSMetrics()
	if err != nil {
		return nil, fmt.Errorf("DNS metrics: %s", err)
	}
	log.Infof("DNS metrics: %#v", dnsMetrics)
	metrics.DNS = *dnsMetrics

	iptv, err := client.getIPTVMetrics()
	if err != nil {
		return nil, fmt.Errorf("IPTV metrics: %s", err)
	}
	log.Infof("IPTV metrics : %#v", iptv)
	metrics.IPTV = *iptv

	return &metrics, nil
}

func (client *Client) Authenticate() error {
	log.Infof("Bbox API perform authentication")
	resp, err := http.Post(
		fmt.Sprintf("%s/login", client.url),
		"application/x-www-form-urlencoded",
		bytes.NewBuffer([]byte(fmt.Sprintf("password=%s", client.password))))
	if err != nil {
		return err
	}
	log.Infof("Login response: %v", resp)
	cookies := resp.Cookies()
	if len(resp.Cookies()) == 0 {
		return fmt.Errorf("Can't retreive Cookie from API response")
	}
	// log.Infof("Cookies : ================== %s", cookies)
	client.cookies = cookies
	return nil
}

func (client *Client) apiRequest(request string, v interface{}) error {
	url := fmt.Sprintf("%s%s", client.url, request)
	log.Infof("Bbox API request : %s", url)

	req, err := http.NewRequest("GET", url, nil)
	// resp, err := http.Get(url)
	if err != nil {
		return err
	}

	req.Header.Set("Cache-Control", "no-cache")
	if client.cookies != nil {
		for _, cookie := range client.cookies {
			req.AddCookie(cookie)
		}
	}

	httpClient := &http.Client{Timeout: time.Second * 10}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil
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

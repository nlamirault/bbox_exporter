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

	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/log"
	// "github.com/nlamirault/bbox_exporter/version"
)

const (
	// acceptHeader = "application/json"
	// mediaType    = "application/json"

	apiVersion = "/api/v1"
)

// var (
// 	application = "bbox-exporter"
// 	userAgent   = fmt.Sprintf("prom/%s", application)
// )

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
	logger   log.Logger
}

func NewClient(endpoint string, password string, logger log.Logger) (*Client, error) {
	url, err := url.Parse(endpoint)
	if err != nil || url.Scheme != "https" {
		return nil, fmt.Errorf("invalid bbox address: %s", err)
	}
	level.Info(logger).Log("msg", "Create client", "endpoint", endpoint)
	return &Client{
		url:      fmt.Sprintf("%s%s", url.String(), apiVersion),
		password: password,
		logger:   logger,
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
	level.Info(client.logger).Log("msg", "Get metrics")

	var metrics Metrics

	deviceMetrics, err := client.getDeviceMetrics()
	if err != nil {
		return nil, fmt.Errorf("device metrics : %s", err)
	}
	level.Info(client.logger).Log("msg", "Device metrics", "metrics", deviceMetrics)
	metrics.Device = *deviceMetrics

	servicesMetrics, err := client.getServicesMetrics()
	if err != nil {
		return nil, fmt.Errorf("services metrics: %s", err)
	}
	level.Info(client.logger).Log("msg", "Services metrics", "metrics", servicesMetrics)
	metrics.Services = *servicesMetrics

	wanMetrics, err := client.getWanMetrics()
	if err != nil {
		return nil, fmt.Errorf("WAN metrics: %s", err)
	}
	level.Info(client.logger).Log("msg", "WAN metrics", "metrics", wanMetrics)
	metrics.Wan = *wanMetrics

	lanMetrics, err := client.getLanMetrics()
	if err != nil {
		return nil, fmt.Errorf("LAN metrics: %s", err)
	}
	level.Info(client.logger).Log("msg", "LAN metrics: %#v", lanMetrics)
	metrics.Lan = *lanMetrics

	wirelessMetrics, err := client.getWirelessMetrics()
	if err != nil {
		return nil, fmt.Errorf("wireless metrics %s", err)
	}
	level.Info(client.logger).Log("msg", "WIFI metrics", "metrics", wirelessMetrics)
	metrics.Wireless = *wirelessMetrics

	dnsMetrics, err := client.getDNSMetrics()
	if err != nil {
		return nil, fmt.Errorf("dns metrics %s", err)
	}
	level.Info(client.logger).Log("msg", "DNS metrics", "metrics", dnsMetrics)
	metrics.DNS = *dnsMetrics

	iptv, err := client.getIPTVMetrics()
	if err != nil {
		return nil, fmt.Errorf("iptv metrics %s", err)
	}
	level.Info(client.logger).Log("msg", "IPTV metrics", "metrics", fmt.Sprintf("%s", iptv))
	metrics.IPTV = *iptv

	return &metrics, nil
}

func (client *Client) Authenticate() error {
	request := fmt.Sprintf("%s/login", client.url)
	level.Info(client.logger).Log("msg", "API request", "api", request)
	resp, err := http.Post(
		request,
		"application/x-www-form-urlencoded",
		bytes.NewBuffer([]byte(fmt.Sprintf("password=%s", client.password))))
	if err != nil {
		return err
	}
	level.Info(client.logger).Log("msg", "API response", "api", request, "code", resp.StatusCode)
	if resp.StatusCode > 300 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("authentication failed: %s", err)
		}
		dec := json.NewDecoder(bytes.NewBuffer(body))
		var apiError APIError
		if err := dec.Decode(&apiError); err != nil {
			return fmt.Errorf("authentication failed: %s", err)
		}
		return fmt.Errorf("authentication failed: %+v", apiError)
	}
	cookies := resp.Cookies()
	if len(resp.Cookies()) == 0 {
		return fmt.Errorf("can't retreive Cookie from API response")
	}
	client.cookies = cookies
	return nil
}

func (client *Client) apiRequest(request string, v interface{}) error {
	url := fmt.Sprintf("%s%s", client.url, request)
	level.Debug(client.logger).Log("msg", "API request", "request", url)

	req, err := http.NewRequest("GET", url, nil)
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
	level.Debug(client.logger).Log("msg", "API response check", "request", url, "code", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	level.Debug(client.logger).Log("msg", "API response value", "request", url, "content", string(body))
	dec := json.NewDecoder(bytes.NewBuffer(body))
	if err := dec.Decode(v); err != nil {
		return err
	}
	level.Info(client.logger).Log("msg", "API entity", "api", fmt.Sprintf("%+v", v))
	return nil
}

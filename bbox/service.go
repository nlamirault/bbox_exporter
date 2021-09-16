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

import "github.com/prometheus/common/log"

type ServicesMetrics struct {
	Informations []ServicesInformations
}

type ServicesInformations struct {
	Services struct {
		Now      string `json:"now"`
		Firewall struct {
			Status  int `json:"status"`
			Enable  int `json:"enable"`
			Nbrules int `json:"nbrules"`
		} `json:"firewall"`
		Dyndns struct {
			State   int `json:"state"`
			Enable  int `json:"enable"`
			Nbrules int `json:"nbrules"`
		} `json:"dyndns"`
		Dhcp struct {
			Status  int `json:"status"`
			Enable  int `json:"enable"`
			Nbrules int `json:"nbrules"`
		} `json:"dhcp"`
		Nat struct {
			Status  int `json:"status"`
			Enable  int `json:"enable"`
			Nbrules int `json:"nbrules"`
		} `json:"nat"`
		Gamermode struct {
			Status int `json:"status"`
			Enable int `json:"enable"`
		} `json:"gamermode"`
		Upnp struct {
			Igd struct {
				Status  int `json:"status"`
				Enable  int `json:"enable"`
				Nbrules int `json:"nbrules"`
			} `json:"igd"`
		} `json:"upnp"`
		Remote struct {
			Proxywol struct {
				Status int    `json:"status"`
				Enable int    `json:"enable"`
				IP     string `json:"ip"`
			} `json:"proxywol"`
			Admin struct {
				Status     int    `json:"status"`
				Enable     int    `json:"enable"`
				Port       int    `json:"port"`
				IP         string `json:"ip"`
				Duration   string `json:"duration"`
				Activable  int    `json:"activable"`
				IP6Address string `json:"ip6address"`
			} `json:"admin"`
		} `json:"remote"`
		Parentalcontrol struct {
			Enable int `json:"enable"`
		} `json:"parentalcontrol"`
		Wifischeduler struct {
			Enable int `json:"enable"`
		} `json:"wifischeduler"`
		Voipscheduler struct {
			Enable int `json:"enable"`
		} `json:"voipscheduler"`
		Notification struct {
			Enable int `json:"enable"`
		} `json:"notification"`
		Hotspot struct {
			Status int `json:"status"`
			Enable int `json:"enable"`
		} `json:"hotspot"`
		Usb struct {
			Samba struct {
				Status int `json:"status"`
				Enable int `json:"enable"`
			} `json:"samba"`
			Printer struct {
				Status int `json:"status"`
				Enable int `json:"enable"`
			} `json:"printer"`
			Dlna struct {
				Status int `json:"status"`
				Enable int `json:"enable"`
			} `json:"dlna"`
		} `json:"usb"`
	} `json:"services"`
}

func (client *Client) getServicesMetrics() (*ServicesMetrics, error) {
	var metrics ServicesMetrics

	informations, err := client.getServicesInformations()
	if err != nil {
		return nil, err
	}
	metrics.Informations = informations

	return &metrics, nil
}

// getServicesInformations returns Services information
// See: https://api.bbox.fr/doc/apirouter/#api-Services-GetServices
func (client *Client) getServicesInformations() ([]ServicesInformations, error) {
	log.Info("Retrieve Services informations from Bbox")
	var informations []ServicesInformations
	if err := client.apiRequest("/services", &informations); err != nil {
		return nil, err
	}
	return informations, nil
}

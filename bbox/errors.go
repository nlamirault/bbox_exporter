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

type APIError struct {
	Exception struct {
		Domain string `json:"domain"`
		Code   string `json:"code"`
		Errors []struct {
			Name   string `json:"name"`
			Reason string `json:"reason"`
		} `json:"errors"`
	} `json:"exception"`
}

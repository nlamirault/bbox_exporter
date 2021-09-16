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
	"encoding/json"
	"strconv"
)

// WTF The bbox API send result in string and/or int :(

// A FlexInt is an int that can be unmarshalled from a JSON field
// that has either a number or a string value.
// E.g. if the json field contains an string "42", the
// FlexInt value will be "42".
type flexInt int

// UnmarshalJSON implements the json.Unmarshaler interface, which
// allows us to ingest values of any json type as an int and run our custom conversion
func (fi *flexInt) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(fi))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*fi = flexInt(i)
	return nil
}

// Copyright 2018 Istio Authors
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

package probe

import (
	"fmt"
	"time"
)

// Options customizes the parameters of a probe.
type Options struct {
	// Path defines the path to the file used for the existence.
	Path string

	// UpdateInterval defines the interval for updating the file's last modified
	// time.
	UpdateInterval time.Duration
}

// IsValid returns true if some values are filled into the options.
func (o *Options) IsValid() bool {
	return o != nil && o.Path != "" && o.UpdateInterval > 0*time.Second
}

// Validate returns true if some values are filled into the options.
func (o *Options) Validate() error {
	if o == nil {
		return fmt.Errorf("option is empty")
	}
	if o.Path == "" {
		return fmt.Errorf("probe-path must be specified")
	}
	if o.UpdateInterval <= 0*time.Second {
		return fmt.Errorf("interval must be greater than zero")
	}
	return nil
}

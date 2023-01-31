// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package config

import "strings"

// Config represents instructions to grm-generate on how to inspect a cloud
// service API model, how to generate the consistent grm model for that API and
// how to generate a resource manager that can manage resources in that API.
type Config struct {
	// Cloud specifies the cloud service that publishes this resource.
	Cloud string `json:"cloud"`
	// Resources contains generator instructions for individual CRDs within an
	// API
	Resources map[string]*ResourceConfig `json:"resources"`
}

// GetResourceConfigs returns the map, keyed by resource name, of
// ResourceConfigs, or an empty map if the config is nil
func (c *Config) GetResourceConfigs() map[string]*ResourceConfig {
	if c == nil {
		return map[string]*ResourceConfig{}
	}
	return c.Resources
}

// GetResourceConfig returns a ResourceConfig matching the supplied resource
// name, using case-insensitive matching.
func (c *Config) GetResourceConfig(search string) *ResourceConfig {
	if c == nil {
		return nil
	}
	for name, rc := range c.Resources {
		if strings.EqualFold(name, search) {
			return rc
		}
	}
	return nil
}

package plugin

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/photoprism/photoprism/pkg/list"
)

const (
	KeyEnable = "enabled"
)

// PluginConfig is a type alias for a key-value based plugin configuration.
type PluginConfig map[string]string

// Enabled checks whether the plugin has been explicitly enabled by the used.
func (c PluginConfig) Enabled() bool {
	if value, ok := c[KeyEnable]; ok {
		return list.Bool[strings.ToLower(value)] == list.True
	}

	// All plugins are disabled per default and only activated if explicitly enabled.
	return false
}

// MandatoryStringParameter reads a mandatory string plugin parameter.
func (c PluginConfig) MandatoryStringParameter(name string) (string, error) {
	if value, ok := c[name]; !ok {
		return "", fmt.Errorf("%s parameter is mandatory", name)
	} else {
		return value, nil
	}
}

// OptionalFloatParameter reads an optional float64 plugin parameter.
func (c PluginConfig) OptionalFloatParameter(name string, defaultValue float64) (float64, error) {
	if value, ok := c[name]; ok {
		if fValue, err := strconv.ParseFloat(value, 64); err != nil {
			return 0, err
		} else {
			return fValue, nil
		}
	} else {
		return defaultValue, nil
	}
}

func loadConfig(p Plugin) PluginConfig {
	var config = make(PluginConfig)

	prefix := fmt.Sprintf("PHOTOPRISM_PLUGIN_%s_", strings.ToUpper(p.Name()))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		key := pair[0]
		value := pair[1]

		if strings.HasPrefix(key, prefix) {
			key = strings.ToLower(key[len(prefix):])

			config[key] = value
		}
	}

	return config
}

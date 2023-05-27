package plugin

import (
	"fmt"
	"strings"
	"os"

	"github.com/photoprism/photoprism/pkg/list"
)

const (
	KeyActive = "active"
)

// PluginConfig is a type alias for a key-value based plugin configuration.
type PluginConfig map[string]string

// Active checks whether the plugin has been explicitly activated by the used.
func (c PluginConfig) Active() bool {
	if value, ok := c[KeyActive]; ok {
		return list.Bool[strings.ToLower(value)] == list.True
	}

	// All plugins are disabled per default and only activated if explicitly enabled.
	return false
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

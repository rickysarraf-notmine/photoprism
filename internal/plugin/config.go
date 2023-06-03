package plugin

import (
	"fmt"
	"os"
	"strings"

	"github.com/creasty/defaults"
	"github.com/mitchellh/mapstructure"
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

// // Decode populates a struct with the configuration data.
func (c PluginConfig) Decode(init any) error {
	if err := defaults.Set(init); err != nil {
		return err
	}

	if err := mapstructure.WeakDecode(c, &init); err != nil {
		return err
	}

	fmt.Printf("CFG OK %#v", init)

	return nil
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

package plugin

import (
	"strings"

	"github.com/photoprism/photoprism/internal/entity"
)

// Plugins is a type alias for a collection of plugins.
type Plugins []Plugin

// Plugin represents a plugin interface to hook into PhotoPrism's internals.
// The interface requires specifying a plugin name, a way to configure the plugin using
// environment variables and various calback hooks.
type Plugin interface {
	Name() string
	Configure(PluginConfig) error
	OnIndex(*entity.File, *entity.Photo) error
}

// OnIndex calls the [OnIndex] hook method for all active plugins.
func OnIndex(file *entity.File, photo *entity.Photo) (changed bool) {
	for _, p := range getPlugins() {
		if err := p.OnIndex(file, photo); err != nil {
			log.Errorf("plugin %s: %s (importing file %s)", strings.ToLower(p.Name()), err, file.FileUID)
		} else {
			changed = true
		}
	}

	return changed
}

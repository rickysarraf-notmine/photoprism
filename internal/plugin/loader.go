package plugin

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
	"strings"
	"sync"
	"sync/atomic"
)

var (
	pluginsInit  sync.Once
	pluginsMutex sync.Mutex
	plugins      atomic.Value
)

func getPlugins() Plugins {
	pluginsMutex.Lock()
	p, _ := plugins.Load().(Plugins)
	pluginsMutex.Unlock()
	return p
}

// LoadPlugins loads plugins from the given directory, looking for all solib files in there.
// Only plugins, which have been activated by configuration parameter will be loaded.
func LoadPlugins(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Infof("plugins: path does not exist %s", path)
		return
	}

	pluginFolder, err := ioutil.ReadDir(path)
	if err != nil {
		log.Errorf("plugins: %s (reading plugin path %s)", err, path)
		return
	}

	for _, entry := range pluginFolder {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".so") {
			return
		}

		pluginPath := filepath.Join(path, entry.Name())
		log.Debugf("plugins: found plugin file %s", pluginPath)

		plug, err := plugin.Open(pluginPath)
		if err != nil {
			log.Errorf("plugins: %s (when opening plugin solib %s)", err, pluginPath)
			continue
		}

		symPlugin, err := plug.Lookup("Plugin")
		if err != nil {
			log.Errorf("plugins: %s (when looking up plugin symbol in solib %s)", err, pluginPath)
			continue
		}

		var plugin Plugin
		plugin, ok := symPlugin.(Plugin)
		if !ok {
			log.Errorf("plugins: %s (unexpected type for plugin symbol in solib %s", err, pluginPath)
			continue
		}

		pluginName := strings.ToLower(plugin.Name())

		config := loadConfig(plugin)
		if !config.Enabled() {
			log.Warnf("plugin %s: plugin was loaded, but is not enabled", pluginName)
			continue
		}

		if err := plugin.Configure(config); err != nil {
			log.Errorf("plugin %s: %s (plugin configuration)", pluginName, err)
			continue
		}

		pluginsMutex.Lock()
		h, _ := plugins.Load().(Plugins)
		plugins.Store(append(h, plugin))
		pluginsMutex.Unlock()

		log.Infof("plugin %s: successfully loaded plugin", pluginName)
	}
}

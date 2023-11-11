package main

import (
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/plugin"
)

type DemoPlugin struct{}

func (p DemoPlugin) Name() string {
	return "demo"
}

func (*DemoPlugin) Configure(plugin.PluginConfig) error {
	return nil
}

func (*DemoPlugin) OnIndex(file *entity.File, photo *entity.Photo) error {
	photo.Details.Notes = "hello from demo plugin"
	photo.Details.NotesSrc = entity.SrcManual

	return nil
}

// Export the plugin.
var Plugin DemoPlugin

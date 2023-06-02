package main

import (
	"encoding/base64"
	"os"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/get"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/plugin"
	"github.com/photoprism/photoprism/pkg/fs"
)

type UpscaleResult struct {
	Image string `json:"image"`
}

type RealESRGANPlugin struct {
	hostname string
	port     string
	scale    float64

	thresholdResolution float64
	thresholdQuality    float64
}

func (p RealESRGANPlugin) Name() string {
	return "realesrgan"
}

func (p RealESRGANPlugin) Hostname() string {
	return p.hostname
}

func (p RealESRGANPlugin) Port() string {
	return p.port
}

func (p *RealESRGANPlugin) Configure(config plugin.PluginConfig) error {
	hostname, err := config.MandatoryStringParameter("hostname")
	if err != nil {
		return err
	}

	port, err := config.MandatoryStringParameter("port")
	if err != nil {
		return err
	}

	scale, err := config.OptionalFloatParameter("scale", 2)
	if err != nil {
		return err
	}

	resolution, err := config.OptionalFloatParameter("threshold_resolution", 3)
	if err != nil {
		return err
	}

	quality, err := config.OptionalFloatParameter("threshold_quality", 3)
	if err != nil {
		return err
	}

	p.hostname = hostname
	p.port = port
	p.scale = scale
	p.thresholdResolution = resolution
	p.thresholdQuality = quality

	return nil
}

func (p *RealESRGANPlugin) OnIndex(file *entity.File, photo *entity.Photo) error {
	if !p.needsUpscaling(photo) {
		return nil
	}

	image, err := p.image(file)
	if err != nil {
		return err
	}

	output, err := p.superscale(image)
	if err != nil {
		return err
	}

	if err := p.save(file, output); err != nil {
		return err
	}

	return nil
}

func (p *RealESRGANPlugin) needsUpscaling(photo *entity.Photo) bool {
	return p.thresholdResolution > float64(photo.PhotoResolution) || p.thresholdQuality > float64(photo.PhotoQuality)
}

func (p *RealESRGANPlugin) image(f *entity.File) (string, error) {
	filePath := photoprism.FileName(f.FileRoot, f.FileName)

	return plugin.ReadImageAsBase64(filePath)
}

func (p *RealESRGANPlugin) superscale(image string) ([]byte, error) {
	payload := map[string]interface{}{"image": image, "scale": p.scale}

	if output, err := plugin.PostJson[UpscaleResult](p, "superscale", payload); err != nil {
		return nil, err
	} else {
		if decoded, err := base64.StdEncoding.DecodeString(output.Image); err != nil {
			return nil, err
		} else {
			return decoded, nil
		}
	}
}

func (p *RealESRGANPlugin) save(f *entity.File, data []byte) error {
	conf := get.Config()

	ext := ".SUPERSCALED" + fs.ExtJPEG
	baseDir := conf.OriginalsPath()
	// if f.InSidecar() {
	// 	baseDir = conf.SidecarPath()
	// }

	imageName := fs.FileName(photoprism.FileName(f.FileRoot, f.FileName), conf.SidecarPath(), baseDir, ext)

	if err := os.WriteFile(imageName, data, 0666); err != nil {
		return err
	}

	return nil
}

// Export the plugin.
var Plugin RealESRGANPlugin

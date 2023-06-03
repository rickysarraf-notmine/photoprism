package main

import (
	"fmt"

	"github.com/photoprism/photoprism/internal/classify"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/plugin"
)

type ClassifyResults map[string]float64
type DetectResults []DetectResult

type DetectResult struct {
	Name       string  `json:"name"`
	Class      int16   `json:"class"`
	Confidence float64 `json:"confidence"`
}

type Config struct {
	Hostname            string
	Port                string
	ThresholdConfidence float64 `default:"0.4"`
}

type Yolo8Plugin struct {
	Config *Config
}

func (p Yolo8Plugin) Name() string {
	return "yolo8"
}

func (p Yolo8Plugin) Hostname() string {
	return p.Config.Hostname
}

func (p Yolo8Plugin) Port() string {
	return p.Config.Port
}

func (p *Yolo8Plugin) Configure(config plugin.PluginConfig) error {
	p.Config = &Config{}

	if err := config.Decode(p.Config); err != nil {
		return err
	}

	return nil
}

func (p *Yolo8Plugin) OnIndex(file *entity.File, photo *entity.Photo) error {
	image, err := p.image(file)
	if err != nil {
		return err
	}

	labels, err := p.classify(image)
	if err != nil {
		return err
	}

	if len(labels) > 0 {
		fmt.Printf("adding labels %#v", labels)
	}

	photo.AddLabels(labels)

	return nil
}

func (p *Yolo8Plugin) image(f *entity.File) (string, error) {
	filePath := photoprism.FileName(f.FileRoot, f.FileName)

	return plugin.ReadImageAsBase64(filePath)
}

func (p *Yolo8Plugin) detect(image string) (classify.Labels, error) {
	payload := map[string]interface{}{"image": image}

	if output, err := plugin.PostJson[DetectResults](p, "detect", payload); err != nil {
		return nil, err
	} else {
		return output.toLabels(p.Config.ThresholdConfidence), nil
	}
}

func (p *Yolo8Plugin) classify(image string) (classify.Labels, error) {
	payload := map[string]interface{}{"image": image}

	if output, err := plugin.PostJson[ClassifyResults](p, "classify", payload); err != nil {
		return nil, err
	} else {
		return output.toLabels(p.Config.ThresholdConfidence), nil
	}
}

func (results DetectResults) toLabels(threshold float64) classify.Labels {
	labels := make(classify.Labels, 0)

	for _, result := range results {
		if result.Confidence > threshold {
			labels = append(labels, classify.Label{
				Name: result.Name,
				// It would be nice to be able to denote that the label comes from the yolo8 plugin,
				// but PhotoPrism does not really like it when custom sources are used.
				// Source: p.Name(),
				Source:      classify.SrcImage,
				Uncertainty: int((1 - result.Confidence) * 100),
				Priority:    0,
			})
		}
	}

	return labels
}

func (results ClassifyResults) toLabels(threshold float64) classify.Labels {
	labels := make(classify.Labels, 0)

	for label, confidence := range results {
		if confidence > threshold {
			labels = append(labels, classify.Label{
				Name: label,
				// It would be nice to be able to denote that the label comes from the yolo8 plugin,
				// but PhotoPrism does not really like it when custom sources are used.
				// Source: p.Name(),
				Source:      classify.SrcImage,
				Uncertainty: int((1 - confidence) * 100),
				Priority:    0,
			})
		}
	}

	return labels
}

// Export the plugin.
var Plugin Yolo8Plugin

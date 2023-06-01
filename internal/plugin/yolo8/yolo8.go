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

type Yolo8Plugin struct {
	hostname      string
	port          string
	confThreshold float64
}

func (p Yolo8Plugin) Name() string {
	return "yolo8"
}

func (p Yolo8Plugin) Hostname() string {
	return p.hostname
}

func (p Yolo8Plugin) Port() string {
	return p.port
}

func (p *Yolo8Plugin) Configure(config plugin.PluginConfig) error {
	hostname, err := config.MandatoryStringParameter("hostname")
	if err != nil {
		return err
	}

	port, err := config.MandatoryStringParameter("port")
	if err != nil {
		return err
	}

	threshold, err := config.OptionalFloatParameter("confidence_threshold", 0.5)
	if err != nil {
		return err
	}

	p.hostname = hostname
	p.port = port
	p.confThreshold = threshold

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

	photo.AddLabels(labels)
	fmt.Printf("adding labels %#v", labels)

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
		return output.toLabels(p.confThreshold), nil
	}
}

func (p *Yolo8Plugin) classify(image string) (classify.Labels, error) {
	payload := map[string]interface{}{"image": image}

	if output, err := plugin.PostJson[ClassifyResults](p, "classify", payload); err != nil {
		return nil, err
	} else {
		return output.toLabels(p.confThreshold), nil
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

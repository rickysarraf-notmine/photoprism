package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/photoprism/photoprism/internal/classify"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/plugin"
	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
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

func (p *Yolo8Plugin) Configure(config plugin.PluginConfig) error {
	hostname, ok := config["hostname"]
	if !ok {
		return fmt.Errorf("hostname parameter is mandatory")
	}

	port, ok := config["port"]
	if !ok {
		return fmt.Errorf("port parameter is mandatory")
	}

	threshold := 0.5
	var err error

	if t, ok := config["confidence_threshold"]; ok {
		if threshold, err = strconv.ParseFloat(t, 64); err != nil {
			return err
		}
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

	if !fs.FileExists(filePath) {
		return "", fmt.Errorf("file %s is missing", clean.Log(f.FileName))
	}

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	return encoded, nil
}

func (p *Yolo8Plugin) detect(image string) (classify.Labels, error) {
	client := &http.Client{Timeout: 60 * time.Second}
	url := fmt.Sprintf("http://%s:%s/detect", p.hostname, p.port)
	payload := map[string]string{"image": image}

	var req *http.Request
	var output *DetectResults

	if j, err := json.Marshal(payload); err != nil {
		return nil, err
	} else if req, err = http.NewRequest(http.MethodPost, url, bytes.NewReader(j)); err != nil {
		return nil, err
	}

	// Add Content-Type header.
	req.Header.Add("Content-Type", "application/json")

	if resp, err := client.Do(req); err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("yolo8 server running at %s:%s, bad status %d\n", p.hostname, p.port, resp.StatusCode)
	} else if body, err := io.ReadAll(resp.Body); err != nil {
		return nil, err
	} else if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	} else {
		return (*output).toLabels(p.confThreshold), nil
	}
}

func (p *Yolo8Plugin) classify(image string) (classify.Labels, error) {
	client := &http.Client{Timeout: 60 * time.Second}
	url := fmt.Sprintf("http://%s:%s/classify", p.hostname, p.port)
	payload := map[string]string{"image": image}

	var req *http.Request
	var output *ClassifyResults

	if j, err := json.Marshal(payload); err != nil {
		return nil, err
	} else if req, err = http.NewRequest(http.MethodPost, url, bytes.NewReader(j)); err != nil {
		return nil, err
	}

	// Add Content-Type header.
	req.Header.Add("Content-Type", "application/json")

	if resp, err := client.Do(req); err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("yolo8 server running at %s:%s, bad status %d\n", p.hostname, p.port, resp.StatusCode)
	} else if body, err := io.ReadAll(resp.Body); err != nil {
		return nil, err
	} else if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	} else {
		return (*output).toLabels(p.confThreshold), nil
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

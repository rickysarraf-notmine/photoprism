package plugin

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
	"github.com/photoprism/photoprism/pkg/fs"
)

// ReadImageAsBase64 reads an image, rotates it if needed and returns it as base64-encoded string.
func ReadImageAsBase64(filePath string) (string, error) {
	if !fs.FileExists(filePath) {
		return "", fmt.Errorf("file %s is missing", filepath.Base(filePath))
	}

	img, err := imaging.Open(filePath, imaging.AutoOrientation(true))
	if err != nil {
		return "", err
	}

	buffer := &bytes.Buffer{}
	err = imaging.Encode(buffer, img, imaging.JPEG)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return encoded, nil
}

// PostJson sends a post request with a json payload to a plugin endpoint and returns a deserialized json output.
func PostJson[T any](p HttpPlugin, endpoint string, payload map[string]interface{}) (T, error) {
	client := &http.Client{Timeout: 60 * time.Second}
	url := fmt.Sprintf("http://%s:%s/%s", p.Hostname(), p.Port(), endpoint)

	var empty T

	var req *http.Request
	var output *T

	if j, err := json.Marshal(payload); err != nil {
		return empty, err
	} else if req, err = http.NewRequest(http.MethodPost, url, bytes.NewReader(j)); err != nil {
		return empty, err
	}

	// Add Content-Type header.
	req.Header.Add("Content-Type", "application/json")

	if resp, err := client.Do(req); err != nil {
		return empty, err
	} else if resp.StatusCode != 200 {
		return empty, fmt.Errorf("%s server running at %s:%s, bad status %d", p.Name(), p.Hostname(), p.Port(), resp.StatusCode)
	} else if body, err := io.ReadAll(resp.Body); err != nil {
		return empty, err
	} else if err := json.Unmarshal(body, &output); err != nil {
		return empty, err
	} else {
		return *output, nil
	}
}

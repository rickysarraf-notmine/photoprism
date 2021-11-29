package overpass

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/pkg/s2"
)

const OverpassStateQuery = `
is_in(%f,%f);
area._[admin_level="4"];
out meta;
`

const OverpassUrl = "https://overpass-api.de/api/interpreter"

var log = event.Log

// OverpassResponse represents the response from the Overpass API.
type OverpassResponse struct {
	Version   float32           `json:"version"`
	Generator string            `json:"generator"`
	Elements  []OverpassElement `json:"elements"`
}

// OverpassElement represents a generic Overpass element.
type OverpassElement struct {
	ID   int               `json:"id"`
	Type string            `json:"type"`
	Tags map[string]string `json:"tags"`
}

// Name returns the native name of the Overpass element.
func (e OverpassElement) Name() string {
	return e.Tags["name"]
}

// InternationalName returns the international name of the Overpass element.
func (e OverpassElement) InternationalName() string {
	return e.Tags["int_name"]
}

// LocalizedNames returns a mapping of the available localized names (language ISO code -> name).
func (e OverpassElement) LocalizedNames() map[string]string {
	names := make(map[string]string)

	for name, value := range e.Tags {
		if strings.HasPrefix(name, "name:") {
			country_code := name[len("name:"):]
			names[country_code] = value
		}
	}

	return names
}

// AdministrativeLevel returns the administrative level of the Overpass element if available.
func (e OverpassElement) AdministrativeLevel() string {
	return e.Tags["admin_level"]
}

// FindState queries the Overpass API to retrieve the state name in the native language for the given s2 cell.
func FindState(token string) (state string) {
	log.Debugf("overpass: quering state for %s s2 cell", token)

	lat, lng := s2.LatLng(token)
	query := fmt.Sprintf(OverpassStateQuery, lat, lng)

	r, err := queryOverpass(query)
	if err != nil {
		log.Errorf("overpass: %s", err)
		return state
	}

	if len(r.Elements) == 0 {
		log.Warnf("overpass: token %s does not have state data", token)
		return state
	}

	// TODO Should we return the "native" name or the international?
	state = r.Elements[0].Name()
	log.Debugf("overpass: found %s state for %s s2 cell", state, token)

	return state
}

// queryOverpass sends the given query to the Overpass API and unmarshals the response.
func queryOverpass(query string) (result OverpassResponse, err error) {
	reader := strings.NewReader(fmt.Sprintf("data=[out:json];%s", query))
	req, err := http.NewRequest(http.MethodPost, OverpassUrl, reader)

	if err != nil {
		return result, err
	}

	client := &http.Client{Timeout: 60 * time.Second}
	r, err := client.Do(req)

	if err != nil {
		return result, fmt.Errorf("%s (http request)", err)
	} else if r.StatusCode >= 400 {
		return result, fmt.Errorf("request failed with code %d", r.StatusCode)
	}

	err = json.NewDecoder(r.Body).Decode(&result)

	if err != nil {
		return result, fmt.Errorf("%s (decode json)", err)
	}

	return result, nil
}

package overpass

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/sethgrid/pester"

	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/pkg/s2"
)

// Different countries have different definitions for state, but generally the state is
// represented by an admin_level [3,6] area. Furthermore we are only interested in the
// "administrative" boundaries, which are the ones that denote the state.
const OverpassQueryState = `
is_in(%f,%f) -> .a;
(
	area.a[admin_level="3"][boundary="administrative"]["border_type"!="territorial"];
	area.a[admin_level="4"][boundary="administrative"];
	area.a[admin_level="5"][boundary="administrative"];
	area.a[admin_level="6"][boundary="administrative"];
);
out;
`

// Overpass query to retrieve all administrative boundaries within a radius of the given location.
// The query will return all possible boundaries - country, state, city, neighborhood.
const OverpassQueryNearbyBoundaries = `
rel["boundary"="administrative"](around:%f,%f,%f);
out tags;
`

const OverpassUrl = "https://overpass-api.de/api/interpreter"

const (
	OverpassTagAdministrativeLevel = "admin_level"
	OverpassTagBoundary            = "boundary"
	OverpassTagInternationalName   = "int_name"
	OverpassTagLocalizedNamePrefix = OverpassTagName + ":"
	OverpassTagName                = "name"
	OverpassTagPlace               = "place"
)

const (
	OverpassPlacesCity = "city"
)

var log = event.Log

// OverpassResponse represents the response from the Overpass API.
type OverpassResponse struct {
	Version   float32          `json:"version"`
	Generator string           `json:"generator"`
	Elements  OverpassElements `json:"elements"`
}

type OverpassElements []OverpassElement

// OverpassElement represents a generic Overpass element.
type OverpassElement struct {
	ID   int               `json:"id"`
	Type string            `json:"type"`
	Tags map[string]string `json:"tags"`
}

// Name returns the native name of the Overpass element.
func (e OverpassElement) Name() string {
	return e.Tags[OverpassTagName]
}

// InternationalName returns the international name of the Overpass element.
func (e OverpassElement) InternationalName() string {
	return e.Tags[OverpassTagInternationalName]
}

// LocalizedNames returns a mapping of the available localized names (language ISO code -> name).
func (e OverpassElement) LocalizedNames() map[string]string {
	names := make(map[string]string)

	for name, value := range e.Tags {
		if strings.HasPrefix(name, OverpassTagLocalizedNamePrefix) {
			country_code := name[len(OverpassTagLocalizedNamePrefix):]
			names[country_code] = value
		}
	}

	return names
}

// AdministrativeLevel returns the administrative level of the Overpass element if available.
func (e OverpassElement) AdministrativeLevel() string {
	return e.Tags[OverpassTagAdministrativeLevel]
}

// Boundary returns the boundary type of the Overpass element if available.
func (e OverpassElement) Boundary() string {
	return e.Tags[OverpassTagBoundary]
}

// IsCity returns whether the Overpass element represents a city.
func (e OverpassElement) IsCity() bool {
	return e.Tags[OverpassTagPlace] == OverpassPlacesCity
}

// FindState attempts to detect a state in the list of Overpass elements and returns the native state name.
func (elements OverpassElements) FindState() (state string) {
	admin_level := 12

	for _, area := range elements {
		area_admin_level, err := strconv.Atoi(area.AdministrativeLevel())
		if err != nil {
			log.Warnf("overpass: area %s has an invalid admin_level %s", area.Name(), area.AdministrativeLevel())
			continue
		}

		// Return the name of the smallest possible administrative boundary,
		// which ideally should represent a "state" for the given country.
		// See: https://wiki.openstreetmap.org/wiki/Tag:boundary%3Dadministrative
		// Return the "native" name instead of the international one to be compatible with the Places API.
		if area_admin_level < admin_level {
			admin_level = area_admin_level
			state = area.Name()
		}
	}

	return state
}

// FindState queries the Overpass API to retrieve the state name in the native language for the given s2 cell.
func FindState(token string) (state string) {
	lat, lng := s2.LatLng(token)
	query := fmt.Sprintf(OverpassQueryState, lat, lng)

	r, err := queryOverpass(query)
	if err != nil {
		log.Errorf("overpass: %s (query state)", err)
		return state
	}

	if len(r.Elements) == 0 {
		log.Warnf("overpass: no return data for token %s (query state)", token)
		return state
	}

	state = r.Elements.FindState()
	log.Debugf("overpass: found %s state for %s s2 cell (query state)", state, token)

	return state
}

// FindNearbyLocation looks for a city and state boundary within 200 meters of the given s2 cell.
func FindNearbyLocation(token string) (city, state string) {
	radius := 200.0
	lat, lng := s2.LatLng(token)

	query := fmt.Sprintf(OverpassQueryNearbyBoundaries, radius, lat, lng)

	r, err := queryOverpass(query)
	if err != nil {
		log.Errorf("overpass: %s (nearby location)", err)
		return city, state
	}

	if len(r.Elements) == 0 {
		log.Warnf("overpass: no return data for token %s (nearby location)", token)
		return city, state
	}

	for _, area := range r.Elements {
		if area.IsCity() {
			city = area.Name()
			break
		}
	}

	state = r.Elements.FindState()

	return city, state
}

// queryOverpass sends the given query to the Overpass API and unmarshals the response.
func queryOverpass(query string) (result OverpassResponse, err error) {
	params := make(url.Values)
	params.Set("data", fmt.Sprintf("[out:json];%s", query))

	client := pester.New()
	client.MaxRetries = 5
	client.Backoff = pester.ExponentialBackoff

	r, err := client.PostForm(OverpassUrl, params)

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

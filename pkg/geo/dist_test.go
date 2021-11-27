package geo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDist(t *testing.T) {
	t.Run("BerlinShanghai", func(t *testing.T) {
		berlin := Position{Name: "Berlin", Lat: 52.5243700, Lng: 13.4105300}
		shanghai := Position{Name: "Shanghai", Lat: 31.2222200, Lng: 121.4580600}

		result := Km(berlin, shanghai)

		assert.Equal(t, 8396, int(result))
	})
}

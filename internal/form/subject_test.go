package form

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSubject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var m = struct {
			SubjName     string `json:"Name"`
			SubjAlias    string `json:"Alias"`
			SubjBio      string `json:"Bio"`
			SubjNotes    string `json:"Notes"`
			SubjFavorite bool   `json:"Favorite"`
			SubjHidden   bool   `json:"Hidden"`
			SubjPrivate  bool   `json:"Private"`
			SubjExcluded bool   `json:"Excluded"`
			SubjThumb    string `json:"Thumb"`
			SubjThumbSrc string `json:"ThumbSrc"`
		}{
			SubjName:     "Foo",
			SubjAlias:    "bar",
			SubjFavorite: true,
			SubjHidden:   true,
			SubjExcluded: false,
			SubjThumb:    "thumb-000",
			SubjThumbSrc: "manual",
		}

		f, err := NewSubject(m)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Foo", f.SubjName)
		assert.Equal(t, "bar", f.SubjAlias)
		assert.Equal(t, true, f.SubjFavorite)
		assert.Equal(t, true, f.SubjHidden)
		assert.Equal(t, false, f.SubjExcluded)
		assert.Equal(t, "thumb-000", f.SubjThumb)
		assert.Equal(t, "manual", f.SubjThumbSrc)
	})
}

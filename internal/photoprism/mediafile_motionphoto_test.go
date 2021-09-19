package photoprism

import (
	"testing"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestMediaFile_IsMotionPhoto(t *testing.T) {
	t.Run("samsung-and-legacy-google-motion-photo", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/samsung_motion_photo.jpg")

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, mediaFile.IsMotionPhoto())
		assert.True(t, mediaFile.MetaData().MicroVideo)
		assert.NotEmpty(t, mediaFile.MetaData().EmbeddedVideoType)
	})

	t.Run("samsung-motion-photo", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/1216505.jpg")

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, mediaFile.IsMotionPhoto())
		assert.NotEmpty(t, mediaFile.MetaData().EmbeddedVideoType)
	})

	t.Run("google-motion-photo", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/PXL_20210506_083558892.MP.jpg")

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, mediaFile.IsMotionPhoto())
		assert.Len(t, mediaFile.MetaData().Directory, 2)
	})
}

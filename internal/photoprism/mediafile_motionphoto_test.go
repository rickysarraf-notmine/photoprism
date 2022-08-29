package photoprism

import (
	"os"
	"testing"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/stretchr/testify/assert"
)

func TestMediaFile_IsMotionPhoto(t *testing.T) {
	t.Run("samsung-and-legacy-google-motion-photo", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/motion_photo.jpg")

		if err != nil {
			t.Fatal(err)
		}

		// check all motion photo related metadata fields
		assert.True(t, mediaFile.IsMotionPhoto())
		assert.False(t, mediaFile.MetaData().MotionPhoto)
		assert.Empty(t, mediaFile.MetaData().Directory)
		assert.True(t, mediaFile.MetaData().MicroVideo)
		assert.Greater(t, mediaFile.MetaData().MicroVideoOffset, 0)
		assert.NotEmpty(t, mediaFile.MetaData().EmbeddedVideoType)
	})

	t.Run("samsung-motion-photo", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/1216505.jpg")

		if err != nil {
			t.Fatal(err)
		}

		// check all motion photo related metadata fields
		assert.True(t, mediaFile.IsMotionPhoto())
		assert.False(t, mediaFile.MetaData().MotionPhoto)
		assert.Empty(t, mediaFile.MetaData().Directory)
		assert.False(t, mediaFile.MetaData().MicroVideo)
		assert.Equal(t, mediaFile.MetaData().MicroVideoOffset, 0)
		assert.NotEmpty(t, mediaFile.MetaData().EmbeddedVideoType)
	})

	t.Run("samsung-heic-motion-photo", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/samsung_mp.heif")

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, mediaFile.IsMotionPhoto())
		assert.Len(t, mediaFile.MetaData().Directory, 2)
		assert.False(t, mediaFile.MetaData().MicroVideo)
		assert.Equal(t, mediaFile.MetaData().MicroVideoOffset, 0)
		assert.Empty(t, mediaFile.MetaData().EmbeddedVideoType)
	})

	t.Run("google-motion-photo", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/PXL_20210506_083558892.MP.jpg")

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, mediaFile.IsMotionPhoto())
		assert.Len(t, mediaFile.MetaData().Directory, 2)
		assert.False(t, mediaFile.MetaData().MicroVideo)
		assert.Equal(t, mediaFile.MetaData().MicroVideoOffset, 0)
		assert.Empty(t, mediaFile.MetaData().EmbeddedVideoType)
	})
}

func TestMediaFile_EmbeddedVideoData(t *testing.T) {
	t.Run("samsung-and-legacy-google-motion-photo", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/motion_photo.jpg")

		if err != nil {
			t.Fatal(err)
		}

		data, err := mediaFile.EmbeddedVideoData()

		if err != nil {
			t.Fatal(err)
		}

		assert.NotEmpty(t, data)
	})

	t.Run("samsung-motion-photo", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/1216505.jpg")

		if err != nil {
			t.Fatal(err)
		}

		data, err := mediaFile.EmbeddedVideoData()

		if err != nil {
			t.Fatal(err)
		}

		assert.NotEmpty(t, data)
	})

	t.Run("google-motion-photo", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/PXL_20210506_083558892.MP.jpg")

		if err != nil {
			t.Fatal(err)
		}

		data, err := mediaFile.EmbeddedVideoData()

		if err != nil {
			t.Fatal(err)
		}

		assert.Nil(t, data)
	})
}

func TestMediaFile_ExtractVideoFromMotionPhoto(t *testing.T) {
	var tests = []struct {
		name string
		file string
	}{
		{"samsung-and-legacy-google-motion-photo", "motion_photo.jpg"},
		{"samsung-motion-photo", "1216505.jpg"},
		{"samsung-heic-motion-photo", "samsung_mp.heif"},
		{"google-motion-photo", "PXL_20210506_083558892.MP.jpg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := config.TestConfig()

			mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/" + tt.file)

			if err != nil {
				t.Fatal(err)
			}

			data, err := mediaFile.ExtractVideoFromMotionPhoto()

			if err != nil {
				t.Fatal(err)
			}

			assert.NotNil(t, data)
			assert.FileExists(t, data.fileName)

			err = os.Remove(data.fileName)

			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestMediaFile_ExtractVideoFromMotionPhoto_BrokenMetadata(t *testing.T) {
	t.Run("google-legacy-motion-photo-with-broken-metadata", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/motion_photo.jpg")

		if err != nil {
			t.Fatal(err)
		}

		mediaFile.MetaData()
		mediaFile.metaData.MicroVideoOffset = 10_000_000_000

		data, err := mediaFile.ExtractVideoFromMotionPhoto()

		if err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, data)
		assert.FileExists(t, data.fileName)

		err = os.Remove(data.fileName)

		if err != nil {
			t.Fatal(err)
		}
	})

	var brokenMetadataTests = []struct {
		nameSuffix string
		length     int
	}{
		{"overflowing-offset", 10_000_000_000},
		{"zero-offset", 0},
	}

	for _, tt := range brokenMetadataTests {
		t.Run("google-v1-motion-photo-with-broken-metadata-"+tt.nameSuffix, func(t *testing.T) {
			conf := config.TestConfig()

			mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/PXL_20210506_083558892.MP.jpg")

			if err != nil {
				t.Fatal(err)
			}

			mediaFile.MetaData()
			assert.True(t, mediaFile.MetaData().MotionPhoto)

			for i, d := range mediaFile.metaData.Directory {
				if d.Semantic == MotionPhotoGoogle && d.Mime == fs.MimeTypeMP4 {
					mediaFile.metaData.Directory[i].Length = tt.length
				}
			}

			data, err := mediaFile.ExtractVideoFromMotionPhoto()

			if err != nil {
				t.Fatal(err)
			}

			assert.NotNil(t, data)
			assert.FileExists(t, data.fileName)

			err = os.Remove(data.fileName)

			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

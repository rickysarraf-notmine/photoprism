package photoprism

import (
	"time"

	"github.com/dustin/go-humanize/english"

	"github.com/photoprism/photoprism/internal/face"
	"github.com/photoprism/photoprism/pkg/clean"
)

// Faces finds faces in JPEG media files and returns them.
func (ind *Index) Faces(jpeg *MediaFile, expected int) face.Faces {
	if jpeg == nil {
		return face.Faces{}
	}

	// Select best thumbnail depending on configured size.
	thumbSize := Config().BestThumbSize()

	thumbName, err := jpeg.Thumbnail(Config().ThumbCachePath(), thumbSize)

	if err != nil {
		log.Debugf("index: %s in %s (faces)", err, clean.Log(jpeg.BaseName()))
		return face.Faces{}
	}

	if thumbName == "" {
		log.Debugf("index: thumb %s not found in %s (faces)", thumbSize, clean.Log(jpeg.BaseName()))
		return face.Faces{}
	}

	start := time.Now()

	faces, err := ind.faceNet.Detect(thumbName, Config().FaceSize(), true, expected)

	if err != nil {
		log.Debugf("%s in %s", err, clean.Log(jpeg.BaseName()))
	}

	if l := len(faces); l > 0 {
		log.Infof("index: found %s in %s [%s]", english.Plural(l, "face", "faces"), clean.Log(jpeg.BaseName()), time.Since(start))
	}

	return faces
}

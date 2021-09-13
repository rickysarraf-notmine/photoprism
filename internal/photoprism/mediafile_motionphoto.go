package photoprism

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/txt"
)

const (
	MotionPhotoSamsung = "MotionPhoto_Data"
	MotionPhotoGoogle  = "MotionPhoto"
)

func (m *MediaFile) IsMotionPhoto() bool {
	if m.MetaData().MotionPhoto {
		// Google MotionPhoto v1
		return true
	} else if m.MetaData().MicroVideo {
		// Google MotionPhoto legacy
		return true
	} else if m.MetaData().EmbeddedVideoType == MotionPhotoSamsung {
		// Samsung MotionPhoto
		return true
	}

	return false
}

// ExtractVideoFromMotionPhoto extracts the embedded motion photo video to the sidecar folder.
func (f *MediaFile) ExtractVideoFromMotionPhoto() (file *MediaFile, err error) {
	conf := Config()

	if f == nil {
		return nil, fmt.Errorf("mp: file is nil - you might have found a bug")
	}

	if !f.Exists() {
		return nil, fmt.Errorf("mp: %s not found", f.RelName(conf.OriginalsPath()))
	}

	mpName := fs.FileName(fs.StripExt(f.FileName()), conf.SidecarPath(), conf.OriginalsPath(), fs.Mp4Ext)
	mediaFile, err := NewMediaFile(mpName)

	if err == nil && mediaFile.IsVideo() {
		log.Debugf("mp: mp4 was already extracted from motion photo %s", txt.Quote(mediaFile.RelName(conf.SidecarPath())))
		return mediaFile, nil
	}

	if conf.DisableExifTool() {
		return nil, fmt.Errorf("mp: disabled in read only mode (%s)", f.RelName(conf.OriginalsPath()))
	}

	fileName := f.RelName(conf.OriginalsPath())
	log.Infof("mp: extracting mp4 from motion photo %s", txt.Quote(fileName))

	event.Publish("index.converting", event.Data{
		"fileType": f.FileType(),
		"fileName": fileName,
		"baseName": filepath.Base(fileName),
		"xmpName":  "",
	})

	if f.MetaData().MotionPhoto {
		log.Infof("mp: detected that %s is a Google motion photo", txt.Quote(fileName))

		var offset int

		for _, v := range f.MetaData().Directory {
			if v.Semantic == MotionPhotoGoogle && v.Mime == fs.MimeTypeMP4 {
				offset = v.Length
			}
		}

		data, err := ioutil.ReadFile(f.FileName())
		if err != nil {
			return nil, err
		}

		if err := ioutil.WriteFile(mpName, data[len(data)-offset:], os.ModePerm); err != nil {
			return nil, err
		}
	} else if f.MetaData().MicroVideo {
		log.Infof("mp: detected that %s is a Google legacy motion photo", txt.Quote(fileName))

		offset := f.MetaData().MicroVideoOffset

		data, err := ioutil.ReadFile(f.FileName())
		if err != nil {
			return nil, err
		}

		if err := ioutil.WriteFile(mpName, data[len(data)-offset:], os.ModePerm); err != nil {
			return nil, err
		}
	} else if f.MetaData().EmbeddedVideoType == MotionPhotoSamsung {
		log.Infof("mp: detected that %s is a Samsung motion photo", txt.Quote(fileName))

		// TODO untested
		// exiftool -b -EmbeddedVideoFile file.jpg >video.mp4
		cmd := exec.Command(conf.ExifToolBin(), "-b", "-EmbeddedVideoFile", f.FileName())

		// Fetch command output.
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr

		// Run convert command.
		if err := cmd.Run(); err != nil {
			if stderr.String() != "" {
				return nil, errors.New(stderr.String())
			} else {
				return nil, err
			}
		}

		// Write output to file.
		if err := ioutil.WriteFile(mpName, []byte(out.String()), os.ModePerm); err != nil {
			return nil, err
		}
	}

	// Check if file exists.
	if !fs.FileExists(mpName) {
		return nil, fmt.Errorf("mp: failed creating motion photo video for %s", filepath.Base(mpName))
	}

	return NewMediaFile(mpName)
}

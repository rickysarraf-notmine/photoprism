package photoprism

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/alfg/mp4"
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
		return nil, fmt.Errorf("file is nil - you might have found a bug")
	}

	if !f.Exists() {
		return nil, fmt.Errorf("file does not exist")
	}

	mpName := fs.FileName(fs.StripExt(f.FileName()), conf.SidecarPath(), conf.OriginalsPath(), fs.ExtMP4)
	mediaFile, err := NewMediaFile(mpName)

	if err == nil && mediaFile.IsVideo() {
		log.Debugf("mp: mp4 was already extracted from motion photo %s", txt.Quote(mediaFile.RelName(conf.SidecarPath())))
		return mediaFile, nil
	}

	if !conf.SidecarWritable() {
		return nil, fmt.Errorf("sidecar location is not writable")
	}

	fileName := f.RelName(conf.OriginalsPath())
	log.Infof("mp: extracting mp4 from motion photo %s", txt.Quote(fileName))

	event.Publish("index.motionphotos", event.Data{
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

		if offset == 0 {
			log.Warnf("mp: mp4 offset in motion photo directory entry is 0 in %s, resetting offset in an attempt to recover", txt.Quote(fileName))
			offset = math.MaxInt32
		}

		if err := f.extractGoogleMotionPhotoVideo(offset, mpName, fileName); err != nil {
			return nil, err
		}
	} else if f.MetaData().MicroVideo {
		log.Infof("mp: detected that %s is a Google legacy motion photo", txt.Quote(fileName))

		offset := f.MetaData().MicroVideoOffset

		if err := f.extractGoogleMotionPhotoVideo(offset, mpName, fileName); err != nil {
			return nil, err
		}
	} else if f.MetaData().EmbeddedVideoType == MotionPhotoSamsung {
		log.Infof("mp: detected that %s is a Samsung motion photo", txt.Quote(fileName))

		if conf.DisableExifTool() {
			data, err := f.EmbeddedVideoData()
			if err != nil {
				return nil, err
			}

			if err := ioutil.WriteFile(mpName, data, os.ModePerm); err != nil {
				return nil, err
			}
		} else {
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
	}

	// Check if file exists.
	if !fs.FileExists(mpName) {
		return nil, fmt.Errorf("motion photo video was not created")
	}

	return NewMediaFile(mpName)
}

func (f *MediaFile) extractGoogleMotionPhotoVideo(offset int, mpName string, fileName string) error {
	data, err := ioutil.ReadFile(f.FileName())
	if err != nil {
		return err
	}

	startIndex := len(data) - offset

	if startIndex < 0 {
		log.Warnf("mp: implausible offset %d in %s, will try to recover the embedded motion photo anyway", offset, txt.Quote(fileName))

		// The metadata is obviously incorrect, so let's make a last ditch effort to find whether
		// there is an mp4 file embedded somewhere in the file.
		for res := range fs.MimeTypeSearch(data) {
			if res.MimeType == fs.MimeTypeMP4 {
				if mp4, err := mp4.OpenFromBytes(data[res.Position:]); err == nil {
					if mp4.Ftyp != nil && mp4.Ftyp.Name == "ftyp" && mp4.Mdat != nil && mp4.Mdat.Name == "mdat" {
						log.Infof("mp: found an mp4 at index %d in %s", res.Position, txt.Quote(fileName))

						startIndex = res.Position
					}
				}
			}
		}

		if startIndex < 0 {
			return fmt.Errorf("motion photo offset %d is bigger than the file size %d and no embedded mp4 file could be found", offset, len(data))
		}
	}

	if err := ioutil.WriteFile(mpName, data[startIndex:], os.ModePerm); err != nil {
		return err
	}

	return nil
}

package meta

import (
	"fmt"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"time"

	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"

	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
)

// XMP parses an XMP file and returns a Data struct.
func XMP(fileName string) (data Data, err error) {
	err = data.XMP(fileName, fs.XmpFile)

	return data, err
}

// XMP parses an XMP file and returns a Data struct.
func (data *Data) XMP(fileName string, fileType fs.Type) (err error) {
	logName := clean.Log(filepath.Base(fileName))

	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("metadata: %s in %s (xmp panic)\nstack: %s", e, logName, debug.Stack())
		}
	}()

	// Resolve file name e.g. in case it's a symlink.
	if fileName, err = fs.Resolve(fileName); err != nil {
		return fmt.Errorf("metadata: %s %s (xmp)", err, clean.Log(filepath.Base(fileName)))
	}

	doc := XmpDocument{}

	// All unsupported types will be silently ignored, otherwise we will polute the log with warnings
	if fileType == fs.XmpFile {
		if err := doc.Load(fileName); err != nil {
			return fmt.Errorf("metadata: cannot read %s (xmp)", logName)
		}
	} else if fileType == fs.ImageJPEG {
		jpegParser := jpegstructure.NewJpegMediaParser()
		intfc, err := jpegParser.ParseFile(fileName)

		if err != nil {
			return fmt.Errorf("metadata: %s in %s (parse jpeg)", err, logName)
		}

		sl := intfc.(*jpegstructure.SegmentList)
		_, s, err := sl.FindXmp()

		if err != nil {
			return fmt.Errorf("metadata: %s in %s (read xmp)", err, logName)
		}

		xmp, err := s.FormattedXmp()

		if err != nil {
			return fmt.Errorf("metadata: %s in %s (formatting xmp)", err, logName)
		}

		if err = doc.LoadFromBytes([]byte(xmp)); err != nil {
			return fmt.Errorf("metadata: can't read xmp string from %s", xmp)
		}
	}

	if doc.Title() != "" {
		data.Title = doc.Title()
	}

	if doc.Artist() != "" {
		data.Artist = doc.Artist()
	}

	if doc.Description() != "" {
		data.Description = doc.Description()
	}

	if doc.Copyright() != "" {
		data.Copyright = doc.Copyright()
	}

	if doc.CameraMake() != "" {
		data.CameraMake = doc.CameraMake()
	}

	if doc.CameraModel() != "" {
		data.CameraModel = doc.CameraModel()
	}

	if doc.LensModel() != "" {
		data.LensModel = doc.LensModel()
	}

	if takenAt := doc.TakenAt(data.TimeZone); !takenAt.IsZero() {
		data.TakenAt = takenAt.UTC()
		if data.TimeZone == "" {
			data.TimeZone = time.UTC.String()
		}
	}

	if len(doc.Keywords()) != 0 {
		data.AddKeywords(doc.Keywords())
	}

	if microVideo := doc.MicroVideo(); microVideo > 0 {
		data.MicroVideo = true
	}

	if microVideoOffset := doc.MicroVideoOffset(); microVideoOffset > 0 {
		data.MicroVideoOffset = microVideoOffset
	}

	if motionPhoto := doc.MotionPhoto(); motionPhoto > 0 {
		data.MotionPhoto = true
	}

	directoryItems := doc.RDF.Description.Directory.Items

	for _, item := range directoryItems {
		var length int
		var padding int

		if length, err = strconv.Atoi(item.Length); err != nil {
			log.Warnf("metadata: invalid xmp %s in %s (non-integer length)", directoryItems, logName)
		}

		if padding, err = strconv.Atoi(item.Padding); err != nil {
			log.Warnf("metadata: invalid xmp %s in %s (non-integer padding)", directoryItems, logName)
		}

		data.Directory = append(data.Directory, DirectoryEntry{
			Mime:     item.Mime,
			Semantic: item.Semantic,
			Length:   length,
			Padding:  padding,
		})
	}

	return nil
}

package meta

import (
	"fmt"
	"path/filepath"
	"runtime/debug"
	"strconv"

	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"

	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/sanitize"
)

// XMP parses an XMP file and returns a Data struct.
func XMP(fileName string) (data Data, err error) {
	err = data.XMP(fileName, fs.FormatXMP)

	return data, err
}

// XMP parses an XMP file and returns a Data struct.
func (data *Data) XMP(fileName string, fileType fs.FileFormat) (err error) {
	logName := sanitize.Log(filepath.Base(fileName))

	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("metadata: %s in %s (xmp panic)\nstack: %s", e, logName, debug.Stack())
		}
	}()

	doc := XmpDocument{}

	// All unsupported types will be silently ignored, otherwise we will polute the log with warnings
	if fileType == fs.FormatXMP {
		if err := doc.Load(fileName); err != nil {
			return fmt.Errorf("metadata: can't read %s (xmp)", logName)
		}
	} else if fileType == fs.FormatJpeg {
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

	if takenAt := doc.TakenAt(); !takenAt.IsZero() {
		data.TakenAt = takenAt
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

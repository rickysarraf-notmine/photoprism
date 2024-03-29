package fs

import (
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/h2non/filetype"
)

const (
	MimeTypeUnknown = ""
	MimeTypeJPEG    = "image/jpeg"
	MimeTypeJPEGXL  = "image/jxl"
	MimeTypePNG     = "image/png"
	MimeTypeAPNG    = "image/vnd.mozilla.apng"
	MimeTypeGIF     = "image/gif"
	MimeTypeBMP     = "image/bmp"
	MimeTypeTIFF    = "image/tiff"
	MimeTypeDNG     = "image/dng"
	MimeTypeAVIF    = "image/avif"
	MimeTypeAVIFS   = "image/avif-sequence"
	MimeTypeHEIC    = "image/heic"
	MimeTypeHEICS   = "image/heic-sequence"
	MimeTypeWebP    = "image/webp"
	MimeTypeMP4     = "video/mp4"
	MimeTypeMOV     = "video/quicktime"
	MimeTypeSVG     = "image/svg+xml"
	MimeTypeAI      = "application/vnd.adobe.illustrator"
	MimeTypePS      = "application/ps"
	MimeTypeEPS     = "image/eps"
	MimeTypeXML     = "text/xml"
	MimeTypeJSON    = "application/json"
)

const (
	mimeTypeBufferSize = 261
)

type MimeTypeSearchResult struct {
	MimeType string
	Position int
}

// MimeType returns the mime type of a file, or an empty string if it could not be detected.
func MimeType(filename string) (mimeType string) {
	if filename == "" {
		return MimeTypeUnknown
	}

	// Workaround for types that cannot be reliably detected.
	switch Extensions[strings.ToLower(filepath.Ext(filename))] {
	case ImageDNG:
		return MimeTypeDNG
	case ImageAVIF:
		return MimeTypeAVIF
	case ImageAVIFS:
		return MimeTypeAVIFS
	case ImageHEIC:
		return MimeTypeHEIC
	case ImageHEICS:
		return MimeTypeHEICS
	case VideoMP4:
		return MimeTypeMP4
	case VideoMOV:
		return MimeTypeMOV
	case VectorSVG:
		return MimeTypeSVG
	case VectorAI:
		return MimeTypeAI
	case VectorPS:
		return MimeTypePS
	case VectorEPS:
		return MimeTypeEPS
	}

	if t, err := mimetype.DetectFile(filename); err != nil {
		return MimeTypeUnknown
	} else {
		mimeType, _, _ = strings.Cut(t.String(), ";")
	}

	return mimeType
}

// MimeTypeSearch searches the given buffer for any valid mime types.
func MimeTypeSearch(buffer []byte) <-chan MimeTypeSearchResult {
	chnl := make(chan MimeTypeSearchResult)

	go func() {
		for i := 0; i < len(buffer)-mimeTypeBufferSize; i++ {
			data := buffer[i : i+mimeTypeBufferSize]

			if t, err := filetype.Get(data); err == nil && t != filetype.Unknown {
				chnl <- MimeTypeSearchResult{
					MimeType: t.MIME.Value,
					Position: i,
				}
			}
		}
		close(chnl)
	}()

	return chnl
}

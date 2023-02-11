package fs

import (
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/h2non/filetype"
)

const (
	MimeTypeUnknown = ""
	MimeTypeJpeg    = "image/jpeg"
	MimeTypePng     = "image/png"
	MimeTypeGif     = "image/gif"
	MimeTypeBitmap  = "image/bmp"
	MimeTypeTiff    = "image/tiff"
	MimeTypeDNG     = "image/dng"
	MimeTypeAVIF    = "image/avif"
	MimeTypeHEIC    = "image/heic"
	MimeTypeWebP    = "image/webp"
	MimeTypeMP4     = "video/mp4"
	MimeTypeMOV     = "video/quicktime"
        MimeTypeSVG     = "image/svg+xml"
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
	// Workaround, since "image/dng" cannot be recognized yet.
	if ext := Extensions[strings.ToLower(filepath.Ext(filename))]; ext == "" {
		// Continue.
	} else if ext == ImageDNG {
		return MimeTypeDNG
	} else if ext == ImageAVIF {
		return MimeTypeAVIF
	} else if ext == VideoMP4 {
		return MimeTypeMP4
	} else if ext == VideoMOV {
		return MimeTypeMOV
        } else if ext == VectorSVG {
                return MimeTypeSVG
        } else if ext == VectorPS {
                return MimeTypePS
        } else if ext == VectorEPS {
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

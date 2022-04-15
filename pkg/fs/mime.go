package fs

import (
	"os"

	"github.com/h2non/filetype"
)

const (
	MimeTypeJpeg   = "image/jpeg"
	MimeTypePng    = "image/png"
	MimeTypeGif    = "image/gif"
	MimeTypeBitmap = "image/bmp"
	MimeTypeWebP   = "image/webp"
	MimeTypeTiff   = "image/tiff"
	MimeTypeHEIF   = "image/heif"
	MimeTypeMP4    = "video/mp4"
)

const (
	mimeTypeBufferSize = 261
)

type MimeTypeSearchResult struct {
	MimeType string
	Position int
}

// MimeType returns the mime type of a file, an empty string if it is unknown.
func MimeType(filename string) string {
	handle, err := os.Open(filename)

	if err != nil {
		return ""
	}

	defer handle.Close()

	// Only the first 261 bytes are used to sniff the content type.
	buffer := make([]byte, mimeTypeBufferSize)

	if _, err := handle.Read(buffer); err != nil {
		return ""
	} else if t, err := filetype.Get(buffer); err == nil && t != filetype.Unknown {
		return t.MIME.Value
	} else if t := filetype.GetType(NormalizedExt(filename)); t != filetype.Unknown {
		return t.MIME.Value
	} else {
		return ""
	}
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

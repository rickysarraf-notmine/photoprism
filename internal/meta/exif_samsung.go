package meta

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/photoprism/photoprism/pkg/fs"
)

const (
	samsungTrailerHeadMarker = "SEFH"
	samsungTrailerTailMarker = "SEFT"
)

const (
	typeEmbeddedImage         = 0x0001
	typeEmbeddedAudioFileName = 0x0100
	typeSurroundShotVideoName = 0x0201
	typeTimeStamp             = 0x0a01
	typeDualCameraImageName   = 0x0a20
	typeEmbeddedVideoType     = 0x0a30
)

var (
	supportedTrailerVersions = map[uint32]struct{}{
		101: {},
		103: {},
		105: {},
		106: {},
	}
)

type trailerEntry struct {
	Type uint16
	Name string
	Data []byte
}

// ExifSamsungTrailer reads and populates the metadata from the trailer section  for images taken with a Samsung device.
//
// Resources:
// 	- https://gist.github.com/HikerBoricua/54aec42ee47f2ecb374ec3208b3a98ee
// 	- https://github.com/exiftool/exiftool/blob/57f44297961839f40e70d682865c41828b7f71b5/lib/Image/ExifTool/JPEG.pm#L277-L279
// 	- https://github.com/exiftool/exiftool/blob/57f44297961839f40e70d682865c41828b7f71b5/lib/Image/ExifTool/Samsung.pm#L1300-L1322
func (metadata *Data) ExifSamsungTrailer(fileName string, fileType fs.FileFormat) (bool, error) {
	if !metadata.isSamsungMedia(fileType) {
		return false, nil
	}

	trailerEntries, err := readTrailerEntries(fileName)

	if err != nil {
		return false, err
	}

	if _, ok := trailerEntries[typeEmbeddedVideoType]; ok {
		metadata.EmbeddedVideoType = trailerEntries[typeEmbeddedVideoType].Name
	}

	return true, nil
}

// EmbeddedVideoData returns the embedded video data from a Samsung motion photo.
func (metadata Data) EmbeddedVideoData(fileName string, fileType fs.FileFormat) ([]byte, error) {
	if !metadata.isSamsungMedia(fileType) {
		return nil, nil
	}

	trailerEntries, err := readTrailerEntries(fileName)

	if err != nil {
		return nil, err
	}

	if _, ok := trailerEntries[typeEmbeddedVideoType]; !ok {
		return nil, fmt.Errorf("no EmbeddedVideoType entry")
	}

	return trailerEntries[typeEmbeddedVideoType].Data, nil
}

func (metadata Data) isSamsungMedia(fileType fs.FileFormat) bool {
	if fileType != fs.FormatJpeg {
		return false
	}

	// Let's assume that the generic metadata is already available.
	// This way we dont have to even bother opening the file.
	if strings.ToLower(metadata.CameraMake) != "samsung" {
		return false
	}

	return true
}

func readTrailer(data []byte) ([]byte, int, error) {
	trailerTail := string(data[len(data)-4:])
	if trailerTail != samsungTrailerTailMarker {
		return nil, 0, nil
	}

	var trailerLength uint32
	err := binary.Read(bytes.NewReader(data[len(data)-8:len(data)-4]), binary.LittleEndian, &trailerLength)

	if err != nil {
		return nil, 0, err
	}

	trailerIndex := len(data) - 8 - int(trailerLength)
	if trailerIndex < 0 {
		return nil, 0, fmt.Errorf("trailer length is invalid %d", trailerLength)
	}

	trailer := data[trailerIndex:]
	trailerHead := string(trailer[:4])
	if trailerHead != samsungTrailerHeadMarker {
		return nil, 0, nil
	}

	var version uint32
	err = binary.Read(bytes.NewReader(trailer[4:8]), binary.LittleEndian, &version)

	if err != nil {
		return nil, 0, err
	}

	_, ok := supportedTrailerVersions[version]
	if !ok {
		return nil, 0, fmt.Errorf("unsupported Samsung trailer version %d", version)
	}

	return trailer, trailerIndex, nil
}

func readTrailerEntries(fileName string) (map[uint16]trailerEntry, error) {
	entries := make(map[uint16]trailerEntry)

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	trailer, trailerIndex, err := readTrailer(data)
	if err != nil {
		return nil, err
	} else if trailer == nil {
		// The media file was not a Samsung motion photo, so there is no need to return an error.
		return nil, nil
	}

	var count uint32
	err = binary.Read(bytes.NewReader(trailer[8:12]), binary.LittleEndian, &count)

	if err != nil {
		return nil, err
	}

	var i uint32
	for i = 0; i < count; i++ {
		entry := 12 + 12*i

		var typ uint16
		err = binary.Read(bytes.NewReader(trailer[entry+2:entry+4]), binary.LittleEndian, &typ)

		if err != nil {
			return nil, err
		}

		var noff uint32
		err = binary.Read(bytes.NewReader(trailer[entry+4:entry+8]), binary.LittleEndian, &noff)

		if err != nil {
			return nil, err
		}

		var size uint32
		err = binary.Read(bytes.NewReader(trailer[entry+8:entry+12]), binary.LittleEndian, &size)

		if err != nil {
			return nil, err
		}

		entryPos := uint32(trailerIndex) - noff
		entryLen := size

		entryPayload := data[entryPos : entryPos+entryLen]

		var entryType uint16
		err = binary.Read(bytes.NewReader(entryPayload[2:4]), binary.LittleEndian, &entryType)

		if err != nil {
			return nil, err
		}

		if typ != entryType {
			return nil, fmt.Errorf("mismatch between type %d and entry type %d", typ, entryType)
		}

		var entryOffset uint32
		err = binary.Read(bytes.NewReader(entryPayload[4:8]), binary.LittleEndian, &entryOffset)

		if err != nil {
			return nil, err
		}

		entryName := string(entryPayload[8 : 8+entryOffset])
		entryData := entryPayload[8+entryOffset:]

		entries[typ] = trailerEntry{
			Type: typ,
			Name: entryName,
			Data: entryData,
		}
	}

	return entries, nil
}

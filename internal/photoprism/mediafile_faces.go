package photoprism

import (
	"math"
	"strconv"
	"strings"

	"github.com/photoprism/photoprism/internal/face"
)

const (
	IPTCShapeRectangle = "rectangle"
	IPTCShapeCircle    = "circle"
	IPTCUnitPixel      = "pixel"
	IPTCUnitRelative   = "relative"
)

// HasFaces returns whether the media contains face region metadata.
func (m *MediaFile) HasFaces() bool {
	if len(m.MetaData().Regions) > 0 {
		// Metadata Working Group (MWG) Format
		return true
	} else if len(m.MetaData().RegionsMP) > 0 {
		// Microsoft Windows Live Photo Gallery (WLPG)
		return true
	} else if len(m.MetaData().RegionsIPTC) > 0 {
		// IPTC
		return true
	}

	return false
}

// Faces returns all unique metadata-based face regions for the given media.
func (m *MediaFile) Faces() face.Faces {
	faces := face.Faces{}

	faces.AppendIfNotContains(m.facesIPTC()...)
	faces.AppendIfNotContains(m.facesMWG()...)
	faces.AppendIfNotContains(m.facesWLPG()...)

	return faces
}

func (m *MediaFile) facesMWG() (faces face.Faces) {
	fittingFn := math.Min // or math.Max

	if len(m.MetaData().Regions) > 0 {
		for _, region := range m.MetaData().Regions {
			if !strings.EqualFold(region.Type, "face") {
				continue
			}

			face := face.Face{
				Rows: m.Height(),
				Cols: m.Width(),
				Area: face.Area{
					Name:  region.Name,
					Row:   int(region.Area.Y * float32(m.Height())),
					Col:   int(region.Area.X * float32(m.Width())),
					Scale: int(fittingFn(float64(region.Area.H)*float64(m.Height()), float64(region.Area.W)*float64(m.Width()))),
				},
			}

			faces = append(faces, face)
		}
	}

	return faces
}

func (m *MediaFile) facesWLPG() (faces face.Faces) {
	fittingFn := math.Min // or math.Max

	if len(m.MetaData().RegionsMP) > 0 {
		for _, region := range m.MetaData().RegionsMP {
			rect := strings.Split(strings.ReplaceAll(region.Rectangle, " ", ""), ",")
			if len(rect) != 4 {
				log.Warnf("faces: face region rectangle '%v' does not contain 4 values (%s)", rect, m.FileName())
				continue
			}

			x, err := strconv.ParseFloat(rect[0], 64)
			if err != nil {
				log.Warnf("faces: face region x %s is not a float (%s)", rect[0], m.FileName())
				continue
			}

			y, err := strconv.ParseFloat(rect[1], 64)
			if err != nil {
				log.Warnf("faces: face region y %s is not a float (%s)", rect[1], m.FileName())
				continue
			}

			w, err := strconv.ParseFloat(rect[2], 64)
			if err != nil {
				log.Warnf("faces: face region w %s is not a float (%s)", rect[2], m.FileName())
				continue
			}

			h, err := strconv.ParseFloat(rect[3], 64)
			if err != nil {
				log.Warnf("faces: face region h %s is not a float (%s)", rect[3], m.FileName())
				continue
			}

			x += w / 2
			y += h / 2

			face := face.Face{
				Rows: m.Height(),
				Cols: m.Width(),
				Area: face.Area{
					Name:  region.PersonDisplayName,
					Row:   int(y * float64(m.Height())),
					Col:   int(x * float64(m.Width())),
					Scale: int(fittingFn(h*float64(m.Height()), w*float64(m.Width()))),
				},
			}

			faces = append(faces, face)
		}
	}

	return faces
}

func (m *MediaFile) facesIPTC() (faces face.Faces) {
	fittingFn := math.Min // or math.Max

	if len(m.MetaData().RegionsIPTC) > 0 {
		for _, region := range m.MetaData().RegionsIPTC {
			if len(region.Person) == 0 {
				continue
			}

			shape := strings.ToLower(region.Boundary.Shape)
			unit := strings.ToLower(region.Boundary.Unit)

			var x, y, w, h, wScale, hScale float64

			switch shape {
			case IPTCShapeRectangle:
				x = region.Boundary.X
				y = region.Boundary.Y
				w = region.Boundary.W
				h = region.Boundary.H

				x += w / 2
				y += h / 2
			case IPTCShapeCircle:
				x = region.Boundary.X
				y = region.Boundary.Y
				w = region.Boundary.Rx * 2
				h = region.Boundary.Rx * 2
			default:
				// Polygon is not supported
				log.Warnf("faces: %s IPTC face regions are not supported (%s)", shape, m.FileName())
				continue
			}

			switch unit {
			case IPTCUnitPixel:
				hScale = 1
				wScale = 1
			case IPTCUnitRelative:
				hScale = float64(m.Height())
				wScale = float64(m.Width())
			default:
				log.Warnf("faces: %s IPTC face regions are not supported (%s)", unit, m.FileName())
				continue
			}

			face := face.Face{
				Rows: m.Height(),
				Cols: m.Width(),
				Area: face.Area{
					Name:  region.Person[0],
					Row:   int(y * hScale),
					Col:   int(x * wScale),
					Scale: int(fittingFn(h*hScale, w*wScale)),
				},
			}

			faces = append(faces, face)
		}
	}

	return faces
}

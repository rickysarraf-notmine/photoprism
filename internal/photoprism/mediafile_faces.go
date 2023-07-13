package photoprism

import (
	"math"
	"strconv"
	"strings"

	"github.com/photoprism/photoprism/internal/face"
	"github.com/photoprism/photoprism/internal/meta"
	"github.com/photoprism/photoprism/pkg/clean"
)

const (
	IPTCShapeRectangle = "rectangle"
	IPTCShapeCircle    = "circle"
	IPTCUnitPixel      = "pixel"
	IPTCUnitRelative   = "relative"

	MWGTypeFace    = "face"
	MWGTypePet     = "pet"
	MWGTypeFocus   = "focus"
	MWGTypeBarcode = "barcode"

	MWGUnitNormalized = "normalized"
)

// area describes a face region bounding box
type area struct {
	x float64 // face center
	y float64 // face center
	w float64 // face width
	h float64 // face height
}

// Normalize takes the image orientation into account and rotates the area accordingly if needed.
func (a *area) Normalize(orientation int) {
	if orientation > 4 {
		a.w, a.h = a.h, a.w
		a.x, a.y = a.y, a.x
	}

	swapX := 0.
	swapY := 0.

	switch orientation {
	case 2, 6:
		swapX = 1
	case 3, 7:
		swapX = 1
		swapY = 1
	case 4, 8:
		swapY = 1
	}

	a.x = math.Abs(a.x - swapX)
	a.y = math.Abs(a.y - swapY)
}

// Row calculates the face region's centerpoint along the Y axis.
func (a area) Row(metadata meta.Data) int {
	return int(a.y * float64(metadata.ActualHeight()))
}

// Col calculates the face region's centerpoint along the X axis.
func (a area) Col(metadata meta.Data) int {
	return int(a.x * float64(metadata.ActualWidth()))
}

// Scale calculates the size of the face region (which is square).
func (a area) Scale(metadata meta.Data) int {
	// Determine how to clip the region rectangle - along the short ot long side
	fittingFn := math.Min // or math.Max

	return int(fittingFn(a.h*float64(metadata.ActualHeight()), a.w*float64(metadata.ActualWidth())))
}

func normalizedAreaFromCoords(x, y, w, h float64, orientation int) area {
	area := area{x, y, w, h}
	area.Normalize(orientation)

	return area
}

func normalizedAreaFromArea(a meta.Area, orientation int) area {
	area := area{a.X, a.Y, a.W, a.H}
	area.Normalize(orientation)

	return area
}

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
// All region coordinates will be converted to normal orientation.
func (m *MediaFile) Faces() face.Faces {
	faces := face.Faces{}

	faces.AppendIfNotContains(m.facesIPTC()...)
	faces.AppendIfNotContains(m.facesMWG()...)
	faces.AppendIfNotContains(m.facesWLPG()...)

	return faces
}

func (m *MediaFile) facesMWG() (faces face.Faces) {
	logName := clean.Log(m.BaseName())

	if len(m.MetaData().Regions) > 0 {
		for _, region := range m.MetaData().Regions {
			person := region.Name

			if !strings.EqualFold(region.Type, MWGTypeFace) {
				log.Warnf("faces: ignoring MWG region for %s, as %s type is not supported (%s)", person, region.Type, logName)
				continue
			}

			if !strings.EqualFold(region.Area.Unit, MWGUnitNormalized) {
				log.Warnf("faces: ignoring MWG region for %s, as %s unit is not supported (%s)", person, region.Area.Unit, logName)
				continue
			}

			area := normalizedAreaFromArea(region.Area, m.Orientation())

			face := face.Face{
				Rows: m.metaData.ActualHeight(),
				Cols: m.metaData.ActualWidth(),
				Area: face.Area{
					Name:  person,
					Row:   area.Row(m.MetaData()),
					Col:   area.Col(m.MetaData()),
					Scale: area.Scale(m.MetaData()),
				},
			}

			faces = append(faces, face)
		}
	}

	return faces
}

func (m *MediaFile) facesWLPG() (faces face.Faces) {
	logName := clean.Log(m.BaseName())

	if len(m.MetaData().RegionsMP) > 0 {
		for _, region := range m.MetaData().RegionsMP {
			rect := strings.Split(strings.ReplaceAll(region.Rectangle, " ", ""), ",")
			if len(rect) != 4 {
				log.Warnf("faces: WLPG face region rectangle '%v' does not contain 4 values (%s)", rect, logName)
				continue
			}

			x, err := strconv.ParseFloat(rect[0], 64)
			if err != nil {
				log.Warnf("faces: WLPG face region x %s is not a float (%s)", rect[0], logName)
				continue
			}

			y, err := strconv.ParseFloat(rect[1], 64)
			if err != nil {
				log.Warnf("faces: WLPG face region y %s is not a float (%s)", rect[1], logName)
				continue
			}

			w, err := strconv.ParseFloat(rect[2], 64)
			if err != nil {
				log.Warnf("faces: WLPG face region w %s is not a float (%s)", rect[2], logName)
				continue
			}

			h, err := strconv.ParseFloat(rect[3], 64)
			if err != nil {
				log.Warnf("faces: WLPG face region h %s is not a float (%s)", rect[3], logName)
				continue
			}

			x += w / 2
			y += h / 2

			area := normalizedAreaFromCoords(x, y, w, h, m.Orientation())

			face := face.Face{
				Rows: m.metaData.ActualHeight(),
				Cols: m.metaData.ActualWidth(),
				Area: face.Area{
					Name:  region.PersonDisplayName,
					Row:   area.Row(m.MetaData()),
					Col:   area.Col(m.MetaData()),
					Scale: area.Scale(m.MetaData()),
				},
			}

			faces = append(faces, face)
		}
	}

	return faces
}

func (m *MediaFile) facesIPTC() (faces face.Faces) {
	logName := clean.Log(m.BaseName())

	if len(m.MetaData().RegionsIPTC) > 0 {
		for _, region := range m.MetaData().RegionsIPTC {
			if len(region.Person) == 0 {
				log.Warnf("faces: IPTC face region does not contain a person (%s)", logName)
				continue
			} else if len(region.Person) > 1 {
				log.Warnf("faces: IPTC face region contains more than one person %s (%s)", region.Person, logName)
			}

			shape := strings.ToLower(region.Boundary.Shape)
			unit := strings.ToLower(region.Boundary.Unit)

			person := region.Person[0]

			var x, y, w, h float64

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
				log.Warnf("faces: ignoring IPTC face region for %s, as %s shape is not supported (%s)", person, shape, logName)
				continue
			}

			switch unit {
			case IPTCUnitPixel:
				width := float64(m.metaData.ActualWidth())
				height := float64(m.metaData.ActualHeight())

				x /= width
				y /= height

				w /= width
				h /= height
			case IPTCUnitRelative:
			default:
				log.Warnf("faces: ignoring IPTC face region for %s, as %s unit is not supported (%s)", person, unit, logName)
				continue
			}

			area := normalizedAreaFromCoords(x, y, w, h, m.Orientation())

			face := face.Face{
				Rows: m.Height(),
				Cols: m.Width(),
				Area: face.Area{
					Name:  person,
					Row:   area.Row(m.MetaData()),
					Col:   area.Col(m.MetaData()),
					Scale: area.Scale(m.MetaData()),
				},
			}

			faces = append(faces, face)
		}
	}

	return faces
}

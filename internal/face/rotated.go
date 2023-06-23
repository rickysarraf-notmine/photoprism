package face

import (
	"github.com/disintegration/imaging"
	"github.com/photoprism/photoprism/internal/crop"
)

// RotatedFace represents a face detected at an angle (rotated).
type RotatedFace struct {
	Face
	Angle float64
}

// RotatedFaces represents a list of faces detected at an angle (rotated).
type RotatedFaces []RotatedFace

// Match finds a sufficiently overlapping existing face region for a given face.
func (faces RotatedFaces) Match(other Face) *RotatedFace {
	cropArea := other.CropArea()

	for _, f := range faces {
		if f.OverlapsAboveThreshold(cropArea) {
			return &f
		}
	}

	return nil
}

// Embeddings computes the embeddings vector for a given angled face region.
func (t *Net) EmbeddingsRotated(fileName string, face RotatedFace) (Embeddings, error) {
	err := t.loadModel()

	if err != nil {
		return nil, err
	}

	img, err := crop.ImageFromThumb(fileName, face.CropArea(), CropSize, false)

	if err != nil {
		return nil, err
	}

	// THERE ARE A LOT OF GAPS IN THE RANGE, COVER THEM ALL OR USE THE GENERIC ROTATE METHOD WITH A SUITABLE BACKGROUND COLOR
	if (face.Angle >= 0.2 && face.Angle <= 0.3) {
		img = imaging.Rotate90(img)
	} else if (face.Angle >= 0.4 && face.Angle <= 0.6) {
		img = imaging.Rotate180(img)
	} else if (face.Angle >= 0.7 && face.Angle <= 0.8) {
		img = imaging.Rotate270(img)
	}

	return t.getEmbeddings(img), nil
}

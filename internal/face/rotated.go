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

// EmbeddingsRotated computes the embeddings vector for a given angled face region.
func (t *Net) EmbeddingsRotated(fileName string, face RotatedFace) (Embeddings, error) {
	err := t.loadModel()

	if err != nil {
		return nil, err
	}

	img, err := crop.ImageFromThumb(fileName, face.CropArea(), CropSize, false)

	if err != nil {
		return nil, err
	}

	// After we have cropped the face region from the image, we need to rotate it accordingly.
	// Weirdly, rotating the cropped image by the same angle, as the face returns worse results, compared to when using binning.
	// img = imaging.Rotate(img, face.Angle * 360, color.Transparent)
	// The binning approach rounds the face angle to the neareast 90° angle, so that the resulting image is always rotated
	// by one of the "standart" orientations. The bins are as follows:
	// [0°, 45°] -> do not rotate
	// (45°, 135°] -> rotate by 90°
	// (135°, 225°] -> rotate by 180°
	// (225°, 315°] -> rotate by 270°
	// (315°, 360°] -> do not rotate
	if face.Angle > 0.125 && face.Angle <= 0.375 {
		img = imaging.Rotate90(img)
	} else if face.Angle > 0.375 && face.Angle <= 0.625 {
		img = imaging.Rotate180(img)
	} else if face.Angle > 0.625 && face.Angle <= 0.875 {
		img = imaging.Rotate270(img)
	}

	return t.getEmbeddings(img), nil
}

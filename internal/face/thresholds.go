package face

import (
	"github.com/photoprism/photoprism/internal/crop"
)

var CropSize = crop.Sizes[crop.Tile160]          // Face image crop size for FaceNet.
var OverlapThreshold = 42                        // Face area overlap threshold in percent.
var OverlapThresholdFloor = OverlapThreshold - 1 // Reduced overlap area to avoid rounding inconsistencies.
var ScoreThreshold = 9.0                         // Min face score.
var ClusterScoreThreshold = 15                   // Min score for faces forming a cluster.
var SizeThreshold = 50                           // Min face size in pixels.
var ClusterSizeThreshold = 80                    // Min size for faces forming a cluster in pixels.
var ClusterDist = 0.64                           // Similarity distance threshold of faces forming a cluster core.
var MatchDist = 0.46                             // Dist offset threshold for matching new faces with clusters.
var ClusterCore = 4                              // Min number of faces forming a cluster core.
var SampleThreshold = 2 * ClusterCore            // Threshold for automatic clustering to start.

// Rotation angles for face regions' face detection (in degrees).
//
// The most probable reason why an exif face region could not be detected by PhotoPrism seems to be that the face region
// is at a weird angle. To support such angled face regions, a fallback has been added, which kicks in when no embeddings
// can be computed for an exif face region. In this case, face detetion is ran for a list of user-configured rotation angles
// and it's checked whether any detected face significatly overlaps with the exif face region.
//
// During testing the following rotation angles have been identified as good candidates: 72°, 252° and 288° (in addition to 0°).
// However per default, the image will not be rotated, as any configured angle will cause additional computational overhead,
// so it is left up to the user to configure meaningful values based on the photos in their collection.
var RegionAngles = []int64{0}

// QualityThreshold returns the scale adjusted quality score threshold.
func QualityThreshold(scale int) (score float32) {
	score = float32(ScoreThreshold)

	// Smaller faces require higher quality.
	switch {
	case scale < 26:
		score += 26.0
	case scale < 32:
		score += 16.0
	case scale < 40:
		score += 11.0
	case scale < 50:
		score += 9.0
	case scale < 80:
		score += 6.0
	case scale < 110:
		score += 2.0
	}

	return score
}

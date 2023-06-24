package face

// DetectorConfig represents face dtector config options.
type DetectorConfig struct {
	FindLandmarks    bool    // Flag whether face landmarks (eyes, mouth) should be computed
	IgnoreLowQuality bool    // Flag whether low quality face detection results should be ignored.
	Angle            float64 // Angle at which the face detection should be done. Value from 0.0 to 1.0, where 0.0 represents 0° and 1.0 - 360°.
}

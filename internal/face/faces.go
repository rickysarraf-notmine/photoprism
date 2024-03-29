package face

// Faces represents a list of faces detected.
type Faces []Face

// Contains returns true if the face conflicts with existing faces.
func (faces Faces) Contains(other Face) bool {
	cropArea := other.CropArea()

	for _, f := range faces {
		if f.OverlapsAboveThreshold(cropArea) {
			return true
		}
	}

	return false
}

// Append adds a face.
func (faces *Faces) Append(f Face) {
	*faces = append(*faces, f)
}

// AppendIfNotContains adds all faces that do not conflict with existing ones.
func (faces *Faces) AppendIfNotContains(others ...Face) {
	for _, other := range others {
		if !faces.Contains(other) {
			faces.Append(other)
		}
	}
}

// Match finds a sufficiently overlapping existing face region for a given face.
func (faces Faces) Match(other Face) *Face {
	cropArea := other.CropArea()

	for _, f := range faces {
		if f.OverlapsAboveThreshold(cropArea) {
			return &f
		}
	}

	return nil
}

// Count returns the number of faces detected.
func (faces Faces) Count() int {
	return len(faces)
}

// Uncertainty return the max face detection uncertainty in percent.
func (faces Faces) Uncertainty() int {
	if len(faces) < 1 {
		return 100
	}

	maxScore := 0

	for _, f := range faces {
		if f.Score > maxScore {
			maxScore = f.Score
		}
	}

	switch {
	case maxScore > 300:
		return 1
	case maxScore > 200:
		return 5
	case maxScore > 100:
		return 10
	case maxScore > 80:
		return 15
	case maxScore > 65:
		return 20
	case maxScore > 50:
		return 25
	case maxScore > 40:
		return 30
	case maxScore > 30:
		return 35
	case maxScore > 20:
		return 40
	case maxScore > 10:
		return 45
	}

	return 50
}

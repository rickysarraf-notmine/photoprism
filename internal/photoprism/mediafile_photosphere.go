package photoprism

// IsPhotosphere returns whether the media is a photo sphere image
func (m *MediaFile) IsPhotosphere() bool {
	if m.MetaData().IsPhotosphere {
		// Google
		return true
	}

	return false
}

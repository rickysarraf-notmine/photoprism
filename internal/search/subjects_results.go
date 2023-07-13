package search

// Subject represents a subject search result.
type Subject struct {
	SubjUID      string `json:"UID"`
	MarkerUID    string `json:"MarkerUID"`
	MarkerSrc    string `json:"MarkerSrc,omitempty"`
	SubjType     string `json:"Type"`
	SubjSlug     string `json:"Slug"`
	SubjName     string `json:"Name"`
	SubjAlias    string `json:"Alias"`
	SubjFavorite bool   `json:"Favorite"`
	SubjHidden   bool   `json:"Hidden"`
	SubjPrivate  bool   `json:"Private"`
	SubjExcluded bool   `json:"Excluded"`
	FileCount    int    `json:"FileCount"`
	PhotoCount   int    `json:"PhotoCount"`
	Thumb        string `json:"Thumb"`
	ThumbSrc     string `json:"ThumbSrc,omitempty"`
}

// SubjectResults represents subject search results.
type SubjectResults []Subject

// Names returns a slice of subject names.
func (subjects SubjectResults) Names() []string {
	result := make([]string, len(subjects))

	for i, el := range subjects {
		result[i] = el.SubjName
	}

	return result
}

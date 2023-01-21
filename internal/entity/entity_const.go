package entity

import (
	"github.com/photoprism/photoprism/pkg/media"
	"github.com/sirupsen/logrus"
)

// Default values.
const (
	Unknown      = ""
	UnknownYear  = -1
	UnknownMonth = -1
	UnknownDay   = -1
	UnknownName  = "Unknown"
	UnknownTitle = UnknownName
	UnknownID    = "zz"
)

// Media content types.
const (
	MediaUnknown  = ""
	MediaImage    = string(media.Image)
	MediaVector   = string(media.Vector)
	MediaAnimated = string(media.Animated)
	MediaLive     = string(media.Live)
	MediaVideo    = string(media.Video)
	MediaRaw      = string(media.Raw)
	MediaText     = string(media.Text)
	MediaSphere   = string(media.Sphere)
)

// Storage root folders.
const (
	RootUnknown   = ""
	RootOriginals = "/"
	RootExamples  = "examples"
	RootSidecar   = "sidecar"
	RootImport    = "import"
	RootPath      = "/"
)

// Event type.
const (
	Created = "created"
	Updated = "updated"
	Deleted = "deleted"
)

// Photo stack states.
const (
	IsStacked   int8 = 1
	IsStackable int8 = 0
	IsUnstacked int8 = -1
)

// Authentication providers.
const (
	ProviderNone     = ""
	ProviderPassword = "password"
)

// Sort options.
const (
	SortOrderDefault   = ""
	SortOrderRelevance = "relevance"
	SortOrderDuration  = "duration"
	SortOrderSize      = "size"
	SortOrderCount     = "count"
	SortOrderAdded     = "added"
	SortOrderImported  = "imported"
	SortOrderEdited    = "edited"
	SortOrderNewest    = "newest"
	SortOrderOldest    = "oldest"
	SortOrderPlace     = "place"
	SortOrderMoment    = "moment"
	SortOrderFavorites = "favorites"
	SortOrderName      = "name"
	SortOrderPath      = "path"
	SortOrderSlug      = "slug"
	SortOrderCategory  = "category"
	SortOrderSimilar   = "similar"
	SortOrderRandom    = "random"
)

// Date Modes

const (
	DateModeLast    = "last"
	DateModeFirst   = "first"
	DateModeAverage = "average"
)

// Log levels.
const (
	PanicLevel logrus.Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

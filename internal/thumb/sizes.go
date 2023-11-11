package thumb

var (
	SizePrecached = 2048
	SizeUncached  = 7680
	Filter        = ResampleLanczos
)

// MaxSize returns the max supported size in pixels.
func MaxSize() int {
	if SizePrecached > SizeUncached {
		return SizePrecached
	}

	return SizeUncached
}

// InvalidSize tests if the size in pixels is invalid.
func InvalidSize(size int) bool {
	return size < 0 || size > MaxSize()
}

// SizeList represents a list of sizes.
type SizeList []Size

// SizeMap maps size names to sizes.
type SizeMap map[Name]Size

// All returns a slice containing all sizes.
func (m SizeMap) All() SizeList {
	result := make(SizeList, 0, len(m))

	for _, s := range m {
		result = append(result, s)
	}

	return result
}

// Sizes contains the properties of all thumbnail sizes.
var Sizes = SizeMap{
	Tile50:   {Tile50, Tile500, "List View", 50, 50, false, false, []ResampleOption{ResampleFillCenter, ResampleDefault}},
	Tile100:  {Tile100, Tile500, "Places View", 100, 100, false, false, []ResampleOption{ResampleFillCenter, ResampleDefault}},
	Tile224:  {Tile224, Tile500, "TensorFlow, Mosaic View", 224, 224, false, false, []ResampleOption{ResampleFillCenter, ResampleDefault}},
	Tile500:  {Tile500, "", "Cards View", 500, 500, false, false, []ResampleOption{ResampleFillCenter, ResampleDefault}},
	Colors:   {Colors, Fit720, "Color Detection", 3, 3, false, false, []ResampleOption{ResampleResize, ResampleNearestNeighbor, ResamplePng}},
	Left224:  {Left224, Fit720, "TensorFlow", 224, 224, false, false, []ResampleOption{ResampleFillTopLeft, ResampleDefault}},
	Right224: {Right224, Fit720, "TensorFlow", 224, 224, false, false, []ResampleOption{ResampleFillBottomRight, ResampleDefault}},
	Fit720:   {Fit720, "", "SD TV, Mobile", 720, 720, true, true, []ResampleOption{ResampleFit, ResampleDefault}},
	Fit1280:  {Fit1280, Fit2048, "HD TV, SXGA", 1280, 1024, true, true, []ResampleOption{ResampleFit, ResampleDefault}},
	Fit1920:  {Fit1920, Fit2048, "Full HD", 1920, 1200, true, true, []ResampleOption{ResampleFit, ResampleDefault}},
	Fit2048:  {Fit2048, "", "DCI 2K, Tablets", 2048, 2048, true, true, []ResampleOption{ResampleFit, ResampleDefault}},
	Fit2560:  {Fit2560, "", "Quad HD, Notebooks", 2560, 1600, true, true, []ResampleOption{ResampleFit, ResampleDefault}},
	Fit3840:  {Fit3840, "", "4K Ultra HD", 3840, 2400, true, true, []ResampleOption{ResampleFit, ResampleDefault}}, // Deprecated in favor of fit_4096
	Fit4096:  {Fit4096, "", "DCI 4K, Retina 4K", 4096, 4096, true, true, []ResampleOption{ResampleFit, ResampleDefault}},
	Fit7680:  {Fit7680, "", "8K Ultra HD 2", 7680, 4320, true, true, []ResampleOption{ResampleFit, ResampleDefault}},
}

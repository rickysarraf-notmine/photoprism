package entity

import (
	"time"
)

var editTime = time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC)
var deleteTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

type PhotoMap map[string]Photo

func (m PhotoMap) Get(name string) Photo {
	if result, ok := m[name]; ok {
		return result
	}

	return Photo{PhotoName: name}
}

func (m PhotoMap) Pointer(name string) *Photo {
	if result, ok := m[name]; ok {
		return &result
	}

	return &Photo{PhotoName: name}
}

var PhotoFixtures = PhotoMap{
	"19800101_000002_D640C559": {
		ID:               1000000,
		PhotoUID:         "pt9jtdre2lvl0yh7",
		TakenAt:          time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenSrc:         "meta",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoDescription: "photo description lake",
		PhotoPath:        "2790/02",
		PhotoName:        "19800101_000002_D640C559",
		PhotoQuality:     4,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     5,
		PhotoExposure:    "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		CameraSerial:     "",
		CameraSrc:        "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoYear:        2790,
		PhotoMonth:       2,
		Details:          DetailsFixtures.Get("lake", 1000000),
		DescriptionSrc:   "",
		LocationID:       UnknownLocation.ID,
		Location:         &UnknownLocation,
		PlaceID:          UnknownPlace.ID,
		Place:            &UnknownPlace,
		PhotoCountry:     UnknownPlace.CountryCode(),
		Links:            []Link{},
		Keywords: []Keyword{
			KeywordFixtures.Get("bridge"),
		},
		Albums: []Album{
			AlbumFixtures.Get("holiday-2030"),
		},
		Files: []File{},
		Labels: []PhotoLabel{
			LabelFixtures.PhotoLabel(1000000, "flower", 38, "image"),
			LabelFixtures.PhotoLabel(1000000, "cake", 38, "manual"),
		},
		CreatedAt: time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:  nil,
		DeletedAt: nil,
	},
	"Photo01": {
		ID:               1000001,
		PhotoUID:         "pt9jtdre2lvl0yh8",
		TakenAt:          time.Date(2006, 1, 1, 2, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Date(2006, 1, 1, 2, 0, 0, 0, time.UTC),
		TakenSrc:         "meta",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoDescription: "photo description blacklist",
		PhotoPath:        "2790/02",
		PhotoName:        "Photo01",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    true,
		PhotoPrivate:     false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		CameraSerial:     "",
		CameraSrc:        "",
		Place:            &UnknownPlace,
		PlaceID:          UnknownPlace.ID,
		Location:         &UnknownLocation,
		LocationID:       UnknownLocation.ID,
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     UnknownPlace.CountryCode(),
		PhotoYear:        2790,
		PhotoMonth:       2,
		Details:          DetailsFixtures.Get("lake", 1000001),
		DescriptionSrc:   "",
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels: []PhotoLabel{
			LabelFixtures.PhotoLabel(1000001, "no-jpeg", 20, "image"),
		},
		CreatedAt: time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:  nil,
		DeletedAt: nil,
	},
	"Photo02": {
		ID:               1000002,
		PhotoUID:         "pt9jtdre2lvl0yh9",
		TakenAt:          time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenSrc:         "meta",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoPath:        "1990/03",
		PhotoName:        "Photo02",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     UnknownPlace.CountryCode(),
		PhotoYear:        1990,
		PhotoMonth:       3,
		Details:          DetailsFixtures.Get("lake", 1000002),
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Location:         &UnknownLocation,
		LocationID:       UnknownLocation.ID,
		Place:            &UnknownPlace,
		PlaceID:          UnknownPlace.ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000002, "cake", 20, "image")},
		CreatedAt:        time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo03": {
		ID:               1000003,
		PhotoUID:         "pt9jtdre2lvl0yh0",
		TakenAt:          time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Time{},
		TakenSrc:         "meta",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoPath:        "1990/04",
		PhotoName:        "Photo03",
		PhotoQuality:     4,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		Place:            LocationFixtures.Pointer("caravan park").Place,
		PlaceID:          LocationFixtures.Pointer("caravan park").Place.ID,
		Location:         LocationFixtures.Pointer("caravan park"),
		LocationID:       LocationFixtures.Pointer("caravan park").ID,
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     LocationFixtures.Pointer("caravan park").Place.CountryCode(),
		PhotoYear:        1990,
		PhotoMonth:       4,
		Details:          DetailsFixtures.Get("bridge", 1000003),
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels: []PhotoLabel{
			LabelFixtures.PhotoLabel(1000003, "cow", 20, "image"),
			LabelFixtures.PhotoLabel(1000003, "updatePhotoLabel", 20, "manual"),
			LabelFixtures.PhotoLabel(1000000, "landscape", 10, "location"),
		},
		CreatedAt: time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:  nil,
		DeletedAt: nil,
	},
	"Photo04": {
		ID:               1000004,
		PhotoUID:         "pt9jtdre2lvl0y11",
		TakenAt:          time.Date(2014, 7, 17, 15, 42, 12, 0, time.UTC),
		TakenAtLocal:     time.Date(2014, 7, 17, 15, 42, 12, 0, time.UTC),
		TakenSrc:         "meta",
		PhotoTitle:       "Neckarbrücke",
		TitleSrc:         "",
		PhotoPath:        "2014/07",
		PhotoName:        "Photo04",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		Place:            PlaceFixtures.Pointer("mexico"),
		PlaceID:          PlaceFixtures.Pointer("mexico").ID,
		Location:         LocationFixtures.Pointer("mexico"),
		LocationID:       LocationFixtures.Pointer("mexico").ID,
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     PlaceFixtures.Pointer("mexico").CountryCode(),
		PhotoYear:        2014,
		PhotoMonth:       7,
		Details:          DetailsFixtures.Get("lake", 1000004),
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Links:            []Link{},
		Keywords: []Keyword{
			KeywordFixtures.Get("bridge"),
		},
		Albums: []Album{
			AlbumFixtures.Get("berlin-2019"),
		},
		Files:     []File{},
		Labels:    []PhotoLabel{LabelFixtures.PhotoLabel(1000004, "batchdelete", 20, "image")},
		CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:  nil,
		DeletedAt: nil,
	},
	"Photo05": {
		ID:               1000005,
		PhotoUID:         "pt9jtdre2lvl0y12",
		TakenAt:          time.Date(2015, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2015, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "meta",
		PhotoTitle:       "Reunion",
		TitleSrc:         "",
		PhotoPath:        "2014/07",
		PhotoName:        "Photo05",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     true,
		PhotoType:        "image",
		PhotoLat:         -21.342636,
		PhotoLng:         55.466944,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "123",
		CameraSrc:        "",
		Location:         LocationFixtures.Pointer("mexico"),
		LocationID:       LocationFixtures.Pointer("mexico").ID,
		LocationSrc:      "",
		Place:            PlaceFixtures.Pointer("mexico"),
		PlaceID:          PlaceFixtures.Pointer("mexico").ID,
		TimeZone:         "",
		PhotoCountry:     PlaceFixtures.Pointer("mexico").CountryCode(),
		PhotoYear:        2014,
		PhotoMonth:       7,
		Details:          DetailsFixtures.Get("lake", 1000005),
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000005, "updateLabel", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo06": {
		ID:               1000006,
		PhotoUID:         "pt9jtdre2lvl0y13",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "meta",
		PhotoTitle:       "ToBeUpdated",
		TitleSrc:         "meta",
		PhotoPath:        "2016/11",
		PhotoName:        "Photo06",
		PhotoQuality:     0,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         -21.342636,
		PhotoLng:         55.466944,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		Location:         &UnknownLocation,
		LocationID:       UnknownLocation.ID,
		LocationSrc:      "",
		Place:            &UnknownPlace,
		PlaceID:          UnknownPlace.ID,
		TimeZone:         "",
		PhotoCountry:     UnknownPlace.CountryCode(),
		PhotoYear:        2014,
		PhotoMonth:       7,
		Details:          DetailsFixtures.Get("lake", 1000006),
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000006, "updatePhotoLabel", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo07": {
		ID:               1000007,
		PhotoUID:         "pt9jtdre2lvl0y14",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "ToBeUpdated",
		TitleSrc:         "meta",
		PhotoPath:        "2016/11",
		PhotoName:        "Photo07",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         -21.342636,
		PhotoLng:         55.466944,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		Location:         &UnknownLocation,
		LocationID:       UnknownLocation.ID,
		LocationSrc:      "",
		Place:            &UnknownPlace,
		PlaceID:          UnknownPlace.ID,
		TimeZone:         "",
		PhotoCountry:     UnknownPlace.CountryCode(),
		PhotoYear:        2014,
		PhotoMonth:       7,
		Details:          DetailsFixtures.Get("lake", 1000007),
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000007, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         &editTime,
		DeletedAt:        nil,
	},
	"Photo08": {
		ID:               1000008,
		PhotoUID:         "pt9jtdre2lvl0y15",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Black beach",
		TitleSrc:         "meta",
		PhotoPath:        "2016/11",
		PhotoName:        "Photo08",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		TimeZone:         "",
		PhotoCountry:     PlaceFixtures.Pointer("mexico").CountryCode(),
		PhotoYear:        2014,
		PhotoMonth:       7,
		Details:          DetailsFixtures.Get("lake", 1000008),
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Location:         LocationFixtures.Pointer("mexico"),
		LocationID:       LocationFixtures.Pointer("mexico").ID,
		LocationSrc:      "manual",
		Place:            PlaceFixtures.Pointer("mexico"),
		PlaceID:          PlaceFixtures.Pointer("mexico").ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000008, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo09": {
		ID:               1000009,
		PhotoUID:         "pt9jtdre2lvl0y16",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Title",
		TitleSrc:         "",
		PhotoPath:        "2016/11",
		PhotoName:        "Photo09",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		TimeZone:         "",
		PhotoCountry:     PlaceFixtures.Pointer("mexico").CountryCode(),
		PhotoYear:        2014,
		PhotoMonth:       7,
		Details:          DetailsFixtures.Get("lake", 1000009),
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Location:         LocationFixtures.Pointer("mexico"),
		LocationID:       LocationFixtures.Pointer("mexico").ID,
		LocationSrc:      "",
		Place:            PlaceFixtures.Pointer("mexico"),
		PlaceID:          PlaceFixtures.Pointer("mexico").ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000008, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo10": {
		ID:               1000010,
		PhotoUID:         "pt9jtdre2lvl0y17",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Title",
		TitleSrc:         "",
		PhotoPath:        "2016/11",
		PhotoName:        "Photo10",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		TimeZone:         "",
		PhotoCountry:     PlaceFixtures.Pointer("holidaypark").CountryCode(),
		PhotoYear:        2014,
		PhotoMonth:       7,
		Details:          DetailsFixtures.Get("lake", 10000010),
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Location:         LocationFixtures.Pointer("hassloch"),
		LocationID:       LocationFixtures.Pointer("hassloch").ID,
		LocationSrc:      "",
		Place:            PlaceFixtures.Pointer("holidaypark"),
		PlaceID:          PlaceFixtures.Pointer("holidaypark").ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000008, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo11": {
		ID:               1000011,
		PhotoUID:         "pt9jtdre2lvl0y18",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Title",
		TitleSrc:         "",
		PhotoPath:        "2016/11",
		PhotoName:        "Photo11",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		TimeZone:         "",
		PhotoCountry:     PlaceFixtures.Pointer("emptyNameLongCity").CountryCode(),
		PhotoYear:        2014,
		PhotoMonth:       7,
		Details:          DetailsFixtures.Get("lake", 10000011),
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Location:         LocationFixtures.Pointer("emptyNameLongCity"),
		LocationID:       LocationFixtures.Pointer("emptyNameLongCity").ID,
		LocationSrc:      "",
		Place:            PlaceFixtures.Pointer("emptyNameLongCity"),
		PlaceID:          PlaceFixtures.Pointer("emptyNameLongCity").ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000008, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo12": {
		ID:               1000012,
		PhotoUID:         "pt9jtdre2lvl0y19",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Title",
		TitleSrc:         "",
		PhotoPath:        "2016/11",
		PhotoName:        "Photo12",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		CameraSerial:     "",
		CameraSrc:        "",
		Location:         LocationFixtures.Pointer("emptyNameShortCity"),
		LocationID:       LocationFixtures.Pointer("emptyNameShortCity").ID,
		LocationSrc:      "",
		Place:            PlaceFixtures.Pointer("emptyNameShortCity"),
		PlaceID:          PlaceFixtures.Pointer("emptyNameShortCity").ID,
		TimeZone:         "",
		PhotoCountry:     PlaceFixtures.Pointer("emptyNameShortCity").CountryCode(),
		PhotoYear:        2014,
		PhotoMonth:       7,
		Details:          Details{},
		DescriptionSrc:   "",
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000008, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo13": {
		ID:               1000013,
		PhotoUID:         "pt9jtdre2lvl0y20",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Title",
		TitleSrc:         "",
		PhotoPath:        "2016/11",
		PhotoName:        "Photo13",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		Location:         LocationFixtures.Pointer("veryLongLocName"),
		LocationID:       LocationFixtures.Pointer("veryLongLocName").ID,
		LocationSrc:      "",
		Place:            PlaceFixtures.Pointer("veryLongLocName"),
		PlaceID:          PlaceFixtures.Pointer("veryLongLocName").ID,
		TimeZone:         "",
		PhotoCountry:     PlaceFixtures.Pointer("veryLongLocName").CountryCode(),
		PhotoYear:        2014,
		PhotoMonth:       7,
		Details:          Details{},
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000008, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo14": {
		ID:               1000014,
		PhotoUID:         "pt9jtdre2lvl0y21",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Title",
		TitleSrc:         "",
		PhotoPath:        "2016/11",
		PhotoName:        "Photo14",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		TimeZone:         "",
		PhotoCountry:     PlaceFixtures.Pointer("mediumLongLocName").CountryCode(),
		PhotoYear:        2014,
		PhotoMonth:       7,
		Details:          Details{},
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Location:         LocationFixtures.Pointer("mediumLongLocName"),
		LocationID:       LocationFixtures.Pointer("mediumLongLocName").ID,
		LocationSrc:      "",
		Place:            PlaceFixtures.Pointer("mediumLongLocName"),
		PlaceID:          PlaceFixtures.Pointer("mediumLongLocName").ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000014, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo15": {
		ID:               1000015,
		PhotoUID:         "pt9jtdre2lvl0y22",
		TakenAt:          time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "location",
		PhotoTitle:       "TitleToBeSet",
		TitleSrc:         "location",
		PhotoDescription: "photo description blacklist",
		PhotoPath:        "1990",
		PhotoName:        "Photo15",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         1.234,
		PhotoLng:         4.321,
		PhotoAltitude:    3,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		Place:            &UnknownPlace,
		Location:         &UnknownLocation,
		PlaceID:          UnknownPlace.ID,
		LocationID:       UnknownLocation.ID,
		LocationSrc:      "location",
		TimeZone:         "",
		PhotoCountry:     UnknownCountry.ID,
		PhotoYear:        0,
		PhotoMonth:       0,
		Details:          DetailsFixtures.Get("blacklist", 1000015),
		DescriptionSrc:   "location",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(10000015, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo16": {
		ID:               1000016,
		PhotoUID:         "pt9jtdre2lvl0y23",
		TakenAt:          time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "ForDeletion",
		TitleSrc:         "",
		PhotoPath:        "1990",
		PhotoName:        "Photo16",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         1.234,
		PhotoLng:         4.321,
		PhotoAltitude:    3,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		CameraSerial:     "",
		CameraSrc:        "",
		Place:            &UnknownPlace,
		Location:         &UnknownLocation,
		PlaceID:          UnknownPlace.ID,
		LocationID:       UnknownLocation.ID,
		LocationSrc:      "location",
		TimeZone:         "",
		PhotoCountry:     UnknownCountry.ID,
		PhotoYear:        0,
		PhotoMonth:       0,
		Details:          DetailsFixtures.Get("lake", 1000015),
		DescriptionSrc:   "location",
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(10000015, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo17": {
		ID:               1000017,
		PhotoUID:         "pt9jtdre2lvl0y24",
		TakenAt:          time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Quality1FavoriteTrue",
		TitleSrc:         "",
		PhotoPath:        "1990/04",
		PhotoName:        "Photo17",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    true,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         1.234,
		PhotoLng:         4.321,
		PhotoAltitude:    3,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Location:         LocationFixtures.Pointer("mexico"),
		LocationID:       LocationFixtures.Pointer("mexico").ID,
		LocationSrc:      "location",
		Place:            PlaceFixtures.Pointer("mexico"),
		PlaceID:          PlaceFixtures.Pointer("mexico").ID,
		TimeZone:         "",
		PhotoCountry:     PlaceFixtures.Pointer("mexico").CountryCode(),
		PhotoYear:        0,
		PhotoMonth:       0,
		Details:          DetailsFixtures.Get("lake", 1000015),
		DescriptionSrc:   "location",
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels: []PhotoLabel{
			LabelFixtures.PhotoLabel(10000015, "landscape", 20, "image"),
			LabelFixtures.PhotoLabel(10000018, "likeLabel", 20, "image")},
		CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:  nil,
		DeletedAt: nil,
	},
	"Photo18": {
		ID:               1000018,
		PhotoUID:         "pt9jtdre2lvl0y25",
		TakenAt:          time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "ArchivedChroma0",
		TitleSrc:         "",
		PhotoPath:        "1990/04",
		PhotoName:        "Photo18",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    true,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         1.234,
		PhotoLng:         4.321,
		PhotoAltitude:    3,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Location:         LocationFixtures.Pointer("mexico"),
		LocationID:       LocationFixtures.Pointer("mexico").ID,
		LocationSrc:      "location",
		Place:            PlaceFixtures.Pointer("mexico"),
		PlaceID:          PlaceFixtures.Pointer("mexico").ID,
		TimeZone:         "",
		PhotoCountry:     PlaceFixtures.Pointer("mexico").CountryCode(),
		PhotoYear:        0,
		PhotoMonth:       0,
		Details:          DetailsFixtures.Get("lake", 1000015),
		DescriptionSrc:   "location",
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels: []PhotoLabel{
			LabelFixtures.PhotoLabel(10000018, "landscape", 20, "image"),
			LabelFixtures.PhotoLabel(10000018, "likeLabel", 20, "image")},
		CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:  nil,
		DeletedAt: &deleteTime,
	},
	"Photo19": {
		ID:               1000019,
		PhotoUID:         "pt9jtxrexxvl0yh0",
		TakenAt:          time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Time{},
		TakenSrc:         "",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoPath:        "1990/04",
		PhotoName:        "Photo03",
		PhotoQuality:     1,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoType:        "image",
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraSerial:     "",
		CameraSrc:        "",
		Place:            &UnknownPlace,
		Location:         &UnknownLocation,
		PlaceID:          UnknownPlace.ID,
		LocationID:       UnknownLocation.ID,
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     UnknownPlace.CountryCode(),
		PhotoYear:        1990,
		PhotoMonth:       4,
		Details:          DetailsFixtures.Get("bridge", 1000019),
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		CameraID:         CameraFixtures.Pointer("canon-eos-6d").ID,
		Lens:             LensFixtures.Pointer("lens-f-380"),
		LensID:           LensFixtures.Pointer("lens-f-380").ID,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{},
		CreatedAt:        time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
}

// CreatePhotoFixtures inserts known entities into the database for testing.
func CreatePhotoFixtures() {
	for _, entity := range PhotoFixtures {
		Db().Create(&entity)
	}
}

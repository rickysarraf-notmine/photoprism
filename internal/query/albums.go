package query

import (
	"fmt"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/internal/search"
)

// Albums returns a slice of albums.
func Albums(offset, limit int) (results entity.Albums, err error) {
	err = UnscopedDb().Table("albums").Select("*").Offset(offset).Limit(limit).Find(&results).Error
	return results, err
}

// AlbumByUID returns a Album based on the UID.
func AlbumByUID(albumUID string) (album entity.Album, err error) {
	if err := Db().Where("album_uid = ?", albumUID).First(&album).Error; err != nil {
		return album, err
	}

	return album, nil
}

// AlbumCoverByUID returns an album cover file based on the uid.
func AlbumCoverByUID(uid string) (file entity.File, err error) {
	a := entity.Album{}

	if err := UnscopedDb().Where("album_uid = ?", uid).First(&a).Error; err != nil {
		return file, err
	} else if a.AlbumType != entity.AlbumDefault { // TODO: Optimize
		f := form.SearchPhotos{Album: a.AlbumUID, Filter: a.AlbumFilter, Order: entity.SortOrderRelevance, Count: 1, Offset: 0, Merged: false}

		if photos, _, err := search.Photos(f); err != nil {
			return file, err
		} else if len(photos) > 0 {
			for _, photo := range photos {
				if err := Db().Where("photo_uid = ? AND file_primary = 1", photo.PhotoUID).First(&file).Error; err != nil {
					return file, err
				} else {
					return file, nil
				}
			}
		}

		return file, fmt.Errorf("no cover found")
	}

	if err := Db().Where("files.file_primary = 1 AND files.file_missing = 0 AND files.file_type = 'jpg' AND files.deleted_at IS NULL").
		Joins("JOIN albums ON albums.album_uid = ?", uid).
		Joins("JOIN photos_albums pa ON pa.album_uid = albums.album_uid AND pa.photo_uid = files.photo_uid AND pa.hidden = 0").
		Joins("JOIN photos ON photos.id = files.photo_id AND photos.photo_private = 0 AND photos.deleted_at IS NULL").
		Order("photos.photo_quality DESC, photos.taken_at DESC").
		First(&file).Error; err != nil {
		return file, err
	}

	return file, nil
}

// UpdateAlbumDates updates album year, month and day based on indexed photo metadata.
func UpdateAlbumDates(mode string) error {
	mutex.Index.Lock()
	defer mutex.Index.Unlock()

	var f string

	switch mode {
	case entity.DateModeFirst:
		f = "MIN(%s)"
		break
	case entity.DateModeLast:
		f = "MAX(%s)"
		break
	case entity.DateModeAverage:
		f = "FROM_UNIXTIME(AVG(UNIX_TIMESTAMP(%s)))"
		break
	default:
		return fmt.Errorf("invalid album date mode %s", mode)
	}

	switch DbDialect() {
	case MySQL:
		return UnscopedDb().Exec(fmt.Sprintf(`UPDATE albums
		INNER JOIN
			(SELECT photo_path, %s AS taken_date
			 FROM photos WHERE taken_src = 'meta' AND photos.photo_quality >= 3 AND photos.deleted_at IS NULL
			 GROUP BY photo_path) AS p ON albums.album_path = p.photo_path
		SET albums.album_year = YEAR(taken_date), albums.album_month = MONTH(taken_date), albums.album_day = DAY(taken_date)
		WHERE albums.album_type = 'folder' AND albums.album_path IS NOT NULL AND p.taken_date IS NOT NULL`, fmt.Sprintf(f, "taken_at_local"))).Error
	default:
		return nil
	}
}

// UpdateMissingAlbumEntries sets a flag for missing photo album entries.
func UpdateMissingAlbumEntries() error {
	mutex.Index.Lock()
	defer mutex.Index.Unlock()

	switch DbDialect() {
	default:
		return UnscopedDb().Exec(`UPDATE photos_albums SET missing = 1 WHERE photo_uid IN
		(SELECT photo_uid FROM photos WHERE deleted_at IS NOT NULL OR photo_quality < 0)`).Error
	}
}

// AlbumEntryFound removes the missing flag from album entries.
func AlbumEntryFound(uid string) error {
	switch DbDialect() {
	default:
		return UnscopedDb().Exec(`UPDATE photos_albums SET missing = 0 WHERE photo_uid = ?`, uid).Error
	}
}

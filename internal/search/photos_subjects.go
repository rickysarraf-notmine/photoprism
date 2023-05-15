package search

import (
	"github.com/photoprism/photoprism/internal/entity"
)

// Subjects finds all subjects associated with the given photo results.
func (photos PhotoResults) Subjects() (results SubjectResults, err error) {
	s := UnscopedDb().Table(entity.Subject{}.TableName()).
		Select("DISTINCT subjects.*").
		Joins("LEFT JOIN markers m on subjects.subj_uid = m.subj_uid").
		Joins("LEFT JOIN files f on m.file_uid = f.file_uid").
		Where("m.marker_type = ? AND f.photo_uid IN (?)", entity.MarkerFace, photos.UIDs())

	if result := s.Scan(&results); result.Error != nil {
		return results, result.Error
	}

	return results, nil
}

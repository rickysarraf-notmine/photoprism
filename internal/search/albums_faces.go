package search

import (
	"github.com/dustin/go-humanize/english"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/pkg/txt"
)

// LoadAlbumsSubjects populates the subjects which have been tagged in photos belonging to the given albums.
// All types of albums are supported - plain ones, folders, moment, state, country, as well as smart albums.
// However be aware that the queries are very suboptimal.
func (albums AlbumResults) LoadAlbumsSubjects(sess *entity.Session) error {
	for i := 0; i < len(albums); i++ {

		f := form.SearchPhotos{
			Album: albums[i].AlbumUID,
			Filter: albums[i].AlbumFilter,
		}

		if err := f.ParseQueryString(); err != nil {
			return err
		}

		photos, count, err := UserPhotos(f, sess)

		if err != nil {
			return err
		}

		if count == 0 {
			return nil
		}

		if subjects, err := photos.Subjects(); err != nil {
			return err
		} else {
			log.Debugf("albums: %s has %s (%s)", albums[i].AlbumTitle, english.Plural(len(subjects), "subject", "subjects"), txt.UniqueNames(subjects.Names()))
			albums[i].Subjects = subjects
		}
	}

	return nil
}

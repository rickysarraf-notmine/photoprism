package query

import (
	"testing"

	"github.com/photoprism/photoprism/pkg/time"
	"github.com/stretchr/testify/assert"
)

func TestAlbumByUID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		if album, err := AlbumByUID("as6sg6bxpogaaba7"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, "Christmas 2030", album.AlbumTitle)
		}

		if cached, err := AlbumByUID("as6sg6bxpogaaba7"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, "Christmas 2030", cached.AlbumTitle)
		}
	})
	t.Run("Missing", func(t *testing.T) {
		album, err := AlbumByUID("as6sg6bxpog00007")
		assert.NotNil(t, album)
		assert.Error(t, err, "record not found")
	})
	t.Run("InvalidUID", func(t *testing.T) {
		album, err := AlbumByUID("3765")
		assert.NotNil(t, album)
		assert.Error(t, err, "invalid album uid")
	})
}

func TestAlbumCoverByUID(t *testing.T) {
	t.Run("existing uid default album", func(t *testing.T) {
		file, err := AlbumCoverByUID("as6sg6bxpogaaba8", true)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "1990/04/bridge2.jpg", file.FileName)
	})

	t.Run("existing uid folder album", func(t *testing.T) {
		file, err := AlbumCoverByUID("as6sg6bipogaaba1", true)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "1990/04/bridge2.jpg", file.FileName)
	})

	t.Run("existing uid empty moment album", func(t *testing.T) {
		file, err := AlbumCoverByUID("as6sg6bitoga0005", true)

		assert.EqualError(t, err, "no cover found", err)
		assert.Equal(t, "", file.FileName)
	})

	t.Run("not existing uid", func(t *testing.T) {
		file, err := AlbumCoverByUID("3765", true)
		assert.Error(t, err, "record not found")
		t.Log(file)
	})

	t.Run("existing uid empty month album", func(t *testing.T) {
		file, err := AlbumCoverByUID("as6sg6bipogaabj9", true)

		assert.EqualError(t, err, "no cover found", err)
		assert.Equal(t, "", file.FileName)
	})
}

func TestUpdateAlbumDates(t *testing.T) {
	t.Run("success with mode first", func(t *testing.T) {
		if err := UpdateAlbumDates(time.DateModeFirst); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("success with mode last", func(t *testing.T) {
		if err := UpdateAlbumDates(time.DateModeLast); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("success with mode average", func(t *testing.T) {
		if err := UpdateAlbumDates(time.DateModeAverage); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("fail with invalid mode", func(t *testing.T) {
		assert.Error(t, UpdateAlbumDates("invalid"))
	})
}

func TestUpdateMissingAlbumEntries(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if err := UpdateMissingAlbumEntries(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestAlbumEntryFound(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if err := AlbumEntryFound("ps6sg6bexxvl0yh0"); err != nil {
			t.Fatal(err)
		}
	})
}

func TestGetAlbums(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r, err := Albums(0, 3)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 3, len(r))
	})
}

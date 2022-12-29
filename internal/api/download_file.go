package api

import (
	"net/http"

	"github.com/photoprism/photoprism/internal/customize"

	"github.com/gin-gonic/gin"

	"github.com/photoprism/photoprism/internal/get"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/query"

	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
)

// TODO: GET /api/v1/dl/file/:hash
// TODO: GET /api/v1/dl/photo/:uid
// TODO: GET /api/v1/dl/album/:uid

// DownloadName returns the download file name type.
func DownloadName(c *gin.Context) customize.DownloadName {
	switch c.Query("name") {
	case "file":
		return customize.DownloadNameFile
	case "share":
		return customize.DownloadNameShare
	case "original":
		return customize.DownloadNameOriginal
	default:
		return get.Config().Settings().Download.Name
	}
}

// GetDownload returns the raw file data.
//
// GET /api/v1/dl/:hash
//
// Parameters:
//
//	hash: string The file hash as returned by the files/photos endpoint
func GetDownload(router *gin.RouterGroup) {
	router.GET("/dl/:hash", func(c *gin.Context) {
		if InvalidDownloadToken(c) {
			c.Data(http.StatusForbidden, "image/svg+xml", brokenIconSvg)
			return
		}

		fileHash := clean.Token(c.Param("hash"))

		f, err := query.FileByHash(fileHash)

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": err.Error()})
			return
		}

		fileName := photoprism.FileName(f.FileRoot, f.FileName)

		if !fs.FileExists(fileName) {
			log.Errorf("download: file %s is missing", clean.Log(f.FileName))
			c.Data(404, "image/svg+xml", brokenIconSvg)

			// Set missing flag so that the file doesn't show up in search results anymore.
			logError("download", f.Update("FileMissing", true))

			return
		}

		c.FileAttachment(fileName, f.DownloadName(DownloadName(c), 0))
	})
}

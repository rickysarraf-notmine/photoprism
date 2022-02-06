package api

import (
	"net/http"

	"github.com/photoprism/photoprism/pkg/sanitize"
	"github.com/photoprism/photoprism/pkg/txt"

	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism/internal/acl"
	"github.com/photoprism/photoprism/internal/face"

	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/query"
	"github.com/photoprism/photoprism/internal/service"
)

// GET /api/v1/photos/:uid/faces
//
// Parameters:
//   uid: string PhotoUID as returned by the API
func GetPhotoFaces(router *gin.RouterGroup) {
	router.GET("/photos/:uid/faces", func(c *gin.Context) {
		s := Auth(SessionID(c), acl.ResourcePhotos, acl.ActionUpdate)

		if s.Invalid() {
			AbortUnauthorized(c)
			return
		}

		f, err := query.FileByPhotoUID(sanitize.IdString(c.Param("uid")))

		if err != nil {
			AbortEntityNotFound(c)
			return
		}

		fileName := photoprism.FileName(f.FileRoot, f.FileName)

		conf := service.Config()
		faces, err := face.DetectAll(fileName, conf.FaceSize())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": txt.UcFirst(err.Error())})
			return
		}

		c.JSON(http.StatusOK, faces)
	})
}

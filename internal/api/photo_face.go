package api

import (
	"net/http"

	"github.com/photoprism/photoprism/pkg/sanitize"
	"github.com/photoprism/photoprism/pkg/txt"

	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism/internal/acl"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
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

		conf := service.Config()
		faces, err := face.DetectAll(photoprism.FileName(f.FileRoot, f.FileName), conf.FaceSize())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": txt.UcFirst(err.Error())})
			return
		}

		c.JSON(http.StatusOK, faces)
	})
}

// POST /api/v1/photos/:uid/faces
//
// Parameters:
//   uid: string PhotoUID as returned by the API
func CreatePhotoFace(router *gin.RouterGroup) {
	router.POST("/photos/:uid/faces", func(c *gin.Context) {
		s := Auth(SessionID(c), acl.ResourcePhotos, acl.ActionUpdate)

		if s.Invalid() {
			AbortUnauthorized(c)
			return
		}

		file, err := query.FileByPhotoUID(sanitize.IdString(c.Param("uid")))

		if err != nil {
			AbortEntityNotFound(c)
			return
		}

		// For convenience sake don't use a form object.
		var f face.Face

		if err := c.BindJSON(&f); err != nil {
			AbortBadRequest(c)
			return
		}

		// Some detections have a very low probability (<1.0), so the resulting face score is 0.
		// This leads to errors down the road when calculating the marker score, so as a workaround
		// we can artificially increase the face score.
		if f.Score == 0 {
			f.Score = 1
		}

		// Calculate the embeddings vector for the given face region.
		net := service.FaceNet()
		embeddings, err := net.Embeddings(photoprism.FileName(file.FileRoot, file.FileName), f)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": txt.UcFirst(err.Error())})
			return
		}

		// Check that we have only one embedding. This duplicates the check in AddFace, but it's important
		// to have it here as well, so that we can report the error to the user.
		if !embeddings.One() {
			log.Errorf("markers: unexpected embeddings count %d for file %s and crop area %s", embeddings.Count(), file.FileUID, f.CropArea())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "region does not contain a face (invalid embeddings)"})
			return
		}

		// Assign the embeddings to the face and add the face to the file, which will create a new marker.
		f.Embeddings = embeddings
		file.AddFace(f, "")

		// It is expected that adding the new face to the file will result in a new unsaved marker.
		if !file.UnsavedMarkers() {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "face marker was not created"})
			return
		}

		// Save the new marker.
		if count, err := file.SaveMarkers(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": txt.UcFirst(err.Error())})
			return
		} else if count == 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "did not save new marker"})
			return
		}

		// Retrieve the newly saved marker.
		marker := file.LatestMarker()

		faces, err := query.Faces(false, false, false)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": txt.UcFirst(err.Error())})
			return
		}

		// Pointer to the matching face.
		var matched_face *entity.Face

		// Distance to the matching face.
		var matched_distance float64

		// Find the closest face match for marker.
		for i, m := range faces {
			if ok, dist := m.Match(marker.Embeddings()); ok && (matched_face == nil || dist < matched_distance) {
				matched_face = &faces[i]
				matched_distance = dist
			}
		}

		// Assign matching face to marker.
		if matched_face != nil {
			_, err := marker.SetFace(matched_face, matched_distance)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": txt.UcFirst(err.Error())})
				return
			}

			// In case the marker could be matched to a subject - use the subject's name as marker name.
			if subj := marker.Subject(); subj != nil {
				marker.MarkerName = subj.SubjName
			}
		}

		PublishPhotoEvent(EntityUpdated, c.Param("uid"), c)

		event.Success("added photo face")

		c.JSON(http.StatusOK, marker)
	})
}

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/photoprism/photoprism/internal/acl"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/internal/get"
	"github.com/photoprism/photoprism/internal/i18n"
	"github.com/photoprism/photoprism/pkg/clean"
)

// Connect confirms external service accounts using a token.
//
// PUT /api/v1/connect/:name
func Connect(router *gin.RouterGroup) {
	router.PUT("/connect/:name", func(c *gin.Context) {
		name := clean.ID(c.Param("name"))

		if name == "" {
			log.Errorf("connect: empty service name")
			AbortBadRequest(c)
			return
		}

		var f form.Connect

		if err := c.BindJSON(&f); err != nil {
			log.Warnf("connect: invalid form values (%s)", clean.Log(name))
			Abort(c, http.StatusBadRequest, i18n.ErrAccountConnect)
			return
		}

		if f.Invalid() {
			log.Warnf("connect: invalid token %s", clean.Log(f.Token))
			Abort(c, http.StatusBadRequest, i18n.ErrAccountConnect)
			return
		}

		conf := get.Config()

		if conf.Public() {
			Abort(c, http.StatusForbidden, i18n.ErrPublic)
			return
		}

		s := Auth(c, acl.ResourceConfig, acl.ActionUpdate)

		if s.Invalid() {
			log.Errorf("connect: %s not authorized", clean.Log(s.User().UserName))
			AbortForbidden(c)
			return
		}

		var err error

		switch name {
		case "hub":
			err = conf.ResyncHub(f.Token)
		default:
			log.Errorf("connect: invalid service %s", clean.Log(name))
			Abort(c, http.StatusBadRequest, i18n.ErrAccountConnect)
			return
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
		}
	})
}

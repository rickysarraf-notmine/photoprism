package api

import (
	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism/internal/acl"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
)

// Auth checks if the user has permission to access the specified resource and returns the session if so.
func Auth(c *gin.Context, resource acl.Resource, grant acl.Permission) *entity.Session {
	return AuthAny(c, resource, acl.Permissions{grant})
}

// AuthAny checks if at least one permission allows access and returns the session in this case.
func AuthAny(c *gin.Context, resource acl.Resource, grants acl.Permissions) (s *entity.Session) {
	// Get client IP address and session ID, if any.
	ip := ClientIP(c)
	sessId := SessionID(c)

	// Find client session.
	if s = Session(sessId); s == nil {
		event.AuditWarn([]string{ip, "unauthenticated", "%s %s as unknown user", "denied"}, grants.String(), string(resource))
		return entity.SessionStatusUnauthorized()
	} else {
		s.SetClientIP(ip)
	}

	// Check authorization.
	if s.User() == nil {
		event.AuditWarn([]string{ip, "session %s", "%s %s as unknown user", "denied"}, s.RefID, grants.String(), string(resource))
		return entity.SessionStatusUnauthorized()
	} else if acl.Resources.DenyAll(resource, s.User().AclRole(), grants) {
		event.AuditErr([]string{ip, "session %s", "%s %s as %s", "denied"}, s.RefID, grants.String(), string(resource), s.User().AclRole().String())
		return entity.SessionStatusForbidden()
	} else {
		event.AuditInfo([]string{ip, "session %s", "%s %s as %s", "granted"}, s.RefID, grants.String(), string(resource), s.User().AclRole().String())
		return s
	}
}

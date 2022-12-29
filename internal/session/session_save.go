package session

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/photoprism/photoprism/internal/entity"
)

// Save updates the client session or creates a new one if needed.
func (s *Session) Save(m *entity.Session) (*entity.Session, error) {
	if m == nil {
		return nil, fmt.Errorf("session is nil")
	}

	// Save session.
	return m.UpdateLastActive(), m.Save()
}

// Create initializes a new client session and returns it.
func (s *Session) Create(u *entity.User, c *gin.Context, data *entity.SessionData) (m *entity.Session, err error) {
	// New session with context, user, and data.
	m = s.New(c).SetUser(u).SetData(data)

	// Create session.
	if err = m.Create(); err != nil {
		m.UpdateLastActive()
	}

	return m, err
}

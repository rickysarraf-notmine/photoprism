package config

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/rnd"
	"github.com/photoprism/photoprism/pkg/txt"
)

const (
	AuthModePublic = "public"
	AuthModePasswd = "password"
)

func isBcrypt(s string) bool {
	b, err := regexp.MatchString(`^\$2[ayb]\$.{56}$`, s)
	if err != nil {
		return false
	}
	return b
}

// SetAuthMode changes the authentication mode (for use in tests only).
func (c *Config) SetAuthMode(mode string) {
	if !c.Debug() {
		return
	}

	switch mode {
	case AuthModePublic:
		c.options.AuthMode = AuthModePublic
		c.options.Public = true
		entity.CheckTokens = false
	default:
		c.options.AuthMode = AuthModePasswd
		c.options.Public = false
		entity.CheckTokens = true
	}
}

// Auth checks if authentication is required.
func (c *Config) Auth() bool {
	return !c.Public()
}

// AuthMode returns the authentication mode.
func (c *Config) AuthMode() string {
	if c.options.Public || c.Demo() {
		return AuthModePublic
	}

	switch c.options.AuthMode {
	case AuthModePublic:
		return AuthModePublic
	default:
		return AuthModePasswd
	}
}

// Public checks if app runs in public mode and requires no authentication.
func (c *Config) Public() bool {
	return c.AuthMode() == AuthModePublic
}

// AdminUser returns the admin auth name.
func (c *Config) AdminUser() string {
	c.options.AdminUser = clean.Username(c.options.AdminUser)

	if c.options.AdminUser == "" {
		c.options.AdminUser = "admin"
	}

	return c.options.AdminUser
}

// AdminPassword returns the initial admin password.
func (c *Config) AdminPassword() string {
	return clean.Password(c.options.AdminPassword)
}

// PasswordLength returns the minimum password length in characters.
func (c *Config) PasswordLength() int {
	if c.Public() {
		return 0
	} else if c.options.PasswordLength < 1 {
		return entity.PasswordLengthDefault
	} else if c.options.PasswordLength > txt.ClipPassword {
		return txt.ClipPassword
	}

	return c.options.PasswordLength
}

// CheckPassword compares given password p with the admin password
func (c *Config) CheckPassword(p string) bool {
	ap := c.AdminPassword()

	if isBcrypt(ap) {
		err := bcrypt.CompareHashAndPassword([]byte(ap), []byte(p))
		return err == nil
	}

	return ap == p
}

// PasswordResetUri returns the password reset page URI, if any.
func (c *Config) PasswordResetUri() string {
	if c.Public() {
		return ""
	}

	return c.options.PasswordResetUri
}

// RegisterUri returns the user registration page URI, if any.
func (c *Config) RegisterUri() string {
	if c.Public() {
		return ""
	}

	return c.options.RegisterUri
}

// LoginUri returns the user authentication page URI.
func (c *Config) LoginUri() string {
	if c.Public() {
		return c.BaseUri("/library/browse")
	}

	if c.options.LoginUri == "" {
		return c.BaseUri("/library/login")
	}

	return c.options.LoginUri
}

// SessionMaxAge returns the time in seconds until API sessions expire automatically.
func (c *Config) SessionMaxAge() int64 {
	if c.options.SessionMaxAge < 0 {
		return 0
	} else if c.options.SessionMaxAge == 0 {
		return DefaultSessionMaxAge
	}

	return c.options.SessionMaxAge
}

// SessionTimeout returns the time in seconds until API sessions expire due to inactivity
func (c *Config) SessionTimeout() int64 {
	if c.options.SessionTimeout < 0 {
		return 0
	} else if c.options.SessionTimeout == 0 {
		return DefaultSessionTimeout
	}

	return c.options.SessionTimeout
}

// DownloadToken returns the DOWNLOAD api token (you can optionally use a static value for permanent caching).
func (c *Config) DownloadToken() string {
	if c.Public() {
		return entity.TokenPublic
	} else if c.options.DownloadToken == "" {
		c.options.DownloadToken = rnd.Base36(8)
	}

	return c.options.DownloadToken
}

// InvalidDownloadToken checks if the token is invalid.
func (c *Config) InvalidDownloadToken(t string) bool {
	return entity.InvalidDownloadToken(t)
}

// PreviewToken returns the preview image api token (based on the unique storage serial by default).
func (c *Config) PreviewToken() string {
	if c.Public() {
		return entity.TokenPublic
	} else if c.options.PreviewToken == "" {
		if c.Serial() == "" {
			return "********"
		} else {
			c.options.PreviewToken = c.SerialChecksum()
		}
	}

	return c.options.PreviewToken
}

// InvalidPreviewToken checks if the preview token is invalid.
func (c *Config) InvalidPreviewToken(t string) bool {
	return entity.InvalidPreviewToken(t)
}

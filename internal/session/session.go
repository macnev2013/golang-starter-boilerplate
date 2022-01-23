package session

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

func New() *scs.SessionManager {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// TODO: add dynamic config here
	session.Cookie.Secure = false

	return session
}

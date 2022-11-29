package session

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieDomain   string
	CookieName     string
	SessionType    string
	CookieSecure   string
}

func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	// how long should sessions last?

	minutes, err := strconv.Atoi(c.CookieLifetime)

	if err != nil {
		minutes = 60
	}

	// should cookies persist
	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	} else {
		persist = false
	}

	//  cookies secure
	if strings.ToLower(c.CookieSecure) == "secure" {
		secure = true
	} else {
		secure = false
	}

	// create session

	session := scs.New()

	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = c.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// which session store

	switch strings.ToLower(c.SessionType) {
	case "redis":

	case "mysql", "mariadb":

	case "postgres", "postgressql":

	default:
		// cookie
	}

	return session
}

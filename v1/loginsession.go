package herohunt

import (
	"time"
)

type LoginSession struct {
	SessionID   string
	PlayerID    string
	Created     *time.Time
	LastUpdated *time.Time
}

var _sessionLifetime time.Duration

func SessionLifeTime() time.Duration {
	if _sessionLifetime == 0 {
		_sessionLifetime = 24 * 60 * time.Minute
	}
	return _sessionLifetime
}

func SetSessionLifeTime(d time.Duration) {
	_sessionLifetime = d
}

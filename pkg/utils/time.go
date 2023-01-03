package utils

import (
	"time"

	cfg "github.com/deall-users/config"
)

func TimeNow() time.Time {
	var loc, _ = time.LoadLocation(cfg.TIME_LOCATION)
	return time.Now().In(loc)
}

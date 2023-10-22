package global

import "time"

const (
	HOST                = "http://111.230.89.67:82"
	TokenExpireDuration = time.Hour * 24 * 30
)

var (
	UserSecret = []byte("254%^FuCo610N!3N")
	Token      = "Bearer "
)

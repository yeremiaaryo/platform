package entity

const (
	ConfigSMTPHost = "smtp.gmail.com"
	ConfigSMTPPort = 587
	ConfigEmail    = "cobayeremia@gmail.com"
	ConfigPassword = "$Tokopedia789"
)

const (
	CookieName = `_SID_HobbyLobby_`
)

const (
	RedisKeyLogin = `Login_%v`
)

const (
	CookieExpireInDays   = 30
	LoginExpireInSeconds = CookieExpireInDays * 24 * 60 * 60
)

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
	RedisKeyLogin               = `Login_%v`
	RedisKeyForgotPasswordToken = `FP_%v`
	RedisKeyVerifyEmail         = `Verify_%v`
)

const (
	CookieExpireInDays          = 30
	LoginExpireInSeconds        = CookieExpireInDays * 24 * 60 * 60
	OTPExpiredInSeconds         = 3 * 60 * 60
	VerifyEmailExpiredInSeconds = 3600

	OTPLength = 6
)

const (
	ContextUserID = `UserID`
	ContextEmail  = `Email`
)

const (
	JWTSecret = "cuvoqhztptiz"
)

const (
	CloudinaryBaseURL   = "https://api.cloudinary.com/v1_1/hobbylobby/image/upload"
	CloudinaryAPIKey    = "882192587435456"
	CloudinaryAPISecret = "EDl6YxD2EM5RDKuJQkpA86_TsBM"
)

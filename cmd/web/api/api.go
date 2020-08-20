package api

import (
	"github.com/yeremiaaryo/go-pkg/router"
	"github.com/yeremiaaryo/platform/cmd/internal"
	"github.com/yeremiaaryo/platform/internal/auth"
	"github.com/yeremiaaryo/platform/internal/usecase"
)

type Options struct {
	Prefix         string
	DefaultTimeout int
	Usecase        *internal.Usecase
}

type API struct {
	options *Options
	userUC  usecase.UserUsecase
	authUC  auth.AuthUsecase
	shopUC  usecase.ShopUsecase
}

func New(o *Options) *API {
	return &API{
		options: o,
		userUC:  o.Usecase.User,
		authUC:  o.Usecase.Auth,
		shopUC:  o.Usecase.Shop,
	}
}

func (a *API) Register() {
	r := router.New(&router.Options{Prefix: a.options.Prefix, Timeout: a.options.DefaultTimeout})
	r.POST("/register", a.RegisterUser)
	r.POST("/login", a.ValidateLogin)
	r.POST("/refresh", a.authUC.Authorize(a.ValidateCookie))
	r.POST("/forgot-password/generate-token", a.authUC.Authorize(a.GenerateForgotPasswordToken))
	r.POST("/forgot-password/validate-token", a.authUC.Authorize(a.ValidateForgotPasswordToken))
	r.POST("/forgot-password/reset", a.authUC.Authorize(a.ResetPassword))
	r.GET("/verify_account", a.VerifyEmail)
	r.GET("/is_verified", a.authUC.Authorize(a.IsVerified))
	r.POST("/resend_verification", a.authUC.Authorize(a.ResendVerification))

	r.GET("/shop/info", a.authUC.Authorize(a.GetShopInfo))
	r.POST("/shop/update", a.authUC.Authorize(a.InsertUpdateShopData))
	r.POST("/shop/upload", a.authUC.Authorize(a.UploadImage))

	r.GET("/inspiration/list", a.authUC.Authorize(a.GetInspirationList))
}

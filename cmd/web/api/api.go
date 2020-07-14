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
}

func New(o *Options) *API {
	return &API{
		options: o,
		userUC:  o.Usecase.User,
		authUC:  o.Usecase.Auth,
	}
}

func (a *API) Register() {
	r := router.New(&router.Options{Prefix: a.options.Prefix, Timeout: a.options.DefaultTimeout})
	r.POST("/register", a.RegisterUser)
	r.POST("/login", a.ValidateLogin)
	r.POST("/refresh", a.authUC.Authorize(a.ValidateCookie))
}

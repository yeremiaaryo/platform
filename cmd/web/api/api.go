package api

import (
	"github.com/yeremiaaryo/go-pkg/router"
	"github.com/yeremiaaryo/platform/cmd/internal"
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
}

func New(o *Options) *API {
	return &API{
		options: o,
		userUC:  o.Usecase.User,
	}
}

func (a *API) Register() {
	r := router.New(&router.Options{Prefix: a.options.Prefix, Timeout: a.options.DefaultTimeout})
	r.GET("/hello", a.GetUserName)
	r.POST("/register", a.GetUserName)
}

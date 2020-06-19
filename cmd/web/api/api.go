package api

import (
	"net/http"

	"github.com/yeremiaaryo/go-pkg/response"
	"github.com/yeremiaaryo/go-pkg/router"
)

type Options struct {
	Prefix         string
	DefaultTimeout int
}

type API struct {
	options *Options
}

func New(o *Options) *API {
	return &API{
		options: o,
	}
}

func (a *API) Register() {
	r := router.New(&router.Options{Prefix: a.options.Prefix, Timeout: a.options.DefaultTimeout})
	r.GET("/hello", a.HelloWorld)
}

func (a *API) HelloWorld(r *http.Request) *response.JSONResponse {
	test := struct {
		Message string `json:"message"`
	}{
		Message: "Hello world!",
	}

	return response.NewJSONResponse().SetData(test)
}

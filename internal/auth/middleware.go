package auth

import (
	"net/http"

	"github.com/yeremiaaryo/go-pkg/response"
	"github.com/yeremiaaryo/go-pkg/router"
)

func (au *authUC) Authorize(h router.Handle) router.Handle {
	return func(r *http.Request) *response.JSONResponse {
		return h(r)
	}
}

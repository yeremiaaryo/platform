package auth

import "github.com/yeremiaaryo/go-pkg/router"

type AuthUsecase interface {
	Authorize(h router.Handle) router.Handle
}

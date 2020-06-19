package web

import (
	"log"
	"net"
	"net/http"
	"platform/cmd/web/api"
	"time"

	"github.com/yeremiaaryo/go-pkg/router"
	"gopkg.in/tylerb/graceful.v1"
)

type Options struct {
	ListenAddress string
}

type Handler struct {
	options     *Options
	listenErrCh chan error
}

func New(o *Options) *Handler {
	handler := &Handler{options: o}

	api.New(&api.Options{
		Prefix:         "/api/v1",
		DefaultTimeout: 15,
	}).Register()

	return handler
}

func (h *Handler) Run() {
	log.Println("API Listening on", h.options.ListenAddress)
	h.listenErrCh <- serve(h.options.ListenAddress)
}

func (h *Handler) ListenError() <-chan error {
	return h.listenErrCh
}

func gracefulServe(listenAddress string) (net.Listener, error) {
	l, err := net.Listen("tcp4", listenAddress)
	if err != nil {
		log.Fatal("Cannot connect to listen address")
	}
	return l, nil
}

func serve(listenAddress string) error {
	srv := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Handler:      router.WrapperHandler(),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
	log.Println("starting serve on", listenAddress)
	l, _ := gracefulServe(listenAddress)
	return srv.Serve(l)
}

package main

import (
	"context"
	"net/http"
	"net/url"

	"github.com/otakakot/giveitashotgo/internal/ogen302/gen/api"
)

var _ api.Handler = (*handler)(nil)

type handler struct{}

// A implements api.Handler.
func (*handler) A(ctx context.Context) (api.ARes, error) {
	res := &api.AFound{}

	uri, _ := url.Parse("http://localhost:8080/b")

	res.SetLocation(api.NewOptURI(*uri))

	return res, nil
}

// B implements api.Handler.
func (*handler) B(ctx context.Context) (api.BRes, error) {
	return &api.BOK{}, nil
}

func main() {
	hdl, err := api.NewServer(&handler{})
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: hdl,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

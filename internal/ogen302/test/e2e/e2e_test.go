package e2e_test

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/otakakot/giveitashotgo/internal/ogen302/gen/api"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:8080"
	}

	// cli, err := api.NewClient(endpoint) // このままだと net/http.Client の実装によって勝手にリダイレクトされてしまう
	cli, err := api.NewClient(endpoint, api.WithClient(&http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("health", func(t *testing.T) {
		t.Parallel()

		res, err := cli.A(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		switch res := res.(type) {
		case *api.AFound:
			expected, _ := url.Parse("http://localhost:8080/b")
			if res.Location.Value != *expected {
				t.Error("unexpected location")
			}
		default:
			t.Fatalf("unexpected response: %#v", res)
		}
	})
}

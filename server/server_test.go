package server_test

import (
	"github.com/bastian-kurz/basic-rest-example/integrationtest"
	is2 "github.com/matryer/is"
	"net/http"
	"testing"
)

func TestServer_Start(t *testing.T) {
	integrationtest.SkipIfShort(t)

	t.Run("starts the server and listens for request", func(t *testing.T) {
		is := is2.New(t)

		cleanup := integrationtest.CreateServer()
		defer cleanup()

		resp, err := http.Get("http://localhost:8085")
		is.NoErr(err)
		is.Equal(http.StatusNotFound, resp.StatusCode)
	})
}

package api_test

import (
	"github.com/bastian-kurz/basic-rest-example/integrationtest"
	is2 "github.com/matryer/is"
	"net/http"
	"testing"
)

func TestHealth(t *testing.T) {
	integrationtest.SkipIfShort(t)

	t.Run("stars the server and test the endpoint", func(t *testing.T) {
		is := is2.New(t)

		cleanup := integrationtest.CreateServer()
		defer cleanup()

		resp, err := http.Get("http://localhost:8085/health")
		is.NoErr(err)
		is.Equal(http.StatusOK, resp.StatusCode)
	})
}

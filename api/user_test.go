package api_test

import (
	"github.com/bastian-kurz/basic-rest-example/integrationtest"
	is2 "github.com/matryer/is"
	"net/http"
	"strings"
	"testing"
)

func TestUser(t *testing.T) {
	integrationtest.SkipIfShort(t)
	cleanup := integrationtest.CreateServer()
	defer cleanup()
	is := is2.New(t)

	t.Run("test valid post request", func(t *testing.T) {

		body := "{\"userName\": \"doe\",  \"email\": \"john.doe@testsubject.de\",  \"firstName\": \"John\",  \"lastName\": \"Doe\",  \"password\":\"foobar\"}"
		r := strings.NewReader(body)
		res, err := http.Post("http://localhost:8085/api/user", "application/json", r)
		is.NoErr(err)
		is.Equal(http.StatusCreated, res.StatusCode)
	})

	t.Run("test invalid post request", func(t *testing.T) {
		body := "{\"userName\": \"!!\",  \"email\": \"john.doe@@testsubject.de\",  \"firstName\": \"John\",  \"lastName\": \"Doe\",  \"password\":\"foobar\"}"
		r := strings.NewReader(body)
		res, err := http.Post("http://localhost:8085/api/user", "application/json", r)
		is.NoErr(err)
		is.Equal(http.StatusBadRequest, res.StatusCode)
	})

	t.Run("test get list request", func(t *testing.T) {
		res, err := http.Get("http://localhost:8085/api/user")
		is.NoErr(err)
		is.Equal(http.StatusOK, res.StatusCode)
	})

	t.Run("test get list request", func(t *testing.T) {
		res, err := http.Get("http://localhost:8085/api/user/10")
		is.NoErr(err)
		is.Equal(http.StatusOK, res.StatusCode)
	})
}

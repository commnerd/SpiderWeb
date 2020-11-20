package api

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/http"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HomeHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect.

	expected := `{"version":"0.0.0","address":"322a1963-2b7f-43d4-b9cf-2fcea27c63da","ip":"192.168.10.10","port":42}`
	assert.Equal(t, expected, rr.Body.String());
}

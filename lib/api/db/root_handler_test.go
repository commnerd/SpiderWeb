package db

import (
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"net/http/httptest"
	"net/http"
	"testing"
	"strings"
	"../util"
	"../id"
	"fmt"
)

func TestRegisterHandlerGET(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    req, err := http.NewRequest(http.MethodGet, "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(RegisterHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)

	// Check the response body is what we expect.
	expected := "405 - GET method not allowed"
	assert.Equal(t, expected, rr.Body.String());
}

func TestRegisterHandlerOPTIONS(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    req, err := http.NewRequest(http.MethodOptions, "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(RegisterHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)


	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check for the appropriate headers
	accessControlMethods := rr.HeaderMap["Access-Control-Allow-Methods"]
	for _, val := range []string{http.MethodOptions} {
		valPresent, _ := util.InArray(val, accessControlMethods)
		assert.True(t, valPresent)
	}

	// Check the response body is what we expect.
	assert.Equal(t, "", rr.Body.String());
}

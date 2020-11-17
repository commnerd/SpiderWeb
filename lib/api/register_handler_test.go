package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"net/http/httptest"
	"net/http"
	"testing"
	"../util"
	"../id"
	"fmt"
)

var mockRequestBody string = fmt.Sprintf(`{id:"%v",type:"%v"}`,
	id.Id(uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da")).String(),
	"new",
)

var mockResponse = RegisterResponse{
	Status: Success,
	AdjustedId: id.Id(uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fceabcd1234")).String(),
	Mask: 25,
}

// // TODO: SHOULD FAIL
// func TestRegisterHandlerGET(t *testing.T) {
// 	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
//     // pass 'nil' as the third parameter.
//     req, err := http.NewRequest(http.MethodGet, "/register", nil)
//     if err != nil {
//         t.Fatal(err)
//     }

//     // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
//     rr := httptest.NewRecorder()
//     handler := http.HandlerFunc(RegisterHandler)

//     // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
//     // directly and pass in our Request and ResponseRecorder.
//     handler.ServeHTTP(rr, req)

// 	// Check the status code is what we expect.
// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	// Check the response body is what we expect.

// 	expected := `{"version":"0.0.0","address":"322a1963-2b7f-43d4-b9cf-2fcea27c63da","ip":"192.168.10.10","port":42}`
// 	assert.Equal(t, expected, rr.Body.String());
// }

// func TestRegisterHandlerPOST(t *testing.T) {
// 	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
//     // pass 'nil' as the third parameter.
//     req, err := http.NewRequest(http.MethodPost, "/register", mockRequestBody)
//     if err != nil {
//         t.Fatal(err)
//     }

//     // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
//     rr := httptest.NewRecorder()
//     handler := http.HandlerFunc(RegisterHandler)

//     // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
//     // directly and pass in our Request and ResponseRecorder.
//     handler.ServeHTTP(rr, req)

// 	// Check the status code is what we expect.
// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	// Check the response body is what we expect.

// 	expected := `{"version":"0.0.0","address":"322a1963-2b7f-43d4-b9cf-2fcea27c63da","ip":"192.168.10.10","port":42}`
// 	assert.Equal(t, expected, rr.Body.String());
// }

// func TestRegisterHandlerPUT(t *testing.T) {
// 	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
//     // pass 'nil' as the third parameter.
//     req, err := http.NewRequest(http.MethodPost, "/register", mockRequestBody)
//     if err != nil {
//         t.Fatal(err)
//     }

//     // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
//     rr := httptest.NewRecorder()
//     handler := http.HandlerFunc(RegisterHandler)

//     // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
//     // directly and pass in our Request and ResponseRecorder.
//     handler.ServeHTTP(rr, req)

// 	// Check the status code is what we expect.
// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	// Check the response body is what we expect.

// 	expected := `{"version":"0.0.0","address":"322a1963-2b7f-43d4-b9cf-2fcea27c63da","ip":"192.168.10.10","port":42}`
// 	assert.Equal(t, expected, rr.Body.String());
// }

func TestRegisterHandlerOPTIONS(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    req, err := http.NewRequest(http.MethodOptions, "/register", nil)
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
	for _, val := range []string{http.MethodOptions, http.MethodPost, http.MethodPut} {
		valPresent, _ := util.InArray(val, accessControlMethods)
		assert.True(t, valPresent)
	}

	// Check the response body is what we expect.
	assert.Equal(t, "", rr.Body.String());
}

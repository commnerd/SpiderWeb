package api

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

var mockRequestBody string = fmt.Sprintf(`{id:"%v"}`,
	id.Id(uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da")).String(),
)

type mockNodeStruct struct{}

func (mns *mockNodeStruct) RegisterChild(child interface{}) (interface{}, error) {
	return RegisterResponse{
		Status: Success,
		Version: "0.0.0",
		AdjustedId: id.Id(uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea1234567")).String(),
		Ip: "192.168.10.10",
		Port: 42,
		Mask: 25,
		PublicRsa: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCsdAZxwU/G50zNoC0C73y+E1o7D16cXk" +
		"YYTG4PyODvihOYyig/6qkZWyDwg5bfQ3J01IBt8DejtHhz+K99gHGor/P+OMQ6p3FvurRci9oJWLfSdv4" +
		"IdQuEOy1ntkhiVm2TT8fhFzPBCwwBWO+9aY+Fgt6vLp9io8EkbBDCGhOHptfHo+INbOc+0fEZ8EkzLDtJ" +
		"iTMgUJDKF7AzXLpK/77BNK59r6v/Iiuq0cC3qL88wEhGV3NiCeJu8oThPkeDEvxBQyjeDDa2SZXBNCKIL" +
		"F2wrKNsQqigk+qZHyWQSqfPcFdmCW5aN4p/sU96W5hmo5NjealrS8578zmCNNMfGj6yY/o1unZ2/371iz" +
		"18AJT4vbhLQH0Mkoh7qK231ZWyFtpsah+yWd0VJl0PbcTMu134V3ON5pzDvMafGZQ5OZfUF0njNffSJ6I" +
		"82N2qfgNcsgGBCF7D1Sn0dqxU4u/5bMtcdc+uA4GiwLvQuMna6xR5/NguRlSiYGDWHYiWCgfEsCFYsGDN" +
		"g3Hq4mW2r5hURSyx57tAn9Zvq4FvqcxlnybSvyYThtojwkXoP/c5hTAGKahQyEXB6JsEv1chgOjGeawDF" +
		"bZmFXs/fUqNiZXz5YmcrlDzkfYZjSiufaFrnNPbfk3FtF5LrrohnuOhnH7zSvzwJ7C6RC4XN2ngYKIl+f" +
		"Ynrw==",
	}, nil
}

func TestRegisterHandlerGET(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    req, err := http.NewRequest(http.MethodGet, "/register", nil)
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

func TestRegisterHandlerPOST(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
	req, err := http.NewRequest(http.MethodPost, "/register", strings.NewReader(mockRequestBody))
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

	// Check the response body is what we expect.

	expected := `{"status":0,"version":"0.0.0","address":"322a1963-2b7f-43d4-b9cf-2fcea1234567",` +
	`"mask":25,"ip":"192.168.10.10","port":42,"pub_rsa":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC` +
	`sdAZxwU/G50zNoC0C73y+E1o7D16cXkYYTG4PyODvihOYyig/6qkZWyDwg5bfQ3J01IBt8DejtHhz+K99gHGor/P+OM` +
	`Q6p3FvurRci9oJWLfSdv4IdQuEOy1ntkhiVm2TT8fhFzPBCwwBWO+9aY+Fgt6vLp9io8EkbBDCGhOHptfHo+INbOc+0` +
	`fEZ8EkzLDtJiTMgUJDKF7AzXLpK/77BNK59r6v/Iiuq0cC3qL88wEhGV3NiCeJu8oThPkeDEvxBQyjeDDa2SZXBNCKI` +
	`LF2wrKNsQqigk+qZHyWQSqfPcFdmCW5aN4p/sU96W5hmo5NjealrS8578zmCNNMfGj6yY/o1unZ2/371iz18AJT4vbh` +
	`LQH0Mkoh7qK231ZWyFtpsah+yWd0VJl0PbcTMu134V3ON5pzDvMafGZQ5OZfUF0njNffSJ6I82N2qfgNcsgGBCF7D1S` +
	`n0dqxU4u/5bMtcdc+uA4GiwLvQuMna6xR5/NguRlSiYGDWHYiWCgfEsCFYsGDNg3Hq4mW2r5hURSyx57tAn9Zvq4Fvq` +
	`cxlnybSvyYThtojwkXoP/c5hTAGKahQyEXB6JsEv1chgOjGeawDFbZmFXs/fUqNiZXz5YmcrlDzkfYZjSiufaFrnNPb` +
	`fk3FtF5LrrohnuOhnH7zSvzwJ7C6RC4XN2ngYKIl+fYnrw=="}`
	assert.Equal(t, expected, rr.Body.String());
}

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
	for _, val := range []string{http.MethodOptions, http.MethodPost} {
		valPresent, _ := util.InArray(val, accessControlMethods)
		assert.True(t, valPresent)
	}

	// Check the response body is what we expect.
	assert.Equal(t, "", rr.Body.String());
}

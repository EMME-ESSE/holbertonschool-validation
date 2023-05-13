package main

import (
  "net/http"
  "net/http/httptest"
  "net/url"
  "testing"
)

func TestIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("Flag `-short` absent: skipping Integration Tests.")
    }

    // create new http server with router
    server := httptest.NewServer(setupRouter())
    defer server.Close()

    tests := []struct {
        name         string
        queryString  string
        responseCode int
        body         string
    }{
        {
            name:         "Grace Hopper",
            queryString:  "name=Grace Hopper",
            responseCode: 200,
            body:         "Hello Grace Hopper!",
        },
        {
            name:         "Rosalind Franklin",
            queryString:  "name=Rosalind Franklin",
            responseCode: 200,
            body:         "Hello Rosalind Franklin!",
        },
        {
            name:         "No name provided",
            queryString:  "",
            responseCode: 400,
            body:         "Bad Request: Missing 'name' parameter",
        },
        // Add more test cases here
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Build an HTTP request for the current test `tt`
            req, err := http.NewRequest("GET", "/hello", nil)
            if err != nil {
                t.Fatal(err)
            }

            queryParams, _ := url.ParseQuery(tt.queryString)
            req.URL.RawQuery = queryParams.Encode()

            // Send HTTP request to the test server
            res, err := http.DefaultClient.Do(req)
            if err != nil {
                t.Fatal(err)
            }

            // Check that the status code is what you expect.
            expectedCode := tt.responseCode
            if res.StatusCode != expectedCode {
                t.Errorf("handler returned wrong status code: got %v want %v", res.StatusCode, expectedCode)
            }

            // Check that the response body is what you expect.
            expectedBody := tt.body
            buf := new(bytes.Buffer)
            _, err = buf.ReadFrom(res.Body)
            if err != nil {
                t.Fatal(err)
            }
            gotBody := buf.String()
            if gotBody != expectedBody {
                t.Errorf("handler returned unexpected body: got %v want %v", gotBody, expectedBody)
            }
        })
    }
}

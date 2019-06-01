package main

import (
	"testing"
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)


func TestMainExecution(t *testing.T) {
    go main()
    resp, err := http.Get("http://127.0.0.1:1982/ping")
    if err != nil {
            t.Fatalf("Cannot make get: %v\n", err)
    }
    bodySb, err := ioutil.ReadAll(resp.Body)
    if err != nil {
            t.Fatalf("Error reading body: %v\n", err)
    }
    body := string(bodySb)
    fmt.Printf("Body: %v\n", body)
    var decodedResponse interface{}
    err = json.Unmarshal(bodySb, &decodedResponse)
    if err != nil {
            t.Fatalf("Cannot decode response <%p> from server. Err: %v", bodySb, err)
    }
    assert.Equal(t, map[string]interface{}{"status": "pong"}, decodedResponse,
            "Should return status:ok")
}
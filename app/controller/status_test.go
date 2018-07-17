package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

// TestStatus should pass but this test does not test routing
func TestStatus(t *testing.T) {
	handler := Status

	req := httptest.NewRequest("GET", "/status", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Status Code should be 200, was %v", resp.StatusCode)
		t.FailNow()
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Content type should be \"application/json\", was %s", resp.Header.Get("Content-Type"))
	}
	var StatusResp statusResponse
	err := json.Unmarshal(body, &StatusResp)
	if err != nil {
		t.Error(err)
	}
	if StatusResp.Status != "OK" {
		t.Errorf("Status should be \"OK\", was %s", StatusResp.Status)
	}
}

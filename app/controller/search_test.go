package controller

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dwin/ohioGeranium/app/model"
)

func TestNewSearch(t *testing.T) {
	// Test Expected Working Query
	reqBody, err := json.Marshal(&model.SearchFields{
		FirstName: "John",
		LastName:  "Brown",
		Limit:     10,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	handler := NewSearch
	req := httptest.NewRequest("POST", "/search", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("Status Code should be 200, was %v", resp.StatusCode)
		t.FailNow()
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Content type should be \"application/json\", was %s", resp.Header.Get("Content-Type"))
	}

	var result searchResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Error("Unable to decode json response error")
		t.FailNow()
	}
	if result.DataResults.ResultCount < 1 {
		t.Error("Query should return results")
	}

	// Test Error Query
	reqBody, err = json.Marshal(&model.SearchFields{
		EnumerationType: "ABC",
		FirstName:       "John",
		LastName:        "Brown",
		Limit:           10,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	handler = NewSearch
	req = httptest.NewRequest("POST", "/search", bytes.NewBuffer(reqBody))
	w = httptest.NewRecorder()
	handler(w, req)

	resp = w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("Status Code should be 200, was %v", resp.StatusCode)
		t.FailNow()
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Content type should be \"application/json\", was %s", resp.Header.Get("Content-Type"))
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Error("Unable to decode json response error")
		t.FailNow()
	}
	if !strings.Contains(result.Message, "Error") {
		t.Error("Request shoud return errors in message")
	}
}

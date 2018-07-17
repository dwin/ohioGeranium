package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type statusResponse struct {
	Status string
}

func Status(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(&statusResponse{
		Status: "OK",
	})
	if err != nil {
		log.Printf("Unable to marshal status response error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Unable to write status response error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}

package model

import (
	"testing"
)

func TestOpenDB(t *testing.T) {
	db := openDB()
	defer db.Close()

	if db.Stat().AvailableConnections < 1 {
		t.Error("DB Connection Error")
	}

}

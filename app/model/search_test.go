package model

import "testing"

func TestCreateSearchRecord(t *testing.T) {
	s := SearchFields{
		FirstName:  "John",
		LastName:   "Doe",
		City:       "Los Angeles",
		State:      "California",
		PostalCode: "90014",
	}

	if err := CreateSearchRecord(s); err != nil {
		t.Logf("Unable to create search record in db, error: %s", err)
		t.FailNow()
	}
}

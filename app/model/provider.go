package model

import (
	"encoding/json"
	"time"
)

type NppesResponse struct {
	ResultCount int              `json:"result_count"`
	Results     []ProviderResult `json:"results"`
}

type ProviderResult struct {
	Number    int `json:"number"`
	Addresses []struct {
		Address1        string `json:"address_1"`
		Address2        string `json:"address_2"`
		AddressPurpose  string `json:"address_purpose"`
		AddressType     string `json:"address_type"`
		City            string `json:"city"`
		CountryCode     string `json:"country_code"`
		CountryName     string `json:"country_name"`
		FaxNumber       string `json:"fax_number"`
		PostalCode      string `json:"postal_code"`
		State           string `json:"state"`
		TelephoneNumber string `json:"telephone_number"`
	} `json:"addresses"`
	Basic struct {
		Credential      string `json:"credential"`
		EnumerationDate string `json:"enumeration_date"`
		FirstName       string `json:"first_name"`
		Gender          string `json:"gender"`
		LastName        string `json:"last_name"`
		LastUpdated     string `json:"last_updated"`
		MiddleName      string `json:"middle_name"`
		Name            string `json:"name"`
		NamePrefix      string `json:"name_prefix"`
		SoleProprietor  string `json:"sole_proprietor"`
		Status          string `json:"status"`
	} `json:"basic"`
	CreatedEpoch    int    `json:"created_epoch"`
	EnumerationType string `json:"enumeration_type"`
	Identifiers     []struct {
		Code       string `json:"code"`
		Desc       string `json:"desc"`
		Identifier string `json:"identifier"`
		Issuer     string `json:"issuer"`
		State      string `json:"state"`
	} `json:"identifiers"`
	LastUpdatedEpoch int `json:"last_updated_epoch"`

	OtherNames []interface{} `json:"other_names"`
	Taxonomies []struct {
		Code    string `json:"code"`
		Desc    string `json:"desc"`
		License string `json:"license"`
		Primary bool   `json:"primary"`
		State   string `json:"state"`
	} `json:"taxonomies"`
}

type ProviderRecord struct {
	ProviderNumber int
	Result         json.RawMessage
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

type ResultMap map[string]interface{}

func CreateProviders(resp NppesResponse) error {
	// Handle as transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback()

	for _, p := range resp.Results {
		_, err = tx.Exec(`INSERT INTO providers(provider_number, data, created_at, updated_at) VALUES($1,$2, 'now()', now()) ON CONFLICT (provider_number) DO NOTHING`, p.Number, p)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

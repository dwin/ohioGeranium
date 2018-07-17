package model

// SearchFields are query fields for NPPES API
type SearchFields struct {
	Number           int    `json:"number"`
	EnumerationType  string `json:"enumerationType"`
	TaxonomyDesc     string `json:"taxonomyDesc"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName "`
	OrganizationName string `json:"organizationName"`
	AddressPurpose   string `json:"addressPurpose"`
	City             string `json:"city"`
	State            string `json:"state"`
	PostalCode       string `json:"postalCode"`
	CountryCode      string `json:"countryCode"`
	Limit            int    `json:"limit"`
	Skip             int    `json:"skip"`
	Pretty           bool   `json:"pretty"`
}

// CreateSearchRecord stores NewSearch query to database
func CreateSearchRecord(s SearchFields) error {
	_, err := db.Query(`INSERT INTO queries(data, created_at, updated_at) VALUES($1,'now()',now())`, s)
	return err
}

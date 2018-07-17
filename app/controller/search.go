package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dwin/ohioGeranium/app/model"
)

const nppesEndpoint = "https://npiregistry.cms.hhs.gov/api"

type searchResponse struct {
	StatusCode  int                 `json:"status_code"`
	DataResults model.NppesResponse `json:"data.results"`
	Message     string              `json:"message"`
}

// NewSearch takes model.SearchFields as json and returns searchResponse as JSON
// from nppesEndpoint using request json input as parameters
func NewSearch(w http.ResponseWriter, r *http.Request) {
	// Process and validate request
	var reqSearch model.SearchFields
	err := json.NewDecoder(r.Body).Decode(&reqSearch)
	if err != nil {
		log.Printf("Unable to unmarshal json request body for NewSearch error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Store Query
	if err := model.CreateSearchRecord(reqSearch); err != nil {
		log.Printf("Unable to store new search query record to db error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Build Querystring as defined by ("https://npiregistry.cms.hhs.gov/registry/help-api")
	reqQueries := buildQuery(reqSearch)

	// Create request to NPPES API endpoint
	client := &http.Client{}
	req, err := http.NewRequest("GET", nppesEndpoint+reqQueries, nil)
	if err != nil {
		log.Printf("Unable to create new request to NPPES Endpoint error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Could not complete API request, error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check API Response OK
	if resp.Status != "200 OK" {
		apiRespBody, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		log.Printf("Unexpected API response. API Resp Body: %s", apiRespBody)

		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(&searchResponse{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("%s", apiRespBody),
		})
		if err != nil {
			log.Printf("Unable to marshal error response error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}

	// Process API response
	var results model.NppesResponse

	apiRespBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// Check API Response Type
	var apiRespJSON map[string]interface{}
	err = json.Unmarshal(apiRespBody, &apiRespJSON)
	if err != nil {
		log.Println("Could not unmarshal API response body to apiRespJSON")
		return
	}

	// Handle Error Response from API is exists and return errors as message to client
	if apiErrors, exists := apiRespJSON["Errors"]; exists {
		var msg string
		for k, v := range apiErrors.(map[string]interface{}) {
			msg += fmt.Sprintf("%s: %s,", k, v)
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&searchResponse{
			StatusCode: 200,
			Message:    msg,
		})
		if err != nil {
			log.Printf("Unable to marshal search error response error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	// Handle typical response from API query
	if err := json.Unmarshal(apiRespBody, &results); err != nil {
		apiRespBody, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		log.Printf("Unexpected API response unable to unmarshal results. API Resp Body: %s", apiRespBody)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Persist results to db
	if err := model.CreateProviders(results); err != nil {
		log.Printf("Unable to persist provider results to db error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Encode and send JSON response to client
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&searchResponse{
		StatusCode:  200,
		DataResults: results,
		Message:     "Data Retrieved from NPPES API",
	})
	if err != nil {
		log.Printf("Unable to marshal search response error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}

func buildQuery(search model.SearchFields) (reqQueries string) {
	// NPI Number & Enumeration Type
	if search.Number > 0 {
		reqQueries = fmt.Sprintf("?number=%v&enumeration_type=%s", search.Number, search.EnumerationType)
	} else {
		reqQueries = fmt.Sprintf("?enumeration_type=%s", search.EnumerationType)
	}
	// TaxonomyDesc & First Name
	reqQueries = reqQueries + fmt.Sprintf("&taxonomy_description=%s&first_name=%s", search.TaxonomyDesc, search.FirstName)
	// Last Name & Org Name
	reqQueries = reqQueries + fmt.Sprintf("&last_name=%s&organization_name=%s", search.LastName, search.OrganizationName)
	// Address Purpose & City
	reqQueries = reqQueries + fmt.Sprintf("&address_purpose=%s&city=%s", search.AddressPurpose, search.City)
	// State & Postal Code
	reqQueries = reqQueries + fmt.Sprintf("&state=%s&postal_code=%s", search.State, search.PostalCode)
	// Country Code & Limit
	if search.Limit > 0 {
		reqQueries = reqQueries + fmt.Sprintf("&country_code=%s&limit=%v", search.CountryCode, search.Limit)
	} else {
		reqQueries = reqQueries + fmt.Sprintf("&country_code=%s", search.CountryCode)
	}
	// Skip
	if search.Skip > 0 {
		reqQueries = reqQueries + fmt.Sprintf("&skip=%v&pretty=off", search.Skip)
	} else {
		reqQueries = reqQueries + "&pretty=off"
	}
	return
}

package model

import (
	"encoding/json"
	"testing"
)

func TestCreateProviders(t *testing.T) {
	var testProviders NppesResponse
	err := json.Unmarshal(testData, &testProviders)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = CreateProviders(testProviders)
	if err != nil {
		t.Errorf("Unable to create providers error: %s", err)
	}
}

var testData = []byte(`{
	"result_count":10, "results":[
	{
		"addresses": [
			{
				"address_1": "15380 BAGLEY RD", 
				"address_2": "", 
				"address_purpose": "LOCATION", 
				"address_type": "DOM", 
				"city": "MIDDLEBURG HEIGHTS", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "440-888-6329", 
				"postal_code": "441304824", 
				"state": "OH", 
				"telephone_number": "440-888-6300"
			}, 
			{
				"address_1": "468 COUNTRYSIDE DR", 
				"address_2": "", 
				"address_purpose": "MAILING", 
				"address_type": "DOM", 
				"city": "BROADVIEW HEIGHTS", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "440-582-0622", 
				"postal_code": "441473413", 
				"state": "OH", 
				"telephone_number": "440-582-0612"
			}
		], 
		"basic": {
			"credential": "D.D.S.", 
			"enumeration_date": "2005-05-27", 
			"first_name": "JOHN", 
			"gender": "M", 
			"last_name": "SMITH", 
			"last_updated": "2007-07-08", 
			"middle_name": "MICHAEL", 
			"name": "SMITH JOHN", 
			"name_prefix": "DR.", 
			"sole_proprietor": "NO", 
			"status": "A"
		}, 
		"created_epoch": 1117152000, 
		"enumeration_type": "NPI-1", 
		"identifiers": [
			{
				"code": "05", 
				"desc": "MEDICAID", 
				"identifier": "0596029", 
				"issuer": "", 
				"state": "OH"
			}, 
			{
				"code": "01", 
				"desc": "Other", 
				"identifier": "483091", 
				"issuer": "UNITED CONCORDIA INSURANC", 
				"state": ""
			}
		], 
		"last_updated_epoch": 1183852800, 
		"number": 1497758122, 
		"other_names": [], 
		"taxonomies": [
			{
				"code": "1223P0221X", 
				"desc": "Dentist Pediatric Dentistry", 
				"license": "17587", 
				"primary": true, 
				"state": "OH"
			}
		]
	},{
		"addresses": [
			{
				"address_1": "5050 NE HOYT ST", 
				"address_2": "STE 256", 
				"address_purpose": "LOCATION", 
				"address_type": "DOM", 
				"city": "PORTLAND", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "503-215-6897", 
				"postal_code": "972132982", 
				"state": "OR", 
				"telephone_number": "503-239-7767"
			}, 
			{
				"address_1": "5050 NE HOYT ST", 
				"address_2": "STE 256", 
				"address_purpose": "MAILING", 
				"address_type": "DOM", 
				"city": "PORTLAND", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "503-215-6897", 
				"postal_code": "972132982", 
				"state": "OR", 
				"telephone_number": "503-239-7767"
			}
		], 
		"basic": {
			"credential": "M.D.", 
			"enumeration_date": "2005-05-27", 
			"first_name": "JOHN", 
			"gender": "M", 
			"last_name": "SMITH", 
			"last_updated": "2012-03-14", 
			"middle_name": "WASHBURN", 
			"name": "SMITH JOHN", 
			"name_prefix": "DR.", 
			"name_suffix": "II", 
			"sole_proprietor": "NO", 
			"status": "A"
		}, 
		"created_epoch": 1117152000, 
		"enumeration_type": "NPI-1", 
		"identifiers": [
			{
				"code": "05", 
				"desc": "MEDICAID", 
				"identifier": "1010081", 
				"issuer": "", 
				"state": "WA"
			}, 
			{
				"code": "05", 
				"desc": "MEDICAID", 
				"identifier": "82339", 
				"issuer": "", 
				"state": "OR"
			}
		], 
		"last_updated_epoch": 1331683200, 
		"number": 1427051077, 
		"other_names": [], 
		"taxonomies": [
			{
				"code": "207RX0202X", 
				"desc": "Internal Medicine Medical Oncology", 
				"license": "MD20151", 
				"primary": true, 
				"state": "OR"
			}
		]
	},{
		"addresses": [
			{
				"address_1": "6800 IH10 WEST", 
				"address_2": "STE 300", 
				"address_purpose": "LOCATION", 
				"address_type": "DOM", 
				"city": "SAN ANTONIO", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "210-616-0231", 
				"postal_code": "782012011", 
				"state": "TX", 
				"telephone_number": "210-616-0008"
			}, 
			{
				"address_1": "204 ZAMBRANO RD", 
				"address_2": "", 
				"address_purpose": "MAILING", 
				"address_type": "DOM", 
				"city": "SAN ANTONIO", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "210-822-7706", 
				"postal_code": "782095459", 
				"state": "TX", 
				"telephone_number": "210-822-9449"
			}
		], 
		"basic": {
			"credential": "MD", 
			"enumeration_date": "2005-06-27", 
			"first_name": "JOHN", 
			"gender": "M", 
			"last_name": "SMITH", 
			"last_updated": "2007-07-08", 
			"middle_name": "MARVIN", 
			"name": "SMITH JOHN", 
			"name_prefix": "DR.", 
			"name_suffix": "III", 
			"sole_proprietor": "NO", 
			"status": "A"
		}, 
		"created_epoch": 1119830400, 
		"enumeration_type": "NPI-1", 
		"identifiers": [
			{
				"code": "01", 
				"desc": "Other", 
				"identifier": "D9666", 
				"issuer": "TX PHYSICIAN PERMIT", 
				"state": "TX"
			}
		], 
		"last_updated_epoch": 1183852800, 
		"number": 1770589525, 
		"other_names": [], 
		"taxonomies": [
			{
				"code": "208G00000X", 
				"desc": "Thoracic Surgery (Cardiothoracic Vascular Surgery)", 
				"license": "D9666", 
				"primary": true, 
				"state": "TX"
			}
		]
	},{
		"addresses": [
			{
				"address_1": "9276 W UNION HILLS DR", 
				"address_2": "STE. A", 
				"address_purpose": "LOCATION", 
				"address_type": "DOM", 
				"city": "PEORIA", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "623-972-6334", 
				"postal_code": "853828158", 
				"state": "AZ", 
				"telephone_number": "623-972-6137"
			}, 
			{
				"address_1": "9276 W UNION HILLS DR", 
				"address_2": "STE. A", 
				"address_purpose": "MAILING", 
				"address_type": "DOM", 
				"city": "PEORIA", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "623-972-6334", 
				"postal_code": "853828158", 
				"state": "AZ", 
				"telephone_number": "623-972-6137"
			}
		], 
		"basic": {
			"credential": "D.M.D.", 
			"enumeration_date": "2005-06-30", 
			"first_name": "JOHNNY", 
			"gender": "M", 
			"last_name": "SMITH", 
			"last_updated": "2007-07-08", 
			"middle_name": "LEE", 
			"name": "SMITH JOHNNY", 
			"name_prefix": "DR.", 
			"sole_proprietor": "NO", 
			"status": "A"
		}, 
		"created_epoch": 1120089600, 
		"enumeration_type": "NPI-1", 
		"identifiers": [], 
		"last_updated_epoch": 1183852800, 
		"number": 1629075130, 
		"other_names": [], 
		"taxonomies": [
			{
				"code": "1223G0001X", 
				"desc": "Dentist General Practice", 
				"license": "4489", 
				"primary": true, 
				"state": "AZ"
			}
		]
	},{
		"addresses": [
			{
				"address_1": "818 N EMPORIA ST", 
				"address_2": "SUITE 200", 
				"address_purpose": "LOCATION", 
				"address_type": "DOM", 
				"city": "WICHITA", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "316-263-9523", 
				"postal_code": "672143729", 
				"state": "KS", 
				"telephone_number": "316-263-0296"
			}, 
			{
				"address_1": "818 N EMPORIA ST", 
				"address_2": "SUITE 200", 
				"address_purpose": "MAILING", 
				"address_type": "DOM", 
				"city": "WICHITA", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "316-263-9523", 
				"postal_code": "672143729", 
				"state": "KS", 
				"telephone_number": "316-263-0296"
			}
		], 
		"basic": {
			"credential": "MD", 
			"enumeration_date": "2005-07-08", 
			"first_name": "JOHN", 
			"gender": "M", 
			"last_name": "SMITH", 
			"last_updated": "2010-12-22", 
			"middle_name": "L", 
			"name": "SMITH JOHN", 
			"name_prefix": "DR.", 
			"sole_proprietor": "NO", 
			"status": "A"
		}, 
		"created_epoch": 1120780800, 
		"enumeration_type": "NPI-1", 
		"identifiers": [
			{
				"code": "05", 
				"desc": "MEDICAID", 
				"identifier": "100162260B", 
				"issuer": "", 
				"state": "KS"
			}
		], 
		"last_updated_epoch": 1292976000, 
		"number": 1669470290, 
		"other_names": [], 
		"taxonomies": [
			{
				"code": "204F00000X", 
				"desc": "Transplant Surgery", 
				"license": "0416074", 
				"primary": true, 
				"state": "KS"
			}
		]
	},{
		"addresses": [
			{
				"address_1": "4013 N RIDGE RD", 
				"address_2": "STE 210", 
				"address_purpose": "LOCATION", 
				"address_type": "DOM", 
				"city": "WICHITA", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "316-945-9131", 
				"postal_code": "672058860", 
				"state": "KS", 
				"telephone_number": "316-945-7309"
			}, 
			{
				"address_1": "PO BOX 905", 
				"address_2": "", 
				"address_purpose": "MAILING", 
				"address_type": "DOM", 
				"city": "WICHITA", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "316-945-9131", 
				"postal_code": "672010905", 
				"state": "KS", 
				"telephone_number": "316-945-7309"
			}
		], 
		"basic": {
			"credential": "DO, FACOS", 
			"enumeration_date": "2005-07-15", 
			"first_name": "JOHN", 
			"gender": "M", 
			"last_name": "SMITH", 
			"last_updated": "2011-11-30", 
			"middle_name": "P", 
			"name": "SMITH JOHN", 
			"name_prefix": "DR.", 
			"sole_proprietor": "NO", 
			"status": "A"
		}, 
		"created_epoch": 1121385600, 
		"enumeration_type": "NPI-1", 
		"identifiers": [
			{
				"code": "05", 
				"desc": "MEDICAID", 
				"identifier": "100189270A", 
				"issuer": "", 
				"state": "KS"
			}
		], 
		"last_updated_epoch": 1322611200, 
		"number": 1902805674, 
		"other_names": [], 
		"taxonomies": [
			{
				"code": "208600000X", 
				"desc": "Surgery", 
				"license": "05-19194", 
				"primary": true, 
				"state": "KS"
			}
		]
	},{
		"addresses": [
			{
				"address_1": "2633 CENTENNIAL BLVD", 
				"address_2": "STE 100", 
				"address_purpose": "LOCATION", 
				"address_type": "DOM", 
				"city": "TALLAHASSEE", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "850-431-4794", 
				"postal_code": "323080585", 
				"state": "FL", 
				"telephone_number": "850-431-5474"
			}, 
			{
				"address_1": "2633 CENTENNIAL BLVD", 
				"address_2": "STE 100", 
				"address_purpose": "MAILING", 
				"address_type": "DOM", 
				"city": "TALLAHASSEE", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "850-431-4794", 
				"postal_code": "323080585", 
				"state": "FL", 
				"telephone_number": "850-431-5474"
			}
		], 
		"basic": {
			"credential": "MD", 
			"enumeration_date": "2005-08-25", 
			"first_name": "JOHN", 
			"gender": "M", 
			"last_name": "SMITH", 
			"last_updated": "2017-03-13", 
			"middle_name": "ORSON", 
			"name": "SMITH JOHN", 
			"name_prefix": "DR.", 
			"sole_proprietor": "NO", 
			"status": "A"
		}, 
		"created_epoch": 1124928000, 
		"enumeration_type": "NPI-1", 
		"identifiers": [
			{
				"code": "01", 
				"desc": "Other", 
				"identifier": "00000", 
				"issuer": "HUMANA CHOICE CARE", 
				"state": "FL"
			}, 
			{
				"code": "01", 
				"desc": "Other", 
				"identifier": "00000", 
				"issuer": "NOVA NET", 
				"state": "FL"
			}, 
			{
				"code": "01", 
				"desc": "Other", 
				"identifier": "00000", 
				"issuer": "SOUTHCARE", 
				"state": "FL"
			}, 
			{
				"code": "05", 
				"desc": "MEDICAID", 
				"identifier": "043511200", 
				"issuer": "", 
				"state": "FL"
			}, 
			{
				"code": "01", 
				"desc": "Other", 
				"identifier": "00000", 
				"issuer": "HUMANA/CHOICE CARE", 
				"state": "FL"
			}, 
			{
				"code": "01", 
				"desc": "Other", 
				"identifier": "00000", 
				"issuer": "BEECH STREET/CAPP CARE", 
				"state": "GA"
			}, 
			{
				"code": "01", 
				"desc": "Other", 
				"identifier": "00000", 
				"issuer": "UNITED HEALTH CARE", 
				"state": "FL"
			}, 
			{
				"code": "01", 
				"desc": "Other", 
				"identifier": "00000", 
				"issuer": "VISTA", 
				"state": "FL"
			}, 
			{
				"code": "01", 
				"desc": "Other", 
				"identifier": "37079", 
				"issuer": "BCBS", 
				"state": "FL"
			}, 
			{
				"code": "05", 
				"desc": "MEDICAID", 
				"identifier": "00055646A", 
				"issuer": "", 
				"state": "GA"
			}
		], 
		"last_updated_epoch": 1489393604, 
		"number": 1467445338, 
		"other_names": [], 
		"taxonomies": [
			{
				"code": "207R00000X", 
				"desc": "Internal Medicine", 
				"license": "ME9938", 
				"primary": true, 
				"state": "FL"
			}, 
			{
				"code": "207R00000X", 
				"desc": "Internal Medicine", 
				"license": "007898", 
				"primary": false, 
				"state": "GA"
			}
		]
	},{
		"addresses": [
			{
				"address_1": "2604 DR MARTIN LUTHER KING JR BLVD", 
				"address_2": "", 
				"address_purpose": "LOCATION", 
				"address_type": "DOM", 
				"city": "NEW BERN", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "252-633-2833", 
				"postal_code": "285624238", 
				"state": "NC", 
				"telephone_number": "252-638-4023"
			}, 
			{
				"address_1": "PO BOX 896206", 
				"address_2": "", 
				"address_purpose": "MAILING", 
				"address_type": "DOM", 
				"city": "CHARLOTTE", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "252-633-2833", 
				"postal_code": "282896206", 
				"state": "NC", 
				"telephone_number": "252-638-4023"
			}
		], 
		"basic": {
			"credential": "MD", 
			"enumeration_date": "2005-08-26", 
			"first_name": "JOHN", 
			"gender": "M", 
			"last_name": "SMITH", 
			"last_updated": "2017-03-17", 
			"middle_name": "MATTHEW", 
			"name": "SMITH JOHN", 
			"sole_proprietor": "NO", 
			"status": "A"
		}, 
		"created_epoch": 1125014400, 
		"enumeration_type": "NPI-1", 
		"identifiers": [
			{
				"code": "05", 
				"desc": "MEDICAID", 
				"identifier": "8977585", 
				"issuer": "", 
				"state": "NC"
			}, 
			{
				"code": "01", 
				"desc": "Other", 
				"identifier": "77585", 
				"issuer": "BLUE CROSS", 
				"state": "NC"
			}
		], 
		"last_updated_epoch": 1489757897, 
		"number": 1932192713, 
		"other_names": [], 
		"taxonomies": [
			{
				"code": "207R00000X", 
				"desc": "Internal Medicine", 
				"license": "9401494", 
				"primary": true, 
				"state": "NC"
			}
		]
	},{
		"addresses": [
			{
				"address_1": "2926 CAPITAL BLVD", 
				"address_2": "", 
				"address_purpose": "LOCATION", 
				"address_type": "DOM", 
				"city": "RALEIGH", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "919-878-8863", 
				"postal_code": "276043235", 
				"state": "NC", 
				"telephone_number": "919-878-8848"
			}, 
			{
				"address_1": "2926 CAPITAL BLVD", 
				"address_2": "", 
				"address_purpose": "MAILING", 
				"address_type": "DOM", 
				"city": "RALEIGH", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "919-878-8863", 
				"postal_code": "276043235", 
				"state": "NC", 
				"telephone_number": "919-878-8848"
			}
		], 
		"basic": {
			"credential": "D.C., CCSP,PA", 
			"enumeration_date": "2005-08-31", 
			"first_name": "JOHN", 
			"gender": "M", 
			"last_name": "SMITH", 
			"last_updated": "2007-07-08", 
			"middle_name": "ALBERT", 
			"name": "SMITH JOHN", 
			"name_prefix": "DR.", 
			"sole_proprietor": "NO", 
			"status": "A"
		}, 
		"created_epoch": 1125446400, 
		"enumeration_type": "NPI-1", 
		"identifiers": [
			{
				"code": "05", 
				"desc": "MEDICAID", 
				"identifier": "8908800", 
				"issuer": "", 
				"state": "NC"
			}
		], 
		"last_updated_epoch": 1183852800, 
		"number": 1992799332, 
		"other_names": [], 
		"taxonomies": [
			{
				"code": "111NS0005X", 
				"desc": "Chiropractor Sports Physician", 
				"license": "1412", 
				"primary": true, 
				"state": "NC"
			}
		]
	},{
		"addresses": [
			{
				"address_1": "12832 NW CENTRAL AVE", 
				"address_2": "", 
				"address_purpose": "LOCATION", 
				"address_type": "DOM", 
				"city": "BRISTOL", 
				"country_code": "US", 
				"country_name": "United States", 
				"fax_number": "850-643-5689", 
				"postal_code": "32424", 
				"state": "FL", 
				"telephone_number": "850-643-2415"
			}, 
			{
				"address_1": "834 PLANTATION DR", 
				"address_2": "", 
				"address_purpose": "MAILING", 
				"address_type": "DOM", 
				"city": "PANAMA CITY", 
				"country_code": "US", 
				"country_name": "United States", 
				"postal_code": "324048629", 
				"state": "FL", 
				"telephone_number": "850-785-4522"
			}
		], 
		"basic": {
			"credential": "DR", 
			"enumeration_date": "2005-09-02", 
			"first_name": "JOHN", 
			"gender": "M", 
			"last_name": "SMITH", 
			"last_updated": "2007-07-08", 
			"middle_name": "DAVID", 
			"name": "SMITH JOHN", 
			"sole_proprietor": "NO", 
			"status": "A"
		}, 
		"created_epoch": 1125619200, 
		"enumeration_type": "NPI-1", 
		"identifiers": [], 
		"last_updated_epoch": 1183852800, 
		"number": 1679567796, 
		"other_names": [], 
		"taxonomies": [
			{
				"code": "122300000X", 
				"desc": "Dentist", 
				"license": "LN3799", 
				"primary": true, 
				"state": "AL"
			}
		]
	}
	]}`)

package chapi

import (
	"encoding/json"
	"fmt"
)

// Company is the structure that models an incorporated company. Models a
// company resource.
// https://developer.companieshouse.gov.uk/api/docs/company/company_number/companyProfile-resource.html
type Company struct {
	Name             string `json:"company_name"`
	Number           string `json:"company_number"`
	Status           string `json:"company_status"`
	CreatedOn        string `json:"date_of_creation"`
	CanFile          bool   `json:"can_file"`
	RegisteredOffice struct {
		Address1 string `json:"address_line_1,omitempty"`
		Address2 string `json:"address_line_2,omitempty"`
		CareOf   string `json:"care_of,omitempty"`
		Country  string `json:"country,omitempty"`
		Locality string `json:"locality,omitempty"`
		POBox    string `json:"po_box,omitempty"`
		PostCode string `json:"postal_code"`
		Premises string `json:"premises,omitempty"`
		Region   string `json:"region,omitempty"`
	} `json:"registered_office_address"`
	// SIC codes categorise the business activities.
	SicCodes []string `json:"sic_codes"`
	Type     string   `json:"type"`
}

func (c *Company) Active() bool {
	return c.Status == "active"
}

func (c Company) String() string {
	return fmt.Sprintf("Company %q (%s)", c.Name, c.Number)
}

// CompanyProfile will return the profile of a sinle company using the
// incorporation number as the identifier.
func CompanyProfile(n string) (c Company, err error) {

	path := fmt.Sprintf("/company/%s", n)

	resp, err := makeRequest("GET", path)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&c)

	return
}

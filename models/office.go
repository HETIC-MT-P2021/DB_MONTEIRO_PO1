package models

// Office : Architecture in db
type Office struct {
	OfficeCode   string     `json:"officeCode"`
	City         string     `json:"city"`
	Phone        string     `json:"phone"`
	AddressLine1 string     `json:"addressLine1"`
	AddressLine2 NullString `json:"addressLine2"`
	State        NullString `json:"state"`
	Country      string     `json:"country"`
	PostalCode   string     `json:"postalCode"`
}

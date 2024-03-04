package controllers

type errorResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type createCustomerAddressRequest struct {
	CustomerId string `json:"customerId"`
	Name       string `json:"name,omitempty"`
	Number     string `json:"number,omitempty"`
	StreetName string `json:"streetName"`
	Locality   string `json:"locality"`
	PostTown   string `json:"postTown"`
	PostCode   string `json:"postCode"`
}

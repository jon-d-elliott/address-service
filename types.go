package main

type address struct {
	Name       string `json:"name,omitempty"`
	Number     string `json:"number,omitempty"`
	StreetName string `json:"streetName"`
	Locality   string `json:"locality"`
	PostTown   string `json:"postTown"`
	PostCode   string `json:"postCode"`
}

type createAddressRequest struct {
	CustomerId string `json:"customerId`
	Name       string `json:"name,omitempty"`
	Number     string `json:"number,omitempty"`
	StreetName string `json:"streetName"`
	Locality   string `json:"locality"`
	PostTown   string `json:"postTown"`
	PostCode   string `json:"postCode"`
}

type getCustomerAddressesResponse struct {
	AddressId  string `json:"addressId"`
	Name       string `json:"name,omitempty"`
	Number     string `json:"number,omitempty"`
	StreetName string `json:"streetName"`
	Locality   string `json:"locality"`
	PostTown   string `json:"postTown"`
	PostCode   string `json:"postCode"`
}

type createAddressResponse struct {
	AddressId string `json:"addressId"`
}

type errorResponse struct {
	Code string `json:"code"`
	Description string `json:"description"`
}

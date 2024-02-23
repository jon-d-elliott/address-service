package main

import (
	"github.com/gin-gonic/gin"
//	"github.com/google/uuid"
	"net/http"
)

type address struct {
	Name string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
	StreetName string `json:"streetName"`
	Locality string `json:"locality"`
	PostTown string `json:"postTown"`
	PostCode string `json:"postCode"`
}

type getCustomerAddressesResponse struct {
	AddressId string `json:"addressId"`
	Name string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
	StreetName string `json:"streetName"`
	Locality string `json:"locality"`
	PostTown string `json:"postTown"`
	PostCode string `json:"postCode"`
}


type creatAddressRequest struct {
	CustomerId string `json:"customerId"`
	Name string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
	StreetName string `json:"streetName"`
	Locality string `json:"locality"`
	PostTown string `json:"postTown"`
	PostCode string `json:"postCode"`
}

type createAddressResponse struct {
	AddressId string `json:"addressId"`
}

type updateAddressRequest struct {
	Name string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
	StreetName string `json:"streetName"`
	Locality string `json:"locality"`
	PostTown string `json:"postTown"`
	PostCode string `json:"postCode"`
}



var addressMap = map[string]map[string]address { 
	"1": {
		"1": {
				Name: "Wychbury",
				Number: "24",
				StreetName: "Yarnborough Hill",
				Locality: "Stourbridge",
				PostTown: "Oldswinford",
				PostCode: "DY8 2EB",
		},
	 	"2": {
				Number: "34",
				StreetName: "Hampton Drive",
				Locality: "Telford",
				PostTown: "Newport",
				PostCode: "TF10 7RE",
		 },
	},
}

func getCustomerAddresses(c *gin.Context) {
	customerId := c.Param("customerId")
	var custAddresses []getCustomerAddressesResponse
	for k,v  := range addressMap {
		if k == customerId {
			for k1,v1 := range v {
				address := getCustomerAddressesResponse{
					AddressId: k1,
					Name: v1.Name,
					Number: v1.Number,
					StreetName: v1.StreetName,
					Locality: v1.Locality,
					PostTown: v1.PostTown,
					PostCode: v1.PostCode,
					}
					custAddresses = append(custAddresses, address)
			}
		}
	}
	if len(custAddresses) > 0 {
		c.IndentedJSON(http.StatusOK, custAddresses)
	}else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no addresses found for customer : " + customerId})
	}
}


func main() {
	router := gin.Default()
	router.GET("/addresses/customers/:customerId", getCustomerAddresses)
	router.Run("localhost:8080")
}
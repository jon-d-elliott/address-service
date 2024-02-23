package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"net/http"
)


var addressMap = map[string]map[string]address{
	"1": {
		"1": {
			Name:       "Wychbury",
			Number:     "24",
			StreetName: "Yarnborough Hill",
			Locality:   "Stourbridge",
			PostTown:   "Oldswinford",
			PostCode:   "DY8 2EB",
		},
		"2": {
			Number:     "34",
			StreetName: "Hampton Drive",
			Locality:   "Telford",
			PostTown:   "Newport",
			PostCode:   "TF10 7RE",
		},
	},
}

func getCustomerAddresses(c *gin.Context) {
	customerId := c.Query("customerId")

	if customerId == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "customer id not supplied"})
		return
	}
	var custAddresses []getCustomerAddressesResponse
	for k, v := range addressMap {
		if k == customerId {
			for k1, v1 := range v {
				address := getCustomerAddressesResponse{
					AddressId:  k1,
					Name:       v1.Name,
					Number:     v1.Number,
					StreetName: v1.StreetName,
					Locality:   v1.Locality,
					PostTown:   v1.PostTown,
					PostCode:   v1.PostCode,
				}
				custAddresses = append(custAddresses, address)
			}
		}
	}
	if len(custAddresses) > 0 {
		c.IndentedJSON(http.StatusOK, custAddresses)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no addresses found for customer : " + customerId})
}

func getAddressById(c *gin.Context) {
	customerId := c.Query("customerId")
	if customerId == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "customer id not supplied"})
		return
	}

	addressId := c.Param("addressId")
	theAddress, exists := addressMap[customerId][addressId]
	if !exists {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Address not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, theAddress)
}

func createAddress(c *gin.Context) {
	var request createAddressRequest


	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customerId := request.CustomerId
	addressId := uuid.NewString()
	addMap, exists := addressMap[customerId]

	if exists {

		addMap[addressId] = address{
			Name: request.Name,
			Number: request.Number,
			StreetName: request.StreetName,
			Locality: request.Locality,
			PostTown: request.PostTown,
			PostCode: request.PostCode,
			}
	}else {
		addressMap[customerId] = map[string]address{
			addressId: {
				Name: request.Name,
				Number: request.Number,
				StreetName: request.StreetName,
				Locality: request.Locality,
				PostTown: request.PostTown,
				PostCode: request.PostCode,
			},
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Address created with id: " + addressId})
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted"})

}

func deleteAddress(c *gin.Context) {
	customerId := c.Query("customerId")
	addressId := c.Param("addressId")
	if customerId == ""{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "Customer Id not provided"})
		return
	}

	delete(addressMap[customerId], addressId)
}

func updateAddress(c *gin.Context) {
	customerId := c.Query("customerId")
	addressId := c.Param("addressId")

	var theAddress address

	if customerId == ""{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "Customer Id not provided"})
		return
	}

	if err := c.BindJSON(&theAddress); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addressMap[customerId][addressId] = theAddress

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Updated"})



}

func main() {
	router := gin.Default()
	router.GET("/addresses", getCustomerAddresses)
	router.GET("/addresses/addressId/:addressId", getAddressById)
	router.POST("/addresses", createAddress)
	router.DELETE("/addresses/addressId/:addressId", deleteAddress)
	router.PUT("/addresses/addressId/:addressId", updateAddress)
	router.Run("localhost:8080")
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"net/http"
)

var addressMap = make(map[string]map[string]address)

func getCustomerAddresses(c *gin.Context) {
	customerId := c.Query("customerId")

	if customerId == "" {
		errorResp := errorResponse{
			Code:        "0001",
			Description: "customer id is mandatory",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResp)
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
	errorResp := errorResponse{
		Code:        "0002",
		Description: "Addresses not found.",
	}

	c.AbortWithStatusJSON(http.StatusNotFound, errorResp)
}

func getAddressById(c *gin.Context) {
	customerId := c.Query("customerId")
	if customerId == "" {
		errorResp := errorResponse{
			Code:        "0001",
			Description: "customer id is mandatory",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResp)
		return
	}

	addressId := c.Param("addressId")
	theAddress, exists := addressMap[customerId][addressId]
	if !exists {
		errorResp := errorResponse{
			Code:        "0002",
			Description: "Address not found.",
		}
		c.AbortWithStatusJSON(http.StatusNotFound, errorResp)
		return
	}

	c.IndentedJSON(http.StatusOK, theAddress)
}

func createAddress(c *gin.Context) {
	var request createAddressRequest

	if err := c.BindJSON(&request); err != nil {
		c.Error(err)
		errorResp := errorResponse{
			Code:        "0004",
			Description: "Invalid request : " + err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusNotFound, errorResp)
		return
	}

	customerId := request.CustomerId
	addressId := uuid.NewString()
	addMap, exists := addressMap[customerId]

	if exists {

		addMap[addressId] = address{
			Name:       request.Name,
			Number:     request.Number,
			StreetName: request.StreetName,
			Locality:   request.Locality,
			PostTown:   request.PostTown,
			PostCode:   request.PostCode,
		}
	} else {
		addressMap[customerId] = map[string]address{
			addressId: {
				Name:       request.Name,
				Number:     request.Number,
				StreetName: request.StreetName,
				Locality:   request.Locality,
				PostTown:   request.PostTown,
				PostCode:   request.PostCode,
			},
		}
	}

	resp := createAddressResponse{
		AddressId: addressId,
	}

	c.IndentedJSON(http.StatusOK, resp)
}

func deleteAddress(c *gin.Context) {
	customerId := c.Query("customerId")
	addressId := c.Param("addressId")
	if customerId == "" {
		errorResp := errorResponse{
			Code:        "0001",
			Description: "Customer id is mandatory",
		}
		c.AbortWithStatusJSON(http.StatusNotFound, errorResp)
		return
	}

	delete(addressMap[customerId], addressId)
}

func updateAddress(c *gin.Context) {
	customerId := c.Query("customerId")
	addressId := c.Param("addressId")

	var theAddress address

	if customerId == "" {
		errorResp := errorResponse{
			Code:        "0001",
			Description: "Customer id is mandatory",
		}
		c.AbortWithStatusJSON(http.StatusNotFound, errorResp)
		return
	}

	if err := c.BindJSON(&theAddress); err != nil {
		c.Error(err)
		errorResp := errorResponse{
			Code:        "0004",
			Description: "Invalid request : " + err.Error(),
			}
			c.AbortWithStatusJSON(http.StatusNotFound, errorResp)
		return
	}

	addressMap[customerId][addressId] = theAddress

	c.IndentedJSON(http.StatusNoContent, nil)

}

func main() {
	router := gin.Default()
	router.GET("/addresses", getCustomerAddresses)
	router.GET("/addresses/addressId/:addressId", getAddressById)
	router.POST("/addresses", createAddress)
	router.DELETE("/addresses/addressId/:addressId", deleteAddress)
	router.PUT("/addresses/addressId/:addressId", updateAddress)
	router.Run("localhost:8082")
}

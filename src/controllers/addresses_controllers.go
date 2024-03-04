package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/jon-d-elliott/address-service/src/models"
)

func GetAllCustomerAddresses(c *gin.Context) {

	customerId, errorResponse := getCustomerId(c)

	if customerId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse)
		return
	}

	addresses, err := models.FetchCustomerAddresses(customerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Addresses fetched successfully", "status": "success", "data": addresses})
}

func GetCustomerAddress(c *gin.Context) {
	customerId, errorResponse := getCustomerId(c)

	if customerId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse)
		return
	}
	addressId := c.Param("addressId")

	address, err := models.FetchAddressById(customerId, addressId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address fetched successfully", "status": "success", "data": address})

}

func UpdateCustomerAddress(c *gin.Context) {
	customerId, errorResponse := getCustomerId(c)

	if customerId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse)
		return
	}

	var input createCustomerAddressRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
	}

	addressId := c.Param("addressId")

	var address *models.Address

	address = &models.Address{
		CustomerId: customerId,
		AddressId:  addressId,
		Name:       input.Name,
		Number:     input.Number,
		StreetName: input.StreetName,
		Locality:   input.Locality,
		PostTown:   input.PostTown,
		PostCode:   input.PostCode,
	}

	address, err := address.UpdateCustomerAddress(customerId, addressId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Address updated successfully", "data": address})

}

func CreateCustomerAddress(c *gin.Context) {

	var input createCustomerAddressRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
	}

	addressId := uuid.NewString()

	address := models.Address{
		CustomerId: input.CustomerId,
		AddressId:  addressId,
		Name:       input.Name,
		Number:     input.Number,
		StreetName: input.StreetName,
		Locality:   input.Locality,
		PostTown:   input.PostTown,
		PostCode:   input.PostCode,
	}

	savedAddress, err := address.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Address Saved succesfully", "data": savedAddress})

}

func DeleteCustomerAddress(c *gin.Context) {
	customerId, errorResponse := getCustomerId(c)

	if customerId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse)
		return
	}

	addressId := c.Param("addressId")

	if err := models.DeleteAddress(customerId, addressId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}
	c.Status(http.StatusNoContent)
}

func getCustomerId(c *gin.Context) (string, errorResponse) {
	log.Print("Entered Get CustomerId ")
	customerId := c.Query("customerId")
	if customerId == "" {
		errorResp := errorResponse{
			Code:        "0001",
			Description: "Customer id is mandatory",
		}
		return "", errorResp
	}
	return customerId, errorResponse{}
}

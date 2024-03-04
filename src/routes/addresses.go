package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jon-d-elliott/address-service/src/controllers"
)

func addressesGroupRouter(baseRouter *gin.RouterGroup) {

	addresses := baseRouter.Group("/addresses")

	addresses.GET("", controllers.GetAllCustomerAddresses)
	addresses.GET("/addressId/:addressId", controllers.GetCustomerAddress)
	addresses.POST("", controllers.CreateCustomerAddress)
	addresses.PATCH("/addressId/:addressId", controllers.UpdateCustomerAddress)
	addresses.DELETE("/addressId/:addressId", controllers.DeleteCustomerAddress)
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	versionRouter := r.Group("/api/v1")
	addressesGroupRouter(versionRouter)
	return r
}

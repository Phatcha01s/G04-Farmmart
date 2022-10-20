package main

import (
	"github.com/B6202385/G04-Farmmart/controller"
	"github.com/B6202385/G04-Farmmart/entity"
	"github.com/B6202385/G04-Farmmart/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// User Routes
			protected.GET("/users", controller.ListUsers)
			protected.GET("/user/:id", controller.GetUser)
			protected.PATCH("/users", controller.UpdateUser)
			protected.DELETE("/users/:id", controller.DeleteUser)

			// order Routes
			protected.GET("/orders", controller.ListOrders)
			protected.GET("/order/:id", controller.GetOrder)
			protected.GET("/order/user/:id", controller.ListOrderbyUser)
			protected.POST("/orders", controller.CreateOrder)
			protected.PATCH("/orders", controller.UpdateOrder)
			protected.DELETE("/orders/:id", controller.DeleteOrder)

			// Paymentmethod Routes
			protected.GET("/paymentmethods", controller.ListPaymentmethods)
			protected.GET("/paymentmethod/:id", controller.GetPaymentmethod)
			protected.POST("/paymentmethods", controller.CreatePaymentmethod)
			protected.PATCH("/paymentmethods", controller.UpdatePaymentmethod)
			protected.DELETE("/paymentmethods/:id", controller.DeletePaymentmethod)

			// Deliverytype Routes
			protected.GET("/deliverytypes", controller.ListDeliverytypes)
			protected.GET("/deliverytype/:id", controller.GetDeliverytype)
			protected.POST("/deliverytypes", controller.CreateDeliverytype)
			protected.PATCH("/deliverytypes", controller.UpdateDeliverytype)
			protected.DELETE("/deliverytypes/:id", controller.DeleteDeliverytype)

			// payment Routes
			protected.GET("payments", controller.ListPayments)
			protected.GET("/payment/:id", controller.GetPayment)
			protected.POST("payments", controller.CreatePayment)
			protected.PATCH("payments", controller.UpdatePayment)
			protected.DELETE("/payments/:id", controller.DeletePayment)

			
		}
	}

	// User Routes
	r.POST("/users", controller.CreateUser)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

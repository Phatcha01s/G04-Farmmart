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

			// Staff Routes
			protected.GET("/staffs", controller.ListStaffs)
			protected.GET("/staff/:id", controller.GetStaff)
			protected.PATCH("/staffs", controller.UpdateStaff)
			protected.DELETE("/staffs/:id", controller.DeleteStaff)

			// ProductType Routes
			protected.GET("/producttypes", controller.ListProductTypes)
			protected.GET("/producttype/:id", controller.GetProductType)
			protected.POST("/producttypes", controller.CreateProductType)
			protected.PATCH("/producttypes", controller.UpdateProductType)
			protected.DELETE("/producttypes/:id", controller.DeleteProductType)

			// Product Routes
			protected.GET("/products", controller.ListProducts)
			protected.GET("/product/:id", controller.GetProduct)
			protected.POST("/products", controller.CreateProduct)
			protected.PATCH("/products", controller.UpdateProduct)
			protected.DELETE("/products/:id", controller.DeleteProduct)

			// Supplier Routes
			protected.GET("/suppliers", controller.ListSuppliers)
			protected.GET("/supplier/:id", controller.GetSupplier)
			protected.POST("/suppliers", controller.CreateSupplier)
			protected.PATCH("/suppliers", controller.UpdateSupplier)
			protected.DELETE("/suppliers/:id", controller.DeleteSupplier)

			// ProductStock Routes
			protected.GET("/product_stocks", controller.ListProductStocks)
			protected.GET("/productstock/:id", controller.GetProductStock)
			protected.POST("/product_stocks", controller.CreateProductStock)
			protected.PATCH("/product_stocks", controller.UpdateProductStock)
			protected.DELETE("/productstocks/:id", controller.DeleteProductStock)

			// Return Routes
			protected.GET("/return_s/:id", controller.ListReturns)
			protected.GET("/return/:id", controller.GetReturn)
			protected.POST("/return_s", controller.CreateReturnod)
			protected.PATCH("/returns", controller.UpdateReturn)
			protected.DELETE("/return_s/:id", controller.DeleteReturn)

		}
	}
	
	// Staff Routes
	r.POST("/staffs", controller.CreateStaff)

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

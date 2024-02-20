package server

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	products "main/src/products/infrastructure"
	"net/http"
)

// @title Ecommerce Joaquin API
// @version 1.0
// @description This is a sample server for cart and products.
// @termsOfService http://swagger.io/terms/

// joaquin.batrez@gmail.com API Support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// Init @host e-commerce.swagger.io
// @BasePath /
func Init() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	productHandler := products.NewProductHandler()
	cartHandler := products.NewCartHandler()

	e.GET("/products", productHandler.ListAllProducts)
	e.GET("/products/:id", productHandler.GetProductById)
	e.DELETE("/products/:id", productHandler.DeleteProductById)
	e.POST("/products", productHandler.CreateOrUpdateProduct)

	e.GET("/cart/:id", cartHandler.GetInfo)
	e.POST("/cart", cartHandler.AddProductToCart)
	e.POST("/cart/:id", cartHandler.AddProductToCart)
	e.POST("/cart/:id/remove", cartHandler.RemoveProductFromCart)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":9998"))
}

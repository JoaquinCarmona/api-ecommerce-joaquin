package infrastructure

import (
	"github.com/labstack/echo/v4"
	"main/src/products/application"
	"net/http"
)

type CartHandler struct {
	Service application.CartService
}

func NewCartHandler() *CartHandler {
	return &CartHandler{
		Service: application.CartService{
			CartInterfaceRepository: &CartRepository{},
		},
	}
}

func (handler *CartHandler) AddProductToCart(c echo.Context) error {

	err := handler.Service.AddProduct(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (handler *CartHandler) RemoveProductFromCart(c echo.Context) error {
	err := handler.Service.RemoveProduct(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil

}

func (handler *CartHandler) GetInfo(c echo.Context) error {
	err := handler.Service.GetInfo(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil
}

package application

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"main/src/products/domain/models"
	"main/src/products/domain/usecases"
	"net/http"
	"time"
)

type CartService struct {
	CartInterfaceRepository usecases.Cart
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}

func (service *CartService) AddProduct(c echo.Context) error {
	id := c.Param("id")

	type RequestBody struct {
		ProductID string `json:"product_id"`
	}

	var requestBody RequestBody

	if err := c.Bind(&requestBody); err != nil {
		response := Response{
			Success: false,
			Msg:     "Unable to parse request body",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	productId, err := uuid.Parse(requestBody.ProductID)
	if err != nil {
		response := Response{
			Success: false,
			Msg:     "Invalid id product (uuid) on adding cart",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	var cartUpdated models.Cart
	var cartRelation models.CartProducts

	if id != "" {

		existentCart, err := service.CartInterfaceRepository.GetById(id)

		if err != nil {
			log.Println("Couldn't add product to no existent cart")
			response := Response{
				Success: false,
				Msg:     err.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		log.Println("Existent Cart")

		cartRelation, exists, _ := service.CartInterfaceRepository.GetExistentRelation(existentCart, productId)

		typeAdd := "add"
		if exists {
			typeAdd = "sum"
		}
		log.Println("type " + typeAdd)
		cartUpdated, err = service.CartInterfaceRepository.AddProduct(existentCart, productId, cartRelation, typeAdd)

		if err != nil {
			response := Response{
				Success: false,
				Msg:     err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, response)
		}
	} else {
		log.Println("Cart not existent")
		newCart, err := service.createCart()
		if err != nil {
			log.Println("Couldn't create new cart")
			response := Response{
				Success: false,
				Msg:     err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, response)
		}
		cartUpdated, err = service.CartInterfaceRepository.AddProduct(newCart, productId, cartRelation, "add")
		if err != nil {
			response := Response{
				Success: false,
				Msg:     err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, response)
		}
	}
	response := Response{
		Success: true,
		Data:    cartUpdated,
		Msg:     "Product added successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (service *CartService) RemoveProduct(c echo.Context) error {
	id := c.Param("id")

	type RequestBody struct {
		ProductID string `json:"product_id"`
	}

	var requestBody RequestBody

	if err := c.Bind(&requestBody); err != nil {
		response := Response{
			Success: false,
			Msg:     "Unable to parse request body",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	productId, err := uuid.Parse(requestBody.ProductID)
	if err != nil {
		response := Response{
			Success: false,
			Msg:     "Invalid id product (uuid) on adding cart",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	var cartUpdated models.Cart

	if id != "" {

		existentCart, errCart := service.CartInterfaceRepository.GetById(id)

		if errCart != nil {
			log.Println("Couldn't remove product from no existent cart")
			response := Response{
				Success: false,
				Msg:     errCart.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		cartRelation, exists, _ := service.CartInterfaceRepository.GetExistentRelation(existentCart, productId)

		if exists == true {
			cartUpdated, _ = service.CartInterfaceRepository.RemoveProduct(existentCart, productId, cartRelation)
			response := Response{
				Success: true,
				Data:    cartUpdated,
				Msg:     "Product removed successfully",
			}
			return c.JSON(http.StatusOK, response)
		} else {
			response := Response{
				Success: false,
				Msg:     "Product not present in cart",
			}
			return c.JSON(http.StatusBadRequest, response)
		}

	} else {
		response := Response{
			Success: false,
			Msg:     "Please provide a valid cart id",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
}

func (service *CartService) GetInfo(c echo.Context) error {

	cartId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response := Response{
			Success: false,
			Msg:     "Invalid id product (uuid) on adding cart",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	cart, err := service.CartInterfaceRepository.GetInfo(cartId)
	response := Response{
		Success: true,
		Data:    cart,
		Msg:     "Product removed successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (service *CartService) createCart() (models.Cart, error) {

	cart := models.Cart{
		TotalPriceInCents: 0,
		Currency:          "USD",
		Status:            "open",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		Products:          nil,
	}

	newCart, err := service.CartInterfaceRepository.CreateCart(cart)

	if err != nil {
		log.Println("could not create cart")
		return newCart, err
	}

	log.Println("Cart created Successfully")
	return newCart, nil
}

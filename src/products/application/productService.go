package application

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"main/src/products/domain/models"
	"main/src/products/domain/usecases"
	"net/http"
	"time"
)

type ProductService struct {
	ProductInterfaceRepository usecases.Product
}

func (service *ProductService) CreateOrUpdate(c echo.Context) error {
	product := models.Product{}
	err := c.Bind(&product)

	if err != nil {
		response := Response{
			Success: false,
			Msg:     "Unable to parse request body",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	newProduct, errdb := service.ProductInterfaceRepository.CreateOrUpdate(product)

	if errdb != nil {
		response := Response{
			Success: false,
			Msg:     errdb.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := Response{
		Success: true,
		Data:    newProduct,
		Msg:     "Product created or updated successfully",
	}
	return c.JSON(http.StatusCreated, response)
}
func (service *ProductService) List(c echo.Context) error {
	products, err := service.ProductInterfaceRepository.List()
	if err != nil {
		response := Response{
			Success: false,
			Msg:     err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := Response{
		Success: true,
		Data:    products,
		Msg:     "Products listed successfully",
	}
	return c.JSON(http.StatusOK, response)
}
func (service *ProductService) DeleteById(c echo.Context) error {
	id := c.Param("id")
	uuidProduct, err := uuid.Parse(id)
	if err != nil {
		response := Response{
			Success: false,
			Msg:     "Invalid id product (uuid)",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	err = service.ProductInterfaceRepository.DeleteById(uuidProduct)
	if err != nil {
		response := Response{
			Success: false,
			Msg:     err.Error(),
		}
		return c.JSON(http.StatusNotFound, response)
	}
	response := Response{
		Success: true,
		Msg:     "Product deleted successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (service *ProductService) GetById(c echo.Context) error {
	id := c.Param("id")
	uuidProduct, err := uuid.Parse(id)
	if err != nil {
		response := Response{
			Success: false,
			Msg:     "Invalid id product (uuid)",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	product, err := service.ProductInterfaceRepository.GetById(uuidProduct)
	if err != nil {
		response := Response{
			Success: false,
			Msg:     err.Error(),
		}
		return c.JSON(http.StatusNotFound, response)
	}
	response := Response{
		Success: true,
		Data:    product,
		Msg:     "Product retrieved successfully",
	}
	return c.JSON(http.StatusOK, response)
}

package infrastructure

import (
	"github.com/labstack/echo/v4"
	"main/src/products/application"
	"net/http"
)

type ProductHandler struct {
	Service application.ProductService
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		Service: application.ProductService{
			ProductInterfaceRepository: &ProductRepository{},
		},
	}
}

func (handler *ProductHandler) CreateOrUpdateProduct(c echo.Context) error {

	err := handler.Service.CreateOrUpdate(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil
}

// ListAllProducts godoc
// @Summary      Show a list of products
// @Description  get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Product
// @Router       /products [get]

func (handler *ProductHandler) ListAllProducts(c echo.Context) error {
	err := handler.Service.List(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil

}

// GetProductById ShowProduct godoc
// @Summary      Show a product
// @Description  get string by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      uuid  true  "Product ID"
// @Success      200  {object}  models.Product
// @Router       /products/{id} [get]
func (handler *ProductHandler) GetProductById(c echo.Context) error {
	err := handler.Service.GetById(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (handler *ProductHandler) DeleteProductById(c echo.Context) error {
	err := handler.Service.DeleteById(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil
}

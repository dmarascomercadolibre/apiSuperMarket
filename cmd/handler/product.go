package handler

import (
	"app/internal/domain"
	"app/internal/product"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productGroup *gin.RouterGroup
	service      *product.ProductService
}

type BodyRequestCreate struct {
	domain.AtributtesProduct
}
type BodyRequestUpdate struct {
	domain.AtributtesProduct
}

// NewProductHandler creates a new instance of ProductHandler with the given productGroup and initializes the service.
// It returns a pointer to the newly created ProductHandler.
func NewProductHandler(productGroup *gin.RouterGroup) *ProductHandler {
	return &ProductHandler{
		productGroup: productGroup,
		service:      product.NewProductService(),
	}
}

// ProductRoutes registers the routes for handling product-related requests.
func (p *ProductHandler) ProductRoutes() {
	p.productGroup.GET("/", p.GetAllProducts())
	p.productGroup.GET("/:id", p.GetByID())
	p.productGroup.POST("/", p.CreateProduct())
	p.productGroup.PUT("/:id", p.UpdateProduct())
}

// GetAllProducts returns a Gin handler function that handles the GET request for retrieving all products.
// It calls the GetAllProducts method of the associated service to fetch all products and responds with a JSON representation of the products.
func (p *ProductHandler) GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products := p.service.GetAllProducts()
		ctx.JSON(http.StatusOK, products)
	}
}

// GetByID returns a Gin handler function that retrieves a product by its ID.
// If the product is found, it will respond with a HTTP 200 OK and the product details in JSON format.
// Otherwise, it will respond with a HTTP 404 Not Found.
func (p *ProductHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		IntId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		product := p.service.GetByID(IntId)
		ctx.JSON(http.StatusOK, product)
	}

}

// CreateProduct is a handler function that creates a new product.
// If the request body is invalid, it returns a 400 Bad Request response.
// Upon successful creation, it returns a 200 OK response with the created product.
func (p *ProductHandler) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
		var body BodyRequestCreate
		err := ctx.ShouldBind(&body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		product := domain.Product{
			AtributtesProduct: body.AtributtesProduct,
		}
		validate, err := ValidateAtributtes(product)
		if !validate {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		product = p.service.CreateProduct(product)
		ctx.JSON(http.StatusCreated, gin.H{"message": "Products created successfully",
			"data": product,
		})
	}
}

// UpdateProduct is a handler function that updates a product.
// If the request is valid, the product is updated and the updated product is returned in the response.
// If any errors occur during the update process, an appropriate error response is returned.
func (p *ProductHandler) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
		var body BodyRequestUpdate
		err := ctx.ShouldBind(&body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		IntId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		product := domain.Product{
			ID:                IntId,
			AtributtesProduct: body.AtributtesProduct,
		}
		validate, err := ValidateAtributtes(product)
		if !validate {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		product, err = p.service.UpdateProduct(product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": "Products updated successfully",
			"data": product,
		})
	}
}

func ValidateAtributtes(product domain.Product) (validate bool, err error) {
	if product.Name == "" {
		return false, errors.New("Name is required")
	}
	if product.Price == 0 {
		return false, errors.New("Price is required")
	}
	if product.Quantity == 0 {
		return false, errors.New("Quantity is required")
	}
	if product.Code_value == "" {
		return false, errors.New("Description is required")
	}
	if product.Expiration == "" {
		return false, errors.New("Expiration is required")
	}
	if product.IsPublished == nil {
		return false, errors.New("IsPublished is required")
	}
	return true, nil
}

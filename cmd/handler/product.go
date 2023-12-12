package handler

import (
	"app/internal/domain"
	"app/internal/product"
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
		var body BodyRequestCreate
		err := ctx.ShouldBind(&body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		product := domain.Product{
			AtributtesProduct: body.AtributtesProduct,
		}
		p.service.CreateProduct(product)
		ctx.JSON(http.StatusOK, product)
	}
}

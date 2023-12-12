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

func NewProductHandler(productGroup *gin.RouterGroup) *ProductHandler {
	return &ProductHandler{
		productGroup: productGroup,
		service:      product.NewProductService(),
	}
}

func (p *ProductHandler) ProductRoutes() {
	p.productGroup.GET("/", p.GetAllProducts())
	p.productGroup.GET("/:id", p.GetByID())
	p.productGroup.POST("/", p.CreateProduct())
}

func (p *ProductHandler) GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products := p.service.GetAllProducts()
		ctx.JSON(http.StatusOK, products)
	}
}

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
func (p *ProductHandler) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body BodyRequestCreate
		err := ctx.BindJSON(&body)
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

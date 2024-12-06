package service

import (
	"context"
	dalproduct "github.com/U1traVeno/tiktok-shop/biz/dal/model"
	query "github.com/U1traVeno/tiktok-shop/biz/dal/query/product"
	"github.com/U1traVeno/tiktok-shop/biz/model/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"strconv"
)

type ProductService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewProductService(ctx context.Context, c *app.RequestContext) *ProductService {
	return &ProductService{
		ctx: ctx,
		c:   c,
	}
}

// GetProduct 根据 id 获取产品
func (s *ProductService) GetProduct(req *product.GetProductReq) (*product.GetProductResp, error) {
	Id := req.GetId()
	productQuery := query.Product
	// 根据 ID 查找产品
	Product, err := productQuery.Where(productQuery.Id.Eq(Id)).First()
	if err != nil {
		return nil, err
	}

	// 构建响应
	resp := &product.GetProductResp{
		StatusCode: 200,
		StatusMsg:  "find product successfully" + strconv.Itoa(int(Id)),
		Product: &product.Product{
			Id:          Product.Id,
			Name:        Product.Name,
			Amount:      int32(Product.Amount),
			Description: Product.Description,
			Picture:     Product.Picture,
			Price:       Product.Price,
			Categories:  Product.Categories,
		},
	}

	return resp, nil
}

func (s *ProductService) convertToProductList(products []*product.Product) []*product.Product {
	var producilist []*product.Product
	for _, productElment := range products {
		producilist = append(producilist, &product.Product{
			Id:          productElment.Id,
			Name:        productElment.Name,
			Amount:      productElment.Amount,
			Description: productElment.Description,
			Picture:     productElment.Picture,
			Price:       productElment.Price,
			Categories:  productElment.Categories,
		})
	}
	return producilist
}

// ListProducts，按类别批量查询商品
func (s *ProductService) ListProducts(req *product.ListProductsReq) (*product.ListProductsResp, error) {
	Categories := pq.StringArray(req.CategoryName)
	products, err := query.Product.Where(query.Product.Categories.Eq(Categories)).Limit(int((req.GetPage()) * (req.GetPageSize()))).Find()
	if err != nil {
		return nil, err
	}

	var producilist []*product.Product
	for _, productElement := range products {
		producilist = append(producilist, &product.Product{
			Id:          productElement.Id,
			Name:        productElement.Name,
			Amount:      int32(productElement.Amount),
			Description: productElement.Description,
			Picture:     productElement.Picture,
			Price:       productElement.Price,
			Categories:  productElement.Categories,
		})
	}
	resp := &product.ListProductsResp{
		StatusMsg:  "find products successfully",
		StatusCode: 200,

		Products: producilist,
	}

	return resp, nil
}

// SearchProducts 根据关键字搜索商品
func (s *ProductService) SearchProducts(req *product.SearchProductsReq) (*product.SearchProductsResp, error) {
	products, err := query.Product.Where(query.Product.Name.Like(req.Query)).Find()
	if err != nil {
		return nil, err
	}

	var productlist []*product.Product
	for _, productElement := range products {
		productlist = append(productlist, &product.Product{
			Id:          productElement.Id,
			Name:        productElement.Name,
			Amount:      int32(productElement.Amount),
			Description: productElement.Description,
			Picture:     productElement.Picture,
			Price:       productElement.Price,
			Categories:  productElement.Categories,
		})
	}
	resp := &product.SearchProductsResp{
		Results:    productlist,
		StatusCode: 200,
		StatusMsg:  "find products successfully" + strconv.Itoa(len(productlist)) + "个商品",
	}

	return resp, nil
}

// CreateProduct 创建商品
func (s *ProductService) CreateProduct(req *product.CreateProductReq) (*product.CreateProductResp, error) {
	newProduct := &dalproduct.Product{
		Id:          uint32(req.Id),
		Name:        req.Name,
		Amount:      int(req.Amount),
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  pq.StringArray(req.Categories),
	}

	productQuery := query.Product
	// 创建商品
	err := productQuery.WithContext(s.ctx).Create(newProduct)
	if err != nil {
		return nil, err
	}

	resp := &product.CreateProductResp{

		Product: &product.Product{
			Id:          uint32(req.Id),
			Name:        req.Name,
			Amount:      req.Amount,
			Description: req.Description,
			Picture:     req.Picture,
			Price:       req.Price,
			Categories:  pq.StringArray(req.Categories),
		},
		StatusCode: 201,
		StatusMsg:  "create product successfully" + strconv.Itoa(int(req.Id)),
	}

	return resp, nil
}

// UpdateProduct 增减商品数量
func (s *ProductService) UpdateProductAmount(Id uint32, UpdateAmount int32) (*product.UpdateProductResp, error) {
	productQuery := query.Product
	preAmount, err := productQuery.WithContext(s.ctx).Where(productQuery.Id.Eq(Id)).Select(productQuery.Amount).First()
	if err != nil {
		return nil, err
	}
	// 更新商品
	_, err = productQuery.WithContext(s.ctx).Where(productQuery.Id.Eq(Id)).Updates(
		dalproduct.Product{
			Amount: int(preAmount.Amount) + int(UpdateAmount),
		})
	if err != nil {
		return nil, err
	}
	resp := &product.UpdateProductResp{
		StatusCode: 200,
		StatusMsg:  "update product successfully" + strconv.Itoa(int(Id)),
		Product: &product.Product{
			Id:     Id,
			Amount: int32(preAmount.Amount + int(UpdateAmount)),
		},
	}

	return resp, nil
}

// UpdateProduct 更新商品
func (s *ProductService) UpdateProduct(req *product.UpdateProductReq) (*product.UpdateProductResp, error) {
	productQuery := query.Product
	// 更新商品
	_, err := productQuery.WithContext(s.ctx).Where(productQuery.Id.Eq(req.Id)).Updates(
		dalproduct.Product{
			Model:       gorm.Model{},
			Id:          req.Id,
			Name:        req.Name,
			Amount:      int(req.Amount),
			Description: req.Description,
			Picture:     req.Picture,
			Price:       req.Price,
			Categories:  req.Categories,
		})
	if err != nil {
		return nil, err
	}

	resp := &product.UpdateProductResp{
		StatusCode: 200,
		StatusMsg:  "update product successfully" + strconv.Itoa(int(req.Id)),
		Product: &product.Product{
			Id:          req.Id,
			Name:        req.Name,
			Amount:      req.Amount,
			Description: req.Description,
			Picture:     req.Picture,
			Price:       req.Price,
			Categories:  pq.StringArray(req.Categories),
		},
	}

	return resp, nil
}

// DeleteProduct 删除商品
func (s *ProductService) DeleteProduct(req *product.DeleteProductReq) (*product.DeleteProductResp, error) {
	productQuery := query.Product
	// 删除商品
	_, err := productQuery.WithContext(s.ctx).Where(productQuery.Id.Eq(req.Id)).Delete()
	if err != nil {
		return nil, err
	}

	resp := &product.DeleteProductResp{
		StatusCode: 200,
		StatusMsg:  "Delete product successfully" + strconv.Itoa(int(req.Id)),
		Message:    "Delete product successfully, id: " + strconv.Itoa(int(req.Id)),
	}

	return resp, nil
}

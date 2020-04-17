package grpcfetcher

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
)

type ProductReader struct {
	conn     connection.CatalogConnection
	connAuth connection.AuthConnection
}

func NewProductReader(conn connection.CatalogConnection, connAuth connection.AuthConnection) *ProductReader {
	return &ProductReader{conn: conn, connAuth: connAuth}
}

func (pr *ProductReader) Search(ctx context.Context, req requests.ListProduct) ([]*data.Product, error) {
	var shopIDs []int64
	var err error

	if req.Province != "" {
		shopIDs, err = pr.getShopIDsByProvince(req, ctx)
		if err != nil {
			return nil, err
		}
		if shopIDs == nil || len(shopIDs) == 0 {
			return nil, nil
		}
	}

	res, err := pr.conn.ListProducts(ctx, &catalogproto.ListProductsReq{
		ShopIDs:   shopIDs,
		Category:  req.Category,
		Search:    req.Search,
		OrderBy:   req.OrderBy,
		OrderType: req.OrderType,
		Pagination: &catalogproto.Pagination{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})
	if err != nil || res == nil {
		return nil, err
	}
	return data.ProductsFromProtos(res.GetProducts()), nil
}

func (pr *ProductReader) Get(ctx context.Context, req requests.GetProduct) (*data.Product, error) {
	res, err := pr.conn.GetProduct(context.Background(), &catalogproto.GetProductReq{
		Id:   req.ID,
		Slug: req.SlugName,
	})
	if err != nil || res == nil {
		return nil, err
	}
	return data.ProductFromProto(res), nil
}

func (pr *ProductReader) GetByShopID(ctx context.Context, req requests.ProductsByShop) ([]*data.Product, error) {
	res, err := pr.conn.ListProducts(ctx, &catalogproto.ListProductsReq{ShopIDs: []int64{int64(req.ShopID)}})
	if err != nil || res == nil {
		return nil, err
	}
	return data.ProductsFromProtos(res.GetProducts()), nil
}

func (pr *ProductReader) getShopIDsByProvince(req requests.ListProduct, ctx context.Context) ([]int64, error) {
	var shopIDs []int64

	shops, err := pr.connAuth.GetShopByProvince(ctx, &authproto.ProvinceReq{Province: req.Province})
	if err != nil {
		return nil, err
	}
	shopIDs = make([]int64, len(shops.GetUsers()))
	for i := range shops.GetUsers() {
		shopIDs[i] = shops.GetUsers()[i].GetId()
	}
	return shopIDs, nil
}

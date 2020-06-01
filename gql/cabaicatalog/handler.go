package cabaicatalog

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/fetcher"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/mutator"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/resolver"
	"github.com/idzharbae/cabai-gqlserver/util"
	"strconv"
)

type CabaiCatalogHandler struct {
	productReader fetcher.ProductReader
	productWriter mutator.ProductWriter
	reviewReader  fetcher.ReviewReader
	reviewWriter  mutator.ReviewWriter
}

func NewCabaiCatalogHandler(productReader fetcher.ProductReader, productWriter mutator.ProductWriter,
	reviewReader fetcher.ReviewReader, reviewWriter mutator.ReviewWriter) *CabaiCatalogHandler {
	return &CabaiCatalogHandler{
		productReader: productReader,
		productWriter: productWriter,
		reviewReader:  reviewReader,
		reviewWriter:  reviewWriter,
	}
}

func (r *CabaiCatalogHandler) Review(ctx context.Context, args struct {
	Params requests.GetReview
}) (*resolver.Review, error) {
	res, err := r.reviewReader.Get(ctx, requests.GetReview{
		ID:         args.Params.ID,
		CustomerID: args.Params.CustomerID,
		ProductID:  args.Params.ProductID,
	})
	if err != nil {
		return nil, err
	}
	return resolver.NewReview(res), nil
}

func (r *CabaiCatalogHandler) Reviews(ctx context.Context, args struct {
	Params requests.ListReview
}) (*[]*resolver.Review, error) {
	res, err := r.reviewReader.List(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	reviews := resolver.NewReviews(res)
	return &reviews, nil
}

func (r *CabaiCatalogHandler) SearchProducts(ctx context.Context, args struct {
	Params requests.ListProduct
}) (*[]*resolver.Product, error) {
	res, err := r.productReader.Search(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	products := resolver.NewProducts(res)
	return &products, nil
}

func (r *CabaiCatalogHandler) ProductsByShop(ctx context.Context, args struct {
	Params requests.ProductsByShop
}) (*[]*resolver.Product, error) {
	res, err := r.productReader.GetByShopID(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	products := resolver.NewProducts(res)
	return &products, nil
}

func (r *CabaiCatalogHandler) Product(ctx context.Context, args struct {
	Params requests.GetProduct
}) (*resolver.Product, error) {
	res, err := r.productReader.Get(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	product := resolver.NewProduct(res)
	return product, nil
}

func (r *CabaiCatalogHandler) CreateProduct(ctx context.Context, args struct {
	Params requests.CreateProduct
}) (*resolver.Product, error) {
	res, err := r.productWriter.CreateProduct(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	product := resolver.NewProduct(res)
	return product, nil
}

func (r *CabaiCatalogHandler) UpdateProduct(ctx context.Context, args struct {
	Params requests.UpdateProduct
}) (*resolver.Product, error) {
	res, err := r.productWriter.UpdateProduct(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	product := resolver.NewProduct(res)
	return product, nil
}

func (r *CabaiCatalogHandler) DeleteProduct(ctx context.Context, args struct {
	Params requests.GetProduct
}) (*resolver.Success, error) {
	err := r.productWriter.DeleteProduct(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	return &resolver.Success{}, nil
}

func (r *CabaiCatalogHandler) CreateReview(ctx context.Context, args struct {
	Params requests.CreateReview
}) (*resolver.Review, error) {
	userID, err := r.getUserID("", ctx)
	if err != nil {
		return nil, err
	}
	args.Params.UserID = strconv.FormatInt(userID, 10)

	res, err := r.reviewWriter.Create(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	return resolver.NewReview(res), nil
}

func (r *CabaiCatalogHandler) UpdateReview(ctx context.Context, args struct {
	Params requests.UpdateReview
}) (*resolver.Review, error) {
	userID, err := r.getUserID("", ctx)
	if err != nil {
		return nil, err
	}
	args.Params.UserID = strconv.FormatInt(userID, 10)

	res, err := r.reviewWriter.Update(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	return resolver.NewReview(res), nil
}

func (r *CabaiCatalogHandler) DeleteReview(ctx context.Context, args struct {
	Params requests.DeleteReview
}) (*resolver.Success, error) {
	userID, err := r.getUserID("", ctx)
	if err != nil {
		return nil, err
	}
	args.Params.UserID = strconv.FormatInt(userID, 10)

	err = r.reviewWriter.Delete(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	return &resolver.Success{}, nil
}

func (h *CabaiCatalogHandler) getUserID(token string, ctx context.Context) (int64, error) {
	var userID int64
	var err error

	if token != "" {
		userID, err = util.UserIDFromToken(token)
	} else {
		userID, err = util.UserIDFromCtx(ctx)
	}
	if err != nil {
		return 0, err
	}
	return userID, nil
}

package grpcmutator

import (
	"context"
	authdata "github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
	"github.com/idzharbae/cabai-gqlserver/util"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/resources/protoresources"
	"io/ioutil"
	"strings"
)

type ProductWriter struct {
	catalog   connection.CatalogConnection
	resources connection.ResourcesConnection
}

func NewProductWriter(catalog connection.CatalogConnection, resources connection.ResourcesConnection) *ProductWriter {
	return &ProductWriter{catalog: catalog, resources: resources}
}

func (pr *ProductWriter) CreateProduct(ctx context.Context, req requests.CreateProduct) (*data.Product, error) {
	photo, err := pr.getPhoto(ctx, req)
	if err != nil {
		return nil, err
	}

	uploadRes, err := pr.resources.UploadPhoto(ctx, &protoresources.UploadPhotoReq{
		File:    photo.Data,
		FileExt: photo.Extension,
		OwnerId: photo.OwnerID,
	})
	if err != nil {
		return nil, err
	}

	res, err := pr.catalog.CreateProduct(ctx, &catalogproto.Product{
		ShopId:     req.ShopID,
		Name:       req.Name,
		Quantity:   req.Quantity,
		PricePerKg: req.PricePerKG,
		StockKg:    float32(req.StockKG),
		Slug:       req.SlugName,
		PhotoUrl:   uploadRes.GetFileUrl(),
	})
	if err != nil {
		return nil, err
	}
	product := data.ProductFromProto(res)
	return product, nil
}

func (pr *ProductWriter) UpdateProduct(ctx context.Context, req requests.UpdateProduct) (*data.Product, error) {
	res, err := pr.catalog.UpdateProduct(context.Background(), &catalogproto.Product{
		Id:         req.ID,
		ShopId:     req.ShopID,
		Name:       req.Name,
		Quantity:   req.Quantity,
		PricePerKg: req.PricePerKG,
		StockKg:    float32(req.StockKG),
		Slug:       req.SlugName,
	})
	if err != nil {
		return nil, err
	}
	product := data.ProductFromProto(res)
	return product, nil
}

func (pr *ProductWriter) DeleteProduct(ctx context.Context, req requests.GetProduct) error {
	_, err := pr.catalog.DeleteProduct(context.Background(), &catalogproto.GetProductReq{
		Id:   req.ID,
		Slug: req.SlugName,
	})
	return err
}

func (pr *ProductWriter) getPhoto(ctx context.Context, req requests.CreateProduct) (data.File, error) {
	photoStream, err := req.Photo.CreateReadStream()
	if err != nil {
		return data.File{}, err
	}
	photoBytes, err := ioutil.ReadAll(photoStream)
	if err != nil {
		return data.File{}, err
	}
	fileName := strings.Split(req.Photo.FileName, ".")
	fileExt := fileName[len(fileName)-1]

	token, err := util.GetTokenFromContext(ctx)
	if err != nil {
		return data.File{}, err
	}
	owner, err := authdata.UserFromToken(token)
	if err != nil {
		return data.File{}, err
	}

	return data.File{
		Data:      photoBytes,
		Extension: fileExt,
		OwnerID:   owner.ID,
	}, nil
}

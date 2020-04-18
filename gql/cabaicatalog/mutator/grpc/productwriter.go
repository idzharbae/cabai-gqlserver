package grpcmutator

import (
	"context"
	"errors"
	"github.com/idzharbae/cabai-gqlserver/globalconstant"
	authdata "github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
	"github.com/idzharbae/cabai-gqlserver/util"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/resources/protoresources"
	graphqlupload "github.com/smithaitufe/go-graphql-upload"
	"io/ioutil"
	"strconv"
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
	var photoURL string
	shop, err := pr.getShop(ctx)
	if err != nil {
		return nil, err
	}

	if req.Photo != nil {
		photo, err := pr.getPhoto(ctx, req.Photo, shop)
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
		photoURL = uploadRes.GetFileUrl()
	}

	res, err := pr.catalog.CreateProduct(ctx, &catalogproto.Product{
		ShopId:      int32(shop.GetID()),
		Name:        req.Name,
		Quantity:    req.Quantity,
		PricePerKg:  req.PricePerKG,
		StockKg:     float32(req.StockKG),
		Slug:        req.SlugName,
		PhotoUrl:    photoURL,
		Description: req.Description,
		Category:    req.Category,
	})
	if err != nil {
		return nil, err
	}
	product := data.ProductFromProto(res)
	return product, nil
}

func (pr *ProductWriter) UpdateProduct(ctx context.Context, req requests.UpdateProduct) (*data.Product, error) {
	var photoURL string
	shop, err := pr.getShop(ctx)
	if err != nil {
		return nil, err
	}

	if req.Photo != nil {
		photo, err := pr.getPhoto(ctx, req.Photo, shop)
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
		photoURL = uploadRes.GetFileUrl()
	}

	res, err := pr.catalog.UpdateProduct(context.Background(), &catalogproto.Product{
		Id:          req.ID,
		ShopId:      int32(shop.GetID()),
		Name:        req.Name,
		Quantity:    req.Quantity,
		PricePerKg:  req.PricePerKG,
		StockKg:     float32(req.StockKG),
		Slug:        req.SlugName,
		PhotoUrl:    photoURL,
		Description: req.Description,
		Category:    req.Category,
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

func (pr *ProductWriter) getPhoto(ctx context.Context, req *graphqlupload.GraphQLUpload, owner authdata.User) (data.File, error) {
	if req == nil {
		return data.File{}, nil
	}
	photoStream, err := req.CreateReadStream()
	if err != nil {
		return data.File{}, err
	}
	photoBytes, err := ioutil.ReadAll(photoStream)
	if err != nil {
		return data.File{}, err
	}
	fileName := strings.Split(req.FileName, ".")
	fileExt := fileName[len(fileName)-1]

	id, _ := strconv.ParseInt(owner.ID, 10, 64)
	return data.File{
		Data:      photoBytes,
		Extension: fileExt,
		OwnerID:   id,
	}, nil
}

func (pr *ProductWriter) getShop(ctx context.Context) (authdata.User, error) {
	token, err := util.GetTokenFromContext(ctx)
	if err != nil {
		return authdata.User{}, err
	}
	owner, err := authdata.UserFromToken(token)
	if err != nil {
		return authdata.User{}, err
	}
	if owner.Type != globalconstant.ShopType {
		return authdata.User{}, errors.New("user type is not shop")
	}
	return owner, nil
}

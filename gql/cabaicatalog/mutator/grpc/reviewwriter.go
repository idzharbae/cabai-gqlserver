package grpcmutator

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/resources/protoresources"
	graphqlupload "github.com/smithaitufe/go-graphql-upload"
	"io/ioutil"
	"strconv"
	"strings"
)

type ReviewWriter struct {
	catalog   connection.CatalogConnection
	resources connection.ResourcesConnection
}

func NewReviewWriter(catalog connection.CatalogConnection, resources connection.ResourcesConnection) *ReviewWriter {
	return &ReviewWriter{catalog: catalog, resources: resources}
}

func (rw *ReviewWriter) Create(ctx context.Context, req requests.CreateReview) (*data.Review, error) {
	userID, productID, err := rw.parseUserIDAndProductID(req.UserID, req.ProductID)
	if err != nil {
		return nil, err
	}
	product, err := rw.catalog.GetProduct(ctx, &catalogproto.GetProductReq{
		Id: int32(productID),
	})
	if err != nil {
		return nil, err
	}
	photoURL, err := rw.getPhotoURL(ctx, req.Photo, userID)
	if err != nil {
		return nil, err
	}

	res, err := rw.catalog.CreateReview(ctx, &catalogproto.Review{
		UserId:    userID,
		ProductId: productID,
		ShopId:    int64(product.ShopId),
		Title:     req.Title,
		Content:   req.Content,
		PhotoUrl:  photoURL,
		Rating:    req.Rating,
	})
	if err != nil {
		return nil, err
	}
	return data.ReviewFromProto(res), nil
}

func (rw *ReviewWriter) Update(ctx context.Context, req requests.UpdateReview) (*data.Review, error) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	userID, productID, err := rw.parseUserIDAndProductID(req.UserID, req.ProductID)
	if err != nil {
		return nil, err
	}

	product, err := rw.catalog.GetProduct(ctx, &catalogproto.GetProductReq{
		Id: int32(productID),
	})
	if err != nil {
		return nil, err
	}

	photoURL, err := rw.getPhotoURL(ctx, req.Photo, userID)
	if err != nil {
		return nil, err
	}

	res, err := rw.catalog.UpdateReview(ctx, &catalogproto.Review{
		Id:        id,
		UserId:    userID,
		ProductId: productID,
		ShopId:    int64(product.ShopId),
		Title:     req.Title,
		Content:   req.Content,
		PhotoUrl:  photoURL,
		Rating:    req.Rating,
	})
	if err != nil {
		return nil, err
	}
	return data.ReviewFromProto(res), nil
}
func (rw *ReviewWriter) Delete(ctx context.Context, req requests.DeleteReview) error {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		return err
	}
	userID, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return err
	}
	review, err := rw.catalog.GetReview(ctx, &catalogproto.GetReviewReq{
		ReviewId: id,
	})
	if err != nil {
		return err
	}
	if review.PhotoUrl != "" {
		_, err = rw.resources.DeletePhoto(ctx, &protoresources.DeletePhotoReq{
			FileUrl: review.PhotoUrl,
			UserId:  userID,
		})
		if err != nil {
			return err
		}
	}
	_, err = rw.catalog.DeleteReview(ctx, &catalogproto.Review{
		Id:     id,
		UserId: userID,
	})
	if err != nil {
		return err
	}
	return err
}

func (rw *ReviewWriter) getPhoto(ctx context.Context, req *graphqlupload.GraphQLUpload, ownerID int64) (data.File, error) {
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

	return data.File{
		Data:      photoBytes,
		Extension: fileExt,
		OwnerID:   ownerID,
	}, nil
}

func (rw *ReviewWriter) parseUserIDAndProductID(userIDstr string, productIDstr string) (int64, int64, error) {
	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	productID, err := strconv.ParseInt(productIDstr, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return userID, productID, nil
}

func (rw *ReviewWriter) getPhotoURL(ctx context.Context, photo *graphqlupload.GraphQLUpload, userID int64) (string, error) {
	photoURL := ""
	if photo != nil {
		photo, err := rw.getPhoto(ctx, photo, userID)
		if err != nil {
			return "", err
		}

		uploadRes, err := rw.resources.UploadPhoto(ctx, &protoresources.UploadPhotoReq{
			File:    photo.Data,
			FileExt: photo.Extension,
			OwnerId: photo.OwnerID,
		})
		if err != nil {
			return "", err
		}
		photoURL = uploadRes.GetFileUrl()
	}
	return photoURL, nil
}

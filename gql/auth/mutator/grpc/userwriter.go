package grpc

import (
	"context"

	"github.com/idzharbae/cabai-gqlserver/gql/auth/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/requests"
	"github.com/idzharbae/cabai-gqlserver/util"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/resources/protoresources"
	graphqlupload "github.com/smithaitufe/go-graphql-upload"
	"io/ioutil"
	"strconv"
	"strings"
)

type UserMutator struct {
	conn      connection.Connection
	catalog   connection.CatalogConnection
	resources connection.ResourcesConnection
}

func NewUserMutator(conn connection.Connection, catalog connection.CatalogConnection, resources connection.ResourcesConnection) *UserMutator {
	return &UserMutator{conn: conn, catalog: catalog, resources: resources}
}

func (um *UserMutator) Register(ctx context.Context, req requests.Register) (*data.User, error) {
	user, err := um.conn.Register(ctx, &authproto.RegisterReq{
		UserName:      req.UserName,
		Email:         req.Email,
		Phone:         req.PhoneNumber,
		Password:      req.Password,
		Type:          req.Role,
		FullName:      req.FullName,
		City:          req.City,
		Province:      req.Province,
		AddressDetail: req.AddressDetail,
		ZipCode:       req.ZipCode,
		Description:   req.Description,
	})
	if err != nil {
		return nil, err
	}
	userData := data.UserFromProto(user)
	return userData, nil
}

func (um *UserMutator) EditProfile(ctx context.Context, req requests.EditProfile) (*data.User, error) {
	var photoURL string
	user, err := um.getUser(ctx)
	if err != nil {
		return nil, err
	}

	if req.Photo != nil {
		photo, err := um.getPhoto(ctx, req.Photo, user)
		if err != nil {
			return nil, err
		}

		uploadRes, err := um.resources.UploadPhoto(ctx, &protoresources.UploadPhotoReq{
			File:    photo.Data,
			FileExt: photo.Extension,
			OwnerId: photo.OwnerID,
		})
		if err != nil {
			return nil, err
		}
		photoURL = uploadRes.GetFileUrl()
	}
	userID, _ := strconv.ParseInt(user.ID, 10, 64)
	updatedUser, err := um.conn.UpdateUser(ctx, &authproto.User{
		Id:            userID,
		Phone:         req.PhoneNumber,
		Password:      req.Password,
		NewPassword:   req.NewPassword,
		Type:          user.Type,
		Name:          req.FullName,
		City:          req.City,
		Province:      req.Province,
		PhotoUrl:      photoURL,
		AddressDetail: req.AddressDetail,
		ZipCode:       req.ZipCode,
		Description:   req.Description,
	})

	if err != nil {
		return nil, err
	}
	userData := data.UserFromProto(updatedUser)
	return userData, nil
}

func (um *UserMutator) getPhoto(ctx context.Context, req *graphqlupload.GraphQLUpload, owner data.User) (data.File, error) {
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

func (um *UserMutator) getUser(ctx context.Context) (data.User, error) {
	token, err := util.GetTokenFromContext(ctx)
	if err != nil {
		return data.User{}, err
	}
	owner, err := data.UserFromToken(token)
	if err != nil {
		return data.User{}, err
	}
	return owner, nil
}

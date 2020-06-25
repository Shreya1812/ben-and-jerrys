package controller

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/convertor"
	"github.com/Shreya1812/ben-and-jerrys/internal/commons"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"github.com/Shreya1812/ben-and-jerrys/internal/proto/icecream"
)

type iceCreamControllerImpl struct {
	s service.IceCreamService
}

func New(c *configs.Config) (IceCreamController, error) {
	s, err := service.New(c)

	if err != nil {
		return nil, err
	}
	return &iceCreamControllerImpl{
		s: s,
	}, nil
}

func (i *iceCreamControllerImpl) Create(ctx context.Context, request *icecream_pb.CreateRequest) (*icecream_pb.CreateResponse, error) {
	m, err := i.s.CreateIceCream(ctx, convertor.PbToModel(request.GetIceCream()))

	if err != nil {
		return nil, commons.GetErrorWithStatus(err)
	}

	return &icecream_pb.CreateResponse{
		IceCream: convertor.ModelToPb(m),
	}, nil
}

func (i *iceCreamControllerImpl) Update(ctx context.Context, request *icecream_pb.UpdateRequest) (*icecream_pb.UpdateResponse, error) {
	m, err := i.s.UpdateIceCream(ctx, convertor.PbToModel(request.GetIceCream()))

	if err != nil {
		return nil, commons.GetErrorWithStatus(err)
	}

	return &icecream_pb.UpdateResponse{IceCream: convertor.ModelToPb(m)}, nil
}

func (i *iceCreamControllerImpl) Delete(ctx context.Context, request *icecream_pb.DeleteRequest) (*icecream_pb.DeleteResponse, error) {
	m, err := i.s.DeleteIceCreamById(ctx, request.GetId())

	if err != nil {
		return nil, commons.GetErrorWithStatus(err)
	}

	return &icecream_pb.DeleteResponse{IceCream: convertor.ModelToPb(m)}, nil
}

func (i *iceCreamControllerImpl) GetById(ctx context.Context, request *icecream_pb.GetByIdRequest) (*icecream_pb.GetByIdResponse, error) {
	m, err := i.s.GetIceCreamById(ctx, request.GetId())

	if err != nil {
		return nil, commons.GetErrorWithStatus(err)
	}

	return &icecream_pb.GetByIdResponse{IceCream: convertor.ModelToPb(m)}, nil
}

func (i *iceCreamControllerImpl) Close() error {
	return i.s.Close()
}

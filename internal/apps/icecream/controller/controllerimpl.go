package controller

import (
	"context"
	"fmt"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/convertor"
	"github.com/Shreya1812/ben-and-jerrys/internal/proto/icecream"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type iceCreamControllerImpl struct {
	s service.IceCreamService
}

func New(s service.IceCreamService) IceCreamController {
	return &iceCreamControllerImpl{
		s: s,
	}
}

func (i *iceCreamControllerImpl) Create(ctx context.Context, request *icecream_pb.CreateRequest) (*icecream_pb.CreateResponse, error) {
	m, err := i.s.CreateIceCream(ctx, convertor.PbToModel(request.GetIceCream()))

	if err != nil {
		return nil, err
	}

	return &icecream_pb.CreateResponse{
		IceCream: convertor.ModelToPb(m),
	}, nil
}

func (i *iceCreamControllerImpl) Update(ctx context.Context, request *icecream_pb.UpdateRequest) (*icecream_pb.UpdateResponse, error) {
	m, err := i.s.UpdateIceCream(ctx, convertor.PbToModel(request.GetIceCream()))

	if err != nil {
		return nil, err
	}

	return &icecream_pb.UpdateResponse{IceCream: convertor.ModelToPb(m)}, nil
}

func (i *iceCreamControllerImpl) Delete(ctx context.Context, request *icecream_pb.DeleteRequest) (*icecream_pb.DeleteResponse, error) {
	m, err := i.s.DeleteIceCreamById(ctx, request.GetId())

	if err != nil {
		return nil, err
	}

	return &icecream_pb.DeleteResponse{IceCream: convertor.ModelToPb(m)}, nil
}

func (i *iceCreamControllerImpl) GetById(ctx context.Context, request *icecream_pb.GetByIdRequest) (*icecream_pb.GetByIdResponse, error) {
	m, err := i.s.GetIceCreamById(ctx, request.GetId())

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}

	return &icecream_pb.GetByIdResponse{IceCream: convertor.ModelToPb(m)}, nil
}

func (i *iceCreamControllerImpl) Close() error {
	return i.s.Close()
}

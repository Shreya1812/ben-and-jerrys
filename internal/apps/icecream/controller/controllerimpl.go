package controller

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/convertor"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/model"
	"github.com/Shreya1812/ben-and-jerrys/internal/commons"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"github.com/Shreya1812/ben-and-jerrys/internal/proto/icecream"
	"github.com/golang/protobuf/proto"
	"log"
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
	log.Println(">>>>> Creating IceCream")
	m, err := i.s.CreateIceCream(ctx, convertor.PbToModel(request.GetIceCream()))

	if err != nil {
		log.Printf("Error creating IceCream %+v", err)
		return nil, commons.GetErrorWithStatus(err)
	}

	log.Println(">>>>> IceCream Created")
	return &icecream_pb.CreateResponse{
		IceCream: convertor.ModelToPb(m),
	}, nil
}

func (i *iceCreamControllerImpl) Update(ctx context.Context, request *icecream_pb.UpdateRequest) (*icecream_pb.UpdateResponse, error) {
	log.Println(">>>>> Updating IceCream")
	m, err := i.s.UpdateIceCream(ctx, convertor.PbToModel(request.GetIceCream()))

	if err != nil {
		fmt.Printf("Error updating IceCream %+v", err)
		return nil, commons.GetErrorWithStatus(err)
	}

	log.Println(">>>>> IceCream Updated")
	return &icecream_pb.UpdateResponse{IceCream: convertor.ModelToPb(m)}, nil
}

func (i *iceCreamControllerImpl) DeleteByProductId(ctx context.Context, request *icecream_pb.DeleteByProductIdRequest) (*icecream_pb.DeleteByProductIdResponse, error) {
	log.Println(">>>>> Deleting IceCream")
	m, err := i.s.DeleteIceCreamByProductId(ctx, request.GetProductId())

	if err != nil {
		fmt.Printf("Error deleting IceCream %+v", err)
		return nil, commons.GetErrorWithStatus(err)
	}

	log.Println(">>>>> IceCream Deleted")
	return &icecream_pb.DeleteByProductIdResponse{IceCream: convertor.ModelToPb(m)}, nil
}

func (i *iceCreamControllerImpl) GetByProductId(ctx context.Context, request *icecream_pb.GetByProductIdRequest) (*icecream_pb.GetByProductIdResponse, error) {
	log.Println(">>>>> Fetching IceCream by ProductId")
	m, err := i.s.GetIceCreamByProductId(ctx, request.GetProductId())

	if err != nil {
		fmt.Printf("Error fetching IceCream %+v", err)
		return nil, commons.GetErrorWithStatus(err)
	}

	log.Println(">>>>> Fetching IceCream by ProductId Completed")
	return &icecream_pb.GetByProductIdResponse{IceCream: convertor.ModelToPb(m)}, nil
}

func (i *iceCreamControllerImpl) GetList(ctx context.Context, request *icecream_pb.ListRequest) (*icecream_pb.ListResponse, error) {
	log.Println(">>>>> Fetching IceCream List")

	// Base64 Decoding and Un-marshaling the lastId
	lastId := ""
	if request.PaginationContext != "" {
		b, err := base64.RawStdEncoding.DecodeString(request.PaginationContext)

		if err != nil {
			log.Printf("Error decoding PaginationContext %+v", err)
			return nil, err
		}

		pc := &icecream_pb.PaginationContext{}
		if err := proto.Unmarshal(b, pc); err != nil {
			log.Printf("Error unmarshaling PaginationContext %+v", err)
			return nil, commons.GetErrorWithStatus(err)
		}
		lastId = pc.LastId
	}

	m, err := i.s.GetIceCreamList(ctx, &model.IceCreamSearchOptions{
		LastId: lastId,
		Limit:  request.Limit,
	})

	if err != nil {
		log.Printf("Error fetching IceCream list %+v", err)
		return nil, commons.GetErrorWithStatus(err)
	}

	res := convertor.ListModelToListPb(m.IceCreams)

	// Marshaling the lastId and base64 Encoding
	paginationCtxStr := ""
	if m.LastId != "" {
		pc := &icecream_pb.PaginationContext{LastId: m.LastId}

		b, err := proto.Marshal(pc)

		if err != nil {
			log.Printf("Error marshaling PaginationContext %+v", err)
			return nil, commons.GetErrorWithStatus(err)
		}

		paginationCtxStr = base64.RawStdEncoding.EncodeToString(b)
	}

	log.Println(">>>>> Fetching IceCream List Completed")
	return &icecream_pb.ListResponse{
		IceCreams:         res,
		PaginationContext: paginationCtxStr,
	}, nil
}

func (i *iceCreamControllerImpl) Close() error {
	return i.s.Close()
}

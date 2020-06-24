package convertor

import (
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/model"
	"github.com/Shreya1812/ben-and-jerrys/internal/proto/icecream"
)

func PbToModel(pb *icecream_pb.IceCream) *model.IceCream {
	return &model.IceCream{
		ProductId:             pb.ProductId,
		Name:                  pb.Name,
		ImageClosed:           pb.ImageClosed,
		ImageOpen:             pb.ImageOpen,
		Description:           pb.Description,
		Story:                 pb.Story,
		SourcingValues:        pb.SourcingValues,
		Ingredients:           pb.Ingredients,
		AllergyInfo:           pb.AllergyInfo,
		DietaryCertifications: pb.DietaryCertifications,
	}
}

func ModelToPb(model *model.IceCream) *icecream_pb.IceCream {
	return &icecream_pb.IceCream{
		ProductId:             model.ProductId,
		Name:                  model.Name,
		ImageClosed:           model.ImageClosed,
		ImageOpen:             model.ImageOpen,
		Description:           model.Description,
		Story:                 model.Story,
		SourcingValues:        model.SourcingValues,
		Ingredients:           model.Ingredients,
		AllergyInfo:           model.AllergyInfo,
		DietaryCertifications: model.DietaryCertifications,
	}
}

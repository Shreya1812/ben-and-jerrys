package convertor

import (
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/model"
	"github.com/Shreya1812/ben-and-jerrys/internal/proto/icecream"
	"strings"
)

func PbToModel(pb *icecream_pb.IceCream) *model.IceCream {
	return &model.IceCream{
		ProductId:             strings.Trim(pb.ProductId, " "),
		Name:                  strings.Trim(pb.Name, " "),
		ImageClosed:           strings.Trim(pb.ImageClosed, " "),
		ImageOpen:             strings.Trim(pb.ImageOpen, " "),
		Description:           strings.Trim(pb.Description, " "),
		Story:                 strings.Trim(pb.Story, " "),
		SourcingValues:        trimStringSlices(pb.SourcingValues),
		Ingredients:           trimStringSlices(pb.Ingredients),
		AllergyInfo:           strings.Trim(pb.AllergyInfo, " "),
		DietaryCertifications: strings.Trim(pb.DietaryCertifications, " "),
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

func ListModelToListPb(models []*model.IceCream) []*icecream_pb.IceCream {
	result := make([]*icecream_pb.IceCream, 0, len(models))

	for _, ic := range models {
		m := ModelToPb(ic)
		result = append(result, m)
	}
	return result
}

func trimStringSlices(s []string) []string {
	var res = make([]string, 0, len(s))

	for _, v := range s {
		trimmed := strings.Trim(v, " ")

		if len(trimmed) > 0 {
			res = append(res, trimmed)
		}
	}

	return res
}

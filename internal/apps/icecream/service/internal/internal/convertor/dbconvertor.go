package convertor

import (
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/internal/data"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/model"
)

func ModelToData(m *model.IceCream) *data.IceCream {
	return &data.IceCream{
		ProductId:             m.ProductId,
		Name:                  m.Name,
		ImageClosed:           m.ImageClosed,
		ImageOpen:             m.ImageOpen,
		Description:           m.Description,
		Story:                 m.Story,
		SourcingValues:        m.SourcingValues,
		Ingredients:           m.Ingredients,
		AllergyInfo:           m.AllergyInfo,
		DietaryCertifications: m.DietaryCertifications,
	}
}

func DataToModel(d *data.IceCream) *model.IceCream {
	return &model.IceCream{
		ProductId:             d.ProductId,
		Name:                  d.Name,
		ImageClosed:           d.ImageClosed,
		ImageOpen:             d.ImageOpen,
		Description:           d.Description,
		Story:                 d.Story,
		SourcingValues:        d.SourcingValues,
		Ingredients:           d.Ingredients,
		AllergyInfo:           d.AllergyInfo,
		DietaryCertifications: d.DietaryCertifications,
	}
}

func ListDataToListModel(d []*data.IceCream) []*model.IceCream {
	result := make([]*model.IceCream, 0, len(d))

	for _, ic := range d {
		m := DataToModel(ic)
		result = append(result, m)
	}
	return result
}

func OptionModelToData(m *model.IceCreamSearchOptions) *data.IceCreamSearchOptions {
	return &data.IceCreamSearchOptions{
		LastId: m.LastId,
		Limit:  m.Limit,
	}
}

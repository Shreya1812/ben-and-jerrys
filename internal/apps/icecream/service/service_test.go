package service

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/dataservice"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/model"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestIceCreamServiceImpl_CreateIceCream(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDataSvc := dataservice.NewMockIceCreamDataService(ctrl)
	svc := &iceCreamServiceImpl{
		ds: mockDataSvc,
		v:  model.GetIceCreamValidator(),
	}

	ic := GenerateIceCreamModel(WithProductId("1234"))

	mockDataSvc.EXPECT().Create(gomock.Any(), ic).Return(ic, nil)

	res, err := svc.CreateIceCream(context.Background(), ic)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(res, ic) {
		t.Errorf("Expected %+v \n Got %+v", ic, res)
	}
}

func TestIceCreamServiceImpl_UpdateIceCream(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDataSvc := dataservice.NewMockIceCreamDataService(ctrl)
	svc := &iceCreamServiceImpl{
		ds: mockDataSvc,
		v:  model.GetIceCreamValidator(),
	}

	ic := GenerateIceCreamModel(
		WithName("Name Updated"),
		WithDescription("Description Updated"),
		WithStory("Story Updated"),
	)

	mockDataSvc.EXPECT().Update(gomock.Any(), ic).Return(ic, nil)

	res, err := svc.UpdateIceCream(context.Background(), ic)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(res, ic) {
		t.Errorf("Expected %+v \n Got %+v", ic, res)
	}
}

func TestIceCreamServiceImpl_DeleteIceCreamById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDataSvc := dataservice.NewMockIceCreamDataService(ctrl)
	svc := &iceCreamServiceImpl{
		ds: mockDataSvc,
		v:  model.GetIceCreamValidator(),
	}

	ic := GenerateIceCreamModel(WithProductId("1234"))
	mockDataSvc.EXPECT().DeleteByProductId(gomock.Any(), "1234").Return(ic, nil)

	res, err := svc.DeleteIceCreamByProductId(context.Background(), "1234")
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(res, ic) {
		t.Errorf("Expected %+v \n Got %+v", ic, res)
	}
}

func TestIceCreamServiceImpl_GetIceCreamById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDataSvc := dataservice.NewMockIceCreamDataService(ctrl)
	svc := &iceCreamServiceImpl{
		ds: mockDataSvc,
		v:  model.GetIceCreamValidator(),
	}

	ic := GenerateIceCreamModel(WithProductId("1234"))

	mockDataSvc.EXPECT().GetByProductId(gomock.Any(), "1234").Return(ic, nil)

	res, err := svc.GetIceCreamByProductId(context.Background(), "1234")
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(res, ic) {
		t.Errorf("Expected %+v \n Got %+v", ic, res)
	}
}

func TestIceCreamServiceImpl_GetIceCreamList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDataSvc := dataservice.NewMockIceCreamDataService(ctrl)
	svc := &iceCreamServiceImpl{
		ds: mockDataSvc,
		v:  model.GetIceCreamValidator(),
	}

	ic1 := GenerateIceCreamModel(WithProductId("1234"), WithName("1234"))
	ic2 := GenerateIceCreamModel(WithProductId("1235"), WithName("1235"))

	so := &model.IceCreamSearchOptions{
		Limit: 2,
	}

	ics := &model.IceCreamListResult{
		IceCreams: []*model.IceCream{ic1, ic2},
		LastId:    "",
	}

	mockDataSvc.EXPECT().GetList(gomock.Any(), so).Return(ics, nil)

	res, err := svc.GetIceCreamList(context.Background(), so)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(res, ics) {
		t.Errorf("Expected %+v \n Got %+v", ics, res)
	}
}

func GenerateIceCreamModel(options ...Option) *model.IceCream {
	ic := &model.IceCream{
		ProductId:             "1234",
		Name:                  "Name",
		ImageClosed:           "ImageClosed",
		ImageOpen:             "ImageOpen",
		Description:           "Description",
		Story:                 "Story",
		SourcingValues:        []string{"SourcingValues"},
		Ingredients:           []string{"Ingredients"},
		AllergyInfo:           "AllergyInfo",
		DietaryCertifications: "DietaryCertifications",
	}
	for _, op := range options {
		ic = op(ic)
	}

	return ic
}

func WithProductId(productId string) Option {
	return func(in *model.IceCream) *model.IceCream {
		in.ProductId = productId
		return in
	}
}

func WithName(name string) Option {
	return func(in *model.IceCream) *model.IceCream {
		in.Name = name
		return in
	}
}

func WithDescription(description string) Option {
	return func(in *model.IceCream) *model.IceCream {
		in.Description = description
		return in
	}
}

func WithStory(story string) Option {
	return func(in *model.IceCream) *model.IceCream {
		in.Story = story
		return in
	}
}

type Option func(in *model.IceCream) *model.IceCream

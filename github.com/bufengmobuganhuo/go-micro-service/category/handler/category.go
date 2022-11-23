package handler

import (
	"context"
	"github.com/bufengmobuganhuo/go-micro-service/category/common"
	"github.com/bufengmobuganhuo/go-micro-service/category/domain/model"
	"github.com/bufengmobuganhuo/go-micro-service/category/domain/service"
	"github.com/bufengmobuganhuo/go-micro-service/category/proto/category"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

func (c Category) CreateCategory(ctx context.Context, req *category.CategoryRequest,
	opts ...client.CallOption) (resp *category.CreateCategoryResponse, err error) {
	category := &model.Category{}
	err = common.SwapTo(req, category)
	if err != nil {
		return nil, err
	}
	categoryId, err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return nil, err
	}
	resp.CategoryId = categoryId
	resp.Message = "分类添加成功"
	return resp, nil
}

func (c Category) UpdateCategory(ctx context.Context, req *category.CategoryRequest,
	opts ...client.CallOption) (resp *category.UpdateCategoryResponse, err error) {
	category := &model.Category{}
	err = common.SwapTo(req, category)
	if err != nil {
		return nil, err
	}
	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil {
		return nil, err
	}
	resp.Message = "分类更新成功"
	return resp, nil
}

func (c Category) DeleteCategory(ctx context.Context, req *category.DeleteCategoryRequest,
	opts ...client.CallOption) (resp *category.DeleteCategoryResponse, err error) {
	err = c.CategoryDataService.DeleteCategory(req.CategoryId)
	if err != nil {
		return nil, err
	}
	resp.Message = "分类删除成功"
	return resp, nil
}

func (c Category) FindCategoryByName(ctx context.Context, req *category.FindByNameRequest,
	opts ...client.CallOption) (resp *category.CategoryResponse, err error) {
	category, err := c.CategoryDataService.FindCategoryByName(req.CategoryName)
	if err != nil {
		return nil, err
	}
	err = common.SwapTo(category, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c Category) FindCategoryById(ctx context.Context, req *category.FindByIdRequest,
	opts ...client.CallOption) (resp *category.CategoryResponse, err error) {
	category, err := c.CategoryDataService.FindCategoryByID(req.CategoryId)
	if err != nil {
		return nil, err
	}
	err = common.SwapTo(category, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c Category) FindCategoryByLevel(ctx context.Context, req *category.FindByLevelRequest,
	opts ...client.CallOption) (resp *category.FindAllResponse, err error) {
	categories, err := c.CategoryDataService.FindCategoryByLevel(req.CategoryLevel)
	if err != nil {
		return nil, err
	}
	categoryToResponse(categories, resp)
	return resp, nil
}

func (c Category) FindCategoryByParent(ctx context.Context, req *category.FindByParentRequest, opts ...client.CallOption) (
	resp *category.FindAllResponse, err error) {
	categories, err := c.CategoryDataService.FindCategoryByParent(req.CategoryParent)
	if err != nil {
		return nil, err
	}
	categoryToResponse(categories, resp)
	return resp, nil
}

func (c Category) FindAllCategory(ctx context.Context, req *category.FindAllRequest, opts ...client.CallOption) (
	resp *category.FindAllResponse, err error) {
	categories, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return nil, err
	}
	categoryToResponse(categories, resp)
	return resp, nil
}

func categoryToResponse(categories []model.Category, resp *category.FindAllResponse) {
	for _, cg := range categories {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		resp.Category = append(resp.Category, cr)
	}
}

package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"mxshop/goods_web/forms"
	"mxshop/goods_web/global"
	"mxshop/goods_web/proto"
	"mxshop/user_web/api"
	"net/http"
	"strconv"
)

func CategoryList(ctx *gin.Context){
	goodsSrvClient := global.GoodsSrvClient
	list, err := goodsSrvClient.GetAllCategorysList(context.Background(), &emptypb.Empty{})
	if err != nil{
		zap.S().Errorw("[CategoryList] Failed to query Goods list")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	res := gin.H{}
	res["total"] = list.Total
	categoryList := make([]interface{}, 0)
	for _, value := range list.Data{
		categoryList = append(categoryList, map[string]interface{}{
			"id": value.Id,
			"name": value.Name,
			"level": value.Level,
			"istab": value.IsTab,
			"parent": value.ParentCategory,
		})
	}
	res["data"] = categoryList
	ctx.JSON(http.StatusOK, res)
}

func CategoryDetail(ctx *gin.Context){
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil{
		ctx.Status(http.StatusNotFound)
		return
	}
	result := gin.H{}
	subCategories := make([]interface{}, 0)
	if lst, err := global.GoodsSrvClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: int32(i),
	});err != nil{
		HandleGrpcErrorToHttp(err, ctx)
		return
	}else{
		for _, value := range lst.SubCategorys{
			subCategories = append(subCategories, map[string]interface{}{
				"id": value.Id,
				"name": value.Name,
				"level": value.Level,
				"parent_category": value.ParentCategory,
				"is_tab": value.IsTab,
			})
		}
		result["id"] = lst.Info.Id
		result["name"] = lst.Info.Name
		result["level"] = lst.Info.Level
		result["parent_category"] = lst.Info.ParentCategory
		result["is_tab"] = lst.Info.IsTab
		result["sub_categorys"] = subCategories
		ctx.JSON(http.StatusOK, result)
	}
	return
}

func NewCategory(ctx *gin.Context){
	categoryForm := forms.CategoryForm{}
	if err := ctx.ShouldBindJSON(&categoryForm);err != nil{
		HandleValidatorError(ctx, err)
		return
	}
	rsp, err := global.GoodsSrvClient.CreateCategory(context.Background(), &proto.CategoryInfoRequest{
		Name:                 categoryForm.Name,
		IsTab:                *categoryForm.IsTab,
		Level: categoryForm.Level,
		ParentCategory: categoryForm.ParentCategory,
	})
	if err != nil{
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	request := make(map[string]interface{})
	request["id"] = rsp.Id
	request["name"] = rsp.Name
	request["parent"] = rsp.ParentCategory
	request["level"] = rsp.Level
	request["is_tab"] = rsp.IsTab

	ctx.JSON(http.StatusOK, request)
}

func DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteCategory(context.Background(), &proto.DeleteCategoryRequest{Id: int32(i)})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}

func UpdateCategory(ctx *gin.Context) {
	categoryForm := forms.UpdateCategoryForm{}
	if err := ctx.ShouldBindJSON(&categoryForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	request := &proto.CategoryInfoRequest{
		Id: int32(i),
		Name: categoryForm.Name,
	}
	if categoryForm.IsTab != nil {
		request.IsTab = *categoryForm.IsTab
	}
	_, err = global.GoodsSrvClient.UpdateCategory(context.Background(), request)
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}
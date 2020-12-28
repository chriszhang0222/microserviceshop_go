package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"mxshop/goods_web/global"
	"net/http"
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

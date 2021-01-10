package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop/userop_web/form"
	"mxshop/userop_web/global"
	"mxshop/userop_web/proto"
	"net/http"
	"strconv"
)

func UserFavList(ctx *gin.Context){
	userId, _ := ctx.Get("userId")
	userFavRsp, err := global.UserFavClient.GetFavList(context.Background(), &proto.UserFavRequest{
		UserId: int32(userId.(uint)),
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ids := make([]int32, 0)
	for _, item := range userFavRsp.Data{
		ids = append(ids, item.GoodsId)
	}
	if len(ids) == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"total":0,
		})
		return
	}
	goods, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: ids,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	reMap := map[string]interface{}{
		"total": userFavRsp.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, item := range userFavRsp.Data{
		data := gin.H{
			"id":item.GoodsId,
		}

		for _, good := range goods.Data {
			if item.GoodsId == good.Id {
				data["name"] = good.Name
				data["shop_price"] = good.ShopPrice
			}
		}

		goodsList = append(goodsList, data)
	}
	reMap["data"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}

func NewUserFav(ctx *gin.Context){
	userId, _ := ctx.Get("userId")
	userFavForm := form.UserFavForm{}
	if err := ctx.ShouldBindJSON(&userFavForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	_, err := global.UserFavClient.AddUserFav(context.Background(), &proto.UserFavRequest{
		UserId: int32(userId.(uint)),
		GoodsId: userFavForm.GoodsId,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})

}

func DeleteUserFav(ctx *gin.Context){
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	userId, _ := ctx.Get("userId")
	_, err = global.UserFavClient.DeleteUserFav(context.Background(), &proto.UserFavRequest{
		UserId: int32(userId.(uint)),
		GoodsId: int32(i),
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func DetailUserFav(ctx *gin.Context){
	goodsId := ctx.Param("id")
	goodsIdInt, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	userId, _ := ctx.Get("userId")
	_, err = global.UserFavClient.GetUserFavDetail(context.Background(), &proto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(goodsIdInt),
	})
	if err != nil {
		zap.S().Errorw("查询收藏状态失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)

}
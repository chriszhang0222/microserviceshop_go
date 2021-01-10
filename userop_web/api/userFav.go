package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"mxshop/userop_web/global"
	"mxshop/userop_web/proto"
	"net/http"
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

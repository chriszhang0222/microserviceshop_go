package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop/goods_web/forms"
	"mxshop/goods_web/global"
	"mxshop/goods_web/proto"
	"net/http"
	"strconv"
)


func List(ctx *gin.Context){
	price_min := ctx.DefaultQuery("pmin", "0")
	priceMin, _ := strconv.Atoi(price_min)
	price_max := ctx.DefaultQuery("pmax", "999999999")
	priceMax, _ := strconv.Atoi(price_max)
	request := &proto.GoodsFilterRequest{
		PriceMin: int32(priceMin),
		PriceMax: int32(priceMax),
	}
	isHot := ctx.DefaultQuery("ih", "0")
	if isHot == "1"{
		request.IsHot = true
	}
	isNew := ctx.DefaultQuery("in", "0")
	if isNew == "1"{
		request.IsNew = true
	}
	isTab := ctx.DefaultQuery("it", "0")
	if isTab == "1"{
		request.IsTab = true
	}
	categoryId := ctx.DefaultQuery("c", "0")
	categoryIdInt, _ := strconv.Atoi(categoryId)
	request.TopCategory = int32(categoryIdInt)

	pages := ctx.DefaultQuery("p", "0")
	pagesInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pagesInt)

	perNums := ctx.DefaultQuery("pnum", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNums = int32(perNumsInt)

	keywords := ctx.DefaultQuery("q", "")
	request.KeyWords = keywords

	brandId := ctx.DefaultQuery("b", "0")
	brandIdInt, _ := strconv.Atoi(brandId)
	request.Brand = int32(brandIdInt)
	goodsSrvClient := global.GoodsSrvClient
	list, err := goodsSrvClient.GoodsList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("[List] Failed to query Goods list")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	res := gin.H{}
	res["total"] = list.Total
	goodsList := make([]interface{}, 0)
	for _, value := range list.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id": value.Id,
			"name":        value.Name,
			"goods_brief": value.GoodsBrief,
			"desc":        value.GoodsDesc,
			"ship_free":   value.ShipFree,
			"images":      value.Images,
			"desc_images": value.DescImages,
			"front_image": value.GoodsFrontImage,
			"shop_price":  value.ShopPrice,
			"category": map[string]interface{}{
				"id":   value.Category.Id,
				"name": value.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   value.Brand.Id,
				"name": value.Brand.Name,
				"logo": value.Brand.Logo,
			},
			"is_hot":  value.IsHot,
			"is_new":  value.IsNew,
			"on_sale": value.OnSale,
		})
	}
	res["data"] = goodsList
	ctx.JSON(http.StatusOK, res)
}

func GoodsDetail(ctx *gin.Context){
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	goodsSrvClient := global.GoodsSrvClient
	r, err := goodsSrvClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: int32(i),
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	rsp := gin.H{
		"id":          r.Id,
		"name":        r.Name,
		"goods_brief": r.GoodsBrief,
		"desc":        r.GoodsDesc,
		"ship_free":   r.ShipFree,
		"images":      r.Images,
		"desc_images": r.DescImages,
		"front_image": r.GoodsFrontImage,
		"shop_price":  r.ShopPrice,
		"category": map[string]interface{}{
			"id":   r.Category.Id,
			"name": r.Category.Name,
		},
		"brand": map[string]interface{}{
			"id":   r.Brand.Id,
			"name": r.Brand.Name,
			"logo": r.Brand.Logo,
		},
		"is_hot":  r.IsHot,
		"is_new":  r.IsNew,
		"on_sale": r.OnSale,
	}
	ctx.JSON(http.StatusOK, rsp)

}


func New(ctx *gin.Context){
	goodsForm := forms.GoodsForm{}
	if err := ctx.ShouldBindJSON(&goodsForm); err != nil{
		HandleValidatorError(ctx, err)
		return
	}
	goodClient := global.GoodsSrvClient
	rsp, err := goodClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		Name: goodsForm.Name,
		GoodsSn: goodsForm.GoodsSn,
		Stocks: goodsForm.Stocks,
		MarketPrice: goodsForm.MarketPrice,
		ShopPrice: goodsForm.ShopPrice,
		GoodsBrief: goodsForm.GoodsBrief,
		ShipFree: *goodsForm.ShipFree,
		Images: goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.BrandId,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, rsp)
}

func Delete(ctx *gin.Context){
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil{
		ctx.Status(http.StatusNotFound)
		return
	}
	goodsClient := global.GoodsSrvClient
	_, err = goodsClient.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{
		Id: int32(i),
	})
	if err != nil{
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
	return
}

func Update(ctx *gin.Context){
	goodsForm := forms.GoodsForm{}
	if err:= ctx.ShouldBindJSON(ctx);err != nil{
		HandleValidatorError(ctx, err)
		return
	}
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if _, err = global.GoodsSrvClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id: int32(i),
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.BrandId,
	});err != nil{
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Update Successfully",
	})
}

func Stocks(ctx *gin.Context){
	id := ctx.Param("id")
	_, err := strconv.ParseInt(id, 10, 32)
	if err != nil{
		ctx.Status(http.StatusNotFound)
		return
	}
	return
}

func UpdateStatus(ctx *gin.Context){
	goodsStatusForm := forms.GoodsStatusForm{}
	if err := ctx.ShouldBindJSON(&goodsStatusForm);err != nil{
		HandleValidatorError(ctx, err)
		return
	}
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if _, err = global.GoodsSrvClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id: int32(i),
		IsHot: *goodsStatusForm.IsNew,
		IsNew: *goodsStatusForm.IsNew,
		OnSale: *goodsStatusForm.OnSale,
	});err != nil{
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Success",
	})
}

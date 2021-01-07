package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"mxshop/order_web/forms"
	"mxshop/order_web/global"
	"mxshop/order_web/proto"
	"net/http"
	"strconv"
)

func OrderList(ctx *gin.Context){
	userId, _ := ctx.Get("userId")
	request := proto.OrderFilterRequest{}
	request.UserId = int32(userId.(uint))
	pages := ctx.DefaultQuery("p", "0")
	pageInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pageInt)

	perNums := ctx.DefaultQuery("pnum", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNums = int32(perNumsInt)
	rsp, err := global.OrderSrvClient.OrderList(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取订单列表失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	reMap := gin.H{
		"total": rsp.Total,
	}
	orderList := make([]interface{}, 0)
	for _, item := range rsp.Data {
		tmpMap := map[string]interface{}{}

		tmpMap["id"] = item.Id
		tmpMap["status"] = item.Status
		tmpMap["pay_type"] = item.PayType
		tmpMap["user"] = item.UserId
		tmpMap["post"] = item.Post
		tmpMap["total"] = item.Total
		tmpMap["address"] = item.Address
		tmpMap["name"] = item.Name
		tmpMap["mobile"] = item.Mobile
		tmpMap["order_sn"] = item.OrderSn
		tmpMap["id"] = item.Id
		tmpMap["add_time"] = item.AddTime

		orderList = append(orderList, tmpMap)
	}
	reMap["data"] = orderList
	ctx.JSON(http.StatusOK, reMap)

}

func NewOrder(ctx *gin.Context){
	orderForm := forms.CreateOrderForm{}
	if err := ctx.ShouldBindJSON(&orderForm); err != nil {
		HandleValidatorError(ctx, err)
	}
	userId, _ := ctx.Get("userId")
	rsp, err := global.OrderSrvClient.CreateOrder(context.Background(), &proto.OrderRequest{
		UserId: int32(userId.(uint)),
		Name: orderForm.Name,
		Mobile: orderForm.Mobile,
		Address: orderForm.Address,
		Post: orderForm.Post,
	})
	if err != nil {
		zap.S().Errorw("新建订单失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	//alipay url generate
	alipayConfig := global.ServerConfig.AliPayInfo
	client, err := alipay.New(alipayConfig.AppID, alipayConfig.PrivateKey, false)
	if err != nil{
		zap.S().Errorw("实例化alipay失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = client.LoadAliPayPublicKey(alipayConfig.AliPublicKey)
	if err != nil {
		zap.S().Errorw("加载alipay公钥失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	var p = alipay.TradePagePay{}
	p.NotifyURL = alipayConfig.NotifyURL
	p.ReturnURL = alipayConfig.ReturnURL
	p.Subject = "MXshop " + rsp.OrderSn
	p.OutTradeNo = rsp.OrderSn
	p.TotalAmount = strconv.FormatFloat(float64(rsp.Total), 'f', 2, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	url, err := client.TradePagePay(p)
	if err != nil {
		zap.S().Errorw("加载alipay支付页面失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id": rsp.Id,
		"alipay_url": url.String(),
	})
}

func OrderDetail(ctx *gin.Context){
	id := ctx.Param("id")
	userId, _ := ctx.Get("userId")
	i, _ := strconv.Atoi(id)
	request := proto.OrderRequest{
		Id: int32(i),
		UserId: int32(userId.(uint)),
	}
	rsp, err := global.OrderSrvClient.OrderDetail(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取订单详情失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	reMap := gin.H{}
	reMap["id"] = rsp.OrderInfo.Id
	reMap["status"] = rsp.OrderInfo.Status
	reMap["user"] = rsp.OrderInfo.UserId
	reMap["post"] = rsp.OrderInfo.Post
	reMap["total"] = rsp.OrderInfo.Total
	reMap["address"] = rsp.OrderInfo.Address
	reMap["name"] = rsp.OrderInfo.Name
	reMap["mobile"] = rsp.OrderInfo.Mobile
	reMap["pay_type"] = rsp.OrderInfo.PayType
	reMap["order_sn"] = rsp.OrderInfo.OrderSn

	goodsList := make([]interface{}, 0)
	for _, item := range rsp.Data {
		tmpMap := gin.H{
			"id": item.GoodsId,
			"name": item.GoodsName,
			"image": item.GoodsImage,
			"price": item.GoodsPrice,
			"nums": item.Nums,
		}

		goodsList = append(goodsList, tmpMap)
	}
	reMap["goods"] = goodsList

	//alipay url generate
	alipayConfig := global.ServerConfig.AliPayInfo
	client, err := alipay.New(alipayConfig.AppID, alipayConfig.PrivateKey, false)
	if err != nil{
		zap.S().Errorw("实例化alipay失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = client.LoadAliPayPublicKey(alipayConfig.AliPublicKey)
	if err != nil {
		zap.S().Errorw("加载alipay公钥失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	var p = alipay.TradePagePay{}
	p.NotifyURL = alipayConfig.NotifyURL
	p.ReturnURL = alipayConfig.ReturnURL
	p.Subject = "MXshop " + rsp.OrderInfo.OrderSn
	p.OutTradeNo = rsp.OrderInfo.OrderSn
	p.TotalAmount = strconv.FormatFloat(float64(rsp.OrderInfo.Total), 'f', 2, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	url, err := client.TradePagePay(p)
	if err != nil {
		zap.S().Errorw("加载alipay支付页面失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	reMap["alipay_url"] = url.String()
	ctx.JSON(http.StatusOK, reMap)
}
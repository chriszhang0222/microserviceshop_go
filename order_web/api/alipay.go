package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"mxshop/order_web/global"
	"mxshop/order_web/proto"
	"net/http"
)

func AlipayNotify(ctx *gin.Context){
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
	noti, err := client.GetTradeNotification(ctx.Request)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
	}
	_ , err = global.OrderSrvClient.UpdateOrderStatus(context.Background(), &proto.OrderStatus{
		OrderSn: noti.OutTradeNo,
		Status: string(noti.TradeStatus),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.String(http.StatusOK, "success")
}

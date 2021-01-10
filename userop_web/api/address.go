package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop/userop_web/form"
	"mxshop/userop_web/global"
	"strconv"

	//"mxshop/userop_web/models"
	"mxshop/userop_web/proto"
	"net/http"
)

func AddressList(ctx * gin.Context){
	request := &proto.AddressRequest{
	}
	//claims, _ := ctx.Get("claims")
	//currentUser := claims.(*models.CustomClaims)
	userId, _ := ctx.Get("userId")
	request.UserId = int32(userId.(uint))
	rsp, err := global.AddressClient.GetAddressList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("Error when get address list")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	reMap := gin.H{
		"total": rsp.Total,
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data{
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["user_id"] = value.UserId
		reMap["province"] = value.Province
		reMap["city"] = value.City
		reMap["district"] = value.District
		reMap["address"] = value.Address
		reMap["signer_name"] = value.SignerName
		reMap["signer_mobile"] = value.SignerMobile

		result = append(result, reMap)
	}
	reMap["data"] = result
	ctx.JSON(http.StatusOK, reMap)
}

func NewAddress(ctx *gin.Context){
	addressForm := form.AddressForm{
	}
	if err := ctx.ShouldBindJSON(&addressForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	userId, _ := ctx.Get("userId")
	rsp, err := global.AddressClient.CreateAddress(context.Background(), &proto.AddressRequest{
		UserId: int32(userId.(uint)),
		Province: addressForm.Province,
		City: addressForm.City,
		District: addressForm.District,
		Address: addressForm.Address,
		SignerName: addressForm.SignerName,
		SignerMobile: addressForm.SignerMobile,
		Country: addressForm.Country,
	})
	if err != nil {
		zap.S().Errorw("新建地址失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":rsp.Id,
	})
}

func DeleteAddr(ctx *gin.Context){
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.AddressClient.DeleteAddress(context.Background(), &proto.AddressRequest{
		Id: int32(i),
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func UpdateAddr(ctx *gin.Context) {
	addressForm := form.AddressForm{}
	if err := ctx.ShouldBindJSON(&addressForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	_, err = global.AddressClient.UpdateAddress(context.Background(), &proto.AddressRequest{
		Id:   int32(i),
		Province: addressForm.Province,
		City: addressForm.City,
		District: addressForm.District,
		Address: addressForm.Address,
		SignerName: addressForm.SignerName,
		SignerMobile: addressForm.SignerMobile,
	})
	if err != nil {
		zap.S().Errorw("更新地址失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}


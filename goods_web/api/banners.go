package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"mxshop/goods_web/forms"
	"mxshop/goods_web/global"
	"mxshop/goods_web/proto"
	"net/http"
	"strconv"
)

func BrandList(ctx *gin.Context){
	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := global.GoodsSrvClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages: int32(pnInt),
		PagePerNums: int32(pSizeInt),
	})
	if err != nil{
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	response := gin.H{}
	response["total"] = rsp.Total
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["name"] = value.Name
		reMap["logo"] = value.Logo
		result = append(result, reMap)
	}
	response["data"] = result
	ctx.JSON(http.StatusOK, response)
}

func NewBrand(ctx *gin.Context){
	brandForm := forms.BrandForm{}
	if err := ctx.ShouldBindJSON(&brandForm);err != nil{
		HandleValidatorError(ctx, err)
		return
	}
	rsp, err := global.GoodsSrvClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	request := gin.H{}
	request["id"] = rsp.Id
	request["name"] = rsp.Name
	request["logo"] = rsp.Logo

	ctx.JSON(http.StatusOK, request)

}
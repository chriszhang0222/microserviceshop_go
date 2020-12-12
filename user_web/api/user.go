package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop/user_web/proto"
	"net/http"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context){
	//grpc code to http status code
	if err != nil{
		if e, ok := status.FromError(err); ok{
			switch e.Code(){
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
				"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "Internal Server Error",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
				"msg": "parameter error",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Other Error",
				})
			}
			return
		}
	}
}

func GetUserList(ctx *gin.Context){
	zap.S().Debug("visit user list")
	ip := "127.0.0.1"
	port := 50058
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil{
		zap.S().Errorw("[GetUserList] connect to user service failed", "msg", err.Error(),)
		return
	}
	userSrvClient := proto.NewUserClient(conn)
	pageInfo := &proto.PageInfo{
		Pn: 0,
		PSize: 0,
	}
	rsp, err := userSrvClient.GetUserList(context.Background(), pageInfo)
	if err != nil{
		zap.S().Errorw("[GetUserList] failed")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		data := make(map[string]interface{})
		data["id"] = value.Id
		data["mobile"] = value.Mobile
		data["nickname"] = value.NickName
		data["birthday"] = value.BirthDay
		data["role"] = value.Role
		data["gender"] = value.Gender
		result = append(result, data)
	}
	ctx.JSON(http.StatusOK, result)

}

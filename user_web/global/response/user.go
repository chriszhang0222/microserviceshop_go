package response

import "time"

type UserResponse struct{
	Id int64 `json:"id"`
	NickName string `json:"nick_name"`
	Mobile string `json:"mobile"`
	Birthday time.Time `json:"birthday"`
	Gender string `json:"gender"`
	Role int32 `json:"role"`
}

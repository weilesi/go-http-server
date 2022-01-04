package professional

import (
	"fmt"
	"net/http"
)

type commonResponse struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func UserIndexHandler(ctx *Context) {
	fmt.Println("user-index-url requesting...")

	ctx.OkJson(&commonResponse{Msg: "User Index..."})
}

func UserLoginHandler(ctx *Context) {
	req := &userReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		_ = ctx.BadRequestJson(&commonResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "bad request:" + err.Error(),
		})
		return
	}

	_ = ctx.OkJson(&commonResponse{
		Data: 1111,
	})
}

type userReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Remark   string `json:"remark"`
}

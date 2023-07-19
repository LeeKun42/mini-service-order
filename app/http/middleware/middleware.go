package middleware

import (
	"errors"
	"github.com/kataras/iris/v12"
	"order/app/rpc/client"
	"strings"
)

func JwtAuthCheck(ctx iris.Context) {
	authorization := ctx.GetHeader("Authorization")
	if authorization == "" {
		ctx.StopWithError(401, errors.New("token无效"))
	}
	authArr := strings.Split(authorization, " ")
	if len(authArr) != 2 {
		ctx.StopWithError(401, errors.New("token无效"))
	}
	tokenString := authArr[1]
	var userId int
	err := client.NewUserClient().CheckJwtToken(tokenString, &userId)
	if err != nil {
		ctx.StopWithError(401, err)
	}
	ctx.Values().Set("user_id", userId)
	ctx.Next()
}

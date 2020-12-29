package middleware

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type res struct {
	Code          string      `json:"code"`
	Data          interface{} `json:"data"`
	Msg           string      `json:"msg"`
	Time          int64       `json:"time"`
	GeneralisedAb string      `json:"generalised_ab"`
}

func ApiResult(ctx echo.Context, err error, code string, msg string, data interface{}) error {
	if code == "" {
		if err != nil {
			code = "S0001"
		} else {
			code = "A0001"
		}
	}
	return ctx.JSON(http.StatusOK, res{
		Code:          code,
		Data:          data,
		Msg:           msg,
		Time:          time.Now().Unix(),
		GeneralisedAb: "",
	})
}

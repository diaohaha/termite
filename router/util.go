package router

import (
	"encoding/json"
	"github.com/labstack/echo"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getIntParam(name string, ctx echo.Context) int64 {
	p := ctx.QueryParam(name)
	if p == "" {
		p = ctx.FormValue(name)
	}
	res, _ := strconv.ParseInt(p, 10, 64)
	return res
}

func getInt64FromCtx(name string, ctx echo.Context) int64 {
	val, _ := GetInt64FromCtx(ctx, name)
	return val
}

func getStringFromCtx(name string, ctx echo.Context) string {
	val, _ := GetStringFromCtx(ctx, name)
	return val
}

func GetInt64FromCtx(ctx echo.Context, key string) (int64, error) {
	val := ctx.Get(key)
	switch val.(type) {
	case json.Number:
		return val.(json.Number).Int64()
	case string:
		return strconv.ParseInt(val.(string), 10, 64)
	case int64:
		return val.(int64), nil
	default:
		//common.Named((common.GetMetaFields(ctx), logs.F{"key": key, "val": val, "type": reflect.TypeOf(val)}).Infoln("GetInt64FromCtx")
	}
	return 0, nil
}

func GetStringFromCtx(ctx echo.Context, key string) (string, error) {
	val := ctx.Get(key)
	switch val.(type) {
	case json.Number:
		return string(val.(json.Number)), nil
	case string:
		return val.(string), nil

	}

	return "", nil
}

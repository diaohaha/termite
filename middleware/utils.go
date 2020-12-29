package middleware

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo"
	"strconv"
)

func getListFromCtx(ctx echo.Context, key string) ([]interface{}, error) {
	val := ctx.Get(key)
	res, _ := val.([]interface{})
	return res, nil
}

func getInt64FromCtx(ctx echo.Context, key string) (int64, error) {
	val := ctx.Get(key)
	switch v := val.(type) {
	case json.Number:
		return v.Int64()
	case string:
		return strconv.ParseInt(v, 10, 64)
	case int64:
		return v, nil
	default:
		return 0, errors.New("not exists")
		//common.Named((common.GetMetaFields(ctx), logs.F{"key": key, "val": val, "type": reflect.TypeOf(val)}).Infoln("GetInt64FromCtx")
	}
}

func getStringFromCtx(ctx echo.Context, key string) (string, error) {
	val := ctx.Get(key)
	switch v := val.(type) {
	case json.Number:
		return string(v), nil
	case string:
		return v, nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	default:
		return "", errors.New("not exists")
	}
}

func GetStringListFromCtx(ctx echo.Context, name string) ([]string, error) {
	val, err := getListFromCtx(ctx, name)
	var res []string = []string{}
	for _, v := range val {
		res = append(res, v.(string))
	}
	return res, err
}

func GetIntListFromCtx(ctx echo.Context, name string) ([]int, error) {
	val, err := getListFromCtx(ctx, name)
	var resStrList []string = []string{}
	var resList []int = []int{}
	for _, v := range val {
		switch v := v.(type) {
		case json.Number:
			resStrList = append(resStrList, string(v))
		default:
			resStrList = append(resStrList, v.(string))
		}
	}
	for _, s := range resStrList {
		resInt, _ := strconv.Atoi(s)
		resList = append(resList, resInt)
	}
	return resList, err
}

func GetInt64ListFromCtx(ctx echo.Context, name string) ([]int64, error) {
	val, err := getListFromCtx(ctx, name)
	var resStrList []string = []string{}
	var resList []int64 = []int64{}
	for _, v := range val {
		switch v := v.(type) {
		case json.Number:
			resStrList = append(resStrList, string(v))
		default:
			resStrList = append(resStrList, v.(string))
		}
	}
	for _, s := range resStrList {
		resInt, _ := strconv.ParseInt(s, 10, 64)
		resList = append(resList, resInt)
	}
	return resList, err
}

func GetInt64FromCtx(ctx echo.Context, name string) int64 {
	val, _ := getInt64FromCtx(ctx, name)
	return val
}

func GetInt64FromCtxDefault(ctx echo.Context, name string, v int64) int64 {
	val, err := getInt64FromCtx(ctx, name)
	if err != nil {
		return v
	}
	return val
}

func GetStringFromCtx(ctx echo.Context, name string) (string, error) {
	val, err := getStringFromCtx(ctx, name)
	return val, err
}

func GetStringFromCtxDefault(ctx echo.Context, name string, v string) string {
	val, err := getStringFromCtx(ctx, name)
	if err != nil {
		return v
	}
	return val
}

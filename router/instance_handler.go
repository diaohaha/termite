package router

import (
	"bufio"
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/middleware"
	"github.com/diaohaha/termite/stable"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type flowInstance struct {
	FlowId     int64             `json:"flow_id"`
	Cid        string            `json:"cid"`
	State      string            `json:"state"`
	Vflow      string            `json:"vflow"`
	Context    map[string]string `json:"context"`
	CreateTime string            `json:"create_time"`
	UpdateTime string            `json:"update_time"`
}

type workInstance struct {
	WorkId     int64                  `json:"work_id"`
	Cid        string                 `json:"cid"`
	State      string                 `json:"state"`
	Vflow      string                 `json:"vflow"`
	Vwork      string                 `json:"vwork"`
	Ouput      map[string]interface{} `json:"output"`
	Error      string                 `json:"error"`
	CreateTime string                 `json:"create_time"`
	UpdateTime string                 `json:"update_time"`
}

func vQueryFlowInstance(ctx echo.Context) error {
	// 参数解析
	var cid string
	var vflow string
	var pageIndex int
	var pageSize int
	var vstates []int
	reqBodyR := ctx.Request().Body
	defer func() {
		_ = reqBodyR.Close()
	}()
	var reqBody []byte
	var err error
	reqBody, err = ioutil.ReadAll(reqBodyR)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	reqDataJs, err := simplejson.NewJson([]byte(reqBody))
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	cid, _ = reqDataJs.Get("cid").String()
	vflow, _ = reqDataJs.Get("vflow").String()
	pageIndex, _ = reqDataJs.Get("page_index").Int()
	pageSize, _ = reqDataJs.Get("page_size").Int()

	vstates, err = middleware.GetIntListFromCtx(ctx, "vstates")
	if err != nil {
		return middleware.ApiResult(ctx, err, "P0001", "参数错误", map[string]int{})
	}

	// 获取
	flows, count, err := db.QueryTermiteFlowInstances(cid, vflow, vstates, pageIndex, pageSize)
	var retFlowInstaces []flowInstance
	for _, flow := range flows {
		c := make(map[string]string, 0)
		_ = json.Unmarshal([]byte(flow.Context), &c)
		retFlowInstaces = append(retFlowInstaces, flowInstance{
			FlowId:     flow.Id,
			Cid:        flow.Cid,
			State:      dal.GetWorkFlowStateZh(flow.Vstate), //strconv.FormatInt(int64(work.Vstate), 10),
			Context:    c,
			Vflow:      flow.Vflow,
			CreateTime: flow.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateTime: flow.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "S0001",
			"msg":  "服务器错误",
		})
	}

	// 返回
	if len(retFlowInstaces) != 0 {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "A0001",
			"data": map[string]interface{}{
				"flows": retFlowInstaces,
				"count": count,
			},
		})
	} else {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "A0001",
			"data": map[string]interface{}{
				"flows": []interface{}{},
				"count": count,
			},
		})
	}
}

func vQueryWorkInstance(ctx echo.Context) error {
	// 参数解析
	var cid string
	var vflow string
	var vwork string
	var pageIndex int
	var pageSize int
	reqBodyR := ctx.Request().Body
	defer func() {
		_ = reqBodyR.Close()
	}()
	var reqBody []byte
	var err error
	reqBody, err = ioutil.ReadAll(reqBodyR)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	reqDataJs, err := simplejson.NewJson([]byte(reqBody))
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	cid, _ = reqDataJs.Get("cid").String()
	vflow, _ = reqDataJs.Get("vflow").String()
	vwork, _ = reqDataJs.Get("vwork").String()
	pageIndex, _ = reqDataJs.Get("page_index").Int()
	pageSize, _ = reqDataJs.Get("page_size").Int()
	if pageSize == 0 {
		pageSize = 100
	}

	// 获取
	works, count, err := db.QueryTermiteWorkInstances(cid, vflow, vwork, pageIndex, pageSize)
	var retWorkInstances []workInstance
	for _, work := range works {
		c := make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(work.Output), &c)
		combineError := work.Error
		if _, ok := c["err"]; ok {
			combineError += c["err"].(string)
		}
		if _, ok := c["error"]; ok {
			combineError += c["error"].(string)
		}
		retWorkInstances = append(retWorkInstances, workInstance{
			WorkId:     work.Id,
			Cid:        work.Cid,
			State:      dal.GetWorkStateZh(work.Vstate), //strconv.FormatInt(int64(work.Vstate), 10),
			Ouput:      c,
			Error:      combineError,
			Vflow:      work.Vflow,
			Vwork:      work.Vwork,
			CreateTime: work.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateTime: work.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "S0001",
			"msg":  "服务器错误",
		})
	}

	// 返回
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": "A0001",
		"data": map[string]interface{}{
			"works": retWorkInstances,
			"count": count,
		},
	})
}

func vAddFlowInstance(ctx echo.Context) (err error) {
	// 批量添加工作流
	var cids []string
	var vflow string
	data := map[string]int{
		"success": 0,
		"fail":    0,
	}
	cids, err = middleware.GetStringListFromCtx(ctx, "cids")
	if err != nil {
		return middleware.ApiResult(ctx, err, "P0001", "参数错误", data)
	}
	vflow, err = middleware.GetStringFromCtx(ctx, "vflow")
	if err != nil {
		return middleware.ApiResult(ctx, err, "P0001", "参数错误", data)
	}

	var vworks []string
	if _, ok := mem.GlobalTermiteConfig.FlowConfigMap[vflow]; ok {
		vworks = mem.GlobalTermiteConfig.FlowConfigMap[vflow].Config.Works
	} else {
		return middleware.ApiResult(ctx, err, "P0001", "not exist", data)
	}

	for _, cid := range cids {
		ierr := db.AddWorkFlowIns(vflow, vworks, cid)
		if ierr != nil {
			data["fail"] = data["fail"] + 1
		} else {
			data["success"] = data["success"] + 1
		}
	}
	return middleware.ApiResult(ctx, err, "", "", data)
}

func vAddFlowInstanceByFile(ctx echo.Context) (err error) {
	var vflow string
	var vworks []string
	data := map[string]int{}
	vflow = ctx.FormValue("vflow")
	if _, ok := mem.GlobalTermiteConfig.FlowConfigMap[vflow]; ok {
		vworks = mem.GlobalTermiteConfig.FlowConfigMap[vflow].Config.Works
	} else {
		return middleware.ApiResult(ctx, err, "P0001", "not exist", data)
	}
	// Source
	log.Println(1)
	file, err := ctx.FormFile("file")
	if err != nil {
		panic(err)
		return err
	}
	log.Println(2)
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func() { _ = src.Close() }()

	log.Println(3)
	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer func() { _ = dst.Close() }()

	log.Println(4)
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	go func() {
		dst, err := os.OpenFile(file.Filename, os.O_RDWR, 0666)
		if err != nil {
			stable.HttpLogger.Error("open file error", zap.Error(err))
		}
		buf := bufio.NewReader(dst)
		for {
			line, err := buf.ReadString('\n')
			line = strings.TrimSpace(line)
			if line != "" {
				err = db.AddWorkFlowIns(vflow, vworks, line)
			}
			fmt.Println(line)
			if err != nil {
				if err == io.EOF {
					fmt.Println("File read ok!")
					break
				} else {
					fmt.Println("Read file error!", err)
				}
			}
		}
	}()

	return middleware.ApiResult(ctx, err, "", "", data)
}

func vDeleteFlowInstance(ctx echo.Context) (err error) {
	// 批量删除工作流
	var flowIds []int64
	data := map[string]int{
		"success": 0,
		"fail":    0,
	}
	flowIds, err = middleware.GetInt64ListFromCtx(ctx, "flow_ids")
	if err != nil {
		return middleware.ApiResult(ctx, err, "P0001", "参数错误", data)
	}

	for _, flowId := range flowIds {
		ierr := db.DeleteFlow(flowId)
		if ierr != nil {
			data["fail"] = data["fail"] + 1
		} else {
			data["success"] = data["success"] + 1
		}
	}

	return middleware.ApiResult(ctx, err, "", "", data)
}

func vRecoverFlowInstance(ctx echo.Context) (err error) {
	// 批量重置工作流
	var flowIds []int64
	data := map[string]int{
		"success": 0,
		"fail":    0,
	}
	flowIds, err = middleware.GetInt64ListFromCtx(ctx, "flow_ids")
	if err != nil {
		return middleware.ApiResult(ctx, err, "P0001", "参数错误", data)
	}

	for _, flowId := range flowIds {
		ierr := db.RecoverWorkFlow(flowId)
		if ierr != nil {
			data["fail"] = data["fail"] + 1
		} else {
			data["success"] = data["success"] + 1
		}
	}

	return middleware.ApiResult(ctx, err, "", "", data)
}

func vRecoverWorkInstance(ctx echo.Context) (err error) {
	// 批量重置工作
	var workIds []int64
	data := map[string]int{
		"success": 0,
		"fail":    0,
	}
	workIds, err = middleware.GetInt64ListFromCtx(ctx, "work_ids")
	if err != nil {
		return middleware.ApiResult(ctx, err, "P0001", "参数错误", data)
	}

	for _, workId := range workIds {
		ierr := db.RecoverWork(workId)
		if ierr != nil {
			data["fail"] = data["fail"] + 1
		} else {
			data["success"] = data["success"] + 1
		}
	}

	return middleware.ApiResult(ctx, err, "", "", data)
}

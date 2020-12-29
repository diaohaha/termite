package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"strings"
)

func BodyParser() echo.MiddlewareFunc {
	// 参数解析
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		resHandler := func(ctx echo.Context) error {
			var err error
			method := ctx.Request().Method

			if method == http.MethodPost {
				m := make(map[string]interface{}, 0)
				// json参数数据
				if buf, err := ioutil.ReadAll(ctx.Request().Body); err == nil {
					// 关键一步
					ctx.Request().Body = ioutil.NopCloser(bytes.NewReader(buf))
					m = make(map[string]interface{}, 0)
					d := json.NewDecoder(strings.NewReader(string(buf)))
					d.UseNumber()
					_ = d.Decode(&m)

					for k, v := range m {
						ctx.Set(k, v)
					}
				}

				err = h(ctx)
			} else {
				err = h(ctx)
			}

			return err
		}
		return resHandler
	}
}

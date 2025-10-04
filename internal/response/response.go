package response

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/NoANameGroup/DAOld-Backend/internal/errorx"
	"github.com/NoANameGroup/DAOld-Backend/pkg/lib"
	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
	"github.com/gin-gonic/gin"
)

// PostProcess 处理http响应, resp要求指针或接口类型
// 在日志中记录本次调用详情, 同时向响应头中注入符合b3规范的链路信息, 主要是trace_id
// 最佳实践:
// - 在controller中调用业务处理, 处理结束后调用PostProcess
func PostProcess(c *gin.Context, req, resp any, err error) {
	log.CtxInfo(c, "[%s] req=%s, resp=%s, err=%v", c.FullPath(), lib.JSONF(req), lib.JSONF(resp), err)

	// 无错, 正常响应
	if err == nil {
		response := makeResponse(resp)
		c.JSON(http.StatusOK, response)
		return
	}

	var ex errorx.Errorx
	if errors.As(err, &ex) { // errorx错误
		StatusCode := http.StatusOK
		c.JSON(StatusCode, &errorx.Errorx{
			Code: ex.Code,
			Msg:  ex.Msg,
		})
	} else { // 常规错误, 状态码500
		log.CtxError(c, "internal error, err=%s", err.Error())
		code := http.StatusInternalServerError
		c.String(code, err.Error())
	}
}

// makeResponse 通过反射构造嵌套格式的响应体
func makeResponse(resp any) map[string]any {
	v := reflect.ValueOf(resp)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil
	}
	// 构建返回数据
	v = v.Elem()
	response := map[string]any{
		"code": v.FieldByName("Code").Int(),
		"msg":  v.FieldByName("Msg").String(),
	}
	data := make(map[string]any)
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldValue := v.Field(i).Interface()

		// 跳过零值字段
		if reflect.ValueOf(fieldValue).IsZero() {
			continue
		}

		// 处理嵌入式结构体
		if field.Anonymous {
			// 如果是嵌入式指针结构体，需要解引用
			if field.Type.Kind() == reflect.Ptr {
				if !v.Field(i).IsNil() {
					embeddedValue := v.Field(i).Elem()
					for j := 0; j < embeddedValue.NumField(); j++ {
						embeddedField := embeddedValue.Type().Field(j)
						// 跳过Code和Msg字段以避免重复
						if embeddedField.Name == "Code" || embeddedField.Name == "Msg" {
							continue
						}
						if jsonTag := embeddedField.Tag.Get("json"); jsonTag != "" {
							data[jsonTag] = embeddedValue.Field(j).Interface()
						} else {
							data[embeddedField.Name] = embeddedValue.Field(j).Interface()
						}
					}
				}
			} else {
				// 非指针嵌入式结构体
				embeddedValue := v.Field(i)
				for j := 0; j < embeddedValue.NumField(); j++ {
					embeddedField := embeddedValue.Type().Field(j)
					// 跳过Code和Msg字段以避免重复
					if embeddedField.Name == "Code" || embeddedField.Name == "Msg" {
						continue
					}
					if jsonTag := embeddedField.Tag.Get("json"); jsonTag != "" {
						data[jsonTag] = embeddedValue.Field(j).Interface()
					} else {
						data[embeddedField.Name] = embeddedValue.Field(j).Interface()
					}
				}
			}
		} else {
			// 处理普通字段
			if jsonTag := field.Tag.Get("json"); jsonTag != "" && field.Name != "Code" && field.Name != "Msg" {
				data[jsonTag] = fieldValue
			}
		}
	}
	if len(data) > 0 {
		response["data"] = data
	}
	return response
}

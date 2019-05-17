package util

import (
	"reflect"
	"zen/model"
)

var (
	CodeOk       = 200
	MessageOk    = "ok"
	CodeError    = 500
	MessageError = "error"
)

func Ok(bean interface{}) map[string]interface{} {
	rst := make(map[string]interface{})
	rst["code"] = CodeOk
	rst["message"] = MessageOk
	rst["data"] = bean
	return rst
}

func Err(bean interface{}) map[string]interface{} {
	rst := make(map[string]interface{})
	rst["code"] = CodeError
	rst["message"] = MessageError
	rst["data"] = bean
	return rst
}

func ErrWithMessage(message string, bean interface{}) map[string]interface{} {
	rst := make(map[string]interface{})
	rst["code"] = CodeError
	rst["message"] = message
	rst["data"] = bean
	return rst
}

func Page(page model.Page, bean interface{}) map[string]interface{} {
	rst := Ok(bean)
	rst["page"] = page
	return rst
}

func PageErr() map[string]interface{} {
	rst := Err(nil)
	rst["page"] = nil
	return rst
}

func BeanToMap(bean interface{}) map[string]interface{} {
	t := reflect.TypeOf(bean)
	v := reflect.ValueOf(bean)
	var data = make(map[string]interface{})

	for idx := 0; idx < t.NumField(); idx++ {
		data[t.Field(idx).Name] = v.Field(idx).Interface()
	}
	return data
}

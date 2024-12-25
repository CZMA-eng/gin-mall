package v1

import (
	"encoding/json"
	"gin_mall_tmp/serializer"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError);ok{
		return serializer.Response{
			Status: 400,
			Msg: "json type doesn't match",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: 400,
		Msg: "parameter error",
		Error: err.Error(),
	}
}
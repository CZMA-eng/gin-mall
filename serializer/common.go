package serializer

type Response struct {
	Status int         `json:"status"` // 状态码
	Data   any `json:"data"`   // 数据内容
	Msg    string      `json:"msg"`    // 消息
	Error  string      `json:"error"`  // 错误信息
}

type TokenData struct {
    User  any `json:"user"`  // 用户信息
    Token string      `json:"token"` // JWT Token
}

type DataList struct {
	Item any `json:"item"`
	Total uint `json:"total"`
}

func  BuildListResponse(items any, total uint) Response {
	return Response{
		Status: 200,
		Data:   DataList{
			Item: items,
			Total: total,
		},
		Msg: "ok",
	}
}
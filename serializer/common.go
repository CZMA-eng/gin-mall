package serializer

type Response struct {
	Status int         `json:"status"` // 状态码
	Data   interface{} `json:"data"`   // 数据内容
	Msg    string      `json:"msg"`    // 消息
	Error  string      `json:"error"`  // 错误信息
}
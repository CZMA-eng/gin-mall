package e

var MsgFlags = map[int] string {
	Success: "ok",
	Error: "fail",
	InvalidParams: "invalid params",

	ErrorExistUser: "user exist", 
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
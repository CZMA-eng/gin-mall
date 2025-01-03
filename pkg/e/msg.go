package e

var MsgFlags = map[int] string {
	Success: "ok",
	Error: "fail",
	InvalidParams: "invalid params",

	ErrorExistUser: "user exist", 
	ErrorFailEncryption: "fail to encrypt",
	ErrorExistUserNotFound: "user not found",
	ErrorNotCompare: "pwd not compare", 
	ErrorAuthToken: "authentication error",
	ErrorAuthCheckTokenTimeout: "token expired",
	ErrorUploadFail: "fail to upload",

	ErrorSendEmail: "fail to send email", 
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
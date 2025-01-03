package e

const (
	Success = 200
	Error = 500
	InvalidParams = 400


	// user module related error
	ErrorExistUser = 30001
	ErrorFailEncryption = 30002
	ErrorExistUserNotFound = 30003
	ErrorNotCompare = 30004
	ErrorAuthToken = 30005
	ErrorAuthCheckTokenTimeout = 30006
	ErrorUploadFail = 30007

	// email 
	ErrorSendEmail = 30008
)
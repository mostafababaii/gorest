package errors

var Messages = map[int]string{
	SUCCESS:               "success",
	ERROR:                 "something went wrong",
	INVALID_PARAMS:        "invalid params",
	AUTHENTICATION_FAILED: "email or password is invalid",
	MISSING_TOKEN:         "missing token",
	INVALID_TOKEN:         "invalid token",
	USER_NOT_FOUND:        "user not found",
}

func GetMessage(code int) string {
	msg, ok := Messages[code]
	if ok {
		return msg
	}

	return Messages[ERROR]
}

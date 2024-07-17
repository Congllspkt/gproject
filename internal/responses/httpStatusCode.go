package responses

const (
	SUCCESS       = 20001
	PARAM_INVALID = 20003
	ERR_INVALID_TOKEN = 30001
)

var MSG = map[int]string{
	SUCCESS:       "success",
	PARAM_INVALID: "email is invalid",
	ERR_INVALID_TOKEN: "error invalid token",
}

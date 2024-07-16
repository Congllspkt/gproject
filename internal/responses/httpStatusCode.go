package responses

const (
	SUCCESS       = 20001
	PARAM_INVALID = 20003
)

var MSG = map[int]string{
	SUCCESS:       "success",
	PARAM_INVALID: "email is invalid",
}

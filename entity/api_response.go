package entity

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const CODE_SUCCESS = 0
const CODE_FAIL = 1000

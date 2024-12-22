package service

import "monitor_server/entity"

type ResponseService struct {
}

var ResponseServ *ResponseService = &ResponseService{}

func (a *ResponseService) Success(data interface{}) (re entity.ApiResponse) {

	re.Message = "success"
	re.Code = entity.CODE_SUCCESS
	re.Data = data
	return
}

func (a *ResponseService) Error(code int, message string) (re entity.ApiResponse) {

	re.Message = message
	re.Code = code
	re.Data = nil
	return
}

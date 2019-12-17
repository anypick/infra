package common

const (
	SuccessCode = 2000
	FailCode    = 5000
)

type ResponseData struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg,omitempty"`
	Success bool        `json:"success"`
	Total   int         `json:"total"`
	Rows    interface{} `json:"rows"`
}

func NewRespSucc() ResponseData {
	return ResponseData{Code: SuccessCode, Success: true}
}

func NewRespSuccWithData(rows interface{}, total int) ResponseData {
	return ResponseData{Code: SuccessCode, Success: true, Total: total, Rows: rows}
}

func NewRespSuccWithMsg(msg string) ResponseData {
	return ResponseData{Code: SuccessCode, Msg: msg, Success: true, Total: 0}
}

func NewRespSuccWithCodeMsg(code int, msg string) ResponseData {
	return ResponseData{Code: code, Msg: msg, Success: true, Total: 0}
}

func NewRespFail() ResponseData {
	return ResponseData{Code: FailCode, Success: false}
}

func NewRespFailWithMsg(msg string) ResponseData {
	return ResponseData{Code: FailCode, Success: false, Msg: msg}
}

func NewRespFailWithCodeMsg(code int, msg string) ResponseData {
	return ResponseData{Code: code, Success: false, Msg: msg}
}

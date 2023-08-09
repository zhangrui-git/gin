package response

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func New(code int, msg string, data any) Response {
	return Response{code, msg, data}
}

var success = New(0, "success", struct{}{})

func DefaultSuccessCode(c int) {
	success.Code = c
}

func DefaultSuccessMsg(m string) {
	success.Msg = m
}

func Success(data any) Response {
	d := success
	d.Data = data
	return d
}

var failed = New(400000, "error", struct{}{})

func DefaultFailedCode(c int) {
	failed.Code = c
}

func DefaultFailedMsg(m string) {
	failed.Msg = m
}

func Failed(code int) Response {
	d := failed
	d.Code = code
	return d
}

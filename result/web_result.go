package result

type WebResult struct {
	Code int8  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

package result

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type WebResult struct {
	Code int16
	Msg  string
	Data interface{}
}

//输出错误信息。
func OutputError(w http.ResponseWriter, msg string) {
	webResult := &WebResult{Code: 400, Msg: msg}
	b, _ := json.Marshal(webResult)
	fmt.Fprintf(w, string(b))
	return
}

//输出成功信息
func OutputSuccess(w http.ResponseWriter, msg string) {
	webResult := &WebResult{Code: 200, Msg: msg}
	b, _ := json.Marshal(webResult)
	fmt.Fprintf(w, string(b))
	return
}

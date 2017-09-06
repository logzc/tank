package index

import (
	"net/http"
	"fmt"
)

func Main(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎来到Tank，你好啊！")
}

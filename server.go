package main

import (
    "time"
	"fmt"
	"net/http"
)

func main() {
	// 路由
	http.HandleFunc("/create", httpOnCreate)
	// 监听端口+阻塞
    http.ListenAndServe(":80", nil)
}

func httpOnCreate(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query();
	orderid, ok := vars["orderid"];
	if (ok) {
		go sleepExec(orderid);
	}
	
}

func sleepExec(orderid []string) {
	// 阻塞
	time.Sleep(time.Second * 10)
	// 业务逻辑
	fmt.Println(orderid)
}

package main

import (
	"time"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/create", httpOnCreate)
	http.ListenAndServe(":80", nil)
}

func httpOnCreate(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query();
	orderid, success := vars["orderid"];
	if (success) {
		go sleepExec(orderid);
	}

	w.Write([]byte("200"))
}


func sleepExec(orderid []string) {
	time.Sleep(time.Second * 10)

	// 业务逻辑
	fmt.Println(orderid)
}
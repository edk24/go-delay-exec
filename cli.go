package main

import (
	"fmt"
	"net/http"
	"strconv"
)
func main(){
	for i := 0; i < 20000; i++ {
		go sendHttp(i)
	}
	select {}
}

func sendHttp(i int) {
	url := "http://localhost:80/create?orderid="+strconv.Itoa(1000+i)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf(strconv.Itoa(1000+i)+" ok\n")
}
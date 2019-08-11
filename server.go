package main

import (
	"time"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
 
func init() {
    db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
    db.SetMaxOpenConns(2000)
    db.SetMaxIdleConns(1000)
    db.Ping()
}

func main() {
	createHttpServer()
}

func createHttpServer() {
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

	// rows, err := db.Query("SELECT * FROM user limit 1")
    // defer rows.Close()
    // checkErr(err)
 
    // columns, _ := rows.Columns()
    // scanArgs := make([]interface{}, len(columns))
    // values := make([]interface{}, len(columns))
    // for j := range values {
    //     scanArgs[j] = &values[j]
    // }
 
    // record := make(map[string]string)
    // for rows.Next() {
    //     //将行数据保存到record字典
    //     err = rows.Scan(scanArgs...)
    //     for i, col := range values {
    //         if col != nil {
    //             record[columns[i]] = string(col.([]byte))
    //         }
    //     }
    // }
 
    // fmt.Println(record)
    // fmt.Fprintln(w, "finish")
}
/*
	webapp hiển thị doanh số bán xe của 1 số công ty ở 1 số quốc gia
*/
package main

import (
	handlers "./handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/sales", handlers.SearchCarSales)
	http.ListenAndServe(":8080", nil)
}

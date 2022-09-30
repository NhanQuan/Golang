/*
	webapp hiển thị doanh số bán xe của 1 số công ty ở 1 số quốc gia
*/
package main

import (
	handlers "./handlers"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	// get biến môi trường
	godotenv.Load("config.env")
	// handler xử lí get car-sales theo company
	http.HandleFunc("/sales", handlers.GetCarSalesByCompany)
	// handler xử lí get thông tin chi tiết car-sales
	http.HandleFunc("/sales/detailCarSales", handlers.GetDetailCarSales)
	// start server http://localhost:HOST
	http.ListenAndServe(os.Getenv("HOST"), nil)
}

/*
	package xử lí yêu cầu từ request
*/
package handlers

import (
	model "../model"
	render "../render"
	errorMessage "../message"
	constant "../constant"
	"net/http"
	"os"
)

// search thông tin doanh số bán xe của các hãng xe 
func GetCarSalesByCompany(w http.ResponseWriter, r *http.Request) {

	// doanh số bán xe
	var carSales model.CarSale
	// kết quả tìm kiếm danh sách doanh số bán xe của các công ty theo các param được input vào
	getCarSalesByCompany, err := carSales.GetCarSalesByCompany(getParam("company", r))

	// show thông báo lỗi
	if err != nil {
		errorMessage.ReturnErrorMessage(err.Error(), w)
		return
	}
	// nếu không tồn tại data thì show view thông báo
	if len(getCarSalesByCompany) == 0 {
		errorMessage.ReturnErrorMessage(constant.NoData, w)
		return
	}
	// show view
	showTemplate(getCarSalesByCompany, "PageResultSearchCarSales", w)
}

// thông tin chi tiết doanh số bán xe của hãng xe 
func GetDetailCarSales(w http.ResponseWriter, r *http.Request) {

	// doanh số bán xe
	var carSales model.CarSale

	// danh sách chi tiết doanh số bán xe của công ty theo param được truyền vào
	detailCarSalesByCompany, err := carSales.GetDetailCarSalesByCompany(getParam("company", r))
	// show thông báo lỗi
	if err != nil {
		errorMessage.ReturnErrorMessage(err.Error(), w)
		return
	}
	// show view
	showTemplate(detailCarSalesByCompany, "PageDetailCarSales", w)
}

// hiển thị thông tin ra màn hình
func showTemplate(detailCarSalesByCompany []model.CarSale, returnPage string, w http.ResponseWriter) {
	templateData := make(map[string]interface{})
	templateData["CarSales"] = detailCarSalesByCompany
	render.RenderTemplate(os.Getenv(returnPage), templateData, w)
}

// get param từ request
func getParam(key string, r *http.Request) string {
	return r.URL.Query().Get(key)
}

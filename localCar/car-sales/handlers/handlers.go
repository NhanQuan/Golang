/*
	package xử lí yêu cầu từ request
*/
package handlers

import (
	model "../model"
	"html/template"
	"net/http"
	"sort"
)

const (
	empty = ""
	noData = "Không tồn tại data"
	templateIndex = "templates/indexPage.html"
	templateError = "templates/messageError.html"
)

type error interface {
	Error() string
}

// search thông tin doanh số bán xe của các hãng xe
func SearchCarSales(w http.ResponseWriter, r *http.Request) {

	// get param company
	paramCompany := r.URL.Query().Get("company")
	// get param continent
	paramContinent := r.URL.Query().Get("continent")
	// lấy ra danh sách tất cả xe đã bán của các hãng 
	allCarSale := model.ShowAllCarSale(w)

	var listCarSale model.AllCarSale
	for _, carSales := range allCarSale.CarSales {

		// input cả continent và company thì get data có giá trị continent và company tương ứng 
		if (carSales.Continent == paramContinent && carSales.Company == paramCompany) {
			listCarSale.CarSales = append(listCarSale.CarSales, carSales)
			continue
		}

		// input param Continent thì get data có giá trị Continent tương ứng
		if (carSales.Continent == paramContinent && paramCompany == empty) {
			listCarSale.CarSales = append(listCarSale.CarSales, carSales)
			continue
		}

		// input param company thì get data có giá trị company tương ứng
		if (carSales.Company == paramCompany && paramContinent == empty) {
			listCarSale.CarSales = append(listCarSale.CarSales, carSales)
			continue
		}

		// không input param nào thì get toàn bộ data tại car-sales.json
		if paramCompany == empty && paramContinent == empty {
			listCarSale.CarSales = append(listCarSale.CarSales, carSales)
		}
	}

	// show data ra view 
	if len(listCarSale.CarSales) != 0 {
		template, err := template.ParseFiles(templateIndex)
		if err != nil {
			model.ReturnErrorMessage(err.Error(), w)
		}
		sortCarSalesBySales(listCarSale)
		template.Execute(w, listCarSale)
		return
	// trường hợp không tồn tại data 
	} 
	model.ReturnErrorMessage(noData, w)
}

// sort danh sách giảm dần theo sales
func sortCarSalesBySales(listCarSale model.AllCarSale) {
	sort.Slice(listCarSale.CarSales, func(i, j int) bool {
		return listCarSale.CarSales[i].Sales > listCarSale.CarSales[j].Sales
	})
}

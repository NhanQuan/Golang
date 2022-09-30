/*
	package xử lí data
*/
package model

import (
	"encoding/json"
	//constant "../constant"
	"io/ioutil"
	"os"
	"sort"
)

// định nghĩa các fields data
type CarSale struct {
	Continent string `json:"Continent"`
	Company string `json:"Company"`
	Sales int `json:"Sales"`
}

// lấy ra tất cả danh sách bán xe của các hãng xe
func getCarSales() ([]CarSale, error) {

	// khai báo biến lưu trữ data
	var carSales []CarSale
	// mở file dữ liệu input vào và gán vào file
	file, err := os.OpenFile(os.Getenv("CarSalesData"), os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return carSales, err
	}

	// đọc file dữ liệu
	carSalesData, err := ioutil.ReadAll(file)
	if err != nil {
		return carSales, err
	}

	// lưu chuỗi json vào mảng data
	json.Unmarshal(carSalesData, &carSales)
	return carSales, err
}

// get thông tin doanh số bán xe của các hãng xe
func (carSale CarSale) GetCarSalesByCompany(company string) ([]CarSale, error) {

	// danh sách kết quả 
	var getCarSalesByCompany []CarSale
	// tổng doanh số bán xe của công ty
	var totalSales int
	// danh sách doanh số bán xe của tất cả các hãng xe của các nước
	carSales, err := getCarSales()
	// return nếu xảy ra lỗi
	if err != nil {
		return getCarSalesByCompany, err
	}
	// show tổng doanh số bán xe của công ty được input
	for _, carSale := range carSales {
		if carSale.Company == company {
			getCarSalesByCompany = append(getCarSalesByCompany, carSale)
			continue
		}
	}

	for _, detailCarSales := range getCarSalesByCompany {
		totalSales += detailCarSales.Sales
		detailCarSales.Sales = totalSales
	}
	// sắp xếp giảm dần theo doanh số bán hàng
	carSale.sortCarSalesBySalesDescending(getCarSalesByCompany)
	return getCarSalesByCompany, err
}

// chi tiết thông tin doanh số bán xe của các hãng xe
func (carSale CarSale) GetDetailCarSalesByCompany(company string) ([]CarSale, error) {

	// danh sách doanh số bán xe của tất cả các hãng xe của các nước
	carSales, err := getCarSales()
	// danh sách kết quả chi tiết
	var detailCarSalesByCompany []CarSale
	// return nếu xảy ra lỗi
	if err != nil {
		return detailCarSalesByCompany, err
	}

	// hiển thị tất cả doanh số bán xe của công ty cần xem chi tiết
	for _, carSale := range carSales {
		if (carSale.Company == company) {
			detailCarSalesByCompany = append(detailCarSalesByCompany, carSale)
			continue
		}
	}
	// sắp xếp giảm dần theo doanh số bán hàng
	carSale.sortCarSalesBySalesDescending(detailCarSalesByCompany)
	return detailCarSalesByCompany, err
}

// sort danh sách giảm dần theo sales
func (carSale CarSale) sortCarSalesBySalesDescending(carSales []CarSale) {
	sort.Slice(carSales, func(i, j int) bool {
		return carSales[i].Sales > carSales[j].Sales
	})
}

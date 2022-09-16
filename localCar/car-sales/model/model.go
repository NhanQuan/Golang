/*
	package xử lí data
*/
package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"net/http"
	"html/template"
)

const (
	templateError = "templates/messageError.html"
	dataCarSales = "list.json"
)

type CarSale struct {
	Continent string  `json:"Continent"`
	Company  string  `json:"Company"`
	Sales   int `json:"Sales"`
}

type AllCarSale struct {
	CarSales []*CarSale
}

// lấy ra tất cả danh sách bán xe của các hãng xe
func ShowAllCarSale(w http.ResponseWriter) (*AllCarSale) {
	file, err := os.OpenFile(dataCarSales, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		ReturnErrorMessage(err.Error(), w)
	}
	carSales, err := ioutil.ReadAll(file)
	if err != nil {
		ReturnErrorMessage(err.Error(), w)
	}
	var allCarSale AllCarSale
	json.Unmarshal(carSales, &allCarSale.CarSales)
	return &allCarSale
}

// trả về thông báo lỗi
func ReturnErrorMessage(errorMessage string, w http.ResponseWriter) {
	template, err := template.ParseFiles(templateError)
	if err != nil {
		template.Execute(w, err.Error())
		return
	}
	template.Execute(w, errorMessage)
	return
}

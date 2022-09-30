/*
	package xử lí hiển thị page
*/
package render

import (
	"net/http"
	constant "../constant"
	"html/template"
	"os"
	"github.com/jackc/numfmt"
)

// format Sales của doanh số bán xe đúng định dạng
func salesFormat(sales int) string {
	formatSales := &numfmt.Formatter{}
	return formatSales.Format(sales)
}

// tạo function cho template
var funcMap = template.FuncMap{
	"salesFormat": salesFormat,
}

// hiển thị view
func RenderTemplate(page_Template string, data map[string]interface{}, w http.ResponseWriter) {
	renderTemplate, err := template.New(constant.Empty).Funcs(funcMap).ParseFiles(os.Getenv("PathTemplate") + page_Template)
	if err != nil {
		renderTemplate.Execute(w, err.Error())
		return
	}
	renderTemplate.ExecuteTemplate(w, page_Template, data)
}

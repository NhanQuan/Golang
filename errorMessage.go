/*
	package thông báo lỗi
*/
package errorMessage

import (
	"os"
	render "../render"
	"net/http"
)

// trả về thông báo lỗi
func ReturnErrorMessage(message string, w http.ResponseWriter) {
	errorMessage := make(map[string]interface{})
	errorMessage["errorMessage"] = message
	render.RenderTemplate(os.Getenv("PageError"), errorMessage, w)
	return
}

package belajar_go_lang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// template layout
// saat kita membuat halaman website, kadang ada beberapa bagian yang selalu sama, misal header dan footer
// jika ada bagian yang selalu sama, disarankan untuk di simpan pada template yang terpisah agar bisa digunakan-
// di template lain (konsep layouting)
// golang template mendukung import dari template lain

// import template
// {{ template "nama" }} : mengimport template tanpa data yang dikirimkan
// {{ template "nama" .Value }} : mengimport template dengan mengirimkan data 'value'
// untuk mengirimkan semua data dari template bisa menggunakan akhiran '.' (titik)

// membuat handler untuk template layout
func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	// meload file dengan parse file ke template
	// mengimport semua file yang berhubungan dengan layout (header, footer, dan layout atau content itu sendiri)
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml", 
		"./templates/footer.gohtml", 
		"./templates/layout.gohtml", 
		))

	// mengeksekusi template (untuk merender file ke template)
	// nama template sekarang menyesuaikan dengan nama file
	// ketika template sudah diberikan define, maka nama define yang sekaran yang dipanggil di execute
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name": "Taufik Hidayat", 
	})
}

// membuat kode uji untuk template layout
func TestTemplateLayout(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateLayout
	TemplateLayout(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	fmt.Println(string(body)) 
}


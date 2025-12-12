package belajar_go_lang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// template data
// saat kita membuat template, kadang kita ingin menambahkan banyak data dinamis
// hal ini bisa kita lakukan dengan cara menggunakan data struct atau map
// namun perlu dilakukan perubahan di dalam text templatenya, kita perlu memberi tahu Field atau Key-
// mana yang akan kita gunakan untuk mengisi data dinamis di template
// kita bisa menyebutkan dengan cara seperti ini {{.NamaField}}
// kalau menggunakan struct, bisa memakai nama attribute. kalau map bisa menggunakan nama key

// membuat handler map
func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	// membuat template
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	// sebelumnya kita mengisikan data dinamis (satu buah) menggunakan string,-
	// kemudian berhubung sekarang ada lebih dari satu, kita bisa menggunakan Map
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name": "Taufik Hidayat",
		"BodyContent": "Haii",
		"Address": map[string]interface{} {
			"Street": "Jalan Jember",
		},
	})
}

// membuat kode uji map
func TestTemplateDataMap(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateDataMap
	TemplateDataMap(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	// kalau sudah di render, maka kurung kurawal pada template akan otomatis hilang
	fmt.Println(string(body)) 
}

// membuat struct
type Address struct {
	Street string
}

type Page struct {
	Title, Name, BodyContent string
	Address Address // menambahkan key baru yang berjenis struct
}

// membuat handler struct
func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	// membuat template
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	// sebelumnya kita mengisikan data dinamis (satu buah) menggunakan string,-
	// kemudian berhubung sekarang ada lebih dari satu, kita bisa menggunakan Struct (nama struct Page)
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name: "Taufik Hidayat",
		BodyContent: "Haii",
		Address: Address{
			Street: "Jalan Jember",
		},
	})
}

// membuat kode uji struct
func TestTemplateDataStruct(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateDataStruct
	TemplateDataStruct(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	// kalau sudah di render, maka kurung kurawal pada template akan otomatis hilang
	fmt.Println(string(body)) 
}


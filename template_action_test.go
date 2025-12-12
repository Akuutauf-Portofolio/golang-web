package belajar_go_lang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// template action
// golang template mendukung perintah action, seperti percabangan, perulangan dan lain lain

// if
// {{ if .Value }} T1 {{ end }}, jika value tidak kosong/true, maka T1 akan di eksekusi. jika kosong/false, maka tidak ada yang dieksekusi
// {{ if .Value }} T1 {{ else }} T2 {{ end }}, jika value tidak kosong/true, maka T1 akan dieksekusi. Jika kosong/false, maka T2 yang akan dieksekusi
// {{ if .Value }} T1 {{ else if .Vallue }} T2 {{ else }} T3 {{ end }}, jika value T1 tidak kosong/true, maka T1 akan dieksekusi.-
// jika value T2 tidak kosong/true, maka T2 yang akan dieksekusi. jika tidak semuanya, maka T3 akan dieksekusi

// selalu akhiri pengkondisian dengan {{ end }}

// membuat handler untuk template action if
func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	// meload file dengan parse file ke template
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	// mengeksekusi template (untuk merender file ke template)
	// nama template sekarang menyesuaikan dengan nama file
	t.ExecuteTemplate(writer, "if.gohtml", map[string]interface{}{
		"Title": "Template Action If",
		"Name": "Taufik Hidayat", // kalau mengirimkan attribute 'Name', maka akan muncul, kalau tidak mengirimkan hanya muncul 'Hello'
	})
}

// membuat kode uji untuk file
func TestTemplateActionIf(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateActionIf
	TemplateActionIf(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	// kalau sudah di render, maka kurung kurawal pada template akan otomatis hilang
	fmt.Println(string(body)) 
}

// operator perbandingan
// golang template juga mendukung operator perbandingan, ini cocok ketika butuh melakukan perbandingan number-
// di if statement, berikut adalah operator nya :
// eq : artinya arg1 == arg2 (equals/sama dengan)
// ne : artinya arg1 != arg2 (not equals/tidak sama dengan)
// lt : artinya arg1 < arg2 (less than/kurang dari)
// le : artinya arg1 <= arg2 (less than equals/kurang dari sama dengan)
// gt : artinya arg1 > arg2 (greater than/lebih dari)
// ge : artinya arg1 >= arg2 (greater than equals/lebih dari sama dengan)

// kenapa operatornya didepan
// hal ini dikarenakan, operator perbandingan diatas adalah sebuah function
// jadi saat kita menggunakan {{ eq First Second }}, sebenarnya dia memanggil function eq-
// dengan parameter First dan Second sebagai parameter. Contoh : eq(First, Second)

// membuat handler untuk template action if else dan operator comparison (operator perbandingan)
func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	// meload file dengan parse file ke template
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	// mengeksekusi template (untuk merender file ke template)
	// nama template sekarang menyesuaikan dengan nama file
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title": "Template Action Comparator",
		"FinalValue": 60,
	})
}

// membuat kode uji untuk operator perbandingan
func TestTemplateActionOperator(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateActionOperator
	TemplateActionOperator(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	// kalau sudah di render, maka kurung kurawal pada template akan otomatis hilang
	fmt.Println(string(body)) 
}

// range
// digunakan untuk melakukan iterasi / perulangan data pada template
// pada golang template tidak ada perulangan seperti for,-
// sehingga perulangan bisa dilakukan dengan range
// perulangan dengan range bisa dilakukan dengan data yang berjenis array, slice, map maupun channel
// {{ range $index, $element := .Value1 }} T1 {{ end }}, jika value memiliki data, maka T1 akan dieksekusi,-
// sebanyak element value, dan kita bisa menggunakan $index untuk mengakses index dan $element untuk mengakses elemen
// {{ range $index, $element := .Value1 }} T1 {{ else }} T2 {{ end }}, sama seperti sebelumnya, namun jika-
// value tidak memiliki elemen apapun, maka T2 yang akan dieksekusi

// membuat handler untuk template action range
func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	// meload file dengan parse file ke template
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	// mengeksekusi template (untuk merender file ke template)
	// nama template sekarang menyesuaikan dengan nama file
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobies": []string{ // membuat array di dalam map
				"Gaming", "Reading", "Coding",
		},
	})
}

// membuat kode uji untuk range
func TestTemplateActionRange(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateActionRange
	TemplateActionRange(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	// kalau sudah di render, maka kurung kurawal pada template akan otomatis hilang
	fmt.Println(string(body)) 
}

// with
// jika kita menggunakan template, kita bisa mengaksesnya menggunakan .Value.NestedValue (nested struct)
// di template terdapat action with, yang bisa digunakan untuk mengubah scope dot menjadi object yang kita mau
// {{ with .Value1 }} T1 {{ end }}, jika value tidak kosong, di T1 semua dot akan merefer ke value
// {{ with .Value1 }} T1 {{ else }} T2 {{ end }}, sama seperti sebelumnya, namun jika value kosong,-
// maka T2 yang akan dieksekusi

// membuat handler untuk template action if else dan operator comparison (operator perbandingan)
func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	// meload file dengan parse file ke template
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))

	// mengeksekusi template (untuk merender file ke template)
	// nama template sekarang menyesuaikan dengan nama file
	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Name": "Taufik Hidayat",
		"Address": map[string]interface{} { // membuat data map nested di dalam map
			"Street": "Jalan Jember",
			"Address": "Banyuwangi Kota",
			"City": "Banyuwangi",
		},
	})
}

// membuat kode uji untuk file
func TestTemplateActionWith(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateActionWith
	TemplateActionWith(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	// kalau sudah di render, maka kurung kurawal pada template akan otomatis hilang
	fmt.Println(string(body)) 
}

// comment
// tempalte juga mendukung komentar
// komentar secara otomatis akan hilang (di ignore) ketika template text di parsing
// untuk membuat komentar bisa menggunakan : {{ /* Contoh Komentar */ }}
// contohnya ada dibaris nomor 9, di file address.gohtml
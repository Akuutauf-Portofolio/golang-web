package belajar_go_lang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// template name
// saat kita membuat template dari file, secara otomatis nama file nya akan menjadi nama template
// namun jika kita bisa mengubah nama templatenya menggunakan perintah : "{{ define "nama_template" }}" TEMPLATE {{ end }},-
// artinya kita membuat alias template dengan kata kunci sesuai pada "nama_template"
// pada satu file template, sebetulnya kita bisa membuat double content (template name), asalkan harus ada lebih dari satu 'define'
// contoh ada pada file 'layout.gohtml'

// template function
// selain mengakses field '{{ .Nama }}', template juga bisa mengakses function
// cara mengakses nya juga sama, namun jika function membutuhkan parameter, kita bisa menambahkan parameter-
// ketika memanggil function di template nya
// {{ .FunctionName }} : memanggil field function name atau FunctionName() (otomatis di panggil dalam bentuk function oleh golang)
// {{ .FunctionName "taufik", "hidayat" }} : memanggil function FunctionName("taufik", "hidayat")

// membuat struct
type MyPage struct {
	Name string
}

// membuat method milik my page
func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

// membuat handler untuk template function
func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	// membuat template langsung dari function
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "taufik" }}`))

	// mengeksekusi template
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "ilham",
	})
}

// membuat kode uji untuk template function
func TestTemplateFunction(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateFunction
	TemplateFunction(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	fmt.Println(string(body)) 
}

// global function
// adalah function yang bisa digunakan secara langsung, tanpa menggunakan template data
// golang template memiliki beberapa global function

// membuat handler untuk template function global (mencoba function global golang yang sudah ada)
func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	// membuat template langsung dari function
	// mencoba function global golang yang sudah ada untuk template (function len)
	t := template.Must(template.New("FUNCTION").Parse(`{{ len .Name }}`))

	// mengeksekusi template
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "taufik",
	})
}

// membuat kode uji untuk template function
func TestTemplateFunctionGlobal(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateFunctionGlobal
	TemplateFunctionGlobal(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	fmt.Println(string(body)) 
}

// menambahkan global function 
// kita juga bisa menambahkan global function sendiri
// untuk menambahkan global function, kita bisa menggunakan method 'Funcs' pada template
// namun syaratnya harus dilakukan sebelum melakukan parsing template

// membuat handler untuk template create function global 
func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	// sebelum diparsing, registrasikan function nya ke dalam template
	// membuat template kosong
	t := template.New("FUNCTION")

	// meregistrasi function ke template 
	// membuat template dengan function upper di dalam template
	t = t.Funcs(map[string]interface{} {
		// nama function nya adalah "upper", yang memiliki 1 parameter string
		"upper": func (value string) string {
			// mengembalikan string dari parameter ke return menjadi uppercase (huruf besar semua)
			return strings.ToUpper(value)
		},
	})

	// membuat template langsung dari function
	// .Parse : untuk parsing dari text
	// .ParseGlob : untuk parsing dari file (file atau direktori)
	// .ParseFS : untuk parsing dari embed (file atau direktori)
	// .ParseFiles : untuk parsing dari beberapa file

	// yang diparsing sekarang adalah t bukan template
	// karena 'upper' sekarang adalah function global, maka tidak perlu pakai titik '.' 
	t = template.Must(t.Parse(`{{ upper .Name }}`))

	// mengeksekusi template
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "taufik hidayat",
	})
}

// membuat kode uji untuk template create function global
func TestTemplateFunctionCreateGlobal(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateFunctionCreateGlobal
	TemplateFunctionCreateGlobal(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	fmt.Println(string(body)) 
}

// function pipelines
// golang template mendukung function pipelines, artinya hasil dari function bisa dikirimkan ke function berikutnya
// untuk mengimplementasikan pipelines, kita bisa menggunakan tanda '|' misal :-
// {{ SayHello .Name | upper }} : artinya akan memanggil global function SayHello(Name) kemudian-
// hasil dari SayHello(Name) akan dikirimkan ke function upper(hasil)
// kita bisa mengirimkan function pipelines lebih dari satu

// membuat handler untuk template function pipelines
func TemplateFunctionPipelines(writer http.ResponseWriter, request *http.Request) {
	// sebelum diparsing, registrasikan function nya ke dalam template
	// membuat template kosong
	t := template.New("FUNCTION")

	// meregistrasi function ke template
	t = t.Funcs(map[string]interface{} {
		// kita bisa membuat function global lebih dari satu pada template
		"sayHello": func (value string) string {
			return "Hello " + value
		},
		"upper": func (value string) string {
			// mengembalikan string dari parameter ke return menjadi uppercase (huruf besar semua)
			return strings.ToUpper(value)
		},
	})

	// yang diparsing sekarang adalah t bukan template
	// implementasi pipelines
	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))

	// mengeksekusi template
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "taufik hidayat",
	})
}

// membuat kode uji untuk template function pipelines
func TestTemplateFunctionPipelines(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateFunctionPipelines
	TemplateFunctionPipelines(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	fmt.Println(string(body)) 
}
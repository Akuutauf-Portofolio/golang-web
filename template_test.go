package belajar_go_lang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// web dinamis
// di golang terdapat fitur HTML Template, yaitu fitur template yang bisa kita gunakan untuk membuat HTML yang dinamis

// html template
// fitur html template terdapat di package html/template
// sebelum menggunakan html template, kita perlu terlebih dahulu membuat template nya-
// tujuannya agar ketika pengembangan aplikasi sudah besar, kita bisa mengetahui bagian setiap templatenya
// template bisa berupa file (umum) atau string
// bagian dinamis pada HTML template, adalah bagian yang menggunakan tanda '{{  }}'

// membuat template
// saat membuat template dengan string, kita perlu memberi tahu nama template nya
// dengan untuk membuat text template, cukup buat text html, dan untuk konten yang dinamis,-
// kita bisa gunakan tanda {{.}}, contoh : <html><body>{{.}}</body></html>

// membuat handler
func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	// membuat template text html
	// simbol '.' digunakan untuk teks dinamis yang akan dikirimkan ke template
	templateText := "<html><body>{{.}}</body></html>" 

	// membuat template
	// jikalau templat dalam bentuk string, maka perlu memberikan nama templatenya
	// kemudian dilakukan parsing
	t, err := template.New("SIMPLE").Parse(templateText) // mengembalikan 2 data (template, dan error)

	// mengecek error template 
	if err != nil {
		panic(err)
	}

	// kode yang serupa namun, tidak mengembalikan error, dan tidak perlu mengecek error lagi
	// t := template.Must(New("SIMPLE").Parse(templateText)) // function Must() : sudah memiliki bawaan pengecekan error

	// mengeksekusi template (untuk merender data ke template)
	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template!")
}

// membuat kode uji 
func TestSimpleHTML(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest SimpleHTML
	SimpleHTML(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	// kalau sudah di render, maka kurung kurawal pada template akan otomatis hilang
	fmt.Println(string(body)) 
}

// template dari file
// selain membuat template dari string, kita juga bisa membuat template langsung dari file
// hal ini mempermudah kita, karena bisa langsung membuat file (contoh html)
// saat membuat template menggunakan file, secara otomatis nama file akan menjadi nama template nya,-
// misal jika kita punya fil html dengan nama 'simple.html', maka nama template nya adalah 'simple.html'

// membuat handler untuk template file
func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	// meload file dengan parse file ke template
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	// mengeksekusi template (untuk merender file ke template)
	// nama template sekarang menyesuaikan dengan nama file
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template!")
}

// membuat kode uji untuk file
func TestSimpleHTMLFile(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest SimpleHTMLFile
	SimpleHTMLFile(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	// kalau sudah di render, maka kurung kurawal pada template akan otomatis hilang
	fmt.Println(string(body)) 
}

// template direktori
// di golang jarang sekali pengembangan template menyebutkan file satu persatu
// sehingga direkomendasikan untuk template di simpan ke dalam satu direktori
// golang template mendukung proses load template dari direktori, sehingga tidak perlu menyebutkan-
// nama file template satu persatu

// membuat handler untuk template direktori
func TemplateDirektory(writer http.ResponseWriter, request *http.Request) {
	// meload beberapa file di direktori template dengan function ParseGlob
	// dengan function ParseGlob : maka akan meload semua file yang berakhiran '*.gohtml'
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	// mengeksekusi template (untuk merender file ke template)
	// nama template sekarang menyesuaikan dengan nama file yang ingin dirender (hanya satu file)
	// kalau nama file tidak ditemukan, maka nanti akan menampilkan kosong
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template!")
}

// membuat kode uji untuk file direktori
func TestTemplateDirektory(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateDirektory
	TemplateDirektory(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	// kalau sudah di render, maka kurung kurawal pada template akan otomatis hilang
	fmt.Println(string(body)) 
}

// tempalte menggunakan golang embed
// direkomendasikan untuk melakukan golang embed untuk meload file, agar tidak mengcopy ulang
// sehingga nanti hasil aplikasi kita sudah mengembed file nya

// meload file (lebih dari satu) menggunakan embed
//go:embed templates/*.gohtml
var templates embed.FS

// membuat handler untuk template embed
func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	// meload beberapa file di direktori template dengan function ParseGlob
	// dengan function ParseGlob : maka akan meload semua file yang berakhiran '*.gohtml'
	// kalau di golang embed tidak ada direktori saat ini ("./templates/*.gohtml"),-
	// maka penulisan path yang benar adalah yang dibawah
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	// mengeksekusi template (untuk merender file ke template)
	// nama template sekarang menyesuaikan dengan nama file yang ingin dirender (hanya satu file)
	// kalau nama file tidak ditemukan, maka nanti akan menampilkan kosong
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template!")
}

// membuat kode uji untuk direktori embed
func TestTemplateEmbed(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateEmbed
	TemplateEmbed(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	// kalau sudah di render, maka kurung kurawal pada template akan otomatis hilang
	fmt.Println(string(body)) 
}

package belajar_go_lang_web

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// template caching
// pada kode sebelumnya sebenarnya tidak efisien, karena setiap handler dipanggil kita selalu melakukan parsing-
// ulang templatenya, sehingga mengakibatkan sistem akan lambat
// disarankan untuk melakukan parsing cukup satu kali diawal pada saat aplikasi dijalankan/berjalan
// kemudian selanjutnya data template akan di caching (di simpan di memory), sehingga tidak perlu melakukan parsing ulang
// dengan demikian akan membuat web menjadi lebih efisien/cepat

// meload file (lebih dari satu) menggunakan embed
//go:embed templates/*.gohtml
var templates_embed embed.FS // di panggil lebih dari satu ada warning

// membuat variabel global untuk format template parsing (di simpan di memory sehingga lebih cepat)
// disarankan dibuat menjadi global variabel agar tidak melakukan parsing setiap kali handler dibuat/dijalankan
var myTemplates = template.Must(template.ParseFS(templates_embed, "templates/*.gohtml"))

// membuat handler yang mengimplementasikan template caching (myTemplates)
func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	// mengeksekusi template dari template caching
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

// membuat kode uji
func TestTemplateCaching(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateCaching
	TemplateCaching(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	fmt.Println(string(body)) 
}
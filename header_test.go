package belajar_go_lang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// header
// selain query parameter, dalam HTTP juga terdapat yang namanya header
// header adalah informasi tambahan yang bisa dikirim dari client ke server atau sebaliknya
// dalam header, tidak hanya pada HTTP request, namun HTTP response pun kita bisa menambahkan informasi header
// saat kita menggunakan browser, biasanya secara otomatis header akan ditambahkan oleh browser,-
// yang berisi informasi seperti browser, jenis tipe content yang dikirim dan diterima oleh browser, dan lain lain

// request header
// Request.Header : digunakan untuk menangkap request header yang dikirim oleh client
// header di simpan dalam bentuk map[string][]string, sama seperti query parameter
// namun kalau query parameter itu bersifat case sensitive, sedangkan header 'key' nya tidaklah case sensitive

// membuat function yang mengimplementasikan kontrak handler
func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	// request header : data HTTP request yang mengandung header, dikirim dari client ke server

	// mengambil header untuk key 'content-type'
	contentType := request.Header.Get("content-type") // kembaliannya adalah data header yang berjenis map

	// menampilkan isi dari contentType
	fmt.Fprint(writer, contentType)
}

// membuat kode uji untuk header
func TestRequestHeader(t *testing.T) {
	// membuat request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)

	// menambahkan header request
	request.Header.Add("content-type", "application/json") // incase sensitive (bebas)

	// membuat recoder (simulasi sebagai writer)
	recorder := httptest.NewRecorder()

	// memanggil endpoint simulasi, yaitu RequestHeader()
	RequestHeader(recorder, request)

	// melihat hasil dari recorder
	response := recorder.Result()

	// melihat isi body dari recorder
	body, _ := io.ReadAll(response.Body)

	// menampilan isi body dari handler
	fmt.Println(string(body))
}

// response header
// ResponseWriter.Header() : digunakan untuk menambahkan header pada response (dari server ke client)

// membuat function yang mengimplementasikan kontrak handler
func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	// response header : data HTTP request yang mengandung header, dikirim dari server ke client

	// menambahkan header untuk key custom 'content-type'
	writer.Header().Add("X-Powerred-By", "Taufik Hidayat")

	// print 
	fmt.Fprint(writer, "OK")
}

// membuat kode uji untuk header
func TestResponseHeader(t *testing.T) {
	// membuat request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)

	// membuat recoder (simulasi sebagai writer)
	recorder := httptest.NewRecorder()

	// memanggil endpoint simulasi, yaitu ResponseHeader()
	ResponseHeader(recorder, request)

	// menampil isi dari header key costum baru, mengambil data dari recorder
	powerredBy := recorder.Header().Get("X-powerred-by") // bersifat incase sensitive (bebas)

	// menampilan isi body dari handler
	fmt.Println(powerredBy)
}

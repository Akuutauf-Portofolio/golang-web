package belajar_go_lang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// http test
// golang memiliki package yang dapat kita gunakan untuk menguji unit test handler
// dengan menggunakan package ini (httptest), kita dapat menguji handler web di digolang
// tanpa harus menjalankan aplikasi web/web servernya
// sehingga kita bisa langsung fokus terhadap handler function yang ingin kita test
// digunakan untuk menguji API / endpoint
// NewRequest(method, url, body) : digunakan untuk membuat request http.Request yang akan dikirim ke http test
// kita bisa mengirimkan method, url, body sebagai simulasi test tanpa perlu mengakses ke aplikasi lain
// juga bisa menamabahkan informasi yang lain seperti header, cookie dan lainnya

// httptst.NewRecorder()
// merupakan function yang digunakan untuk membuat ResponseRecorder (sebagai writer pada http test)

// menyamakan kontrak dengan function handler
func HelloHandler(writter http.ResponseWriter, request *http.Request) {
	// menampilkan pesan
	fmt.Fprint(writter, "Hello World")
}

// mebuat kode uji
func TestHttp(t *testing.T) {
	// tidak perlu membuat server, karena kita ingin menguji unit test http test
	// sehingga kita bisa menggunakan NewRequest
	// kalau takut keliru mendefinisikan method seperti contoh "GET", maka bisa menggunakan 'http.MethodGet'
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil) // untuk body berikan nilai nil

	// membuat recoder (simulasi sebagai writer)
	recorder := httptest.NewRecorder()

	// memanggil endpoint simulasi, yaitu Hello Handler
	HelloHandler(recorder, request)

	// melihat hasil dari recorder
	response := recorder.Result()

	// melihat isi body dari recorder, harus menggunakan io (karena outputnya berjenis io)
	body, _ := io.ReadAll(response.Body) // mengambalikan 2 nilai (body berjenis []byte, dan err)

	// menampilan isi body dari handler
	fmt.Println(string(body))
}
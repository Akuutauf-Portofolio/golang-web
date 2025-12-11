package belajar_go_lang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// response code
// adalah representasi kode response
// dari response code ini kita bisa melihat apakah sebuah request yang kita kirim itu sukses diproses-
// oleh server atau gagal
// response code success : 200 - 299
// response code redirect : 300 - 399
// response code client error : 400 - 499
// response code server error : 500 - 599

// mengubah response code
// secara default, jika kita tidak menyebutkan response code, maka response code bawaan adalah 200 OK
// function ResponseWriter.Write.Header(int) : untuk mengubah response code
// semua data status code sudah disediakan di golang, sehingga kita bisa menggunakan variabel agar tidak keliru
// ini dokumentasi untuk response code : https://go.dev/src/net/http/status.go

// membuat handler
func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	// mengambil data query parameter name
	name := request.URL.Query().Get("name")

	// mengecek jikalau data name nya kosong, maka akan mengembalikan response code 400 (client error)
	if name == "" {
		// mengirimkan response code ke writer (pengguna)
		// status bad request (400)
		writer.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(writer, "Name is empty")
	} else {
		// mengirimkan response code ke writer (pengguna)
		// status ok (200)
		writer.WriteHeader(http.StatusOK)

		fmt.Fprintf(writer, "Hello %s", name)
	}
}

// membuat kode uji response code
func TestResponseCode(t *testing.T) {
	// mengambil data request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recoder (simulasi sebagai writer)
	recorder := httptest.NewRecorder()

	// memanggil endpoint simulasi, yaitu ResponseCode()
	ResponseCode(recorder, request)

	// melihat hasil dari recorder
	response := recorder.Result()

	// melihat isi body dari recorder
	body, _ := io.ReadAll(response.Body)

	// menampilan isi body dari handler
	fmt.Println(string(body))

	// menampilkan status code dan statusnya apa
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}

func TestResponseCodeValid(t *testing.T) {
	// mengambil data request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Taufik", nil)

	// membuat recoder (simulasi sebagai writer)
	recorder := httptest.NewRecorder()

	// memanggil endpoint simulasi, yaitu ResponseCode()
	ResponseCode(recorder, request)

	// melihat hasil dari recorder
	response := recorder.Result()

	// melihat isi body dari recorder
	body, _ := io.ReadAll(response.Body)

	// menampilan isi body dari handler
	fmt.Println(string(body))

	// menampilkan status code dan statusnya apa
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}
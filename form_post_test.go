package belajar_go_lang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// form post
// saat membuat form, kita bisa mengirimkan submit data nya dengan method GET atau POST
// sebetulnya tidak ada bedanya, cuman kalau pakai method GET data nya dirkimkan melalui query parameter-
// sedangkan method POST datanya dikirimkan melalui body HTTP request

// Request.PostForm
// semua data form post yang dikirim dari client, secara otomatis akan di simpan dalam atribute, 'Request.PostForm'
// namun sebelum bisa mengambil data di attribute PostForm, maka sebelumnya wajib melakukan parsing data dengan method-
// Request.ParseForm() : digunakan untuk melakukan parsing data body, apakah bisa diparsing menjadi form data atau tidak,-
// jika tidak bisa diparsing maka akan menyebabkan error

// membuat handler
func FormPost(writer http.ResponseWriter, request *http.Request) {
	// melakukan parsing data dengan ParseForm()
	err := request.ParseForm()

	// melakukan pengecekan, apakah data body form nya sudah sesuai atau belum
	if err != nil {
		panic(err)
	}

	// jika tidak error, kita bisa mengambil data melalui PostForm, dan Get untuk mengambil data (seperti query parameter)
	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	// tampilkan ke writter nya
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

// membuat kode uji untuk FormPost
func TestFormPost(t *testing.T) {
	// membuat request body langsung menggunakan strings, yang nantinya akan dikonversi menjadi reader
	// penulisannya hampir sama dengan query parameter
	requestBody := strings.NewReader("first_name=Taufik&last_name=Hidayat")

	// membuat request post
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)

	// jangan lupa menambahkan header untuk content type, dimana standar nya menggunakan-
	// application/x-www-form-urlencoded 
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded ")

	// membuat recoder (simulasi sebagai writer)
	recorder := httptest.NewRecorder()

	// memanggil endpoint simulasi, yaitu FormPost()
	FormPost(recorder, request)

	// melihat hasil dari recorder
	response := recorder.Result()

	// melihat isi body dari recorder
	body, _ := io.ReadAll(response.Body)

	// menampilan isi body dari handler
	fmt.Println(string(body))
}
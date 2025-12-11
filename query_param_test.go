package belajar_go_lang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// query parameter
// digunakan untuk mengirimkan data melalui parameter endpoint dari client ke server
// query parameter ditempatkan di URL
// untuk menambahkan query parameter, kita bisa menggunakan contoh '?nama=value' pada urlnya

// url.URL
// dalam parameter request, terdapat atribut URL yang berisi data url.URL
// dari data url ini, kita bisa mengambil data query parameter (data parameter) yang dikirim dari-
// client ke server dengan menggunakan method Query() yang akan mengembalikan map

// membuat handler yang akan diuji sebagai simulai
func SayHello(writer http.ResponseWriter, request *http.Request) {
	// request.URI : digunakan untuk mendapatkan data url dalam bentuk string (alamat url nya saja)
	// request.URL : digunakan untuk mendapatkan full data url mulai parametet, method dan lain lain
	
	// mendapatkan data url untuk mengambil data parameter dengan function Query (mengembalikan values/map)
	// query parameter : adalah data yang dikirimkan ke url endpoint melalui atribut nya contoh nya 'name'
	// function GET() : hanya mengambil satu nilai dari query parameter
	name := request.URL.Query().Get("name")

	// melakukan pengecekan
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		// %s adalah salah satu subtitusi format untuk menampilkan data dalam bentuk string dengan fprintf
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

// membuat kode uji
func TestQueryParam(t *testing.T) {
	// membuat request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=taufik", nil)

	// membuat recoder (simulasi sebagai writer)
	recorder := httptest.NewRecorder()

	// memanggil endpoint simulasi, yaitu SayHello()
	SayHello(recorder, request)

	// melihat hasil dari recorder
	response := recorder.Result()

	// melihat isi body dari recorder
	body, _ := io.ReadAll(response.Body)

	// menampilan isi body dari handler
	fmt.Println(string(body))
}

// mutiple query parameter
// dalam spesifikasi url, kita bisa menambahkan lebih dari satu query parameter dalam endpoint
// cocok sekali, ketika kita ingin mengirim banyak data ke server, cukup tambahkan query parameter lainnya
// untuk menambahkan lebih dari satu query parameter, kita bisa gunakan tanda '&' lalu diikuti dengan-
// query parameter berikutnya
// kalau kita mengirimkan data yang tidak tercantum pada query parameter nya, maka akan secara otomatis datanya adalah kosong

// membuat function mutiple query parameter sebagai simulasi handler
func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	// mendapatkan query parameter
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

// membuat kode uji untuk mutiple query parameter
func TestMultipleQueryParam(t *testing.T) {
	// membuat request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Taufik&last_name=Hidayat", nil)

	// membuat recoder (simulasi sebagai writer)
	recorder := httptest.NewRecorder()

	// memanggil endpoint simulasi, yaitu MultipleQueryParameter()
	MultipleQueryParameter(recorder, request)

	// melihat hasil dari recorder
	response := recorder.Result()

	// melihat isi body dari recorder
	body, _ := io.ReadAll(response.Body)

	// menampilan isi body dari handler
	fmt.Println(string(body))
}

// mutiple value query parameter
// sebenarnya url melakukan parsing query parameter dan menyimpannya dalam map[string][]string (key string, value string)
// artinya, dalam satu key query parameter, kita bisa memasukkan beberapa value
// caranya kita bisa menambahkan query parameter dengan nama query parameter yang sama, namun-
// value nya berbeda, contoh: name=taufik&name=hidayat, dalam satu url endpoint yang sama
// kalau menggunakan function Get(), maka akan sellau mendapatkan data query parameter yang pertama (meski value nya lebih dari satu)

// membuat function handler baru untuk simulasi
func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	// mendapatkan query parameter yang nilainya lebih dari satu
	query := request.URL.Query() // cukup sampai query

	// .GET(), mendapatkan data paling pertama dari setiap query parameter
	// untuk mendapatkan semua data/value yang lebih dari satu pada query parameter bisa menggunakan-
	names := query["nama"]

	// menampilkan hasil dari kumpulan value di query parameter nama
	// mengkombinasikan string dengan function strings.Join(), separator (pemisah) menggunakan spasi
	fmt.Fprintln(writer, strings.Join(names, " "))
}

// membuat kode uji baru
func TestMultipleParamerValues(t *testing.T) {
	// membuat request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?nama=Mohammad&nama=Ilham&nama=Nurizky", nil)

	// membuat recoder (simulasi sebagai writer)
	recorder := httptest.NewRecorder()

	// memanggil endpoint simulasi, yaitu MultipleParameterValues()
	MultipleParameterValues(recorder, request)

	// melihat hasil dari recorder
	response := recorder.Result()

	// melihat isi body dari recorder
	body, _ := io.ReadAll(response.Body)

	// menampilan isi body dari handler
	fmt.Println(string(body))
}
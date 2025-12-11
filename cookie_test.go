package belajar_go_lang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// stateless
// adalah server tidak menyimpan data apapun untuk mengingat setiap request dari client, yang bertujuan-
// agar mudah melakukan skalabilitas di sisi server (agar tidak mengingat ke satu server tertentu)
// http merupakan stateless antara client dan server
// namun hal ini perlu dilakukan agar data tetap di simpan di server, seperti contoh di kasus login-
// maka hal ini bisa ditangani dari cookie (data di simpan dari sisi client)

// cookie
// adalah fitur di HTTP dimana server bisa memberi response cookie (yang berisi key-value seperti query parameter)-
// dan client akan menyimpan cookie tersebut di web browser
// maka request selanjutnya, client akan membawa cookie tersebut secara otomatis
// dan server secara otomatis juga akan selalu menerima data cookie yang dibawa oleh client setiap kali client mengirim request

// membuat cookie
// cookie merupakan data yang dibuat di server dan sengaja agar data di simpan di web browser
// untuk membuat cookie di server, kita bisa menggunakan function http.SetCookie()
// maka dari itu jangan terlalu banyak, membuat data cookie untuk client, agar proses request tidak lambat

// membuat handler cookie
func SetCookie(writer http.ResponseWriter, request *http.Request) {
	// membuat cookie
	cookie := new(http.Cookie)

	// menambahkan key baru untuk cookie
	cookie.Name = "X-PZN-Name"

	// mengisikan nilai value untuk key name dari request name
	cookie.Value = request.URL.Query().Get("name")

	// mengaktifkan path yang di izinkan untuk cookie dibawa terus oleh client
	// jika menggunakan path '/', maka cookie akan selalu dibawa oleh client di setiap requestnya
	cookie.Path = "/"

	// memasukkan cookie ke browser (client)
	http.SetCookie(writer, cookie)

	fmt.Fprint(writer, "Succes create cookie")
}

// membuat handler untuk membaca/mengambil cookie
func GetCookie(writer http.ResponseWriter, request *http.Request) {
	// mengambil cookie
	// request.Cookies() : untuk mengambil semua data cookies yang di miliki oleh client
	// request.Cookie() : untuk mengambil data cookie berdasarkan nama key nya
	cookie, err := request.Cookie("X-PZN-Name") // mengembalikan 2 nilai (cookie, dan error)

	// mengecek apakah cookie nya ada atau tidak
	if err != nil {
		fmt.Fprint(writer, "Cookie not found")
	} else {
		// menyimpan data cookie ke dalam variabel
		name := cookie.Value // karena cookie bersifat (key-value), maka cara mengambil nilainya dengan .Value
		fmt.Fprintf(writer, "Hello %s", name) 
	}
}

// membuat kode uji cookie
func TestCookie(t *testing.T) {
	// membuat server mux
	mux := http.NewServeMux()

	// membuat endpoint
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	// menjalankan server
	err := server.ListenAndServe()

	// mengecek apakah error pada saat running server
	if err != nil {
		panic(err)
	}
}

// mengambil cookie dari sisi pengguna
func TestSetCookie(t *testing.T) {
	// membuat request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Taufik", nil)

	// membuat recorder 
	recorder := httptest.NewRecorder()

	// memanggil set cookie
	SetCookie(recorder, request)

	// menampilkan hasil recorder untuk mendapatkan cookies (output nya adalah []cookie)
	cookies := recorder.Result().Cookies()

	// menampilkan cookies dengan perulangan
	for _, cookie := range cookies {
		fmt.Printf("Cookie %s : %s", cookie.Name, cookie.Value)
	}
}

// mengirim cookie ke server
func TestGetCookie(t *testing.T) {
	// membuat request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat cookie
	cookie := new(http.Cookie)
	
	// set cookie name dan value
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Taufik Hidayat"

	// menambahkan cookie ke request server
	request.AddCookie(cookie)

	// membuat recorder 
	recorder := httptest.NewRecorder()

	// memanggil get cookie
	GetCookie(recorder, request)

	// mengambil isi nya
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan isi body
	fmt.Println(string(body))
}
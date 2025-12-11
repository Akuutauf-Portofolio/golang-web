package belajar_go_lang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

// serve file
// kadang kita ingin menggunakan static file sesuai dengan yang kita inginkan
// hal ini bisa dilakukan dengan menggunakan function http.ServeFile()
// dengan menggunakan function tersebut, kita bisa menentukan file mana yang ingin kita tulis ke http response

// membuat handler untuk mengambil file berdasarkan query parameter
func ServeFile(writer http.ResponseWriter, request *http.Request) {
	// mengecek query parameter

	// jika query parameter tidak kosong
	if request.URL.Query().Get("name") != "" {
		// maka akan menampilkan file tertentu
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		// jika file kosong akan menampilkan file notfound.html
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

// membuat kode uji
func TestServeFileServer(t *testing.T) {
	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile), // menggunakan satu handler yaitu handler ServeFile
	}

	// menjalankan server 
	err := server.ListenAndServe()

	// mengecek error pada saat menjalankan server
	if err != nil {
		panic(err)
	}
}

// golang embed (studi kasus ServeFile)
// parameter function http.ServeFile hanya berisi string file name, sehingga tidak bisa menggunakan golang embed
// namun bukan berarti kita tidak bisa menggnakan golang embed, karena jika untuk melakukan load file,-
// maka kita hanya butuh package fmt, dan ResponseWriter saja

// membuat golang embed string
//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

// membuat kode uji embed yang implementasinya mirip dengan ServeFile (ServeFile tidak bisa menggunakan embed)-
// maka alternatif nya juga bisa pakai embed string seperti ini


func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	// mengecek query parameter

	// jika query parameter tidak kosong
	if request.URL.Query().Get("name") != "" {
		// jika pakai embed untuk studi kasus ServeFile satu persatu bisa langsung menggunakan Fprint
		fmt.Fprint(writer, resourceOk)
	} else {
		// jika pakai embed untuk studi kasus ServeFile satu persatu bisa langsung menggunakan Fprint
		fmt.Fprint(writer, resourceNotFound)
	}
}

// membuat kode uji (embed)
func TestServeFileServerEmbed(t *testing.T) {
	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed), // menggunakan satu handler yaitu handler ServeFile
	}

	// menjalankan server 
	err := server.ListenAndServe()

	// mengecek error pada saat menjalankan server
	if err != nil {
		panic(err)
	}
}
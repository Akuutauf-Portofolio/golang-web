package belajar_go_lang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// download file
// selain upload file, golang juga menyediakan fitur untuk mendownload file
// sebenarnya di golang sudah menyediakan fitur file server atau ServeFile-
// namun dalam implementasinya, masih diharuskan untuk melakukan rendering file (menampilkan)
// sehingga jika kita ingin langsung mengunduh fila tanpa di render, kita bisa ubah header-
// dari Content-Dispositoion menjadi 'attachment'

// mebuat handler untuk pencarian file (render)
func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	// mengambil nama file berdasarkan query parameter user
	namaFile := request.URL.Query().Get("file")

	// mengecek jika nama file kosong dari query parameter
	if namaFile == "" {
		// mengembalikan respon ke writer (pengguna), dengan bad request
		writer.WriteHeader(http.StatusBadRequest) // menulis code 400 pada header

		fmt.Fprint(writer, "Bad Request")

		// kalau misalnya sudah dicek, maka jangan dilanjutkan untuk dieksekusi ke kode berikutnya
		return
	}

	// jika ada nama file, maka akan dicarikan filenya, untuk dilakukan render (ditampilkan)
	// jika ketemu file yang sesuai dengan nama di query parameter, maka akan ditampilkan
	// namun jika tidak ada, maka akan menampilkan not found (404)
	http.ServeFile(writer, request, "./resources/" + namaFile)
}

// membuat kode uji untuk download
func TestDownloadFile(t *testing.T) {
	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	// menjalankan server
	err := server.ListenAndServe()

	// mengecek error
	if err != nil {
		panic(err)
	}
}

// mebuat handler untuk download file (download)
func DownloadFileHeader(writer http.ResponseWriter, request *http.Request) {
	// mengambil nama file berdasarkan query parameter user
	namaFile := request.URL.Query().Get("file")

	// mengecek jika nama file kosong dari query parameter
	if namaFile == "" {
		// mengembalikan respon ke writer (pengguna), dengan bad request
		writer.WriteHeader(http.StatusBadRequest) // menulis code 400 pada header

		fmt.Fprint(writer, "Bad Request")

		// kalau misalnya sudah dicek, maka jangan dilanjutkan untuk dieksekusi ke kode berikutnya
		return
	}

	// menambahkan header untuk memaksa pengguna agar melakukan download (tidak akan dirender)
	writer.Header().Add("Content-Disposition", "attachment; filename=\"" + namaFile + "\"")

	// jika ada nama file, maka akan dicarikan filenya, untuk dilakukan download
	// jika ketemu file yang sesuai dengan nama di query parameter, maka akan di download
	// namun jika tidak ada, maka akan menampilkan not found (404)
	http.ServeFile(writer, request, "./resources/" + namaFile)
}

// membuat kode uji untuk download
func TestDownloadFileHeader(t *testing.T) {
	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(DownloadFileHeader),
	}

	// menjalankan server
	err := server.ListenAndServe()

	// mengecek error
	if err != nil {
		panic(err)
	}
}
package belajar_go_lang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// file server
// golang memiliki sebauah fitur bernama file server
// file server adalah handler, digunakan untuk load file secara otomatis (tidak manual lagi)
// file server sebagai static file, bisa kita tambahkan ke dalam http.Serve atau http.ServeMux

// func pengujian baru
func TestFileServer(t *testing.T) {
	// membuat direktori yang digunakan sebagai tempat penyimpanan file
	direktory := http.Dir("./resources")

	// membuat file server dengan menambahkan direktori
	fileServer := http.FileServer(direktory) // output nya berupa handler

	// membuat server
	mux := http.NewServeMux()

	// menambahkan endpoint
	// mux.Handle("/static/", fileServer) // tanpa strip prefix
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // pakai strip prefix: untuk menghapus prefix '/static'

	// ketika dijalankan dengan memanggil endpoint '/static/index.html' maupun '/static/', maka-
	// file tidak ditemukan, karena by default, file harus tersimpan di folder seperti ini 'resource/static/index.html'
	// nah jikalau kita tidak mau untuk menambahkan folder lagi di dalam resources, maka tambahkan kode ini-
	// http.StripPrefix() : untuk menghapus prefix di url

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	// menjalankan server 
	err := server.ListenAndServe()

	// mengecek error saat runnning server
	if err != nil {
		panic(err)
	}
}

// golang embed (studi kasus file server)
// di golang versi 1.16 ke atas, terdapat fitur yang bernama go lang embed
// dalam golang embed, kita bisa embed file ke dalam binary distribution file (direktori file server),-
// agar memermudah dalam proses uploading ke server nantinya (tidak perlu mengcopy statis file lagi)
// golang embed juga memiliki fitur bernama embed.FS, fitur ini bisa diintegrasikan dengan file server

// melakukan embed untuk resources
//go:embed resources
var resources embed.FS // semua file akan masuk ke embed variabel resources (bertipe file sistem)

// membuat kode uji
func TestFileServerGolangEmbed(t *testing.T) {
	// namun jikalau ingin menghapus '/resources' maka bisa menggunakan fs.Sub
	// membuat direktori dengam embed, dan menghapus sub direktori ('/resources')
	direktori, _ := fs.Sub(resources, "resources")

	// membuat file server dengan menambahkan direktori
	// mengkonversi embed resources (FS) 
	// kalau menggunakan embed, nanti prefix nya akan berubah seperti ini :-
	// http://localhost:8080/static/resources/ (secara default)

	// jadi jika menggunakan fs.sub, maka direktori yang dipanggil adalah variabel direktori baru bukan resource (embed)
	fileServer := http.FileServer(http.FS(direktori)) // output nya berupa handler

	// membuat server
	mux := http.NewServeMux()

	// menambahkan endpoint
	// mux.Handle("/static/", fileServer) // tanpa strip prefix
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	// menjalankan server 
	err := server.ListenAndServe()

	// mengecek error saat runnning server
	if err != nil {
		panic(err)
	}
}


package belajar_go_lang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// redirect
// biasanya digunakan untuk melekukan peralihan halaman ke halaman lain setelah melakukan proses tertentu
// contoh : setelah selesai login, kita lakukan redirect ke halaman dashboard
// redirect sendiri sebenarnya sudah standar di HTTP
// kita perlu membuat response code 3xx dan menambah header location
// namun di golang, ada function bisa kita gunakan untuk mempermudah hal ini-
// (sehingga tidak manual set response code dan header location)

// membuat handler sebagai alamat tujuan yang menampilkan hasil teks
func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	// menampilkan ke pengguna mau redirect ke mana
	fmt.Fprint(writer, "Hello Redirect")
}

// membuat handler redirect ke alamat tujuan secara langsung
func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	// melakukan redirect dengan function http.Redirect
	// untuk status redirect ada 2, permanen dan temporal
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect) // contoh redirect ke website sendiri
}

func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	// melakukan redirect dengan function http.Redirect
	// untuk status redirect ada 2, permanen dan temporal
	http.Redirect(writer, request, "https://www.google.com", http.StatusPermanentRedirect) // contoh redirect ke website lain
}

// membuat pengujian redirect
func TestRedirect(t *testing.T) {
	// membuat server mux, dan mendaftarkan handler ke endpoint redirect
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
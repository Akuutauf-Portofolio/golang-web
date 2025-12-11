package belajar_go_lang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// server
// adalah struct yang terdapat pada package net/http yang digunakan sebagai representasi web server di golang
// untuk membuat web, kita wajib membuat/menjalankan server
// saat membuat server, kita perlu menentkan host (localhost), port (8080)
// direkomendasikan untuk port server menggunakan 4 digit, jangan 2 digit (80), karena harus run as admin
// ListenAndServe() : untuk menjalankan server
// fmt.Fprint() : untuk menampilkan pesan ke sisi end user 
// fmt.Fprintf() : untuk menampilkan pesan ke sisi end user (menggunakan golang formatter)

// membuat function pengujian
func TestServer(t *testing.T) {
	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
	}

	// menjalankan 
	// ketika unit test dijalankan/server dirunning, maka unit test nya tidak akan berhenti
	// untuk memastikan server berjalan bisa membuka browser dan mengetikkan address localhost dan port nya
	err := server.ListenAndServe()
	
	// pengecekan error
	if err != nil {
		panic(err)
	}
}

// handler
// digunakan untuk menerima HTTP request yang masuk ke server
// handler di golang direpresentasikan dalam interface, dimana dalam kontraknya terdapat function-
// bernama ServeHTTP() : digunakan sebagai function yang akan di eksekusi ketika menerima HTTP request

// handler function
// salahs satu implementasi dari interface handler
// digunakan untuk membuat function handler HTTP (agar lebih mudah)

// membuat pengujian handler
func TestHandler(t *testing.T) {
	// membuat handler (dengan anonymous function)
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		// menggunakan fprint untuk menampilkan hasilnya ke writer
		// bisa menggunakan error untuk outputnya
		fmt.Fprint(writer, "Hello World")
	}

	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}
	
	// menjalankan server
	err := server.ListenAndServe()
	
	// pengecekan error
	if err != nil {
		panic(err)
	}
}

// serveMux
// adalah implementasi handler yang bisa mendukung mutiple endpoint
// juga merupakan alternatif implementasi dari handler 
// sehingga nanti serveMux dapat memiliki banyak handler yang bisa kita atur ke sebuah enpoint
// serveMux sama seperti router

func TestServeMux(t *testing.T) {
	// membuat serveMux
	mux := http.NewServeMux()

	// membuat handler dengan endpoint
	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		// aksi yang mau kita tampilkan ke pengguna, hasil endpoint
		fmt.Fprint(writer, "Hello World")
	})
	
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, r *http.Request) {
		// aksi yang mau kita tampilkan ke pengguna, hasil endpoint
		fmt.Fprint(writer, "Hi!")
	})

	// priorias url servemux
	// 1. ketika endpoint '/images/' di panggil, maka akan muncul image. Namun jika dibelakang endopoint-
	// ditambahkan seperti ini '/images/taufik', maka akan tetap memanggil image (bukan not found)
	// 2. ketika endpoint '/images/thumbnails/' dipangggil, maka akan muncul thumbnail. Namun jika dibelakang-
	// endpoint typo seperti ini '/images/thumbnailssss', maka yang dipanggil akan image
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, r *http.Request) {
		// aksi yang mau kita tampilkan ke pengguna, hasil endpoint
		fmt.Fprint(writer, "Image")
	})
	
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, r *http.Request) {
		// aksi yang mau kita tampilkan ke pengguna, hasil endpoint
		fmt.Fprint(writer, "Thumbnail")
	})

	// pastikan endpoint harus unik satu sama lain, jikalau ada yang double, maka akan tertimpa dengan-
	// endpoint yang paling baru

	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux, // mux adalah satu satu implementasi yang kontraknya mengikuti handler
	}
	
	// menjalankan server
	err := server.ListenAndServe()
	
	// pengecekan error
	if err != nil {
		panic(err)
	}
}

// request
// adalah struct yang merepresentasikan HTTP request yang dikirim oleh web browser
// semua informasi request yang dikirim bisa kita dapatkan di request
// informasi seperti url, http method, http header, http body, dan lain lain

func TestRequest(t *testing.T) {
	// membuat handler
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// mengambil method http request (GET, POST, PUT)
		fmt.Fprint(writer, request.Method)

		// mendapatkan request url/uri
		fmt.Fprint(writer, request.RequestURI)
	}

	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	// menjalankan server
	err := server.ListenAndServe()

	// mengecek error
	if err != nil {
		panic(err)
	}	
}
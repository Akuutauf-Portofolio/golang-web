package belajar_go_lang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// middleware
// adalah sebuah fitur dimana kita bisa menambahkan kode sebelum dan setelah sebuah handler di eksekusi
// dalam pembuatan web, ada konsep yang bernama middleware atau filter atau interceptor
// contoh studi kasus login, untuk validasi user yang sudah login, akan di berikan akses atau tidak
// namun di golang tidak ada istilah middleware jadi hanya handler
// tetapi karena struktur handler yang baik menggunakan interface, maka kita bisa membuat middleware sendiri-
// dengan menggunakan handler

// membuat log dengan middleware
// membuat struct log middleware
type LogMiddleware struct {
	// membuat atribut handler, merujuk ke function handler yang asli
	// jadi nanti jika ada request masuk, maka akan masuk ke middleware terlebih dahulu (handler bawaan)-
	// kemudian dilanjutkan ke handler yang sesuai dengan request, kemudian dikembalikan lagi
	Handler http.Handler
}

// membuat handler (method) milik struct LogMiddleware
// membuat handler yang mengikuti kontrak function ServeHTTP
func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// menampilkan log ke terminal sebelum request
	fmt.Println("Before Execute Handler")

	// memanggil middleware (struct), untuk atribut Handler (handler asli)-
	// menggunakan fungsi bawaan dari handler yang asli
	middleware.Handler.ServeHTTP(writer, request)
	
	// menampilkan log ke temrinal setelah request
	fmt.Println("After Execute Handler")
}

// error handler
// kadang middleware juga bisa digunakan untuk melakukan error handler
// sehingga ketika terjadi panic di handler, kita bisa melakukan recover di middleware, dan-
// mengubah panic tersebut menjadi error response
// dengan hal ini, kita bisa menjaga aplikasi kita agak tidak berhenti berjalan
// karena by default ketika golang menjalankan panic, maka aplikasi akan berhenti

// membuat error handler (struct)
type ErrorHandler struct {
	Handler http.Handler
}

// membuat method milik struct erorr handler
// membuat handler yang mengikuti kontrak function ServeHTTP
func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// sebelum menjalankan isi kontrak error handler, lakukan pengecekan error handler
	defer func ()  {
		// melakukan recover error
		err := recover()

		// kalau misalnya terjadi error, maka akan dilakukan recover
		if err != nil {
			// menampilkan pemberitahuan error ke terminal
			fmt.Println("Terjadi Error")
			
			// menambahkan kode error
			writer.WriteHeader(http.StatusInternalServerError)

			// tampilkan error nya apa, dalam bentuk format
			fmt.Fprintf(writer, "Error: %s", err)
		}
	} ()
	
	// jalankan kontrak ketika tidak ada error
	// isi kontak serve http
	errorHandler.Handler.ServeHTTP(writer, request)
}

// membuat kode uji untuk middleware
func TestMiddleware(t *testing.T) {
	// membuat server mux
	mux := http.NewServeMux()

	// mendaftarkan endpoint 1 ke server mux
	mux.HandleFunc("/", func (writer http.ResponseWriter, request *http.Request)  {
		// menampilkan ke terminal
		fmt.Println("Handler Executed")

		// menampilkan ke pengguna
		fmt.Fprint(writer, "Hello Middleware")

		// ketika endpoint ini dijalankan, maka nanti akan memanggil ke terminal 2x request (memanggil fav.icon dan localhost)
	})

	// mendaftarkan endpoint 2 ke server mux
	mux.HandleFunc("/foo", func (writer http.ResponseWriter, request *http.Request)  {
		// menampilkan ke terminal
		fmt.Println("Foo Executed")

		// menampilkan ke pengguna
		fmt.Fprint(writer, "Hello Foo")

		// ketika endpoint ini dijalankan, maka nanti akan memanggil ke terminal 2x request (memanggil fav.icon dan localhost)
	})
	
	// mendaftarkan endpoint 3 ke server mux (khusus studi kasus panic)
	mux.HandleFunc("/panic", func (writer http.ResponseWriter, request *http.Request)  {
		// menampilkan ke terminal
		fmt.Println("Panic Executed")
		panic("ups")
	})

	// membuat object middleware baru
	logMiddleware := new(LogMiddleware)

	// mendaftarkan server mux ke handler milik log middleware
	logMiddleware.Handler = mux

	// membuat object error handler
	errorHandler := new(ErrorHandler)

	// mendaftarkan logMiddleware ke dalam error handler
	errorHandler.Handler = logMiddleware

	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: errorHandler, // karena error handler merupakan handler paling atas.

		// sehingga request nya dilewatkan dari server > error handler > log middleware > handler
	}

	// menjalankan server
	err := server.ListenAndServe()

	// mengecek error
	if err != nil {
		panic(err)
	}
}
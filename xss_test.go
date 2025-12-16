package belajar_go_lang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// xss (cross site scripting)
// xss adalah salah satu security issue yang terjadi ketika membuat web
// xss adalah celah keamanan, dimana orang bisa secara sengaja memasukkan parameter yang mengandung-
// javascript agar dirender oleh halaman website kita
// biasanya tujuan dari xss adalah mencuri cookie browser pengguna yang sedang mengakses website kita
// xss bisa menyebabkan account pengguna kita diambil (account take over) alih jika tidak ditangani dengan baik

// auto escape
// berbeda dengan bahasa pemrograman lain seperti PHP, pada golang template masalah XSS sudah diatasi secara otomatis
// golang template memiliki fitur auto escape, dimana dia bisa mendeteksi data yang perlu ditampilkan di template,
// jika mengandung tag tag html atau script, secara otomatis akan di escape
// dokumentasi auto escape golang : https://go.dev/src/html/template/escape.go
// kalau menggunakan text template maka tidak bisa melakukan auto escape, namun jika menggunakan html template maka bisa

// membuat handler auto escape
func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	// menggunakan parsing my template dari var global sebelumnya di file (template_caching_test.go)
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{} {
		// tambahkan data yang ingin di render
		"Title": "Template Auto Escape",

		// kalau data yang di render berupa tag html seperti dibawah ini, maka outputnya seperti ini (escape):-
		// &lt;p&gt;Ini adalah body paragraf&lt;/p&gt;
		// akan di escape, sehingga secara default golang akan otomatis melakukan escape untuk script yang tidak di izinkan
		"Body": "<p>Ini adalah body paragraf<script>alert('Anda di Hack!')</script></p>",
	})
}

// membuat kode uji
func TestTemplateAutoEscape(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateAutoEscape
	TemplateAutoEscape(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	fmt.Println(string(body)) 
}

// membuat pengujian request
func TestTemplateAutoEscapeServer(t *testing.T) {
	// membuat server
	// menggunakan handler dari template auto escape sebelumnya
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	// menjalankan server
	// ketika dijalankan akan keluar tag p sebagai string, bukan berupa paragraf
	err := server.ListenAndServe()

	// mengecek error
	if err != nil {
		panic(err)
	}
}

// mematikan auto escape
// jika kita mau, auto escape juga bisa dimatikan (tidak secara global, hanya untuk kode tertentu yang kita inginkan)
// namun, kita perlu memberi tahu template secara eksplisit (jelas), ketika menambahkan template bahwasannya data mengandung html
// data yang kita kirimkan bisa berupa selain html, dan bisa menggunakan penerapan template seperti berikut:-
// template.HTML : jika data adalah html
// template.CSS : jika data adalah css
// template.JS : jika data adalah javascript

// membuat handler auto escape (dengan kode html)
func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	// menggunakan template dari var global
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{} {
		// tambahkan data yang ingin di render
		"Title": "Template Auto Escape",

		// dengan penggunaan data template html, maka kode html di dalam nya akan di render secara html
		"Body": template.HTML("<p>Ini adalah body paragraf<script>alert('Anda di Hack!')</script></p>"),
	})
}

// membuat kode uji
func TestTemplateAutoEscapeDisabled(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateAutoEscapeDisabled
	TemplateAutoEscapeDisabled(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	fmt.Println(string(body)) 
}

// membuat pengujian request
func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	// membuat server
	// menggunakan handler dari template auto escape disabled
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	// menjalankan server
	// ketika dijalankan akan keluar tag p sebagai html sebagaimana mestinya
	err := server.ListenAndServe()

	// mengecek error
	if err != nil {
		panic(err)
	}
}

// masalah xss
// saat kita mematikan fitur auto escape (untuk salah satu kode), bisa dipastikan masalah XSS akan mengintai kita
// jadi pastikan kita benar benar percaya terhadap sumber data yang kita matikan auto escape nya-
// contoh yang kita percaya adalah di query dan database, jangan yang dikirim oleh user ditampilkan secara bulat-bulat

// membuat handler auto escape (kasus menerima data dari input pengguna melalui query parameter)
func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	// menggunakan template dari var global
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{} {
		// tambahkan data yang ingin di render
		"Title": "Template Auto Escape",

		// tetap menggunakan template HTML untuk merender html, dengan data dari pengguna
		// jangan pernah percaya terhadap data yang dikirimkan oleh pengguna (berbahaya)
		// contoh jika kita percaya dan memasukkan data dari pengguna
		// mengambil data bulat bulat yang dikirim oleh user

		// kalau ingin menguji website dengan xss, jika kita mengirimkan parameter di query muncul tag html yang kita buat,-
		// maka website tersebut restan terhadap xss
		"Body": template.HTML(request.URL.Query().Get("body")),
	})
}

// membuat kode uji
func TestTemplateXSS(t *testing.T) {
	// membuat request baru
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?body=<p>alert</p>", nil)

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest TemplateXSS
	TemplateXSS(recorder, request)

	// memabaca hasil request
	body, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	fmt.Println(string(body)) 
}

// membuat pengujian request
func TestTemplateXSSServer(t *testing.T) {
	// membuat server
	// menggunakan handler dari template xss
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	// menjalankan server
	// ketika dijalankan akan keluar tag p sebagai html sebagaimana mestinya
	err := server.ListenAndServe()

	// mengecek error
	if err != nil {
		panic(err)
	}
}
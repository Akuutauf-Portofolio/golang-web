package belajar_go_lang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// upload file
// saat membuat web, selain menerima input berupa form atau query parameter,-
// kadang kita juga menerima input data berupa file dari client
// golang web sudah memiliki fitur untuk manajemen upload file
// hal ini memudahkan kita ketika butuh membuat web yang menerima input file upload (gambar, doc, dll)

// multipart
// kalau upload file biasanya inputnya berupa multipart
// saat kita ingin menerima upload file, kita perlu melakukan parsing terlebih dahulu menggunakan-
// Request.ParseMultipartForm(size), atau kita bisa langsung ambil data file nya menggunakan
// Request.FormFile(name), didalam function ini secara otomatis melakukan parsing terlebih dahulu
// hasilnya merupakan data data yang terdapat pada package multipart, seperti
// multipart.File sebagai representasi filenya, dan multipart.FileHeader sebagai informasi file nya

// membuat handler untuk render form (tanpa mengirimkan data)
func UploadForm(writer http.ResponseWriter, request *http.Request) {
	// merender form
	err := myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)

	// melakukan pengecekan error
	if err != nil {
		panic(err)
	}
}

// membuat handler untuk menangani upload file dari user
func Upload(writer http.ResponseWriter, request *http.Request) {
	// Tahap 1 : mengambil file dari input data
	// mendapatkan data menggunakan Request.FormFile (parsing cukup dilakukan sekali tidak berkali kali oleh FormFile)
	// FormFile memiliki default memory (32 Mb), bisa diubah ukuran maksimal yang bisa user unggah untuk filenya
	// FormFile memiliki return value berupa file, file header dan error
	file, fileHeader, err := request.FormFile("file")

	// mengecek error
	if err != nil {
		panic(err)
	}

	// Tahap 2 : mengecek folder yang digunakan untk menyimpan sudah ada atau belum
	// mengecek folder resources sudah ada 
	err = os.MkdirAll("./resources", os.ModePerm)

	// mengecek error
	if err != nil {
		panic(err)
	}

	// Tahap 3 : membuat destinasi file nya
	// menyimpan file dengan function os.Create
	// dengan nama bawaan file yang diupload oleh user, dengan return value berupa file dan error
	// fileDestination, err := os.Create("./resources" + fileHeader.Filename)

	// gunakan nama asli file
	filename := fileHeader.Filename
	filePath := "./resources/" + filename // digabungkan nama asli file dengan resource

	fileDestination, err := os.Create(filePath)

	// mengecek error
	if err != nil {
		panic(err)
	}

	// Tahap 4 : simpan semua file nya ke destinasi yang sudah dibuat
	// menyalin file dengan function io.Copy
	// dengan melampirkan parameter fileDestination sebagai destination writer, dan file sebagai source io reader
	// dan return valuenya adalah size (abaikan), dan error
	_, err = io.Copy(fileDestination, file)

	// mengecek error
	if err != nil {
		panic(err)
	}

	// mengambil data yang bukan file (text)
	// dengan function PostFormValue, dengan return nya adalah string
	name := request.PostFormValue("name")

	// membuat template
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{} {
		"Name": name,

		// file akan di simpan dan diakses di static
		"File": "/static/" + fileHeader.Filename,
	})
}

// membuat server pengujian upload form
func TestUploadForm(t *testing.T) {
	// membuat server mux, dan mendaftarkan endpoint ke server mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	
	// menambahkan upload dan static
	mux.HandleFunc("/upload", Upload)

	// menggunakan handle karena untuk membaca file
	// menggunakan strip prefix, agar di hapus path 'static' nya. static hanya berfungsi sebagia pengalihan dari resources
	// menggunakan file server dengan direktori resources, karena tadi file yang asli di simpan di sana
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	// menjalankan server
	err := server.ListenAndServe()

	// mengecek error
	if err != nil {
		panic(err)
	}
}

// penggunaan embed untuk memudahkan konversi binary code pada file
//go:embed resources/top-logo-white.png
var uploadFiletest []byte

// membuat kode uji untuk upload file
func TestUploadFile(t *testing.T) {
	// membuat body kosong
	// karena nanti akan mengirimkan data berupa file, maka data harus berupa binary-
	// nah hal ini bisa lebih mudah dilakukan dengan menggunakan package multipart
	body := new(bytes.Buffer)

	// mengisi body kosong menggunakan writer
	writer := multipart.NewWriter(body)

	// mengisi data name, karena berupa text bisa di isi menggunakan function WriteField
	// mengembalikan error
	writer.WriteField("name", "Taufik Hidayat")

	// mengupload file dalam body, menggunakan function CreateFromFile
	// mengembalikan output berupa io writer, file, dan error
	file, _ := writer.CreateFormFile("file", "SAMPLE.png")

	// melakukan upload file oleh writer, dengan bantuan embed untuk konversi binary code
	file.Write(uploadFiletest)
	writer.Close() // jangan lupa tutup writer, jika sudah selesai digunakan
	
	// membuat request baru
	// path upload sebenarnya tidak perlu, karena nanti akan langsung mengirimkan data ke body nya
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body) 

	// tambahkan header untuk content type, agar data tidak gagal pada saat di kirimkan untuk pengujian
	// kalau lupa memasukkan 'multipart/form-data' bisa menggunakan ini 'writer.FormDataContentType()'
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// membuat recorder
	recorder := httptest.NewRecorder()

	// memanggil function untuk merequest Upload
	Upload(recorder, request)

	// memabaca hasil request
	bodyResponse, _ := io.ReadAll(recorder.Result().Body)

	// menampilkan hasil isi body
	fmt.Println(string(bodyResponse)) 
}

package belajar_go_lang_web

// routing library
// walaupun golang sudah menyediakan serve mux sebagai handler yang bisa menghandle beberapa endpoint-
// atau istilahnya adalah routing
// tapi kebanyakan programmer golang biasanya akan menggunakan library untuk melakukan routing
// karena serve mux tidak memiliki advanced fitur seperti path variabel, auto binding parameter dan middleware
// banyak alternatif lain yang bisa kita gunakan untuk library routing selain serve mux, contoh:-
// http router, gorilla mux dan lain lain. referensi : https://github.com/julienschmidt/go-http-routing-benchmark

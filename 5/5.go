package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	target, _ := url.Parse("http://backend-service:8080")
// 	proxy := httputil.NewSingleHostReverseProxy(target)
// 	proxy.ServeHTTP(w, r)
// }

func handler(w http.ResponseWriter, r *http.Request) {
	a := "http://backend-service:8080"
	target, err := url.Parse(a)
	if err != nil {
		http.Error(w, "Invalid backend URL", http.StatusInternalServerError)
		log.Println("URL Parse Error:", err)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Proxy berjalan di :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

//Jawaban B
//Karena dari code sebelumnya itu mengabaikan sebuah kesalahan kecil namun berakibat fatal jika url.parse() nya gagal
//Karena tidak valid maka bisa menyebabkan nil yang dapat menyebabkan panic, code sebelumnya memang bisa di running tapi jika hasil dari url.parse() nya itu gagal
//Maka bisa menyebabkan nil

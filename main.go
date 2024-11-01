package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Listning on 0.0.0.0 with port 80")
	http.HandleFunc("/", handler)
	http.HandleFunc("/.well-known/atproto-did", handle_did)
	http.ListenAndServe("", nil)
}

func handle_did(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	subdomain := strings.Split(host, ".")[0]

	if subdomain != "www" && subdomain != "artism.place" && subdomain != "localhost" {
		fmt.Fprintf(w, get_usr(subdomain).did)
	} else {
		fmt.Fprint(w, "Welcome to the main domain")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	subdomain := strings.Split(host, ".")[0]

	if subdomain != "www" && subdomain != "artism.place" && subdomain != "localhost" {
		fmt.Fprintf(w, "welcome to the %s subdomain", subdomain)
	} else {
		fmt.Fprint(w, "Welcome to the main domain")
	}
}

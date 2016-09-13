package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"time"
)

func HashHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		password := r.PostFormValue("password")
		fmt.Printf(password + "\n")

		if password != "" {
			hash := sha512.Sum512([]byte(password))
			sha512_hash := base64.StdEncoding.EncodeToString(hash[:])
			time.Sleep(5 * time.Second)
			fmt.Printf(sha512_hash + "\n")

			fmt.Fprintf(w, "%s", sha512_hash)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	}

}

func main() {
	server := &http.Server {Addr: ":8080"}
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}

	http.HandleFunc("/", HashHandler)
	err = server.Serve(listener)
	if err != nil {
		fmt.Println("Error Serving:", err.Error())
	}
}

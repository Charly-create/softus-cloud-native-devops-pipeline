package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Softus cloud-native-devops-pipeline"))
	})
	http.ListenAndServe(":8080", nil)
}
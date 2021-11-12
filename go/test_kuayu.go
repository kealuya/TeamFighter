package main

import (
	"fmt"
	"net/http"
)

func mdain() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		allowOrigin := []string{"http://localhost:3000", ""}
		h := request.Header
		origin := h["Origin"]
		fmt.Println(origin)
		fmt.Println(request.Header)
		for _, v := range allowOrigin {
			if v == origin[0] {
				writer.Header().Set("Access-Control-Allow-Origin", v)
				break
			}
		}
		writer.Write([]byte("is ok"))
	})

	http.ListenAndServe(":8000", nil)

}

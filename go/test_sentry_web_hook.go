package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func maifn() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		body, _ := ioutil.ReadAll(request.Body)
		var b []byte
		bf := bytes.NewBuffer(b)
		json.Indent(bf, body, "", "    ")

		fmt.Println(string(bf.String()))
		writer.WriteHeader(200)
	})
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err)
	}

}

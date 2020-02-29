package main

import (
	"fmt"
	"net/http"

	"github.com/510909033/ucm_dev/config"
)

func h(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var s = ""

	fmt.Printf("projectCode=%s\n", r.FormValue("projectCode"))

	switch r.FormValue("projectCode") {
	case "Golang_Database":
		s = config.DatabaseConfig

	case "Golang_Memcache":
		s = config.MemcacheConfig
	case "Golang_Domain":
		s = config.DomainConfig
	case "go_empty":
		s = config.GoEmptyConfig

	}

	w.Write([]byte(s))
}

func main() {

	http.HandleFunc("/go_ucm/api/ucm/get_config", h)

	fmt.Println(http.ListenAndServe("0.0.0.0:9602", nil))
}
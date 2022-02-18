package main

import (
	"fmt"
	"log"
	"net/http"

)

func main() {
	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request){
		newHandler(res, req)
	})


	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func newHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("GET params were:", req.URL.Query())

	// if only one expected
	param1 := req.URL.Query().Get("param1")
	if param1 != "" {
		// ... process it, will be the first (only) if multiple were given
		// note: if they pass in like ?param1=&param2= param1 will also be "" :|
	}

	// if multiples possible, or to process empty values like param1 in
	// ?param1=&param2=something
	param1s := req.URL.Query()["param1"]
	if len(param1s) > 0 {
		for _, number := range param1s {
			Add()
		}
		// ... process them ... or you could just iterate over them without a check
		// this way you can also tell if they passed in the parameter as the empty string
		// it will be an element of the array that is the empty string
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorld02Response struct {
	Message string `json:"message"` // Result with tag json : {"message":"HelloWorld"}
	//Unfortunately we can't, as in Go, lowercase properties are not exported, Marshal will ignore these and will not include them in
	//the output.if I prefer to use camel case and would rather see "message",All is not lost as the encoding/json package implements struct field attributes that allow us to change the output for the
	//property to anything we choose.
	//Result without tag json : {"Message":"HelloWorld"}
}

/*
we can use field tags to control the output even further. We can convert object types and even
ignore a field altogether if we need to:

type helloWorld02Response struct {
	// change the output field to be "message"
	Message string `json:"message"`
	// do not output this field
	Author string `json:"-"`
	// do not output the field if the value is empty
	Date string `json:",omitempty"`
	// convert output to a string and rename "id"
	Id int `json:"id, string"`
}

Channel, complex types, and functions cannot be encoded in JSON; attempting to encode these types will result in an
UnsupportedTypeError being returned by the Marshal function.
It also can't represent cyclic data structures; if your stuct contains a circular reference then Marshal will result in an infinite
recursion, which is never a good thing for a web request.
*/

func helloWorld02MarshalHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorld02Response{Message: "HelloWorld"}
	//If we want to export our JSON prettily formatted with indentation, we can use the MarshallIndent function, this allows you to
	//pass an additional parameter of string to specify what you would like the indent to be.
	//func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
	//json.MarshalIndent(response, "", "    ")
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}

func helloWorld02EncoderHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorld02Response{Message: "HelloWorld"}

	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld1", helloWorld02MarshalHandler)
	//json.NewEncoder(w) is Faster than json.Marshal(response)
	//The astute reader might have noticed that we are decoding our struct into a byte array and then writing that to the response
	//stream, this does not seem to be particularly efficient and in fact it is not. Go provides Encoders and Decoders, which can write
	//directly to a stream, since we already have a stream with the ResponseWriter then let's do just that.
	http.HandleFunc("/helloworld2", helloWorld02EncoderHandler)

	log.Printf("Server starting on port %v\n", 8080)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

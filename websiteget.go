
package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
)

type User struct {
	Id      int
	Name    string
	Age     int
}

var allUsers = []User {
	{Id: 1,Name: "john", Age: 20},
	{Id: 2,Name: "mary", Age: 18},
	{Id: 3,Name: "mickey", Age: 100},

}

func plainHandler_1 (writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.Write([]byte("Id: 1,Name: john, Age: 20"))
}

func plainHandler_2 (writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.Write([]byte("Id: 2,Name: mary, Age: 18"))
}

func plainHandler_3 (writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.Write([]byte("Id: 3,Name: mickey, Age: 100"))
}

func xmlHandler_1 (writer http.ResponseWriter, request *http.Request) {
	x, err := xml.MarshalIndent(allUsers[0], "", "  ")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/xml")
	writer.Write(x)
}
func xmlHandler_2 (writer http.ResponseWriter, request *http.Request) {
	x, err := xml.MarshalIndent(allUsers[1], "", "  ")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/xml")
	writer.Write(x)
}
func xmlHandler_3 (writer http.ResponseWriter, request *http.Request) {
	x, err := xml.MarshalIndent(allUsers[2], "", "  ")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/xml")
	writer.Write(x)
}
func jsonHandler_1 (writer http.ResponseWriter, request *http.Request) {
	js, err := json.Marshal(allUsers[0])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(js)
}
func jsonHandler_2 (writer http.ResponseWriter, request *http.Request) {
	js, err := json.Marshal(allUsers[1])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(js)
}
func jsonHandler_3 (writer http.ResponseWriter, request *http.Request) {
	js, err := json.Marshal(allUsers[2])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(js)
}


func main() {
	http.HandleFunc("/plain/1", plainHandler_1)
	http.HandleFunc("/plain/2", plainHandler_2)
	http.HandleFunc("/plain/3", plainHandler_3)

	http.HandleFunc("/xml/1", xmlHandler_1)
	http.HandleFunc("/xml/2", xmlHandler_2)
	http.HandleFunc("/xml/3", xmlHandler_3)

	http.HandleFunc("/json/1", jsonHandler_1)
	http.HandleFunc("/json/2", jsonHandler_2)
	http.HandleFunc("/json/3", jsonHandler_3)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

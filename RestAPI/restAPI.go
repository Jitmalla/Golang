package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(w, "Hello World")
	// data := []byte("Hello World")
	// w.Write(data)
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, person := range people {
		if person.ID == params["id"] {
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for i, person := range people {
		if person.ID == params["id"] {
			people = append(people[:i], people[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

// UpdatePerson updates a particular record of a person.
func UpdatePerson(w http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	id := mux.Vars(req)["id"]
	var updatedPerson Person
	_ = json.NewDecoder(req.Body).Decode(&updatedPerson)
	// reqBody, _ := ioutil.ReadAll(req.Body)
	// json.Unmarshal(reqBody, &updatedPerson)
	for i, person := range people {
		if person.ID == id {
			person.FirstName = updatedPerson.FirstName
			person.LastName = updatedPerson.LastName
			people = append(people[:i], updatedPerson)
			json.NewEncoder(w).Encode(people)
		}
	}
}

func main() {
	
	// Mock data
	people = append(people, Person{ID: "1", FirstName: "James", LastName: "Smith", Address: &Address{City: "Dallas", State: "Texas"}})
	people = append(people, Person{ID: "2", FirstName: "Tony", LastName: "Lowe"})

	router := mux.NewRouter()

	router.HandleFunc("/home/people", GetPeople).Methods("GET")
	router.HandleFunc("/home/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/home/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/home/person/{id}", DeletePerson).Methods("DELETE")
	router.HandleFunc("/home/person/{id}", UpdatePerson).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8092", router))
}

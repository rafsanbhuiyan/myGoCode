package main          //package main to run all go Executables

/*
Coding Challenge:

Set up a single file that creates an http server exercising get, post, put, and delete. 
Within the language you can set up unit tests to exercise your code to make sure it works correctly.  
Document the functions and a small writeup as to how it works so I can run it.  
Bonus points for taking a stab at how you might approach securing your endpoints.
*/

//Rafsan Bhuiyan


//log to handle errors, net/http for restful Api
//https github site link
//Using Gorilla Mux library to use ne/http built in methods
import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

//Creating get function, to write and request
func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")  //header and set built in function
    w.WriteHeader(http.StatusOK) //write header built in function
    w.Write([]byte(`{"message": "get called"}`))
}

//creating a post function, with write and request
//header and set built in function
//write header built in function
func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post called"}`))
}

//creating a put function
//header and set built in function
//write header built in function
func put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`{"message": "put called"}`))
}

//Delete function
//header and set built in function
//write header built in function
func delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "delete called"}`))
}

//function not found
//header and set built in function
//write header built in function
//relay map with key "message" and value "not found"
func notFound(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(`{"message": "not found"}`))
}

//Main function and initiating the objects of the methods
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", get).Methods(http.MethodGet) //Handling get method
    r.HandleFunc("/", post).Methods(http.MethodPost) //Handling post
    r.HandleFunc("/", put).Methods(http.MethodPut) //Handling put
    r.HandleFunc("/", delete).Methods(http.MethodDelete) //Handling delete
    r.HandleFunc("/", notFound) //handling notfound
    log.Fatal(http.ListenAndServe(":8080", r)) //handling error on port 8080 for https
}

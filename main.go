package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meghashyamc/dbconnect/dbapi"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/addtodb", dbapi.Add).Methods("POST")
	router.HandleFunc("/getfromdb", dbapi.Get).Methods("POST")
	fmt.Println("db listening on port", dbapi.Port)
	http.ListenAndServe(":"+dbapi.Port, router)

}

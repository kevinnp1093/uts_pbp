package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/modul2/controllers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")

	router.HandleFunc("/roomsDetailed", controllers.GetDetailRooms).Methods("GET")

	router.HandleFunc("/rooms", controllers.InsertRooms).Methods("POST")

	router.HandleFunc("/rooms", controllers.LeaveRoom).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

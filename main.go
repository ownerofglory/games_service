package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/ownerofglory/games_service/handler"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/games", handler.GetGamesHandler).Methods(http.MethodGet)
	router.HandleFunc("/games/{id:[0-9]+}", handler.GetGameHandler).Methods(http.MethodGet)
	router.HandleFunc("/games", handler.AddGameHandler).Methods(http.MethodPost)
	router.HandleFunc("/games/{id:[0-9]+}", handler.DeleteGameHandler).Methods(http.MethodDelete)

	log.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

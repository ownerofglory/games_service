package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ownerofglory/games_service/model"
	"github.com/ownerofglory/games_service/service"
)

var srvs service.GameServiceOps

func init() {
	srvs = service.GameService{}
}

func GetGamesHandler(w http.ResponseWriter, r *http.Request) {
	games, err := srvs.GetAllGames()
	if err != nil {
		hanleError(err, &w)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(games)
}

func GetGameHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		hanleError(err, &w)
	}
	var game *model.Game
	game, err = srvs.GetGameById(id)
	if err != nil {
		hanleError(err, &w)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(game)
}

func AddGameHandler(w http.ResponseWriter, r *http.Request) {
	var game model.Game
	err := json.NewDecoder(r.Body).Decode(&game)
	defer r.Body.Close()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = srvs.CreateGame(&game)
	if err != nil {
		hanleError(err, &w)
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteGameHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		hanleError(err, &w)
	}

	err = srvs.DeleteGame(id)
	if err != nil {
		hanleError(err, &w)
	}

	w.WriteHeader(http.StatusOK)
}

func hanleError(err error, w *http.ResponseWriter) {
	json.NewEncoder(*w).Encode(err)
	log.Println(err)
	(*w).WriteHeader(http.StatusBadRequest)
}

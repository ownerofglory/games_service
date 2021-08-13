package service

import (
	"github.com/ownerofglory/games_service/model"
	"github.com/ownerofglory/games_service/repository"
)

var repo repository.CrudRepoOps

func init() {
	repo = repository.CrudRepo{}
}

type GameServiceOps interface {
	GetAllGames() ([]model.Game, error)
	GetGameById(id int64) (*model.Game, error)
	CreateGame(game *model.Game) error
	DeleteGame(id int64) error
}

type GameService struct{}

// service implememntation
func (s GameService) GetAllGames() ([]model.Game, error) {
	return repo.FindAllGames()
}

func (s GameService) GetGameById(id int64) (*model.Game, error) {
	return repo.FindGameById(id)
}

func (s GameService) CreateGame(game *model.Game) error {
	return repo.SaveGame(game)
}

func (s GameService) DeleteGame(id int64) error {
	return repo.DeleteGame(id)
}

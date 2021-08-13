package service

import (
	"testing"

	"github.com/ownerofglory/games_service/model"
)

type crudRepoMock struct{}

func init() {
	repo = crudRepoMock{}
}

var serv GameService

func TestGetAllGames(t *testing.T) {
	_, err := serv.GetAllGames()
	if err != nil {
		t.Error("Error occured", err)
	}
}

func TestGetGameById(t *testing.T) {
	testId := 10000
	_, err := serv.GetGameById(int64(testId))
	if err != nil {
		t.Error("Error occured", err)
	}
}

func TestCreateGame(t *testing.T) {
	game := model.Game{ID: 11111, Title: "Test Game"}
	err := serv.CreateGame(&game)
	if err != nil {
		t.Error("Error occured", err)
	}
}

func TestDeleteGame(t *testing.T) {
	testId := 10000
	err := serv.DeleteGame(int64(testId))
	if err != nil {
		t.Error("Error occured", err)
	}
}

// mock implementation
func (r crudRepoMock) FindAllGames() ([]model.Game, error) {
	return nil, nil
}

func (r crudRepoMock) FindGameById(id int64) (*model.Game, error) {
	return nil, nil
}

func (r crudRepoMock) SaveGame(game *model.Game) error {
	return nil
}

func (r crudRepoMock) DeleteGame(id int64) error {
	return nil
}

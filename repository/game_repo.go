package repository

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/ownerofglory/games_service/model"
)

type CrudRepo struct{}

type CrudRepoOps interface {
	FindAllGames() ([]model.Game, error)
	FindGameById(id int64) (*model.Game, error)
	SaveGame(game *model.Game) error
	DeleteGame(id int64) error
}

func init() {
	err := connectDb()
	if err != nil {
		log.Panic(err)
	}
}

func (r CrudRepo) FindAllGames() ([]model.Game, error) {
	games := make([]model.Game, 0)

	rows, err := db.Query("SELECT * FROM games g ORDER BY g.id")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var game model.Game
		rows.Scan(
			&game.ID,
			&game.Title,
		)

		games = append(games, game)
	}

	return games, nil

}

func (r CrudRepo) FindGameById(id int64) (*model.Game, error) {
	query := "SELECT * FROM games g WHERE g.id = $1"
	row := db.QueryRow(query, id)
	var game model.Game
	err := row.Scan(
		&game.ID,
		&game.Title,
	)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func (r CrudRepo) SaveGame(game *model.Game) error {
	var lastInsertedId int
	err := db.QueryRow("INSERT INTO games(id, title) VALUES($1, $2)", game.ID, game.Title).Scan(&lastInsertedId)
	return err
}

func (r CrudRepo) DeleteGame(id int64) error {
	err := db.QueryRow("DELETE FROM games g WHERE g.id = $1", id).Scan()
	return err
}

package beer_test

import (
	"database/sql"
	"testing"

	"github.com/WesGtoX/beer-go-api/core/beer"
	_ "github.com/mattn/go-sqlite3"
)

func TestStore(t *testing.T) {
	b := &beer.Beer{
		ID:    1,
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}
	defer db.Close()
	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro limpando o banco de dados: %s", err.Error())
	}
	service := beer.NewService(db)
	err = service.Store(b)
	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}
	saved, err := service.Get(1)
	if err != nil {
		t.Fatalf("Erro buscando do banco de dados: %s", err.Error())
	}
	if saved.ID != 1 {
		t.Fatalf("Dados inválidos. Esperado %d, recebido %d", 1, saved.ID)
	}
}

func clearDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from beer")
	tx.Commit()
	return err
}

//@todo implementar testes para as outras funções

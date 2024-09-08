package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	id          string
	name        string
	description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err = c.db.Exec("INSET INTO Categories (id,name,description) VALUES($1,$2,$3)", id, name, description)

	if err != nil {
		return Category{}, err
	}
	return Category{id: id, name: name, description: description}, nil

}

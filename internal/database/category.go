package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	Id          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO Category (id,name,description) VALUES($1,$2,$3)", id, name, description)

	if err != nil {
		return Category{}, err
	}
	return Category{Id: id, Name: name, Description: description}, nil

}

func (c *Category) Query() ([]Category, error) {
	Categories := []Category{}
	row, err := c.db.Query("SELECT id,name,descrition FROM Category")
	if err != nil {
		return nil, err
	}
	row.Close()

	for row.Next() {
		var id, name, description string
		err := row.Scan(&id, &name, &description)
		if err != nil {
			return nil, err
		}
		Categories = append(Categories, Category{Id: id, Name: name, Description: description})
	}
	return Categories, nil
}

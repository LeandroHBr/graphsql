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
	row, err := c.db.Query("SELECT id,name,description FROM Category")
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var id, name, description string
		err := row.Scan(&id, &name, &description)
		if err != nil {
			return nil, err
		}
		Categories = append(Categories, Category{Id: id, Name: name, Description: description})
	}
	row.Close()
	return Categories, nil
}

func (c *Category) Query_CategoryFindByCourse(CourseID string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow("SELECT c.id, c.name, c.description FROM Category c JOIN Courses co ON c.id = co.category WHERE co.id = $1", CourseID).Scan(&id, &name, &description)

	if err != nil {
		return Category{}, err
	}

	return Category{Id: id, Name: name, Description: description}, nil

}

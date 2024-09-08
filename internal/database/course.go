package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	Id          string
	Name        string
	Description string
	IDCategory  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) CreateCourse(name string, description string, IDCategory string) (Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO Courses (id,name,description,category) VALUES ($1,$2,$3,$4)", id, name, description, IDCategory)
	if err != nil {
		return Course{}, err
	}
	return Course{Id: id, Name: name, Description: description, IDCategory: IDCategory}, nil
}

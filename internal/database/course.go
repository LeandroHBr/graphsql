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

func (c *Course) Query_Course() ([]Course, error) {
	row, err := c.db.Query("SELECT id,name,description,category FROM Courses")
	if err != nil {
		return nil, err
	}
	courses := []Course{}
	for row.Next() {
		var id, name, description, id_categories string
		err2 := row.Scan(&id, &name, &description, &id_categories)
		if err2 != nil {
			return nil, err
		}
		courses = append(courses, Course{Id: id, Name: name, Description: description, IDCategory: id_categories})
	}
	row.Close()
	return courses, nil
}

func (c *Course) Query_CourseFindByCategory(category_id string) ([]Course, error) {
	row, err := c.db.Query("SELECT id,name,description,category FROM Courses WHERE category=$1", category_id)
	if err != nil {
		return nil, err
	}
	courses := []Course{}
	for row.Next() {
		var id, name, description, id_categories string
		err2 := row.Scan(&id, &name, &description, &id_categories)
		if err2 != nil {
			return nil, err
		}
		courses = append(courses, Course{Id: id, Name: name, Description: description, IDCategory: id_categories})
	}
	row.Close()
	return courses, nil
}

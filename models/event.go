package models

import (
	"go-project/database"
)

type Event struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	User_Id     int64  `json:"user_id"`
}

var data []Event

func (e *Event) Save() error {
	//perform to db
	query := `INSERT INTO events (name, location, description , user_id) VALUES (?, ?, ?, ?)`
	//Prepare is like to Exec and Query
	//but Prepare save Query on memory
	stmt, err := database.Db.Prepare(query)
	if err != nil {
		print("|" + err.Error())
		return err
	}
	//defer to close
	defer stmt.Close()
	//stmt wait for data (?)
	result, err := stmt.Exec(e.Name, e.Location, e.Description, e.User_Id)
	if err != nil {
		print("2" + err.Error())

		return err
	}

	//get the Id that autoGenerate
	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAll() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := database.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	// like itrator
	for rows.Next() {
		var event Event
		//scan the rows Like "fmt"
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.User_Id); err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id =?`

	//this method return 1 Row Instead of all Rows
	// id for the id that set (?)
	row := database.Db.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.User_Id)

	if err != nil {
		return &Event{}, err
	}

	return &event, nil

}

func (e Event) Update() error {

	query := `
	Update events
	set name =? , description =?,location =?
	where id =?
	`

	stmt, err := database.Db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.ID)

	return err
}

func (e Event) Delete() error {

	query := `
	DELETE FROM events 
	WHERE id = ?
	`

	stmt, err := database.Db.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

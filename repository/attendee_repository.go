package repository

import (
	"context"
	"database/sql"
	"jirbthagoras/event-management/domain/model"
	"jirbthagoras/event-management/exception"
)

type AttendeeRepository interface {
	Save(ctx context.Context, tx *sql.Tx, attendee *model.Attendee) (*model.Attendee, error)
	Update(ctx context.Context, tx *sql.Tx, attendee *model.Attendee) (*model.Attendee, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (*model.Attendee, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Attendee, error)
}

type AttendeeRepositoryImpl struct {
}

func NewAttendeeRepositoryImpl() *AttendeeRepositoryImpl {
	return &AttendeeRepositoryImpl{}
}

func (e AttendeeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, attendee *model.Attendee) (*model.Attendee, error) {
	query := "INSERT INTO attendees(name, email) values(?, ?)"

	result, err := tx.ExecContext(ctx, query, attendee.Name, attendee.Email)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	attendee.Id = int(id)

	return attendee, nil
}

func (e AttendeeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, attendee *model.Attendee) (*model.Attendee, error) {
	query := "UPDATE attendees SET name = ?, email = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, attendee.Name, attendee.Email, attendee.Id)
	if err != nil {
		return nil, err
	}
	return attendee, nil
}

func (e AttendeeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM attendees WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return err
}

func (e AttendeeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (*model.Attendee, error) {
	query := "SELECT name, email FROM attendees WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	attendee := model.Attendee{}

	if rows.Next() {
		err := rows.Scan(&attendee.Name, &attendee.Email)
		if err != nil {
			return nil, err
		}
		return &attendee, nil
	}
	return nil, exception.NewNotFoundError("Attendee Not Found")
}

func (e AttendeeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Attendee, error) {
	query := "SELECT id, name, email FROM attendees"
	rows, err := tx.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var attendees []*model.Attendee

	for rows.Next() {
		attendee := model.Attendee{}
		err = rows.Scan()
		if err != nil {
			return nil, err
		}
		attendees = append(attendees, &attendee)
	}

	return attendees, nil
}

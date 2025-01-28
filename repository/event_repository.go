package repository

import (
	"context"
	"database/sql"
	"jirbthagoras/event-management/domain/model"
	"jirbthagoras/event-management/exception"
)

type EventRepository interface {
	Save(ctx context.Context, tx *sql.Tx, event *model.Event) (*model.Event, error)
	Update(ctx context.Context, tx *sql.Tx, event *model.Event) (*model.Event, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (*model.Event, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Event, error)
}

type EventRepositoryImpl struct {
}

func NewEventRepositoryImpl() *EventRepositoryImpl {
	return &EventRepositoryImpl{}
}

func (e EventRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, event *model.Event) (*model.Event, error) {
	query := "INSERT INTO events(name, description, start_time, end_time) values(?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, query, event.Name, event.Description, event.StartTime, event.EndTime)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	event.Id = int(id)

	return event, nil
}

func (e EventRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, event *model.Event) (*model.Event, error) {
	query := "UPDATE events SET name = ?, description = ?, start_time = ?, end_time = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, event.Name, event.Description, event.StartTime, event.EndTime, event.Id)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (e EventRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM events WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return err
}

func (e EventRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (*model.Event, error) {
	query := "SELECT id, name, description, start_time, end_time FROM events WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	event := model.Event{}

	if rows.Next() {
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.StartTime, &event.EndTime)
		if err != nil {
			return nil, err
		}
		return &event, nil
	}
	return nil, exception.NewNotFoundError("Event Not Found")
}

func (e EventRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Event, error) {
	query := "SELECT id, name, description, start_time, end_time FROM events"
	rows, err := tx.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var events []*model.Event

	for rows.Next() {
		event := model.Event{}
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.StartTime, &event.EndTime)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
		return events, nil
	}

	return events, nil
}

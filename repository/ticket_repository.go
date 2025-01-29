package repository

import (
	"context"
	"database/sql"
	"jirbthagoras/event-management/domain/model"
	"jirbthagoras/event-management/exception"
)

type TicketRepository interface {
	Save(ctx context.Context, tx *sql.Tx, ticket *model.Ticket) (*model.Ticket, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (*model.Ticket, error)
	Update(ctx context.Context, tx *sql.Tx, ticket *model.Ticket) (*model.Ticket, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Ticket, error)
}

type TicketRepositoryImpl struct {
}

func NewTicketRepositoryImpl() *TicketRepositoryImpl {
	return &TicketRepositoryImpl{}
}

func (t TicketRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, ticket *model.Ticket) (*model.Ticket, error) {
	query := "INSERT INTO tickets(event_id, attendee_id) values(?, ?)"
	result, err := tx.ExecContext(ctx, query, ticket.Event.Id, ticket.Attendee.Id)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	ticket.Id = int(id)

	return ticket, nil
}

func (t TicketRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (*model.Ticket, error) {
	query := `
		SELECT 
			t.id, t.status,
			e.id, e.name, e.description, e.start_time, e.end_time,
			a.id, a.name, a.email
		FROM tickets t
		JOIN events e ON t.event_id = e.id
		JOIN attendees a ON t.attendee_id = a.id
		WHERE t.id = ?
	`
	row, err := tx.QueryContext(ctx, query, id)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	ticket := model.Ticket{}
	ticket.Event = &model.Event{}
	ticket.Attendee = &model.Attendee{}

	if row.Next() {
		err := row.Scan(
			&ticket.Id, &ticket.Status,
			&ticket.Event.Id, &ticket.Event.Name, &ticket.Event.Description, &ticket.Event.StartTime, &ticket.Event.EndTime,
			&ticket.Attendee.Id, &ticket.Attendee.Name, &ticket.Attendee.Email,
		)
		if err != nil {
			return nil, err
		}
		return &ticket, nil
	}

	return nil, exception.NewNotFoundError("Ticket Not Found")
}

func (t TicketRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, ticket *model.Ticket) (*model.Ticket, error) {
	query := "UPDATE tickets SET status = 'canceled' WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, ticket.Id)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (t TicketRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Ticket, error) {
	query := `
		SELECT 
			t.id, t.status,
			e.id, e.name, e.description, e.start_time, e.end_time,
			a.id, a.name, a.email
		FROM tickets t
		JOIN events e ON t.event_id = e.id
		JOIN attendees a ON t.attendee_id = a.id
	`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []*model.Ticket

	for rows.Next() {
		ticket := model.Ticket{}
		ticket.Event = &model.Event{}
		ticket.Attendee = &model.Attendee{}

		err := rows.Scan(
			&ticket.Id, &ticket.Status,
			&ticket.Event.Id, &ticket.Event.Name, &ticket.Event.Description, &ticket.Event.StartTime, &ticket.Event.EndTime,
			&ticket.Attendee.Id, &ticket.Attendee.Name, &ticket.Attendee.Email,
		)

		if err != nil {
			return nil, err
		}

		tickets = append(tickets, &ticket)
	}

	return tickets, nil
}

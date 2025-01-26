CREATE TABLE tickets(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    event_id INT NOT NULL,
    attendee_id INT NOT NULL,
    status ENUM('booked', 'canceled'),
    FOREIGN KEY (event_id) REFERENCES events(id),
    FOREIGN KEY (attendee_id) REFERENCES attendees(id)
)
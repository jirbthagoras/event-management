CREATE TABLE admin (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(200) NOT NULL,
    password VARCHAR(200) NOT NULL,
    token VARCHAR(36) NOT NULL DEFAULT (UUID())
);

INSERT INTO admin(email, password) VALUES('admin@admin.com', 'admin')
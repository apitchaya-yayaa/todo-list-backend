CREATE TABLE tasks (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    active tinyint
);

INSERT INTO tasks(title, active) VALUES
('A', 1),
('B', 1),
('C', 1);
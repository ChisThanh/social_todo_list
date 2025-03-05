CREATE TABLE todo_items (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    image JSON DEFAULT NULL,
    description TEXT DEFAULT NULL,
    status ENUM('Doing', 'Done', 'Deleted') NULL DEFAULT 'Doing',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX status_index ON todo_items (status);

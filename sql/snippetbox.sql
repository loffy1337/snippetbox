/* Создание базы данных snippetbox */
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE snippetbox;

/* Создание таблицы для заметок и индекса на поле created */
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);
CREATE INDEX idx_snippets_created ON snippets(created);

/* Добавление нескольких тестовых записей */
INSERT INTO snippets (title, content, created, expires) VALUES (
    'Не имей сто рублей',
    'Не имей сто рублей,\nа имей сто друзей.',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
INSERT INTO snippets (title, content, created, expires) VALUES (
    'Лучше один раз увидеть',
    'Лучше один раз увидеть,\nчем сто раз услышать.',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
INSERT INTO snippets (title, content, created, expires) VALUES (
    'Не откладывай на завтра',
    'Не откладывай на завтра,\nчто можешь сделать сегодня.',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
CREATE TABLE IF NOT EXISTS comment (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	id_post INT,
	author TEXT,
	comment TEXT,
	likes INT DEFAULT 0,
	dislikes INT DEFAULT 0,
	created_at DATE DEFAULT (datetime('now','localtime')),
	FOREIGN KEY (author) REFERENCES user(username)
);
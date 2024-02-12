CREATE TABLE IF NOT EXISTS post (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	author TEXT,
	title TEXT,
	description TEXT,
	imageURL TEXT,
	likes INT DEFAULT 0,
	dislikes INT DEFAULT 0,
	category TEXT,
	created_at DATE DEFAULT (datetime('now','localtime')),
	FOREIGN KEY (author) REFERENCES user(username)
);
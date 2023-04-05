CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY,
	forname TEXT NOT NULL,
	surname TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	birthday DATETIME,
	nickname TEXT,
	about TEXT,
	image_path TEXT,
	created_at DATETIME,
	is_public BOOL NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS followers(
	id INTEGER PRIMARY KEY,
	following_id INTEGER NOT NULL,
	follower_id INTEGER NOT NULL,
	accepted bool NOT NULL,
	active bool NOT NULL,
		FOREIGN KEY (following_id)
			REFERENCES users(id)	
		FOREIGN KEY (follower_id)
			REFERENCES users(id)
);



CREATE TABLE IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY,
	user_id INTEGER NOT NULL,
	privacy_type_id INTEGER NOT NULL,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME NOT NULL,
	image_path TEXT,
		FOREIGN KEY (user_id) 
			REFERENCES users(id)
		FOREIGN KEY (privacy_type_id) 
			REFERENCES privacy_type(id)
);


CREATE TABLE IF NOT EXISTS privacy_type (
	id integer PRIMARY KEY NOT NULL,
	name text UNIQUE NOT NULL,
	description text
);

INSERT INTO privacy_type (id, name)
VALUES 
(0, "public"),
(1, "private"),
(2, "sub-private");


CREATE TABLE IF NOT EXISTS allowed_private_posts (
	id INTEGER PRIMARY KEY,
	post_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) 
		REFERENCES users(id)
	FOREIGN KEY (post_id) 
		REFERENCES posts(id)
);


CREATE TABLE IF NOT EXISTS comments (
	id INTEGER PRIMARY KEY,
	post_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	content TEXT NOT NULL,
	image_path TEXT,
	created_at DATETIME NOT NULL,
	FOREIGN KEY (user_id) 
		REFERENCES users (id)
	FOREIGN KEY (post_id) 
		REFERENCES posts (id)
);

CREATE TABLE IF NOT EXISTS groups(
	id INTEGER PRIMARY KEY,
	creator_id INTEGER NOT NULL,
	title TEXT NOT NULL,
	description TEXT,
	created_at DATETIME NOT NULL,
		FOREIGN KEY (creator_id)
			REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS user_groups (
	id INTEGER PRIMARY KEY,
	user_id INTEGER NOT NULL,
	group_id INTEGER NOT NULL,
	joined_at DATETIME NOT NULL,
	FOREIGN KEY (user_id) 
		REFERENCES users (id)
	FOREIGN KEY (group_id) 
		REFERENCES groups (id)
);

CREATE TABLE IF NOT EXISTS group_events (
	id INTEGER PRIMARY KEY,
	group_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	created_at DATETIME NOT NULL,
	event_time DATETIME NOT NULL,
	timespan INTEGER,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	FOREIGN KEY (user_id) 
		REFERENCES users (id)
	FOREIGN KEY (group_id) 
		REFERENCES groups (id)
);


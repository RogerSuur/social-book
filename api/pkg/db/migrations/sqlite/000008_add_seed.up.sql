CREATE TABLE IF NOT EXISTS seed (
	is_seeded BOOL NOT NULL DEFAULT false
);

INSERT INTO seed
VALUES (false);
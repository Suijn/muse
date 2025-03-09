-- migrate:up
CREATE TABLE IF NOT EXISTS users (
	id uuid PRIMARY KEY,
	username varchar(255),
	password varchar(255)
);

-- migrate:down
DROP TABLE users;


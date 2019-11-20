CREATE TABLE users
(
    id TEXT PRIMARY KEY NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    username TEXT NOT NULL,
    avatar BYTEA
);

CREATE TABLE favorite
(
    id TEXT PRIMARY KEY NOT NULL,
    user_id TEXT NOT NULL,
    gif_id TEXT NOT NULL,
    gif_data BYTEA,
    gif_name TEXT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE hubs
(
    id TEXT PRIMARY KEY NOT NULL,
    logo BYTEA,
    user_id TEXT NOT NULL,
    is_private BOOLEAN NOT NULL,
    is_close BOOLEAN NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE chat_msg
(
    id TEXT PRIMARY KEY NOT NULL,
    hub_id TEXT NOT NULL,
    gif_data BYTEA,
    create_at TIMESTAMP,
    user_id TEXT NOT NULL,

    FOREIGN KEY (hub_id) REFERENCES hubs (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE hub_user
(
    id TEXT PRIMARY KEY NOT NULL,
    hub_id TEXT NOT NULL,
    user_id TEXT NOT NULL,

    FOREIGN KEY (hub_id) REFERENCES hubs (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
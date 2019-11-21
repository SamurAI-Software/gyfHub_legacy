CREATE TABLE users
(
    id TEXT PRIMARY KEY NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash BYTEA NOT NULL,
    username TEXT UNIQUE NOT NULL,
    mobile TEXT UNIQUE NOT NULL,
    verify_token TEXT UNIQUE NOT NULL,
    reset_pass_token TEXT UNIQUE NOT NULL,
    verified BOOLEAN NOT NULL,
    avatar BYTEA
);

CREATE TABLE gifs
(
    id TEXT PRIMARY KEY NOT NULL,
    gif_data BYTEA NOT NULL
);

CREATE TABLE user_category
(
    id TEXT PRIMARY KEY NOT NULL,
    user_id TEXT NOT NULL,
    category TEXT NOT NULL,

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

CREATE TABLE user_gif
(
    id TEXT PRIMARY KEY NOT NULL,
    user_id TEXT NOT NULL,
    gif_id TEXT UNIQUE NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (gif_id) REFERENCES gifs (id)
);

CREATE TABLE followers
(
    id TEXT PRIMARY KEY NOT NULL,
    user_id TEXT NOT NULL,
    follower_id TEXT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (follower_id) REFERENCES users (id)
);

CREATE TABLE user_favorite
(
    id TEXT PRIMARY KEY NOT NULL,
    user_id TEXT NOT NULL,
    gif_id TEXT NOT NULL,
    category TEXT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (gif_id) REFERENCES gifs (id)
);

CREATE TABLE chat_msg
(
    id TEXT PRIMARY KEY NOT NULL,
    hub_id TEXT NOT NULL,
    gif_id TEXT NOT NULL,
    create_at TIMESTAMP,
    user_id TEXT NOT NULL,

    FOREIGN KEY (hub_id) REFERENCES hubs (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (gif_id) REFERENCES gifs (id)
);

CREATE TABLE hub_user
(
    id TEXT PRIMARY KEY NOT NULL,
    hub_id TEXT NOT NULL,
    user_id TEXT NOT NULL,

    FOREIGN KEY (hub_id) REFERENCES hubs (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
DROP TABLE sessions;
DROP TABLE tasks;
DROP TABLE users;

CREATE TABLE users
(
    id         serial primary key,
    uuid       varchar(64)  not null unique,
    name       varchar(255) not null,
    email      varchar(255) not null unique,
    password   varchar(255) not null,
    created_at timestamp    not null
);

CREATE TABLE tasks
(
    id         serial primary key,
    uuid       varchar(64)  not null unique,
    body       varchar(255) not null,
    user_id    integer references users (id),
    created_at timestamp    not null,
    updated_at timestamp    not null
);

CREATE TABLE sessions
(
    id         serial primary key,
    uuid       varchar(64) not null unique,
    user_id    integer references users (id) unique,
    created_at timestamp   not null
);
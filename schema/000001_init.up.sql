CREATE TABLE users (
    id serial not null unique,
    username varchar(255) not null,
    email varchar(255) not null unique,
    password varchar(255) not null
);
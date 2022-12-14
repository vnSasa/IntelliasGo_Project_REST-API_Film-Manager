CREATE TABLE users
(
    id serial not null unique,
    login varchar(255) not null unique,
    password_hash varchar(255) not null,
    age varchar(255) not null
);

CREATE TABLE director
(
    id serial not null unique,
    name varchar(255) not null unique,
    date_of_birth varchar(255) not null
);

CREATE TABLE film
(
    id serial not null unique,
    name varchar(255) not null unique,
    genre varchar(255) not null,
    director_id int references director (id) on delete cascade not null,
    rate float(5) not null,
    year int not null,
    minutes float(5) not null
);

CREATE TABLE favourite
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    film_id int references film (id) on delete cascade not null
);

CREATE TABLE wishlist
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    film_id int references film (id) on delete cascade not null
);
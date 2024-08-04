CREATE TABLE users (
    id serial not null unique ,
    name varchar(255) not null ,
    username varchar(255) not null unique ,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists(
    id serial not null unique ,
    title varchar(255) not null ,
    description varchar(255)
);

CREATE TABLE todo_item(
    id SERIAL not null unique ,
    title varchar(255) not null ,
    description varchar(255) not null,
    done boolean not null default false
);

CREATE TABLE user_lists(
    id serial not null unique ,
    user_id int references users(id) on delete cascade not null ,
    list_id int references todo_lists(id)  on delete cascade not null
);

create TABLE list_items(
    id serial not null unique ,
    list_id int references todo_lists(id)  on delete cascade not null ,
    item_id int references todo_item(id)  on delete cascade not null
);
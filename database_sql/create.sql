create table users (
    user_id serial primary key
);

create table segments (
    segment_id serial primary key,
    segment_name varchar(50) unique
);

create table history (
    history_id serial primary key,
    user_id integer not null,
    segment_name varchar(50) not null,
    added_at timestamp with time zone not null,
    deleted_at timestamp with time zone
);

create table adding (
    adding_id serial primary key,
    user_id integer references users(user_id),
    added_at timestamp with time zone not null,
    segment_id integer references segments(segment_id) on delete cascade,
    rm_at interval
)
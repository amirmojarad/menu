-- +goose Up
create table users(
    id serial primary key,
    first_name varchar(255) not null default '',
    last_name varchar(255) not null default '',
    email varchar(255) not null default '',
    hashed_password varchar(255) not null default '',
    "created_at" timestamptz default current_timestamp,
    "updated_at" timestamptz default current_timestamp,
    "deleted_at" timestamptz default null
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

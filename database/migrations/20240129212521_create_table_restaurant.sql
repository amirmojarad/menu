-- +goose Up
create table restaurant(
                           id serial primary key,
                           title varchar(255) not null,
                           is_active bool default true,
                           user_id bigint not null,
                           "created_at" timestamptz default current_timestamp,
                           "updated_at" timestamptz default current_timestamp,
                           "deleted_at" timestamptz default null,
                           constraint fk_user foreign key (user_id) references users(id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

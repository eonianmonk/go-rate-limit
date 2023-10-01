-- +migrate Up
create table if not exists rate (
    id integer primary key,
    hits smallint,
    tstamp integer
);

-- +migrate Down
drop table rate;
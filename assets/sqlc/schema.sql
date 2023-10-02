create table if not exists rate (
    id integer primary key,
    hits smallint NOT NULL,
    tstamp timestamp NOT NULL DEFAULT NOW()
);
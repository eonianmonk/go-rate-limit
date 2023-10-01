// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: queries.sql

package data

import (
	"context"
	"database/sql"
)

const deleteOld = `-- name: DeleteOld :exec
delete from rate where tstamp < now() - interval '1 minute'
`

func (q *Queries) DeleteOld(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteOld)
	return err
}

const hitRate = `-- name: HitRate :execresult
insert into rate (id, hits, tstamp) values ($1, 0, now())
on conflict(id) do 
UPDATE SET hits = CASE 
    WHEN rate.hits < 50 THEN rate.hits + 1
    ELSE rate.hits
  END
returning id, hits, tstamp
`

func (q *Queries) HitRate(ctx context.Context, id int32) (sql.Result, error) {
	return q.db.ExecContext(ctx, hitRate, id)
}

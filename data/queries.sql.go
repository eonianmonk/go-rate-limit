// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: queries.sql

package data

import (
	"context"
)

const deleteOld = `-- name: DeleteOld :exec
delete from rate where tstamp < now() - interval '1 minute'
`

func (q *Queries) DeleteOld(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteOld)
	return err
}

const hitRate = `-- name: HitRate :one
insert into rate (id, hits, tstamp) values ($1, 0, now())
on conflict(id) do UPDATE 
SET hits = CASE 
    WHEN rate.hits < $2 THEN rate.hits + 1
    ELSE rate.hits
  END
returning hits
`

type HitRateParams struct {
	ID   int32
	Hits int16
}

func (q *Queries) HitRate(ctx context.Context, arg HitRateParams) (int16, error) {
	row := q.db.QueryRowContext(ctx, hitRate, arg.ID, arg.Hits)
	var hits int16
	err := row.Scan(&hits)
	return hits, err
}

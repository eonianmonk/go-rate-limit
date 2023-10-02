-- name: HitRate :one
insert into rate (id, hits, tstamp) values ($1, 0, now())
on conflict(id) do UPDATE 
SET hits = CASE 
    WHEN rate.hits < $2 THEN rate.hits + 1
    ELSE rate.hits
  END
returning hits;

-- name: DeleteOld :exec
delete from rate where tstamp < now() - interval '1 minute';

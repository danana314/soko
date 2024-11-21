-- name: GetTrip :one
SELECT trip_id, name, start_date, end_date
FROM trips
WHERE trip_id=?;

-- name: GetUsers :many
SELECT user_id, name
FROM users
WHERE trip_id=?;

-- name: GetSchedule :many
SELECT s.date, s.user_id, u.name
FROM schedule s
	INNER JOIN users u on s.user_id = u.user_id
WHERE s.trip_id=?;

-- name: SaveTripDetails :execresult
INSERT INTO trips(
    trip_id, name, start_date, end_date
) VALUES (
    ?, ?, ?, ?
) ON CONFLICT(
    trip_id
) DO UPDATE SET
	name=excluded.name,
	start_date=excluded.start_date,
	end_date=excluded.end_date;

-- name: AddUser :execresult
INSERT INTO users(
    user_id, trip_id, name
) VALUES (
    ?, ?, ?
);

-- name: DeleteSchedule :exec
DELETE FROM schedule
WHERE trip_id=?;

-- name: AddSchedule :execresult
INSERT INTO schedule(trip_id, user_id, date)
VALUES (?, ?, ?);

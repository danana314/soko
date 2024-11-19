-- name: GetTrip :one
SELECT tripId, name, startDate, endDate
FROM trips
WHERE tripId=?;

-- name: GetUsers :many
SELECT userId, name
FROM users
WHERE tripId=?;

-- name: GetSchedule :many
SELECT s.date, s.userId, u.name
FROM schedule s
	INNER JOIN users u on s.userId = u.userId
WHERE s.tripId=?;

-- name: SaveTripDetails :execresult
INSERT INTO trips(
    tripId, name, startDate, endDate
) VALUES (
    ?, ?, ?, ?
) ON CONFLICT(
    tripId
) DO UPDATE SET
	name=excluded.name,
	startDate=excluded.startDate,
	endDate=excluded.endDate;

-- name: AddUser :execresult
INSERT INTO users(
    userId, tripId, name
) VALUES (
    ?, ?, ?
);

-- name: DeleteSchedule :exec
DELETE FROM schedule
WHERE tripId=?;

-- name: AddSchedule :execresult
INSERT INTO schedule(tripId, userId, date)
VALUES (?, ?, ?);

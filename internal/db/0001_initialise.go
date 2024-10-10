package db

var init_db = []string{create_trips, create_users, create_schedule}

const create_trips string = `CREATE TABLE IF NOT EXISTS trips (
    id integer primary key autoincrement,
    ref text unique,
    name text,
    start_date date,
    end_date date
);`

const create_users string = `CREATE TABLE IF NOT EXISTS users (
    id integer primary key autoincrement,
    trip_id integer,
    name text,
    foreign key(trip_id) references trips(id)
);`

const create_schedule string = `CREATE TABLE IF NOT EXISTS schedule (
    pk integer primary key autoincrement,
    trip_id integer,
    user_id integer,
    date date,
    foreign key(trip_id) references trips(id),
    foreign key(user_id) references users(id)
);`

var seed_db = []string{insert_trip, insert_users}

const insert_trip string = `
	INSERT INTO trips (id, ref, name, start_date, end_date)
	VALUES (1, 'test', 'test', '2024-01-14', '2024-02-10');
`

const insert_users string = `
	INSERT INTO users (id, trip_id, name)
	VALUES
		(1, 1, 'John Smith'),
		(2, 1, 'Jane Smith'),
		(3, 1, 'Will I Am');
`

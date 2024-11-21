CREATE TABLE IF NOT EXISTS trips (
    trip_id text primary key,
    name text,
    start_date date,
    end_date date
);

CREATE TABLE IF NOT EXISTS users (
    user_id text primary key,
    trip_id text,
    name text,
    foreign key(trip_id) references trips(trip_id)
);

CREATE TABLE IF NOT EXISTS schedule (
    -- pk integer primary key autoincrement,
    trip_id text,
    user_id text,
    date date,
    foreign key(trip_id) references trips(trip_id),
    foreign key(user_id) references users(user_id),
    unique(trip_id, user_id, date)
);

CREATE TABLE IF NOT EXISTS expenses (
	-- pk integer primary key autoincrement,
	trip_id text,
	date date,
	description text,
	amount decimal(10,2),
	paid_by_user_id text,
	participants blob,
	foreign key(trip_id) references trips(trip_id),
	foreign key(paid_by_user_id) references users(user_id)
);

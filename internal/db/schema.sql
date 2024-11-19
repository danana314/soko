CREATE TABLE IF NOT EXISTS trips (
    tripId text primary key,
    name text,
    startDate date,
    endDate date
);

CREATE TABLE IF NOT EXISTS users (
    userId text primary key,
    tripId text,
    name text,
    foreign key(tripId) references trips(tripId)
);

CREATE TABLE IF NOT EXISTS schedule (
    -- pk integer primary key autoincrement,
    tripId text,
    userId text,
    date date,
    foreign key(tripId) references trips(tripId),
    foreign key(userId) references users(userId),
    unique(tripId, userId, date)
);

CREATE TABLE IF NOT EXISTS expenses (
	-- pk integer primary key autoincrement,
	tripId text,
	date date,
	description text,
	amount decimal(10,2),
	paidByUserId text,
	participants blob,
	foreign key(tripId) references trips(tripId),
	foreign key(paidByUserId) references users(userId)
);

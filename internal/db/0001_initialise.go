package db

var init_db = []string{createTrips, createUsers, createSchedule}

const createTrips string = `CREATE TABLE IF NOT EXISTS trips (
    tripId text primary key,
    name text,
    startDate date,
    endDate date
);`

const createUsers string = `CREATE TABLE IF NOT EXISTS users (
    userId text primary key,
    tripId text,
    name text,
    foreign key(tripId) references trips(tripId)
);`

const createSchedule string = `CREATE TABLE IF NOT EXISTS schedule (
    pk integer primary key autoincrement,
    tripId integer,
    userId integer,
    date date,
    foreign key(tripId) references trips(tripId),
    foreign key(userId) references users(userId),
    unique(tripId, userId, date)
);`

var seed_db = []string{insertTrip, insertUsers}

const insertTrip string = `
	INSERT INTO trips (tripId, name, startDate, endDate)
	VALUES ('test', 'test', '2024-01-14', '2024-02-10');
`

const insertUsers string = `
	INSERT INTO users (userId, tripId, name)
	VALUES
		('testuser1', 'test', 'John Smith'),
		('testuser2', 'test', 'Jane Smith'),
		('testuser3', 'test', 'Will I Am');
`

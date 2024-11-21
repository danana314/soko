package store

var seed_db = []string{insertTrip, insertUsers, insertExpenses}

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
const insertExpenses string = `
	INSERT INTO expenses (tripId, date, description, amount, paidByUserId, participants)
	VALUES ('test', '2024-01-15', 'food at restaurant', 143.50, 'testuser3', '');
`

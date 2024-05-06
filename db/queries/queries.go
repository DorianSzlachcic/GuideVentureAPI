package queries

const CREATE = `CREATE TABLE IF NOT EXISTS activities (
				id INTEGER NOT NULL PRIMARY KEY,
				time DATETIME NOT NULL,
				description TEXT
				);`

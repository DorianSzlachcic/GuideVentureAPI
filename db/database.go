package db

type Database interface {
	NewDb() error
}

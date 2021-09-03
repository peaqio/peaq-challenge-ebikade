package storage

// Type defines available storage types
type Type string

const (
	// MYSQL will store data in Mysql Database
	MYSQL Type = "mysql"
	// Memory will store data in memory
	Memory Type = "memory"
)

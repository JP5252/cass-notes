package database

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func (db *Database) GetDbConnection() error {

	return nil
}

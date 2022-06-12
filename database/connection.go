package database

var connection string

func init() {
	connection = "Mysql"
}

func getDatabase() string {
	return connection
}

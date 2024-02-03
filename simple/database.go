package simple

type Database struct {
	Name string
}

type DatabasePostreSQL Database
type DatabaseMongoDB Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

func NewDatabasePostreSQL() *DatabasePostreSQL {
	return (*DatabasePostreSQL)(&Database{Name: "PostgreSQL"})

}

type DatabaseRepository struct {
	DatabasePostreSQL *DatabasePostreSQL
	DatabaseMongoDB   *DatabaseMongoDB
}

func NewDatabaseRepository(databasePostreSQL *DatabasePostreSQL, databaseMongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostreSQL: databasePostreSQL,
		DatabaseMongoDB:   databaseMongoDB,
	}
}

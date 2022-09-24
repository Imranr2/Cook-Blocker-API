package database

type DBConfig struct {
	username string
	password string
	net      string
	port     string
	dbName   string
}

func NewDBConfig(username string, password string, net string, port string, dbName string) *DBConfig {
	return &DBConfig{
		username: username,
		password: password,
		net:      net,
		port:     port,
		dbName:   dbName,
	}
}

func (dbConfig *DBConfig) GetUsername() string {
	return dbConfig.username
}

func (dbConfig *DBConfig) GetPassword() string {
	return dbConfig.password
}

func (dbConfig *DBConfig) GetNet() string {
	return dbConfig.net
}

func (dbConfig *DBConfig) GetPort() string {
	return dbConfig.port
}

func (dbConfig *DBConfig) GetDBName() string {
	return dbConfig.dbName
}

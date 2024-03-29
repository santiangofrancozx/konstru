package config

var DB_DSN string

func SetDSN(dsn string) {
	DB_DSN = dsn
}

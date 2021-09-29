package infrastructure

import "os"

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
		Test struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {

	c := new(Config)

	c.DB.Production.Host = os.Getenv("DB_HOST")
	c.DB.Production.Username = os.Getenv("DB_USERNAME")
	c.DB.Production.Password = os.Getenv("DB_USERPASS")
	c.DB.Production.DBName = os.Getenv("DB_NAME")

	c.DB.Test.Host = "repgram_db"
	c.DB.Test.Username = "user"
	c.DB.Test.Password = "password"
	c.DB.Test.DBName = "mydb"

	c.Routing.Port = ":3001"

	return c
}

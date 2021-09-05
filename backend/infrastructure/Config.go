package infrastructure

type Config struct {
    DB struct {
        Production struct {
            Host string
            Username string
            Password string
            DBName string
        }
        Test struct {
            Host string
            Username string
            Password string
            DBName string
        }
    }
    Routing struct {
        Port string
    }
}

func NewConfig() *Config {

    c := new(Config)

    c.DB.Production.Host = "repgram_db"
    c.DB.Production.Username = "root"
    c.DB.Production.Password = "password"
    c.DB.Production.DBName = "mydb"

    c.DB.Test.Host = "repgram_db"
    c.DB.Test.Username = "root"
    c.DB.Test.Password = "password"
    c.DB.Test.DBName = "mydb"

    c.Routing.Port = ":3001"

    return c
}

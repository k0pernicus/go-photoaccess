package app

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	// C is the internal service configuration variable
	C Configuration
	// DB is the PostgreSQL connection pointer
	DB *pgxpool.Pool
)

// AppConfiguration handles the proper configuration of the service
type AppConfiguration struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// String formats and returns the URL of the service
func (c AppConfiguration) String() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

// DBConfiguration handles the informations about the configuration of the database
type DBConfiguration struct {
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
	Password string `yaml:"pswd"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

// String formats and returns the URL of the DB (for pgxpool)
func (c DBConfiguration) String() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.Name)
}

// Configuration handles all the informations for creation & launch
type Configuration struct {
	App AppConfiguration `yaml:"app"`
	DB  DBConfiguration  `yaml:"db"`
}

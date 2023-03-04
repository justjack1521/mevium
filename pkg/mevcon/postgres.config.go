package mevcon

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Dialect  string `default:"postgres"`
	Username string `required:"true"`
	Password string `required:"true"`
	Host     string `required:"true"`
	Name     string `required:"true"`
	Port     int
	SSLMode  bool
}

func (c PostgresConfig) getSource() string {
	conn := fmt.Sprintf("port=%d host=%s user=%s "+"password=%s dbname=%s sslmode=require", c.Port, c.Host, c.Username, c.Password, c.Name)
	return conn
}

func (c PostgresConfig) NewPostgresConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.getSource()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

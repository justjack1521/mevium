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

func (pc PostgresConfig) getSource() string {
	conn := fmt.Sprintf("port=%d host=%s user=%s "+"password=%s dbname=%s sslmode=require", pc.Port, pc.Host, pc.Username, pc.Password, pc.Name)
	return conn
}

func NewPostgresConnection(config PostgresConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.getSource()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

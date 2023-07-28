package db

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Db struct {
	Postgres struct {
		DatabaseName string `yaml:"databaseName"`
		DriverName   string `yaml:"driverName"`
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
	} `yaml:"postgres"`
	Redis struct {
		DatabaseName string `yaml:"databaseName"`
		DriverName   string `yaml:"driverName"`
	} `yaml:"redis"`
}

func New() *Db {
	return &Db{}
}

func (db *Db) LoadConfig(filename string) error {

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, db)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) Connect(driverName string) (*sql.DB, error) {

	err := db.LoadConfig("config/main.yml")
	if err != nil {
		fmt.Printf("Error load config %v", err)
	}

	param := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.Postgres.Host, db.Postgres.Port, db.Postgres.User, db.Postgres.Password, db.Postgres.DatabaseName)

	connect, err := sql.Open(driverName, param)
	if err != nil {
		return nil, err
	}

	return connect, nil
}

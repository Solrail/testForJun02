package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testForJun02/internal/config/db"
	"testForJun02/internal/models"
)

func SelectUser(ctx context.Context, num int) (models.User, error) {

	var user models.User

	db := db.New()

	err := db.LoadConfig("config/main.yml")
	if err != nil {
		fmt.Printf("Error load config %v", err)
	}

	conn, err := db.Connect(db.Postgres.DriverName)
	if err != nil {
		return user, fmt.Errorf("problem with connection: %v", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			err = fmt.Errorf("connection not closed %w", err)
			log.Println(err.Error())
		}
	}(conn)

	err = conn.QueryRowContext(ctx, `SELECT * FROM users WHERE id = $1`, num).Scan(&user.Id, &user.Name, &user.Surname, &user.Birthday)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}

func SelectUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	db := db.New()

	err := db.LoadConfig("config/main.yml")
	if err != nil {
		fmt.Printf("Error load config %v", err)
	}

	conn, err := db.Connect(db.Postgres.DriverName)
	if err != nil {
		return users, errors.New("problem with connection")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			err = fmt.Errorf("connection not closed %w", err)
			log.Println(err.Error())
		}
	}(conn)

	rows, err := conn.QueryContext(ctx, `SELECT * FROM users`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return users, errors.New("nothing found")
		}
		return users, err
	}

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Surname, &user.Birthday)
		if err != nil {
			err = fmt.Errorf("something wrong %w", err)
			log.Println(err.Error())
		}
		users = append(users, user)
	}

	return users, nil
}

func AddUser(ctx context.Context, user models.User) (string, error) {
	//var user models.User

	db := db.New()

	err := db.LoadConfig("config/main.yml")
	if err != nil {
		fmt.Printf("Error load config %v", err)
	}

	conn, err := db.Connect(db.Postgres.DriverName)
	if err != nil {
		return "", errors.New("problem with connection")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			err = fmt.Errorf("connection not closed %w", err)
			log.Println(err.Error())
		}
	}(conn)

	_, err = conn.ExecContext(ctx, "INSERT INTO users (name, surname, birthday) VALUES ($1, $2, $3)", user.Name, user.Surname, user.Birthday)
	if err != nil {
		return "", errors.New("problem with add new user")
	}
	return "Inserted  successfully", nil
}

func DelUser(ctx context.Context, id int) (string, error) {
	db := db.New()

	err := db.LoadConfig("config/main.yml")
	if err != nil {
		fmt.Printf("Error load config %v", err)
	}

	conn, err := db.Connect(db.Postgres.DriverName)
	if err != nil {
		return "", errors.New("problem with connection")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			err = fmt.Errorf("connection not closed %w", err)
			log.Println(err.Error())
		}
	}(conn)

	_, err = conn.ExecContext(ctx, "DELETE FROM users where id = $1", id)
	if err != nil {
		return "", err
	}

	return "Delete  successfully", nil

}

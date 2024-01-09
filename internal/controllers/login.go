package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"testForJun02/internal/config/db"
	"testForJun02/internal/models"
	"time"
)

var secretKey = "qweasdzxc"

func generateJWT() string {
	exp := time.Now().Add(time.Hour * 72).Unix()
	claim := jwt.MapClaims{
		"exp": exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		err = fmt.Errorf("something wrong %w", err)
		log.Println(err.Error())
	}
	return tokenString
}

func CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header["Token"] != nil {
			token, err := jwt.Parse(c.Request().Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Some error")
				}
				return []byte(secretKey), nil
			})
			if err != nil {
				return c.String(http.StatusUnauthorized, err.Error())
			}

			if token.Valid {
				return next(c)
			}
		}
		return c.String(http.StatusUnauthorized, "Not Authorized")
	}
}

func CheckLogin(ctx context.Context, login models.Login) (bool, error) {
	db := db.New()

	err := db.LoadConfig("config/main.yml")
	if err != nil {
		fmt.Printf("Error load config %v", err)
	}

	conn, err := db.Connect(db.Postgres.DriverName)
	if err != nil {
		return false, errors.New("problem with connection")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			err = fmt.Errorf("connection not closed %w", err)
			log.Println(err.Error())
		}
	}(conn)

	var row models.Login
	err = conn.QueryRowContext(ctx, "SELECT password FROM login where name = $1", login.Name).Scan(&row.Password)
	if err != nil {
		return false, errors.New("problem with select")
	}

	if row.Password != login.Password {
		return false, errors.New("not correct password")
	}

	return true, nil
}

func Login(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Content-Type", "application/json")
	var login models.Login
	err := json.NewDecoder(c.Request().Body).Decode(&login)
	if err != nil {
		errMsg := fmt.Errorf("wrong operation: %w", err)
		return errMsg
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := CheckLogin(ctx, login)
	if err != nil {
		errMsg := fmt.Errorf("error of authorization: %w", err)
		json.NewEncoder(c.Response()).Encode("error of authorization")
		return errMsg
	}
	if res {
		validToken := generateJWT()
		err := json.NewEncoder(c.Response()).Encode(validToken)
		if err != nil {
			err = fmt.Errorf("something wrong with validToken %w", err)
			log.Println(err.Error())
		}
	}

	return nil
}

// Package controllers
package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"testForJun02/internal/db"
	"testForJun02/internal/models"
	"time"
)

// Add new User in base
// @Summary      Add user into db
// @Description  add user
// @Tags         add
// @Accept       json
// @Produce      json
// @Param        user body models.User true "Add user" exclude:id
// @Success      200  {object} models.User
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router       /adduser [post]
func Add(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		http.Error(c.Response(), err.Error(), http.StatusBadRequest)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			err = fmt.Errorf("body not closed %w", err)
			log.Println(err.Error())
		}
	}(c.Request().Body)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.AddUser(ctx, user)
	if err != nil {
		err = fmt.Errorf("%w", err)
		log.Println(err.Error())
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(c.Response()).Encode(res)
	if err != nil {
		err = fmt.Errorf("something wrong %w", err)
		log.Println(err.Error())
	}
	return nil
}

// GetById returns the user by ID
// @Summary Get user by ID
// @Description Get user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id body int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user [get]
func GetById(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Content-Type", "application/json")
	var data map[string]int

	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		errMsg := fmt.Errorf("wrong operation: %w", err)
		return errMsg
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			err = fmt.Errorf("body not closed %w", err)
			log.Println(err.Error())
		}
	}(c.Request().Body)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.SelectUser(ctx, data["id"])
	if err != nil {
		err = fmt.Errorf("%w", err)
		log.Println(err.Error())
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
		return err
	}

	err = json.NewEncoder(c.Response()).Encode(res)
	if err != nil {
		err = fmt.Errorf("something wrong %w", err)
		log.Println(err.Error())
	}
	return nil
}

// GetAll returns all users
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} string
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [get]
func GetAll(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.SelectUsers(ctx)
	if err != nil {
		err = fmt.Errorf("something wrong %w", err)
		log.Println(err.Error())
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(c.Response()).Encode(res)
	if err != nil {
		err = fmt.Errorf("something wrong %w", err)
		log.Println(err.Error())
	}
	return nil
}

// DelById deletes a user by ID
// @Summary Delete user by ID
// @Description Delete user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id body integer true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /deluser [delete]
func DelById(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Content-Type", "application/json")

	var data map[string]int

	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		errMsg := fmt.Errorf("wrong operation: %w", err)
		return errMsg
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			err = fmt.Errorf("body not closed %w", err)
			log.Println(err.Error())
		}
	}(c.Request().Body)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.DelUser(ctx, data["id"])
	if err != nil {
		err = fmt.Errorf("%w", err)
		log.Println(err.Error())
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(c.Response()).Encode(res)
	if err != nil {
		err = fmt.Errorf("something wrong %w", err)
		log.Println(err.Error())
	}
	return nil
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

	res, err := db.CheckLogin(ctx, login)
	if err != nil {
		fmt.Println("3")
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

var secretKey = []byte("sdk;sd13axx")

func generateJWT() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(secretKey)
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
				return secretKey, nil
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

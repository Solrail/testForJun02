// Package controllers
package controllers

import (
	"context"
	"encoding/json"
	"fmt"
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

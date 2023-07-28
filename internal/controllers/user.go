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

// Add new user in base
// @Summary      Add user into db
// @Description  add user
// @Tags         add
// @Accept       json
// @Produce      json
// @Param        user body  models.User true "Add user"
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

func DelById(c echo.Context) error {

	/*	w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")*/

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

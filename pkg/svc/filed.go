package svc

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Filed struct {
	FirstName string
	LastName  string
	Age       int
}

func FiledService(c *gin.Context) {
	filed, err := NewFiled(c.PostForm("FirstName"), c.PostForm("LastName"), c.PostForm("Age"))

	if err != nil {
		FieldExceptionHandler(err)
		c.String(500, "Exception: %s", FieldExceptionHandler(err))
		return
	}

	jsonBytes, err := json.Marshal(filed)
	if err != nil {
		c.String(500, "Exception: %s", FieldExceptionHandler(err))
		return
	}

	slog.Info("Value of 'field' is: " + string(jsonBytes))
	c.String(200, "Value of 'field' is: %s", string(jsonBytes))
}

func NewFiled(firstName, lastName, ages string) (*Filed, error) {
	age, err := strconv.Atoi(ages)
	filed := new(Filed)

	if err != nil {
		return filed, err
	}

	filed.FirstName = firstName
	filed.LastName = lastName
	filed.Age = age

	return filed, nil
}

func FieldExceptionHandler(err error) string {
	return fmt.Sprintf("Error: %s", err)
}

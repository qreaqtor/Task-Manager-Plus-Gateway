package utils

import (
	"net/http"
	"os"
)

const (
	Get    = "GET"
	Post   = "POST"
	Patch  = "PATCH"
	Delete = "DELETE"
)

var (
	client        *http.Client
	API_SECRET    string
	USERS_BACKEND string
	TASKS_BACKEND string
)

func init() {
	client = &http.Client{}

	API_SECRET = os.Getenv("API_SECRET")
	USERS_BACKEND = os.Getenv("USERS_BACKEND")
	TASKS_BACKEND = os.Getenv("TASKS_BACKEND")
}

package utils

import "os"

var API_SECRET string

func init() {
	API_SECRET = os.Getenv("API_SECRET")
}

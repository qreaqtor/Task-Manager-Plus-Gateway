package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func SendRequest(method, endpoint string, jsonData []byte) (map[string]interface{}, int) {
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return map[string]interface{}{"message": "invalid request"}, http.StatusBadRequest
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return map[string]interface{}{"message": "something gone wrong"}, http.StatusInternalServerError
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{"message": "invalid data"}, http.StatusBadRequest
	}

	var outputData map[string]interface{}
	err = json.Unmarshal(respData, &outputData)
	if err != nil {
		return map[string]interface{}{"message": "invalid data"}, http.StatusBadRequest
	}

	return outputData, resp.StatusCode
}

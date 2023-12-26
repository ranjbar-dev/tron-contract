package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "http://localhost:9090"

func post(path string, payload map[string]any) (map[string]any, error) {

	json_data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(baseUrl+path, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	var res map[string]interface{}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(body))

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func get(path string) (map[string]any, error) {

	resp, err := http.Get(baseUrl + path)
	if err != nil {
		return nil, err
	}

	var res map[string]any

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

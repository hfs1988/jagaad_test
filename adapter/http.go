package adapter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type errorResponse struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

type httpAdapter struct{}

func GetHTTPAdapter() *httpAdapter {
	return &httpAdapter{}
}

func (a *httpAdapter) Get(url string, data interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	respData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		var errResp errorResponse
		err = json.Unmarshal(respData, &errResp)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", errResp.Error.Message)
	}

	err = json.Unmarshal(respData, &data)
	if err != nil {
		return err
	}

	return nil
}

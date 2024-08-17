package helper

import (
	"fmt"
	"io"
	"net/http"
)

func GetBody(param string) ([]byte, error) {
	res, err := http.Get(param)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	body, err := io.ReadAll(res.Body)
	err = res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return body, nil
}

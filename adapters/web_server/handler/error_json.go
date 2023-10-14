package handler

import "encoding/json"

func jsonError(message string) []byte {
	_error := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}

	res, err := json.Marshal(_error)
	if err != nil {
		return []byte(err.Error())
	}

	return res
}

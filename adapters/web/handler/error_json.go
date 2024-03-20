package handler

import "encoding/json"

func jsonError(msg string) []byte {
	bs, err := json.Marshal(map[string]string{
		"message": msg,
	})
	if err != nil {
		return []byte(err.Error())
	}

	return bs
}

package utils

import "encoding/base64"

func GenerateApi(key string) string {
	data := []byte(key)
	apiKey := base64.StdEncoding.EncodeToString(data)
	return apiKey
}

package models

import (
	"encoding/json"
	"os"
)

type Config struct {
	Host         string       `json:"host"`
	Port         string       `json:"port"`
	Db           DbConnection `json:"db"`
	TemplatePath string       `json:"templateFolder"`
	StaticFolder string       `json:"staticFolder"`
	Policies     string       `json:"policies"`
}

type DbConnection struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func MapJson(file string, obj any) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, obj)
}

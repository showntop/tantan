package models

import (
	"fmt"
	// "github.com/showntop/tantan/server"
)

var _ = fmt.Println

type User struct {
	Id   int    `json:"user_id"`
	Name string `json:"name"`
}

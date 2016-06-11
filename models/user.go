package models

import (
	"fmt"
)

var _ = fmt.Println

type User struct {
	Id   int    `json:"id,string" sql:",pk"`
	Name string `json:"name"`
	Type string `json:"type"`
}

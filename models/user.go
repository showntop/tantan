package models

import (
	"fmt"
	// "github.com/showntop/tantan/server"
)

var _ = fmt.Println

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

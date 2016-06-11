package main

import (
	"fmt"

	"github.com/showntop/tantan/handlers"
	// "github.com/showntop/tantan/server"
	// _ "github.com/showntop/tantan/models"
	// "github.com/showntop/tantan/stores"
)

func main() {
	fmt.Println(".....................................start server.......................................")
	InitDb(Configure["database"].(map[string]string))
	handlers.Setup(Configure)
}

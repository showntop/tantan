package main

import (
	"log"

	. "github.com/showntop/tantan/config"
	"github.com/showntop/tantan/handlers"
	"github.com/showntop/tantan/schema"
)

func main() {
	log.Println("start server")
	schema.InitDb(Configure["database"].(map[string]string))
	handlers.Setup(Configure)
}
